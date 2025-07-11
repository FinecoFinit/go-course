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
mainLoop:
	for {
		var currency decimal.Decimal

		println("Enter value and currency(USD, EUR), example: 1 USD or 1 EUR: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			print(fmt.Errorf("error reading input: %w", err))
			break
		}
		input = strings.TrimSpace(input)

		if func(str string) bool {
			for _, r := range str {
				if unicode.IsLetter(r) {
					return true
				}
			}
			return false
		}(strings.Split(input, " ")[0]) {
			println("error converting input (value):", strings.Split(input, " ")[0])
			break
		}

		value, _ := decimal.NewFromString(strings.Split(input, " ")[0])

		switch strings.Split(strings.TrimSpace(input), " ")[1] {
		case "USD":
			currency, _ = decimal.NewFromString(USD)
		case "EUR":
			currency, _ = decimal.NewFromString(EUR)
		default:
			println("error converting input to decimal (wrong currency):", strings.Split(input, " ")[1])
			continue mainLoop
		}

		output := value.Mul(currency).IntPart()
		println(output, "KZT")
	}
}
