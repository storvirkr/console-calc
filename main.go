package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c Calculator) calculate(input string) interface{} {
	elements := strings.Fields(input)
	if len(elements) != 3 {
		panic("Invalid input format")
	}

	a := parseNumber(elements[0])
	op := elements[1]
	b := parseNumber(elements[2])

	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("Division by zero")
		}
		result = a / b
	default:
		panic("Invalid operation")
	}

	if isRoman(elements[0]) && isRoman(elements[2]) {
		return arabicToRoman(result)
	}
	return result
}

func parseNumber(str string) int {
	romanToArabic := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if num, ok := romanToArabic[str]; ok {
		return num
	}

	num, err := strconv.Atoi(str)
	if err != nil || num < 1 || num > 10 {
		panic("Invalid number")
	}

	return num
}

func isRoman(str string) bool {
	romanChars := "IVXLCDM"
	for _, char := range str {
		if !strings.ContainsRune(romanChars, char) {
			return false
		}
	}
	return true
}

func arabicToRoman(n int) string {
	if n <= 0 {
		panic("Result cannot be represented in Roman numerals")
	}

	romanSymbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	arabicValues := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}

	result := ""
	for i := len(romanSymbols) - 1; i >= 0; i-- {
		for n >= arabicValues[i] {
			n -= arabicValues[i]
			result += romanSymbols[i]
		}
	}

	return result
}

func main() {
	calculator := Calculator{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter expression (e.g., 3 + 5):")
	for scanner.Scan() {
		input := scanner.Text()
		result := calculator.calculate(input)
		fmt.Println("Result:", result)
		fmt.Println("Enter expression (e.g., 3 + 5):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
