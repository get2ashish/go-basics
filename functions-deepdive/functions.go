package main

import (
	"fmt"
)

func main() {
	//we need to double every number in the slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = doubleNumbers(&numbers)
	fmt.Println("updated slice ", numbers)
	doubleNumbersTest := transformNumbers(&numbers, double) //doot write double() as that will execute just give the functiom name
	fmt.Println("Functions as parameters ", doubleNumbersTest)

	//Anonymous functions
	// Since transformNumbers takes another function we create an Anonymous on the fly and pass it like this
	//Remember that you donot need to specify the function name
	tripleNumbersTest := transformNumbers(&numbers, func(number int) int {
		return number * 3
	})
	fmt.Println("Anonymous triple ", tripleNumbersTest)

	//Recusion
	fmt.Println("Factorial of 5 is ", factorialWithoutRecusion(5))
	fmt.Println("Factorial of 5 using recursion ", factorial(5))

	//Variadic functions : these accept any number of argument of a particular type
	fmt.Println("Sum of 1,2,3,4,5,6,7 is", sumUp(1, 2, 3, 4, 5, 6, 7))
	fmt.Println("Sum of 199,1 is", sumUp(199, 1))
}

func doubleNumbers(slice *[]int) []int {
	dNumbers := []int{}
	for _, value := range *slice {
		dNumbers = append(dNumbers, value*2)
	}
	return dNumbers
}

//tomorrow if we neeed a function that triples the value we need to write another function
//It would be good if we can write a transform function that takes slice and a function as argument which describes what it needs to do
//Functions are just like variables in golang

func transformNumbers(slice *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *slice {
		fmt.Println("Tranfoming value", value)
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func factorialWithoutRecusion(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func sumUp(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}
