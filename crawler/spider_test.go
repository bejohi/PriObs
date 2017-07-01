package crawler

import (
	"golang.org/x/net/html"
	"io"
	"strings"
	"testing"
)

func getMockWebsite() io.Reader {
	var exampleWebsite string = "<html>" +
		"<head>" +
		"<title class='a-class'>Example Title</title>" +
		"</head>" +
		"<body id='my-body'>" +
		"<p>Para1</p>" +
		"<a class='a-class b-class'></a>" +
		"<p>Para1</p>" +
		"<p>Para1</p>" +
		"</body>" +
		"</html>"

	reader := strings.NewReader(exampleWebsite)
	return reader
}

func TestGetNodesBySelector(t *testing.T) {

	// ARRANGE

	root, _ := html.Parse(getMockWebsite())

	// ACT
	resultClasses := GetNodesBySelector(".a-class", root)
	resultIds := GetNodesBySelector("#my-body", root)
	resultTagNames := GetNodesBySelector("p", root)

	// ASSERT
	if len(resultClasses) != 2 {
		t.Errorf("The wrong number of classes was found: %d!", len(resultClasses))
	}
	if len(resultIds) != 1 {
		t.Errorf("The wrong number of ids was found: %d!", len(resultIds))
	}
	if len(resultTagNames) != 3 {
		t.Errorf("The wrong number of tags was found: %d!", len(resultTagNames))
	}

}
