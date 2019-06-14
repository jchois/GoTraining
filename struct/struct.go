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

func main() {
	var firtName, lastName, birthdate string

	currentDate := time.Now()

	fmt.Println("\n")
	fmt.Println("Create a new user: \n")

	fmt.Println("Firt Name:")
	fmt.Scanf("%v \n", &firtName)
	fmt.Println("Last Name:")
	fmt.Scanf("%v \n", &lastName)
	fmt.Println("Birthdate: (dd/mm/yyyy)")
	fmt.Scanf("%v \n", &birthdate)

	layout := "02-01-2006"
	str := birthdate
	t, err := time.Parse(layout, str)
	
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(t)

	age := age.AgeAt(t, currentDate)

	u := &User{
		FirstName: firtName,
		LastName: lastName,
		Birthdate: t,
		ExpectedAge: age,
	}

	date := u.Birthdate
	
	fmt.Println("\nName: ", u.FirstName, u.LastName)
	fmt.Println("Birthdate: ", date.Format("Jan 02"))
	fmt.Println("Age: ", u.ExpectedAge)
	fmt.Println("\n")

}