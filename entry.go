package main

import (	
	"fmt"
	"go-practice/keygen"
)

func main() {
	fmt.Println("___________________________________")
	fmt.Println("___Password BCrypt Key Generator___")
	fmt.Println("___________________________________")

	keygen.Execute()
}