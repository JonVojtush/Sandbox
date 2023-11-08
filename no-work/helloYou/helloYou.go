package main

import (
	cm "customModules"
	_ "database/sql"
	"fmt"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/example/hello/reverse"
)

func fetchDB() string {
	name := "Jon"
	return name
}

func helloYou(fullName string) {
	phrase := "Hello " + fullName + "!"
	fmt.Println("Phrase: " + phrase)
	fmt.Println("Upper Case (local pkg): " + cm.HelloYouToUpper(phrase))
	fmt.Println("Reverse (pkg from web): " + reverse.String(phrase))
}

func main() {
	name := fetchDB()
	helloYou(name)
}
