package main

import (
	"fmt"
)

func main() {
	a := [3][3]int{
			{4, 3, 9},
			{1, 6, 7},
			{8, 5, 2},
	}
	fmt.Println(a)

	for  i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i+j) == 2{
				fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
			}
		}
	 }
}
