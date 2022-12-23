package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(GoQuote())
}

func GoQuote() string {
	return quote.Go()
}
