package main

import (
	"fmt"
	"os"

	"niklaskorz.de/nklang/lexer"
	"niklaskorz.de/nklang/parser"
)

func main() {
	fmt.Println("nklang version 0.1")
	f, err := os.Open("example.nk")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := lexer.NewScanner(f)
	p, err := parser.Parse(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
	p.Execute()
}
