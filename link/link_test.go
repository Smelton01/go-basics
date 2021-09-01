package link_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/smelton01/go-basics/link"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		file	string
		expect []link.Link
		
	}{
		{
			desc: "ex1",
			file: "examples/ex1.html",
			expect: []link.Link{{Href: "/other-page", Text: "A link to another page"}},
			
		},
		{
			desc: "ex2",
			file: "examples/ex2.html",
			expect: []link.Link{{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}, {Href: "https://github.com/gophercises", Text: "Gophercises is on Github !"}},
			
		},
		{
			desc: "ex3",
			file: "examples/ex3.html",
			expect: []link.Link{{Href: "#", Text: "Login"}, {Href: "/lost", Text: "Lost? Need help?"}, {Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"}},
			
		},
		{
			desc: "ex4",
			file: "examples/ex4.html",
			expect: []link.Link{{Href: "/dog-cat", Text: "dog cat"}},
			
		},
		{
			desc: "ex6",
			file: "examples/ex6.html",
			expect: []link.Link{{Href: "https://www.fukuoka-now.com/en/classified/archive/", Text: "Fukuoka Now"}, {Href: "https://github.com/Smelton01/Site-tracker", Text: "here"}},
			
		},
		
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			got := link.LinkFunc(func () []byte {
				html, err := ioutil.ReadFile(tC.file)
				if err != nil {
					fmt.Println(err)
				}
				return html
			}()); 
			if len(got) != len(tC.expect) {
				t.Errorf("Link func got: %v\n expected %v", got, tC.expect)
			}
			for i := range got {
				if got[i] != tC.expect[i] {
					t.Errorf("Link func %v got: %v\n expected %v", i, got[i], tC.expect[i])
				}
			}
		})
	}
}