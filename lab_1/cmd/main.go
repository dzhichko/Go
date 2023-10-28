package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	a := inputValue()

	sign := inputSign()

	b := inputValue()

	calculate(a, b, sign)
}

func inputValue() float32 {
	var errIn = errors.New("invalid input type, it must be float")
	var x float32
	fmt.Print("Enter value: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(errIn)
		os.Exit(1)
	}
	return x
}
func inputSign() string {
	var errIn = errors.New("invalid operation sign, use +, -, *, /")
	var x string
	fmt.Print("Enter operation: ")
	_, err := fmt.Scan(&x)
	if err != nil || len(x) > 1 {
		fmt.Println(errIn)
		os.Exit(1)
	}
	return x
}
func calculate(x float32, y float32, oper string) {
	var res float32
	switch oper {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		if y == 0 {
			fmt.Println("Cannot divide by zero")
			os.Exit(1)
		}
		res = x / y
	default:
		fmt.Println("Unknown operation")
		os.Exit(1)
	}

	fmt.Println("Result: ", res)
}
