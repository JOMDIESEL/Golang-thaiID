package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validateThaiID("1959900459873"))
	fmt.Println(validateThaiID("1419385944635"))
}

func validateThaiID(id string) error {
	if len(id) != 13 {
		return errors.New("ID length incorrect")
	}
	totalValue := calculateSum(id)
	result := calculateResult(totalValue)

	lastID := conLastDigit(id)

	if result != lastID  {
		return errors.New("ID incorrect")
	}else{
		return nil
	}
}
func conLastDigit(id string)int { 
	splited := strings.Split(id, "")
	lastID, _ := strconv.Atoi(splited[len(splited)-1])
	return lastID
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
