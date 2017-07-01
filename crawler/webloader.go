package crawler

import (
	"golang.org/x/net/html"
	"net/http"
)

// Loads a website from the network and returns the root of the html tree.
// Returns an error if the connection failed,
// or if the received html document was not valid formed.
func LoadWebRoot(url string) (*html.Node,error){
	website, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return html.Parse(website.Body)
}