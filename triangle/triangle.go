package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Println("\n")
	fmt.Println("Write the sides of a triangle: \n")

	fmt.Println("First side:")
	fmt.Scanf("%v \n", &a)
	fmt.Println("Second side:")
	fmt.Scanf("%v \n", &b)
	fmt.Println("Third side:")
	fmt.Scanf("%v \n", &c)

	if (a != 0 || b != 0 || c != 0) {

		if ((a+b) > c && (a+c) > b && (b+c) > a) {
			
			if (a == b && b == c ) {
				fmt.Println("\n")
				fmt.Println("It's an Equilateral Triangle")
			}
			if ((a == b && b != c && a != c ) || (b == c && b != a && c != a) || (a == c && c != b && a != b)) {
				fmt.Println("\n")
				fmt.Println("It's an Isosceles Triangle")
			}
			if (a != b && b != c && c != a) {
				fmt.Println("\n")
				fmt.Println("It's a Scalene Triangle")
			}
		}else{
			fmt.Println("With these values you can't form a triangle.")
		}
		
	}else{
		fmt.Println("With these values you can't form a triangle.")
	}
}