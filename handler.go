package main

import (
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (c *ComputeHandler) Compute() error {
	content, err := io.ReadAll(c.Input)
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}
	inputExpression := strings.TrimSpace(string(content))

	result, err := EvaluatePostfix(inputExpression)
	if err != nil {
		return fmt.Errorf("error evaluating expression: %v", err)
	}

	_, err = fmt.Fprintf(c.Output, "The result of the postfix expression '%s' is: %.2f\n", inputExpression, result)
	if err != nil {
		return fmt.Errorf("error writing output: %v", err)
	}

	return nil
}
