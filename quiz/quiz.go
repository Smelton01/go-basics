package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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
	var fileFlag = flag.String("csv", "problems.csv", "File path for the quiz csv in the form, 'question, answer'.")
	var timeFlag = flag.Int("limit", 30, "Time limit in seconds for quiz. Default: 30secs.")
	var shuffleFlag = flag.Bool("shuffle", false, "Set to true to shuffle the question order.")
	flag.Parse()

	quiz := quiz{
		file: *fileFlag,
		correct: 0,
		total: 0,
		limit: *timeFlag,
	}

	f, err := os.Open(quiz.file)
	if err != nil {
		log.Fatal("Unable to read input file: " + quiz.file, "\n", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
    quizLines, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV: " + quiz.file, err)
    }
	quiz.total = len(quizLines)

	if *shuffleFlag {
		quizLines = shuffle(quizLines)
	}

	c := make(chan int)
	go questions(quizLines, &quiz, c)

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

func shuffle(questions [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	return questions
}

func questions(lines [][]string, quiz *quiz, c chan int) {
	// Send signal through channel to start quiz
	var temp []byte
	fmt.Println("Press enter to start quiz....")
	n, _ := fmt.Scanln(&temp)
	c <- n

	// Iterate through each question and take answers htrough std input
	for i , line := range lines {
		question, expectedAns := line[0], line[1]

		expectedAns = strings.ToLower(strings.TrimSpace(expectedAns))

		fmt.Printf("Problem #%v: %s = ", i+1, question)

		var userAns []byte 
		_, err := fmt.Scanln(&userAns)

		if err != nil {
			fmt.Println("Incorrect input: ", err)
			continue
		}
		
		if strings.ToLower(string(userAns)) == expectedAns{
			quiz.correct += 1
		} 
	}
	c <- 1
	close(c)
}