package main

import (
	"fmt"
	"os"

	"niklaskorz.de/nklang/lexer"
	"niklaskorz.de/nklang/parser"
	"niklaskorz.de/nklang/semantics"
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
	if err := semantics.AnalyzeLookups(p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p)
	if err := p.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
