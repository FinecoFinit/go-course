package utils

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strings"
	"unicode"
)

const (
	USD = "519.09"
	EUR = "608.44"
)

func Convert() {
	for {
		var currency decimal.Decimal

		fmt.Println("Enter value and currency(USD, EUR), example: 1 USD or 1 EUR: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(fmt.Errorf("error reading input: %w", err))
			continue
		}
		input = strings.TrimSpace(input)

		if len(strings.Split(input, " ")) != 2 {
			fmt.Println("invalid input:", input)
			continue
		}

		if func(str string) bool {
			for _, r := range str {
				if unicode.IsLetter(r) {
					return true
				}
			}
			return false
		}(strings.Split(input, " ")[0]) {
			fmt.Println("error converting input (value):", strings.Split(input, " ")[0])
			continue
		}

		value, _ := decimal.NewFromString(strings.Split(input, " ")[0])

		switch strings.Split(input, " ")[1] {
		case "USD":
			currency, _ = decimal.NewFromString(USD)
		case "EUR":
			currency, _ = decimal.NewFromString(EUR)
		default:
			fmt.Println("error converting input to decimal (wrong currency):", strings.Split(input, " ")[1])
			continue
		}
		fmt.Println(value.Mul(currency), "KZT")
	}
}
