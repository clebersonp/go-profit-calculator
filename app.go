package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// get revenue
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	// get expenses
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	// get taxRate
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// earnings before tax
	ebt := revenue - expenses

	// earnings after tax
	profit := ebt * (1 - taxRate/100)

	// ratio EBT / profit
	ratio := ebt / profit

	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}
