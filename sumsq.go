package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int //number of test cases
	print("num cases??\n")
	fmt.Scanln(&N)
	count := N
	op := []string{}
	inter(&count, &op)
	//fmt.Printf("%T", op)
	//fmt.Println(sum)
	output(op)

	//var X int //number of
}
func inter(count_addr *int, op *[]string) int {
	if *count_addr != 0 {
		temp := input(count_addr) // Get each test case input
		x := strconv.Itoa(temp)
		*op = append(*op, x)
		//return temp +
		inter(count_addr, op) // not fine
		//in[i] = input(&count, &i)
	}
	return 0
}

func input(count *int) int {
	// input: count - number of remaining test cases
	// Output: sum1 - sum of squares for each test case
	var X int
	print("Num of elements: \n")
	fmt.Scanln(&X)
	fmt.Printf("Enter %d elements:", X)

	reader := bufio.NewReader(os.Stdin)
	//read a lineof input as string "23 11 44 55"
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	// Split string into array [23 11 44 55]
	arr := strings.Fields(text)

	sum1 := summ(arr, X)
	*count--
	return sum1
}

func summ(arr []string, X int) int {
	//input: arr - string slice of test case elements, X int number of elements
	//output sum of squares
	if X-1 >= 0 {
		//access text[X]
		num, err := strconv.Atoi(arr[X-1]) // refactor this same as output
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if num < 0 { //return 0 for negative numbers
			return 0 + summ(arr[:X-1], X-1)
		}
		return num*num + summ(arr[:X-1], X-1)
		//fmt.Println(num)
	}
	return 0
}
func output(arr []string) {
	// input: arr - slice of sum of squares for all test cases to print
	i := len(arr)
	if i == 0 {
		return
	} else if i == 1 {
		fmt.Println(arr[0])
		return
	}
	fmt.Println(arr[0])
	output(arr[1:])
	return
}
