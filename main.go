package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack []float64

func (s *Stack) Push(v float64) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (float64, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, true
}

func evaluatePostfix(expression string) (float64, error) {
	tokens := strings.Fields(expression)
	stack := Stack{}

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(num)
		} else {
			b, okB := stack.Pop()
			a, okA := stack.Pop()

			if !okA || !okB {
				return 0, fmt.Errorf("invalid postfix expression")
			}

			switch token {
			case "+":
				stack.Push(a + b)
			case "-":
				stack.Push(a - b)
			case "*":
				stack.Push(a * b)
			case "/":
				stack.Push(a / b)
			case "^":
				stack.Push(math.Pow(a, b))
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	res, _ := stack.Pop()
	return res, nil
}

func main() {
	expression := "3 8 + 2 * 2 + 3 ^"
	result, err := evaluatePostfix(expression)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The result of the postfix expression '%s' is: %.2f\n", expression, result)
	}
}
