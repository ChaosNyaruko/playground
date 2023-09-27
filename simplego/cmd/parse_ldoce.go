package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// doc, err := html.Parse(strings.NewReader(s))
	fd, err := os.Open("./doctor_ldoce.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	doc, err := html.ParseWithOptions(fd, html.ParseOptionEnableScripting(false))
	if err != nil {
		log.Fatal(err)
	}
	// Type      NodeType
	// DataAtom  atom.Atom
	// Data      string
	// Namespace string
	// Attr      []Attribute
	var f func(*html.Node)
	f = func(n *html.Node) {
		log.Printf("Type: [%#v], DataAtom: [%s], Data: [%#v], Namespace: [%#v], Attr: [%#v]", n.Type, n.DataAtom, n.Data, n.Namespace, n.Attr)
		readElement(n, "div", "dictionary")
		dictionary(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	// log.Printf("result: %v", readText(doc))
	f(doc)
}

func oneIter(n *html.Node, hit func(*html.Node), miss func(*html.Node)) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {

	}
}

func dictionary(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if readElement(c, "span", "ldoceEntry Entry") {
			for c := c.FirstChild; c != nil; c = c.NextSibling {
				// read "frequent head" for PRON
				if readElement(c, "span", "frequent Head") {
					for c := c.FirstChild; c != nil; c = c.NextSibling {
						if readElement(c, "span", "PronCodes") {
							for c := c.FirstChild; c != nil; c = c.NextSibling {
								if readElement(c, "span", "PRON") {
									fmt.Printf("result: PRON [%s]\n", readText(c))
								}
							}
						}
					}
				}
				// read Sense for DEF
				if readElement(c, "span", "Sense") {
					for c := c.FirstChild; c != nil; c = c.NextSibling {
						if readElement(c, "span", "DEF") {
							fmt.Printf("result: DEF [%s]\n", readText(c))
						}
					}
				}
			}
		}
	}
}

func readElement(n *html.Node, ele string, class string) bool {
	if n.Type == html.ElementNode && n.DataAtom.String() == ele {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == class {
				log.Printf("[wft] readElement good %v, %v, %#v", ele, class, n.Data)
				return true
			}
		}
	}
	return false
}

func readText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var s string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s += readText(c)
	}
	return s
}
