package html2markdown

import (
	"strings"

	"golang.org/x/net/html"
)

func Convert(htmlString string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return "", err
	}
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		s := ""
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s += f(c)
		}

		switch n.Type {
		case html.TextNode:
			return n.Data
		case html.ElementNode:
			switch n.Data {
			case "b":
				s = "**" + s + "**"
			case "em":
				s = "*" + s + "*"
			case "h1":
				s = "# " + s
			case "h2":
				s = "## " + s
			case "h3":
				s = "### " + s
			case "h4":
				s = "#### " + s
			case "h5":
				s = "##### " + s
			case "h6":
				s = "###### " + s
			case "hr":
				s = "***"
			}
		}
		return s
	}

	return f(doc), nil
}
