package sitemap

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/smelton01/go-basics/link"
)
const queueSize = 100

type node struct {
	url string
	level int
	children []node
}

type Map struct {
	rootURL string
	visitedURL map[string]bool
	toVisit chan *node
	siteMapRoot node

}

func SiteMap(rootURL string){
	fmt.Println("Going throu urlz")
	u, err := url.Parse(rootURL)
	CheckError(err)

	sitemap := Map{
		rootURL: u.Hostname(),
		visitedURL: make(map[string]bool),
		toVisit: make(chan *node, queueSize),
		siteMapRoot: node{
			url: rootURL,
			level: 0,
			children: []node{},
		},
	}

	sitemap.toVisit <- &sitemap.siteMapRoot

	go sitemap.timer()

	for nextNode := range sitemap.toVisit {
		// nextNode := <-sitemap.toVisit
		fmt.Println(nextNode.url)
		children, err := sitemap.getChildren(nextNode.url)
		CheckError(err)

		sitemap.addChildren(nextNode, children)
	}

	fmt.Println(sitemap)
}

func (m *Map) timer(){
	time.Sleep(time.Duration(15)*time.Second)
	fmt.Println("Times up!!!!")
	close(m.toVisit)
}

func (m *Map) getChildren(parentURL string) ([]string, error) {
	resp, err := http.Get(parentURL)
	CheckError(err)

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	children := link.LinkFunc(html)
	childrenURL := []string{}

	for _, pair := range children {
		childrenURL = append(childrenURL, pair.Href)
	}

	return childrenURL, nil
}

func (m *Map) addChildren(currNode *node, children []string){
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

		currNode.children = append(currNode.children, node{
			url: fullURL,
			level: currNode.level + 1,
			children: []node{},
		})
		// mark URL as viisted
		m.visitedURL[child] = true
		
		if currNode.level > 3 {
			log.Println("Max depth reached")
			return
		}
		// add chiled node to visit queue
		m.toVisit <- &currNode.children[len(currNode.children)-1]
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
	if URL[0] == '/' {
		return "https://" + m.rootURL + URL, nil
	}

	return "", errors.New("Invalid URL: " + URL)
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}