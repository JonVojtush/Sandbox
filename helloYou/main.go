package main

import (
	"fmt"

	cm "github.com/JonVojtush/GoLang-Sandbox/CustomModules"

	"golang.org/x/example/hello/reverse"
)

func fetchData() string {
	name := "Jon"
	return name
}

func helloYou(fullName string) {
	phrase := "Hello " + fullName + "!"
	fmt.Println("Phrase: " + phrase)
	fmt.Println("Upper Case (local pkg): " + cm.UpperCaseString(phrase))
	fmt.Println("Reverse (pkg from web): " + reverse.String(phrase))
	// Import XML & TXT file for data display
}

func main() {
	name := fetchData()
	helloYou(name)
}
