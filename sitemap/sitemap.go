package sitemap

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smelton01/go-basics/link"
)


func SiteMap(url string){
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	layer1 := link.LinkFunc(html)

	for _, pair := range layer1 {
		fmt.Println(pair)
	}
}