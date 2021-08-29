package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	Text string
}


func main(){
	file := flag.String("file", "ex1.html", "Path to html file to parse.")
	flag.Parse()

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(data)
	doc,err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	links := []Link{}

	dfs(doc, &links)
	fmt.Println(links)

}

func dfs (n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				current := Link{
					Href: a.Val,
					Text: n.FirstChild.Data,
				}
				*links = append(*links, current)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}