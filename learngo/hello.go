package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	PrintHello()
	PrintQuote()
}

func PrintHello() {
	fmt.Println("Hello, world!")
}

func PrintQuote() {
	fmt.Println(quote.Go())
}
