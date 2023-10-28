package main

import (
	"fmt"
	"log"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("input.png")
	text, _ := client.Text()
	fmt.Println(text)
	f, err := os.Create("lastread.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(text)
	if err2 != nil {
		log.Fatal(err2)
	}
	HOCRText, _ := client.HOCRText()
	f2, err := os.Create("lastread.hocr")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	write, err := f2.WriteString(HOCRText)
	if err != nil {
		log.Fatal(write)
	}
	// Hello, World!
}
