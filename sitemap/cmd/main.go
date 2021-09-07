package main

import (
	"flag"

	"github.com/smelton01/go-basics/sitemap"
)

func main(){
	url := flag.String("url", "example.com", "the url of the page to build sitemap from")
	flag.Parse()

	sitemap.SiteMap(*url)
}