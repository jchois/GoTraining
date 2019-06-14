package main

import (
	"time"
	"fmt"
	
	// age "github.com/bearbin/go-age"
)

func main()  {

	// var birth time.Time

	// fmt.Println("-------------")
	// fmt.Scanf("%v \n", birth)

	// layout := "2006-01-02" //T15:04:05.000Z
	// str := birth
	// t, err := time.Parse(layout, str)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%s\n",t)

	dt := time.Now()
    //Format MM-DD-YYYY
    fmt.Println(dt)
}