package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("input.png")
	text, _ := client.HOCRText()
	fmt.Println(text)
	// Hello, World!
}
