package sitemap

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/smelton01/go-basics/link"
)

type Node struct {
	url string
	level int
	children []Node
}

type Map struct {
	rootURL string
	visitedURL map[string]bool
	toVisit  []*Node
	siteMapRoot Node

}

func SiteMap(rootURL string){
	fmt.Println("Going throu urlz")
	u, err := url.Parse(rootURL)
	CheckError(err)

	sitemap := Map{
		rootURL: u.Hostname(),
		visitedURL: make(map[string]bool),
		toVisit: []*Node{},
		siteMapRoot: Node{
			url: rootURL,
			level: 0,
			children: []Node{},
		},
	}

	sitemap.push(&sitemap.siteMapRoot)

	// var nextNode *node
	for {
		nextNode := sitemap.pop()
		if nextNode == nil {
			break
		}
		children, err := sitemap.getChildren(nextNode.url)
		CheckError(err)

		sitemap.addChildren(nextNode, children)
	}

	fmt.Println(sitemap.siteMapRoot)
}

func (m *Map) push(currNode *Node) {
	m.toVisit = append(m.toVisit, currNode)
}

func (m *Map) pop() *Node {
	if len(m.toVisit) == 0 {
		return nil
	}
	head := m.toVisit[0]
	m.toVisit = m.toVisit[1:]
	return head
}

func (m *Map) getChildren(parentURL string) ([]string, error) {
	resp, err := http.Get(parentURL)
	CheckError(err)

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	r := bytes.NewReader(html)

	children := link.Parse(r)
	childrenURL := []string{}

	for _, pair := range children {
		childrenURL = append(childrenURL, pair.Href)
	}

	return childrenURL, nil
}

func (m *Map) addChildren(currNode *Node, children []string){
	// Add all valid children of node note yet visited to the node and add the node to toVisit list
	if len(children) == 0 {
		return
	}

	for _, child := range children {

		fullURL, err := m.checkURL(child)
		if err != nil {
			log.Println(err)
			continue
		}

		if _,ok := m.visitedURL[fullURL]; ok {
			continue
		}

		currNode.children = append(currNode.children, Node{
			url: fullURL,
			level: currNode.level + 1,
			children: []Node{},
		})
		// mark URL as viisted
		m.visitedURL[child] = true
		
		if currNode.level > 1 {
			log.Println("Max depth reached")
			return
		}
		// add tail child node to visit queue
		m.push(&currNode.children[len(currNode.children)-1])
	}
	
}

func (m Map) checkURL(URL string) (string, error) {
	// Check if provided url is of valid format and in same domain as root url
	if strings.HasPrefix(URL, "http") && strings.Contains(URL, func (root string) string {
		if strings.HasPrefix(root,"www")  {
			return root[4:]
		}
		return root
	}(m.rootURL)) {
		return URL, nil
	}
	if len(URL) >0 && URL[0] == '/' {
		return "https://" + m.rootURL + URL, nil
	}

	return "", errors.New("Invalid URL: " + URL)
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}