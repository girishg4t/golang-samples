package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

func main() {
	l := []link{}
	wg := &sync.WaitGroup{}
	dat, err := ioutil.ReadFile("./ex1.html")

	doc, err := html.Parse(bytes.NewBuffer(dat))
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			href := n.Attr[0].Val
			wg.Add(1)
			c := make(chan string)
			crawl(n.FirstChild, wg, c)
			text := <-c
			l = append(l, link{Href: href, Text: text})
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println(l)
	wg.Wait()
}

func crawl(node *html.Node, wg *sync.WaitGroup, c chan string) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode || n.Type == html.TextNode {
			getData(n, c)
		}
		if n.NextSibling == nil {
			defer wg.Done()
			return
		}
		f(n.NextSibling)
	}
	go f(node)
}

func getData(n *html.Node, c chan string) {
	if n.Type == html.TextNode {
		c <- remoUnwatedChar(n.Data)
	} else if n.FirstChild != nil {
		c <- remoUnwatedChar(n.FirstChild.Data)
	}
}

func remoUnwatedChar(s string) string {
	chars := []string{"\n"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	s = re.ReplaceAllString(s, "")
	return strings.TrimLeft(s, " ")
}
