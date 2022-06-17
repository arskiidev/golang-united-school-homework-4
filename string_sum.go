package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	digits, err := validateInput(input)
	if err != nil {
		return "", err
	}
	values := getInts(digits)
	if len(values) == 0 {
		return "", fmt.Errorf("Empty string error: %w", errorEmptyInput)
	}
	if len(values) != 2 {
		return "", fmt.Errorf("More or less then two operands: %w", errorNotTwoOperands)
	}
	return strconv.Itoa(values[0] + values[1]), nil
}

func validateInput(input string) (digits []string, err error) {
	for _, digit := range strings.Split(input, "") {
		switch digit {
		case "+", "-":
			digits = append(digits, digit)
		case " ":
			continue
		default:
			if _, err := strconv.Atoi(digit); err != nil {
				return nil, fmt.Errorf("Input error: %w", err)
			}
			digits = append(digits, digit)
		}
	}
	digits = append(digits, "+")
	return digits, err
}

func getInts(digits []string) (values []int) {
	integer := []string{digits[0]}
	for i := 1; i < len(digits); i++ {
		if digits[i] == "+" || digits[i] == "-" {
			value, _ := makeInt(integer)
			values = append(values, value)
			integer = append(integer[0:0], digits[i])
			continue
		}
		integer = append(integer, digits[i])
	}
	return values
}

func makeInt(integer []string) (int, error) {
	var sb strings.Builder
	for _, i := range integer {
		sb.WriteString(i)
	}
	return strconv.Atoi(sb.String())
}
