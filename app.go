package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// 		=> Show error message & exit if invalid input is provided
// 		- no negative numbers
// 		- not 0
// 2) Store calculated results into file

const financialsFile = "financials.txt"

func writeFinancialsToFile(ebt, profit, ratio float64) error {
	financials := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	return os.WriteFile(financialsFile, []byte(financials), 0644)
}

func main() {
	revenue, err := getInputValue("Revenue: ")
	printError(err)

	expenses, err := getInputValue("Expenses: ")
	printError(err)

	taxRate, err := getInputValue("Tax Rate: ")
	printError(err)

	ebt, profit, ratio := calculateEbtProfitAndRatio(revenue, expenses, taxRate)

	// store calculated results
	err = writeFinancialsToFile(ebt, profit, ratio)
	printError(err)

	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

func printError(err error) {
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		panic("You can't continue.")
	}
}

func getInputValue(text string) (float64, error) {
	fmt.Print(text)
	var value float64
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("invalid number. Must be greater than 0")
	}
	return value, nil
}

func calculateEbtProfitAndRatio(revenue, expanses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expanses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
