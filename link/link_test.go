package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		file	string
		expect []Link
		
	}{
		{
			desc: "ex1",
			file: "ex1.html",
			expect: []Link{{Href: "/other-page", Text: "A link to another page"}},
			
		},
		{
			desc: "ex2",
			file: "ex2.html",
			expect: []Link{{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}, {Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"}},
			
		},
		{
			desc: "ex3",
			file: "ex3.html",
			expect: []Link{{Href: "#", Text: "Login"}, {Href: "Lost", Text: "Lost? Need help?"}, {Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"}},
			
		},
		{
			desc: "ex4",
			file: "ex4.html",
			expect: []Link{{Href: "/dog-cat", Text: "dog cat"}},
			
		},
		
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := LinkFunc(tC.file); 
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