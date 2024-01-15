package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age is :", age)
	fmt.Println("Age pointer is ", agePointer)
	fmt.Println("Age pointer dereferenced ", *agePointer)
	adultYears := getAdultYearsUsingPointers(&age)
	fmt.Println("Adult years are : ", adultYears)
	fmt.Println("Age value after :", age)
}

/*
func getAdultYears(age int) int {
	return age - 18
}
*/

func getAdultYearsUsingPointers(age *int) int {
	//This will not change the age value.
	//If we want to directly update it then we can use something like
	*age = *age - 18
	//return *age - 18
	return 0
}

//Rememeber that fmt.Scan uses pointer we pass the variable name with &name and it internally dereferences it and mutates the value
