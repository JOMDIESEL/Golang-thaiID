package main

import (
	"errors"
	"fmt"
	"strconv"

)

func main() {
	fmt.Println(validateThaiID("1419385944635"))
}

func validateThaiID(id string) error {
	if len(id) != 13 {
		return errors.New("ID length incorrect")
	}
	totalValue := calculateSum(id)
	result := calculateResult(totalValue)
	if result == 0  {
		return errors.New("ID incorrect")
	}
	return nil
}

func calculateSum(id string) int {
	sum := 0
	for i, j := 0, 13; j > 1; i, j = i+1, j-1 {
		val, _ := strconv.Atoi(string(id[i]))
		sum += val * j
	}
	return sum
}

func calculateResult(sum int) int {
	mod := sum % 11
	result := 11 - mod
	return result
}

func isValidLastDigit(result int) int {
	value  := result%10 
	return value
}
