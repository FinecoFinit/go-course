package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func WeeklyCalc() {
	var sum decimal.Decimal
	transactions := []map[int]decimal.Decimal{
		{1: decimal.NewFromFloat(25000.00)},
		{1: decimal.NewFromFloat(20000.00)},
		{2: decimal.NewFromFloat(-9800.00)},
		{3: decimal.NewFromFloat(-1222.22)},
		{4: decimal.NewFromFloat(-1500.07)},
		{5: decimal.NewFromFloat(1201.37)},
		{6: decimal.NewFromFloat(-100.32)},
		{7: decimal.NewFromFloat(-523.33)},
	}
	for _, t := range transactions {
		for key, val := range t {
			fmt.Println(time.Weekday(func(val int) int {
				if val == 7 {
					return 0
				}
				return val
			}(key)).String())
			if val.IsNegative() {
				fmt.Println("Withdrawal:", strings.Trim(val.String(), "-"), "\n")
			} else {
				fmt.Println("Inflow:", val, "\n")
			}
			sum = decimal.Sum(sum, val)
		}
	}
	fmt.Println("Weekly summary:", sum)
}
