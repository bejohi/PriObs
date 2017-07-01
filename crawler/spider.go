package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// GetElementsByTagName returns an array of html.Node's where all nodes
// have the given selector, e.g. <body>, #my_ud or .class-a
func GetNodesBySelector(selector string, root *html.Node) []*html.Node {
	doc := goquery.NewDocumentFromNode(root)
	return doc.Find(selector).Nodes
}
