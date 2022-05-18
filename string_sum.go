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
	result := 0
	digitSign := true
	ints := []int{}
	sBuffer := []string{}
	nums := strings.Split(input, "")[:]
	nums = append(nums, " ")
	for _, v := range nums {
		_, err := strconv.Atoi(v)
		if err != nil {
			if len(sBuffer) != 0 {
				var sb strings.Builder
				for _, val := range sBuffer {
					sb.WriteString(val)
				}
				i, _ := strconv.Atoi(sb.String())
				if !digitSign {
					i = i * -1
				}
				ints = append(ints, i)
				result += i
				sBuffer = sBuffer[:0]
			}
			if v == " " {
				continue
			} else if v == "-" {
				digitSign = false
				continue
			} else if v == "+" {
				digitSign = true
				continue
			} else {
				return "", fmt.Errorf("Input error: %w", err)
			}
		}
		sBuffer = append(sBuffer, v)
	}
	if len(ints) == 0 {
		return "", fmt.Errorf("Empty string error: %w", errorEmptyInput)
	} else if len(ints) != 2 {
		return "", fmt.Errorf("More or less then two operands: %w", errorNotTwoOperands)
	} else {
		return strconv.Itoa(result), nil
	}
}
