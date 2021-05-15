package main

import "fmt"

func main() {
	i := 55
	x := fact(i)
	f := fib(i)
	fmt.Printf("Calculating %d! it is: %d\n", i, x)
	fmt.Printf("Calculating fib number %d, it is %d\n", i, f)
	fmt.Println(sumofsq(2, 4, 5, 7, 3))
}

func sumofsq(args ...int) int { //
	sum := 0 
	for _, val := range args {
		sum += val * val
	}
	return sum
}
func fact(i int) int {
	fact := 1
	for n := 1; n <= i; n++ {
		fact *= n
	}
	return fact
}
func fact_r(i int) int {
	if i == 0 {
		return 1
	} else {
		return i * fact(i-1)
	}
}

func fib(num int) int {
	a := 1
	b := 1
	var fib int = 0
	if num < 0 {
		return 0
	} else if num == 2 {
		return 1
	} else {
		i := 2
		for i < num {
			fib = a + b
			a = b
			b = fib
			i++
		}
		return fib
	}
}
