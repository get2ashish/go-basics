package main

import (
	"fmt"

	"github.com/get2ashish/bank-application/fileops"
)

const fileName = "accountbalance.txt"

func main() {
	var accountBalance, error = fileops.GetFloatValueFromFile(fileName)
	if error != nil {
		fmt.Println("Error occurred")
		fmt.Println(error)
		fmt.Println("------------------------------")
		//panic(error)
	}
	fmt.Println("Hello and welcome to Go-Bank Application")
	for {
		presentOptions()
		var choice int
		fmt.Scan(&choice)
		fmt.Printf("Your choice is %v \n", choice)

		/*
			if choice == 1 {
				fmt.Printf("Your balace is %v \n", accountBalance)
			} else if choice == 2 {
				fmt.Printf("Enter the deposit amount")
				var depositAmount float64
				fmt.Scan(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Deposit amount cannot be less than 0")
					continue
				}
				accountBalance = accountBalance + depositAmount
				fmt.Printf("Updated balance is %v \n", accountBalance)
			} else if choice == 3 {
				fmt.Printf("Enter the withdrawl amount \n")
				var withdrawlAmount float64
				fmt.Scan(&withdrawlAmount)

				if withdrawlAmount <= 0 {
					fmt.Println("Withdrawl amount cannot be less than 0")
					continue
				}

				if withdrawlAmount > accountBalance {
					fmt.Println("Withdrawl amount cannot be more than account balance")
					continue
				}
				accountBalance = accountBalance - withdrawlAmount
				fmt.Printf("Updated balance is %v", accountBalance)
			} else {
				fmt.Println("Ending..")
				break
			}
		*/

		switch choice {
		case 1:
			fmt.Printf("Your balace is %v \n", accountBalance)
		case 2:
			fmt.Printf("Enter the deposit amount")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount cannot be less than 0")
				continue
			}
			accountBalance = accountBalance + depositAmount
			fileops.WriteFloatValueToFile(accountBalance, fileName)
			fmt.Printf("Updated balance is %v \n", accountBalance)
		case 3:
			fmt.Printf("Enter the withdrawl amount \n")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Withdrawl amount cannot be less than 0")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Withdrawl amount cannot be more than account balance")
				continue
			}
			accountBalance = accountBalance - withdrawlAmount
			fileops.WriteFloatValueToFile(accountBalance, fileName)
			fmt.Printf("Updated balance is %v \n", accountBalance)
		case 4:
			fmt.Println("Ending the banking application")
			return
		default:
			fmt.Println("Invalid input must be betweeb 1 to 4")
		}

	}
}

/*
func writeFloatValueToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func getFloatValueFromFile(fileName string) (float64, error) {
	data, error := os.ReadFile(fileName)
	if error != nil {
		//panic(error)
		error := fmt.Sprintf("Error while reading the file %v", fileName)
		return 1000.0, errors.New(error)
	}
	floatText := string(data)
	floatValue, err := strconv.ParseFloat(floatText, 64)
	if err != nil {
		//panic(err)
		return 1000.0, errors.New("failed to parse stored balance value")
	}
	return floatValue, nil
}
*/
