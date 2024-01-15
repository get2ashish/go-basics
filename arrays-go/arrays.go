package main

import "fmt"

func main() {
	var productNames [4]string
	fmt.Println("Product Names ", productNames)
	var storeNames = [4]string{"Store 1"}
	fmt.Println("store Names ", storeNames)
	var prices = [5]float64{10.00, 12.50, 14.99, 16.88, 18.25}
	prices[0] = 99.99
	fmt.Println("Price array is ", prices)
	fmt.Println("3rd element in array is ", prices[2])

	//Slices can contain a part of the array --> it a window into an array
	// starting from element 1 upto element 3 [not including element 3]
	var featuredPrice = prices[1:3]
	fmt.Println("Slice is ", featuredPrice)

	//Variations
	fmt.Println(featuredPrice[:3]) //starting upto 3rd index
	fmt.Println(prices[1:])        //1st element till end

	//IMP on modyfying a slice the main array also gets changed
	//so we can saw slices are like pointer reference to an array as it can mutate the original array
	featuredPrice[0] = 77.77
	fmt.Println("Modified array ", prices)

	//Length of array or slice
	fmt.Println("Length is", len(featuredPrice))
	fmt.Println("Length is", len(prices))

}
