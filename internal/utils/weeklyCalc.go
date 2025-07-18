package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func WeeklyCalc() {
	var sum decimal.Decimal
	transactions := [][]decimal.Decimal{
		1: {decimal.NewFromFloat(25000.00), decimal.NewFromFloat(20000.00)},
		2: {decimal.NewFromFloat(-9800.00)},
		3: {decimal.NewFromFloat(-1222.22)},
		4: {decimal.NewFromFloat(-1500.07)},
		5: {decimal.NewFromFloat(1201.37)},
		6: {decimal.NewFromFloat(-100.32)},
		7: {decimal.NewFromFloat(-523.33)},
	}
	for day, t := range transactions {
		for _, val := range t {
			fmt.Println(time.Weekday(func(day int) int {
				if day == 7 {
					return 0
				}
				return day
			}(day)).String())
			if val.IsNegative() {
				fmt.Println("Withdrawal:", strings.Trim(val.String(), "-"), "\n")
			} else {
				fmt.Println("Inflow:", val, "\n")
			}
		}
		sum = decimal.Sum(sum, transactions[day]...)
	}
	fmt.Println("Weekly summary:", sum)
}
