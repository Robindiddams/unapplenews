package main

import (
	"fmt"
	"os"

	"github.com/Robindiddams/unapplenews"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "please provide an apple news url")
		return
	}
	url := os.Args[1]
	if url == "" {
		fmt.Fprintln(os.Stderr, "please provide an apple news url")
		return
	}
	newURL, err := unapplenews.UnAppleNews(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error unwrapping url:", err)
		return
	}
	fmt.Println(newURL)
}
