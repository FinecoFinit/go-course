package main

import (
	"bufio"
	"fmt"
	//"github.com/shopspring/decimal"
	"os"
	//"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Enter value and currency(USD, EUR), example: 1 USD or 1 EUR: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		print(fmt.Errorf("error reading input: %w", err))
	}
	sliceOfInput := strings.Split(strings.TrimSpace(input), " ")
	print(sliceOfInput[0])
}
