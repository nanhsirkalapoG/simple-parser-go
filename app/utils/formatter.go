package utils

import (
	"strconv"
	"strings"
)

// Format the salary - Remove special characters(dollar($) and command(,) and format
func GetFormattedSalary(salary string) (float64, error) {
	wage := strings.TrimSpace(salary)
	wage = strings.Replace(wage, "$", "", 1)
	wage = strings.ReplaceAll(wage, ",", "")
	return strconv.ParseFloat(wage, 64)
}
