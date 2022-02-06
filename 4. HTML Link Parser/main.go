package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

/*
    A
  B   E
 C D   F


  div
		div
			a
        <span>....</span> + <div>...</div>
			a
        <span>....</span> + <div>...</div>
		a

*/

type Link struct {
	Href string
	Text string
}

func getHref(node *html.Node) string {
	for _, a := range node.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func getInnerText(node *html.Node, text *string) {
	if node == nil {
		return
	}

	if node.Type != html.ElementNode && node.Type != html.CommentNode {
		*text += strings.Trim(node.Data, "\n")
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		getInnerText(n, text)
	}
}

func dfs(node *html.Node, arr *[]Link) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		link := Link{Href: getHref(node)}
		getInnerText(node, &link.Text)
		link.Text = strings.Trim(link.Text, " ")
		*arr = append(*arr, link)
	}

	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		dfs(nextNode, arr)
	}
}

func printLinks(anchors []Link, idx int) {
	formatLine := "----------------------------------------"
	fmt.Println("\nTESTCASE: ", idx, " Links Parsed->", len(anchors), "\n")
	for i, anchor := range anchors {
		fmt.Printf("href: %s\ntext: %s\n", anchor.Href, anchor.Text)

		if i != len(anchors)-1 {
			fmt.Println()
		}
	}
	fmt.Println(formatLine)
}

func driver(idx int, path string) {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	node, _ := html.Parse(bytes.NewReader(buffer))
	var anchors []Link

	// for currNode := node; currNode != nil; currNode = currNode.NextSibling {
	dfs(node, &anchors)
	printLinks(anchors, idx)
}

func main() {
	testCases := []string{
		"./test/tc1.html",
		"./test/tc2.html",
		"./test/tc3.html",
		"./test/tc4.html",
		"./test/tc5.html",
	}

	for idx, test := range testCases {
		driver(idx+1, test)
	}
}
