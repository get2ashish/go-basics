package main

import "fmt"

func main() {
	//This initializes an empty map
	websites1 := map[string]string{}
	//Another way is
	websites2 := map[string]string{
		"google": "https://www.google.com",
		"amazon": "www.amazon.com",
	}

	//adding a new key value
	websites2["yahoo"] = "https://www.yahoo.com"

	//updating values
	websites2["amazon"] = "www.amazon.in"

	//deleting a key value pair

	delete(websites2, "yahoo")

	fmt.Println("Maps is ", websites1)
	fmt.Println("Map is ", websites2)

	fmt.Println("google website is ", websites2["google"])

	//Structs vs Maps
	//structs are used when the contract is written in stone cant be changes or deleted where as we can adda or modify new keys in maps

}
