package main

import (
	"time"
	"fmt"
	
	age "github.com/bearbin/go-age"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate  time.Time
	ExpectedAge int
}

var firtName, lastName, birthdate string

func info()  {
	fmt.Println("\n")
	fmt.Println("Create a new user: \n")

	fmt.Println("Firt Name:")
	fmt.Scanf("%v \n", &firtName)
	fmt.Println("Last Name:")
	fmt.Scanf("%v \n", &lastName)
	fmt.Println("Birthdate: (dd-mm-yyyy)")
	fmt.Scanf("%v \n", &birthdate)
}

func process() *User {
	currentDate := time.Now()

	layout := "02-01-2006"
	str := birthdate
	t, err := time.Parse(layout, str)
	
	if err != nil {
		fmt.Println(err)
	}

	age := age.AgeAt(t, currentDate)

	u := &User{
		FirstName: firtName,
		LastName: lastName,
		Birthdate: t,
		ExpectedAge: age,
	}

	return u
}

func print(u *User)  {

	fmt.Println("\nName: ", u.FirstName, u.LastName)
	fmt.Println("Birthdate: ", u.Birthdate.Format("Jan 02"))
	fmt.Println("Age: ", u.ExpectedAge)
	fmt.Println("\n")
}

func main() {
	
	info()
	u := process()
	print(u)
	
}