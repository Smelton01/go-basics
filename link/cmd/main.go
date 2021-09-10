package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smelton01/go-basics/link"
)

func main(){
	file := flag.String("file", "", "Path to html file to parse.")
	url := flag.String("url", "", "URL of page to parse")
	flag.Parse()
	
	var html []byte
	if *url != "" {
		resp, err := http.Get(*url)
		if err != nil {
			log.Fatal("Url Error: ", err)
		}
		defer resp.Body.Close()

		html, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}else if *file != "" {
		var err error
		html, err = ioutil.ReadFile(*file)
		if err != nil {
			fmt.Println(err)
		}
	}else {
		log.Fatal("Please provide url (-url) or file (-file) to parse.")
	}

	r := bytes.NewReader(html)

	output := link.Parse(r)
	fmt.Println(output)

}
