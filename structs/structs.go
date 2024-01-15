package main

import (
	"fmt"

	"github.com/get2ashish/golang-structs/user"
)

/*
	type User struct {
		firstName   string
		lastName    string
		dateOfBirth string
		createdAt   time.Time
	}

// Attaching a function to a struct which becomes a method
// (appUser User) after function is called a receiver function
// receiver function only receives a copy of the value hence if we update the value it wont get changed

	func (appUser User) structMethodToPrintAllValues() {
		fmt.Println("First Name is ", appUser.firstName)
		fmt.Println("Last Name is ", appUser.lastName)
		fmt.Println("Date of birth is ", appUser.dateOfBirth)
		fmt.Println("Created at is ", appUser.createdAt)
	}

// If we need to use a method that mutates the object then we should use pointer as shown below

	func (appUser *User) mutateStructValues() {
		appUser.firstName = ""
		appUser.lastName = ""
		appUser.dateOfBirth = ""
	}
*/

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	/*var appUser User = User{
		firstName:   userFirstName,
		lastName:    userLastName,
		dateOfBirth: userBirthdate,
		createdAt:   time.Now(),
	}*/

	appUser := user.NewUser(userFirstName, userLastName, userBirthdate)

	outputUserDetails(appUser)

	var admin = user.NewAdmin("username", "password")
	fmt.Println("Admin Struct is ", admin)

	//If order is same as struct definition then we can omit the key name
	/*appUser = User{
		"Ashish-new",
		"Shuklaji",
		"01/20/1988",
		time.Now(),
	}*/

	//outputUserDetailsViaPointer(&appUser)
	appUser.StructMethodToPrintAllValues()
	appUser.MutateStructValues()
	fmt.Println("Printing all values after mutate method")
	//appUser.structMethodToPrintAllValues()

	// makes life easy not necessary
	var newUserObject = user.NewUser("ashish", "shukla", "01/01/2020")
	var newUserObjectPtr *user.User = user.NewUserPointer("ashish", "shukla", "01/01/2020")
	newUserObjectPtr.StructMethodToPrintAllValues()
	newUserObject.StructMethodToPrintAllValues()
	newUserObjectWithValidation, error := user.NewUserPointerWithValidations("ashish", "shukla", "")
	if error != nil {
		fmt.Println("Invalid argument passed")
		panic(error)
	} else {
		newUserObjectWithValidation.StructMethodToPrintAllValues()
	}

}

// If we use User type as type then a copy of app user is passed we can use pointer to avoid such things
func outputUserDetails(appUser user.User) {
	fmt.Println("First Name is ", appUser.FirstName)
	fmt.Println("Last Name is ", appUser.LastName)
	fmt.Println("Date of birth is ", appUser.DateOfBirth)
	fmt.Println("Created at is ", appUser.CreatedAt)
}

/*func outputUserDetailsViaPointer(appUser *user.User) {
	//This works without dereferencing because go allows that
	//Technically correct way is
	fmt.Println("First Name is ", (*appUser).FirstName)
	//This format is also allowed by go for ease of use
	fmt.Println("Last Name is ", appUser.LastName)
	fmt.Println("Date of birth is ", appUser.DateOfBirth)
	fmt.Println("Created at is ", appUser.CreatedAt)
}*/

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

/*
// There is a convention (not a rule for Go) to create a new Struct object we can use a constructor function like below
// To make this utility more better we can return a pointer
// This is also known as Contructor pattern
func newUser(firstName, lastName, birthdate string) User {
	return User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: birthdate,
		createdAt:   time.Now(),
	}
}

// This is also known as Contructor pattern that returns a pointer this is more optimized as copy is not passed
func newUserPointer(firstName, lastName, birthdate string) *User {
	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: birthdate,
		createdAt:   time.Now(),
	}
}

// Adding validation logic as well
func newUserPointerWithValidations(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birth date are required and cannot be empty")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: birthdate,
		createdAt:   time.Now(),
	}, nil
}
*/
