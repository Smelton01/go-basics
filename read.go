package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	fmt.Print("-> ")

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	fmt.Printf("%T\n", text)
	fmt.Println(strings.Fields(text)) //split to slice by whitespace
	arr := strings.Fields(text)
	num, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(num)
	//if strings.Compare("hi", text) == 0 {
	//	fmt.Println("hello, Yourself")
	//}

}
