package link

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	Text string
}

func Parse(r io.Reader) []Link{
	doc,err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	links := getLinks(doc)
	return links
}

func getLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := extractText(n)
				current := Link{
					Href: strings.TrimSpace(a.Val),
					Text: strings.TrimSpace(text),
				}
				links = append(links, current)
		}}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		allLinks := getLinks(c)
		links = append(links, allLinks...)
	}
	return links
}

func extractText(n *html.Node) string {
	var text string
	if n.Type != html.ElementNode && n.Type != html.CommentNode {
		text = n.Data
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		text += extractText(child)
	}
	return strings.Trim(text, "\n")
}