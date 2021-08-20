package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const file = "problems.csv"

type quiz struct{
	file string;
	correct int;
	total int;
	limit int
}

func main() {
	var fileFlag = flag.String("csv", "problems.csv", "File name for the quiz csv")
	var timeFlag = flag.Int("limit", 30, "Time limit in seconds for quiz. Defaults to 30secs.")
	flag.Parse()

	quiz := quiz{
		file: *fileFlag,
		correct: 0,
		total: 0,
		limit: *timeFlag,

	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Unable to read input file: " + file, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
    lines, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV: " + file, err)
    }
	quiz.total = len(lines)

	c := make(chan int)
	go questions(lines, &quiz, c)

	// Wait for signal to start timer
	startFlag := <-c
	_ = startFlag

	// Start timer
	go timer(&quiz, c)
	
	endFlag := <-c
	_ = endFlag
	fmt.Println("\nTotal correct:  ", quiz.correct, "/", quiz.total)

}

func timer(quiz *quiz, c chan int){
	time.Sleep(time.Second*time.Duration(quiz.limit))
	fmt.Println("\nTime's up:")
	c <- 1
	close(c)
}

func questions(lines [][]string, quiz *quiz, c chan int) {
	// Send signal through channel to start quiz
	var temp []byte
	fmt.Println("Press enter to start quiz....")
	n, _ := fmt.Scanln(&temp)
	c <- n

	// Iterate through each question and take answers htrough std input
	for i , line := range lines {
		fmt.Printf("Problem #%v: %s = ", i+1, line[0])

		var ans []byte 
		_, err := fmt.Scanln(&ans)

		if err != nil {
			fmt.Println("Incorrect input: ", err)
			continue
		}
		
		if string(ans) == line[1]{
			quiz.correct += 1
		} 
	}
	c <- 1
	close(c)
}