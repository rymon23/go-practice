package main

import (	
	"fmt"
	//"go-practice/keygen"
	//"go-practice/gowiki"
	// "go-practice/webserver"
	"go-practice/webserverfinal"	
	// "path/filepath"
)

func main() {
	fmt.Println("\n\n___________________________________")
	fmt.Println("___Password BCrypt Key Generator___")
	fmt.Print("___________________________________\n\n")

	// var srcPath, _ = filepath.Abs(filepath.Dir("go-practice/src/"))
	// fmt.Println(srcPath)

	//keygen.Execute()
	//gowiki.Execute()
	// webserver.Execute(":8080")
	// webserver.ExecuteWiki(":8080")
	//webserver.RunTest()
	webserverfinal.Run()
}