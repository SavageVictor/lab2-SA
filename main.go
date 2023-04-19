package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	expression := flag.String("e", "", "Postfix expression to evaluate")
	inputFile := flag.String("f", "", "Input file containing a postfix expression")
	outputFile := flag.String("o", "", "Output file to write the result")

	flag.Parse()

	if *expression == "" && *inputFile == "" {
		fmt.Println("Error: Please provide an expression using -e or an input file using -f")
		os.Exit(1)
	}

	if *expression != "" && *inputFile != "" {
		fmt.Println("Error: Please provide either an expression using -e or an input file using -f, not both")
		os.Exit(1)
	}

	var input io.Reader
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = strings.NewReader(*expression)
	}

	var output io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
