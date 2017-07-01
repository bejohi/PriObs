package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
)

// GetElementsByTagName returns an array of html.Node's where all nodes
// have the given selector, e.g. <body>, #my_ud or .class-a
func GetNodesBySelector(selector string, root *html.Node) []*html.Node {
	doc := goquery.NewDocumentFromNode(root)
	return doc.Find(selector).Nodes
}

func GetAllChildes(root *html.Node) []*html.Node {
	return GetNodesBySelector("*", root)
}

func GetNodesWithClasses(classes []string, root *html.Node) []*html.Node {

	resultArray := make([]*html.Node,0)

	if len(classes) < 1{
		return resultArray
	}

	allNodes := GetAllChildes(root)

	for _, node := range allNodes {
		for _, attr := range node.Attr{
			if attr.Key == "class"{
				allInside := true
				for _, htmlClass := range classes {
					if !strings.Contains(attr.Val,htmlClass){
						allInside = false
						break
					}
				}
				if allInside{
					allNodes = append(allNodes,node)
				}
				break
			}
		}
	}
	return allNodes
}