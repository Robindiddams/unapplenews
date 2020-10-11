package unapplenews

import (
	"io"

	"golang.org/x/net/html"
)

func findArticleURL(r io.Reader) (string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}
	var url string
	// TODO(robin): use the tokenizer and remove the unnecessary recursion
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			child := n.FirstChild
			getHref := func() string {
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						return attr.Val
					}
				}
				return ""
			}
			if child != nil && child.Data == "span" && len(child.Attr) > 0 {
				var found bool
				for _, attr := range child.Attr {
					if attr.Key == "class" && attr.Val == "click-here" {
						found = true
						break
					}
				}
				if found {
					href := getHref()
					url = href
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return url, nil
}
