package sitemap

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/smelton01/go-basics/link"
)

type node struct {
	url string
	level int
	children []node
}

type Map struct {
	rootURL string
	visitedURL map[string]bool
	siteMap node

}

func SiteMap(rootURL string){
	// u, err := url.parse(rootURL)
	// CheckError(err)
	fmt.Println("Going throu urlz")

	sitemap := Map{
		rootURL: rootURL,
		visitedURL: make(map[string]bool),
		siteMap: node{
			url: rootURL,
			level: 0,
			children: []node{},
		},
	}

	children, err := sitemap.getChildren()
	CheckError(err)
	// fmt.Println(children)

	sitemap.addChildren(&sitemap.siteMap, children)

	fmt.Println(sitemap)
}

func (m *Map) getChildren() ([]string, error) {
	resp, err := http.Get(m.rootURL)
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
	for _,child := range children {
		// re := regexp.MustCompile(m.rootURL)

		if _,ok := m.visitedURL[child]; ok {
			continue
		}

		currNode.children = append(currNode.children, node{
			url: child,
			level: currNode.level + 1,
			children: []node{},
		})
		m.visitedURL[child] = true
	}
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}