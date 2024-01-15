package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName   string    `json:"first_name"` //these struct tags that hold metadata and can be used for providing additional details, when marshalling the object to json it will be used by json package to assign the keys
	LastName    string    `json:"last_name"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

type Admin struct {
	Username string
	Password string
	//you can use User anoynomously or add a variable name
	User User
}

// Attaching a function to a struct which becomes a method
// (appUser User) after function is called a receiver function
// receiver function only receives a copy of the value hence if we update the value it wont get changed
func (appUser User) StructMethodToPrintAllValues() {
	fmt.Println("First Name is ", appUser.FirstName)
	fmt.Println("Last Name is ", appUser.LastName)
	fmt.Println("Date of birth is ", appUser.DateOfBirth)
	fmt.Println("Created at is ", appUser.CreatedAt)
}

// If we need to use a method that mutates the object then we should use pointer as shown below
func (appUser *User) MutateStructValues() {
	appUser.FirstName = ""
	appUser.LastName = ""
	appUser.DateOfBirth = ""
}

// There is a convention (not a rule for Go) to create a new Struct object we can use a constructor function like below
// To make this utility more better we can return a pointer
// This is also known as Contructor pattern
func NewUser(firstName, lastName, birthdate string) User {
	return User{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: birthdate,
		CreatedAt:   time.Now(),
	}
}

func NewAdmin(username, password string) Admin {
	return Admin{
		Username: username,
		Password: password,
		User: User{
			FirstName:   "ADMIN",
			LastName:    "ADMIN-LASTNAME",
			DateOfBirth: "MM-DD-YYYY",
			CreatedAt:   time.Now(),
		},
	}
}

// This is also known as Contructor pattern that returns a pointer this is more optimized as copy is not passed
func NewUserPointer(firstName, lastName, birthdate string) *User {
	return &User{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: birthdate,
		CreatedAt:   time.Now(),
	}
}

// Adding validation logic as well
func NewUserPointerWithValidations(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birth date are required and cannot be empty")
	}

	return &User{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: birthdate,
		CreatedAt:   time.Now(),
	}, nil
}
