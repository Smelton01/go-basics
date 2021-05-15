package main

import "fmt"

func main(){
	//fmt.Println("This seem s hard")
	var A = 23
	var C int = 32
	var B string = "Random strhiu34w"
	D, E := 64,46
	var isbool bool = true
	flag := false
	fmt.Println(B,A+C,"Some joining sentence\n" ,len(B), A*2, B+" "+B, D+E)
	fmt.Println(flag == isbool,"\n", flag || isbool)
	//Pointers 
	var x = 34
	fmt.Println("x and its address is:", x, &x)
	changeV(&x)
	fmt.Println("x and its address is:", x, &x)
	fmt.Printf("type of flag: %T\n", flag)
	fmt.Printf("Bool of x: %t\n", x)
	fmt.Printf("Val of A/D: %4.3f\n", float32(A))
	fmt.Printf("binary %b, character code %c, %c, hex code %x\n", 23, 92,"#", 23)
}

func changeV(x *int){
	*x = 7
}

