package main

import (	
	"fmt"
	// "go-practice/keygen"
	//"go-practice/gowiki"
	"go-practice/webserver"
	
)

func main() {
	fmt.Println("\n\n___________________________________")
	fmt.Println("___Password BCrypt Key Generator___")
	fmt.Print("___________________________________\n\n")

	//keygen.Execute()
	//gowiki.Execute()
	// webserver.Execute(":8080")
	webserver.ExecuteViewHandler()
}