package main

import "fmt"

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	// get revenue
	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)
	revenue := getInputValue("Revenue: ")

	// get expenses
	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)
	expenses := getInputValue("Expenses: ")

	// get taxRate
	//fmt.Print("Tax Rate: ")
	//fmt.Scan(&taxRate)
	taxRate := getInputValue("Tax Rate: ")

	// earnings before tax
	//ebt := revenue - expenses

	// earnings after tax
	//profit := ebt * (1 - taxRate/100)

	// ratio EBT / profit
	//ratio := ebt / profit

	ebt, profit, ratio := calculateEbtProfitAndRatio(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

func getInputValue(text string) float64 {
	fmt.Print(text)
	var value float64
	fmt.Scan(&value)
	return value
}

func calculateEbtProfitAndRatio(revenue, expanses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expanses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
