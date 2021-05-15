package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//defer stuff
	//r1 := Rectangle{length: 4, width: 2}
	//fmt.Scanln(&r1.length)
	//fmt.Println("Area of Rectangle r1 is ", r1.area())
	//lines()
	words()

}

type Rectangle struct {
	length float64
	width  float64
}

func (rec *Rectangle) area() float64 {
	return rec.length * rec.width
}
func lines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func words() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(os.Stdin)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
	// Output: 15
}
