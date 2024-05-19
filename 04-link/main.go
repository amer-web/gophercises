package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type link struct {
	Href string
	Text string
}

func main() {
	flagForHtmlFile := flag.String("html", "ex1.html", "specific file")
	flag.Parse()
	fileOpen, err := os.Open(*flagForHtmlFile)
	defer fileOpen.Close()
	if err != nil {
		log.Fatalf("can't open this file")
	}
	doc, err := html.Parse(fileOpen)

	ch := make(chan *html.Node)
	go findAnchors(doc, ch)
	var links []link
	for a := range ch {
		links = append(links, link{Text: findText(a), Href: findHref(a)})
	}
	fmt.Println(links)

}
func findText(n *html.Node) string {
	var text string
	//if n.Type == html.TextNode {
	//	text = n.Data
	//}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
			continue
		}
		text += findText(c)
	}
	return strings.TrimSpace(text)
}
func findHref(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, atr := range n.Attr {
			if atr.Key == "href" {
				return atr.Val
			}
		}
	}
	return ""
}
func findAnchors(n *html.Node, am chan *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		am <- n
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchors(c, am)
	}
	if n.Parent == nil {
		close(am)
	}
}
