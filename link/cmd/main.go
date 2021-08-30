package main

import (
	"flag"
	"fmt"

	"github.com/smelton01/go-basics/link"
)




func main(){
	file := flag.String("file", "ex1.html", "Path to html file to parse.")
	flag.Parse()

	output := link.LinkFunc(*file)
	fmt.Println(output)

}
