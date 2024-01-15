package main

import "fmt"

func main() {
	prices := []float64{1.0, 2.0}
	//you can modify the elemets that were defined when creating the slice in this case 1.0 and 2.0
	//we cannot add thrird value by using [2], use append function to add more values
	prices = append(prices, 3.0)
	//Sinces slices work on arrays internally go creates a new array copies values and returns the new slice from it
	fmt.Println(prices)
	//please note there are no functions to remove elements from slice you need to use append along with [start:end-1] logic

	//adding a new slice to a new slice

	discountedPrices := []float64{3.0, 4.0}
	//you need to unpack the list/slic as append takes values
	prices = append(prices, discountedPrices...)
	fmt.Println("Added list ", prices)

}
