package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	fmt.Println("Hello GoLang, Lets start calculating")

	/*
		If you want to go automatically infer the type of the variable based on the value then you should use like
		investmentAmount := 1000.00
		In this case no need to use var keyword also no variable type is specified

		var investmentAmount float64 = 1000
		var yearsOfInvestment float64 = 10

		If you want to delcare multiple variable in same line thats also allowed like
		It all comes down to personal preference as this might become a little more prone to bugs if you miss align
	*/
	var investmentAmount, yearsOfInvestment, expectedRateOfInterest float64

	/*
		var dummyVar1, dummyVar2 = 10, "Ashish"
		fmt.Println(dummyVar1)
		fmt.Println(dummyVar2)
	*/

	fmt.Println("Enter the investment amount")
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter the expectedRateOfInterest value")
	fmt.Scan(&expectedRateOfInterest)

	fmt.Println("Enter the yearsOfInvestment value")
	fmt.Scan(&yearsOfInvestment)

	var futureValue, inflationValue = calculateFutureValues(investmentAmount, yearsOfInvestment, expectedRateOfInterest)
	fmt.Printf("futureValue is %v is if of Type %T and inflationValue is %v and is of Type %T \n", inflationValue, inflationValue, inflationValue, inflationValue)
	futureValue, inflationValue = calculateFutureValues2(investmentAmount, yearsOfInvestment, expectedRateOfInterest)
	/*
		fmt.Printf("Future value is %v \nInflation value is %v", futureValue, inflationValue)
		fmt.Printf("Future value is %.2f \nInflation value is %.2f", futureValue, inflationValue)
	*/
	printValue := fmt.Sprintf("Future value is %.2f \nInflation value is %.2f", futureValue, inflationValue)
	fmt.Println(printValue)
	//fmt.Println("inflationValue : ", inflationValue)
}

func calculateFutureValues(investmentAmount float64, yearsOfInvestment float64, expectedRateOfInterest float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedRateOfInterest/100), yearsOfInvestment)
	rfv := fv / math.Pow((1+inflationRate/100), yearsOfInvestment)
	return fv, rfv
}

// alternative syntax
func calculateFutureValues2(investmentAmount float64, yearsOfInvestment float64, expectedRateOfInterest float64) (fv float64, rfv float64) {
	//fv and rfc are now declared so we cannot use := syntax hence use = to get the values
	//just saying return will actaully return those values which are specified in the method signature
	fv = investmentAmount * math.Pow((1+expectedRateOfInterest/100), yearsOfInvestment)
	rfv = fv / math.Pow((1+inflationRate/100), yearsOfInvestment)
	//return fv, rfv
	return
}
