package Link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	ch := make(chan *html.Node)
	go findAnchors(doc, ch)
	var links []Link
	for a := range ch {
		links = append(links, Link{Text: findText(a), Href: findHref(a)})
	}
	return links, nil

}
func findText(n *html.Node) string {
	var text string
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
