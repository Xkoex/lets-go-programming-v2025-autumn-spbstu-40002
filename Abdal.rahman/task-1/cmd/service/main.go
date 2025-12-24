package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first, _ := reader.ReadString('\n')
	first = strings.TrimSpace(first)
	a, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	second, _ := reader.ReadString('\n')
	second = strings.TrimSpace(second)
	b, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}
