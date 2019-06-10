package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	word := bufio.NewScanner(os.Stdin)

	var s, aux int

	fmt.Print("Write anything: ")
	word.Scan()
	value := word.Text()

	for char := len(value) - 1; char >= 0; char-- {
		if value[char] == value[aux] {
			s++
		}
		aux++
	}
	if len(value) == s {
		fmt.Println("it is a palindromo!")
	} else {
		fmt.Println("it isn't a palindromo")
	}
}
