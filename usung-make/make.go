package main

import "fmt"

type floatMap map[string]float64

func main() {
	//this is a slice
	userNames := []string{}
	userNames = append(userNames, "ashish")
	//Here intenally go uses an array and copies data to slice
	//Always remember slice is nothing but a window into an array
	fmt.Println("usernames list ", userNames)

	//Using Make we specify the type the inital length and capacity, so go will allocate memeory till 20 elements and anothing above that new memory will be allocated
	updatedUserName := make([]string, 5, 20)
	updatedUserName = append(updatedUserName, "shukla")
	//using append the value will be appended at the end and you will see the empty 2 spaces in front to use the empty front space use [0] [1]
	fmt.Println("updated names are ", updatedUserName)

	//Creating maps with make
	//50 is the intended length of map
	courses := make(map[string]string, 50)
	courses["go"] = "Go Lang"
	fmt.Println(courses)

	//Creating maps using alias this useful when using such long types
	myFloatmap := make(floatMap, 20)
	myFloatmap["price"] = 200.00
	fmt.Println(myFloatmap)

	//Looping over slices/arrays
	for index, value := range updatedUserName {
		fmt.Println("Index ", index)
		fmt.Println("Value ", value)
	}

	//Looping over maps
	for key, value := range courses {
		fmt.Println("map key :", key)
		fmt.Println("map value :", value)
	}
}
