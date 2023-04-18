package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ComputeHandler struct {
	Input       string
	Output      string
	IsFileInput bool
}

func (c *ComputeHandler) Compute() error {
	var inputExpression string

	if c.IsFileInput {
		content, err := ioutil.ReadFile(c.Input)
		if err != nil {
			return fmt.Errorf("error reading input file: %v", err)
		}
		inputExpression = strings.TrimSpace(string(content))
	} else {
		inputExpression = c.Input
	}

	result, err := EvaluatePostfix(inputExpression)
	if err != nil {
		return fmt.Errorf("error evaluating expression: %v", err)
	}

	if c.Output != "" {
		err := ioutil.WriteFile(c.Output, []byte(fmt.Sprintf("%.2f", result)), 0644)
		if err != nil {
			return fmt.Errorf("error writing output file: %v", err)
		}
	} else {
		fmt.Printf("The result of the postfix expression '%s' is: %.2f\n", inputExpression, result)
	}

	return nil
}

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

	handler := &ComputeHandler{
		Output: *outputFile,
	}

	if *inputFile != "" {
		handler.Input = *inputFile
		handler.IsFileInput = true
	} else {
		handler.Input = *expression
		handler.IsFileInput = false
	}

	err := handler.Compute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if handler.Output != "" {
		fmt.Printf("The result has been written to %s\n", handler.Output)
	}
}
