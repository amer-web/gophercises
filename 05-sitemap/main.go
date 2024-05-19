package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sitemap/Link"
	"strings"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URL     []URl    `xml:"url"`
}
type URl struct {
	Loc string `xml:"loc"`
}

func main() {
	url := flag.String("url", "", "url u need tp parse")
	flag.Parse()
	if *url == "" {
		log.Fatal("u should pass url")
	}
	links, _ := buildSitemap(*url)

	generateXmlSiteMap(links)
}

func generateXmlSiteMap(links []string) {
	urls := []URl{}
	urlXml := Urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, link := range links {
		urls = append(urls, URl{Loc: link})
	}
	urlXml.URL = urls
	xmlData, err := xml.MarshalIndent(urlXml, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data := append([]byte(xml.Header), xmlData...)

	if os.WriteFile("sitemap.xml", data, os.ModePerm) != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("XML file created successfully.")
}
func buildSitemap(baseDomain string) ([]string, error) {
	urls := []string{baseDomain}
	urlMap := make(map[string]bool)
	for i := 0; i < 2; i++ {
		var newUrls []string
		for _, url := range urls {
			url = strings.TrimSuffix(url, "/")
			subUrl, _ := sitemap(url)
			for _, sUrl := range subUrl {
				if urlMap[sUrl] {
					continue
				}
				urlMap[sUrl] = true
			}
			newUrls = append(newUrls, subUrl...)
		}
		urls = newUrls
	}
	return urls, nil
}
func sitemap(url string) ([]string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("can't response")
	}
	links, er := Link.Parse(res.Body)
	if er != nil {
		return nil, er
	}
	urls := getUrls(links, url)
	return urls, nil
}

func getUrls(links []Link.Link, baseDomain string) []string {
	var urls []string
	for _, link := range links {
		urls = append(urls, link.Href)
	}
	var filterUrls []string
	for _, href := range urls {
		if strings.HasPrefix(href, "http") &&
			!strings.HasPrefix(href, baseDomain) {
			continue
		}
		if strings.HasPrefix(href, baseDomain) {
			filterUrls = append(filterUrls, href)
			continue
		}
		if strings.Contains(href, "@") {
			continue
		}
		if href == "" || href[0] != '/' {
			href = "/" + href
		}

		filterUrls = append(filterUrls, baseDomain+href)
	}
	return filterUrls
}
