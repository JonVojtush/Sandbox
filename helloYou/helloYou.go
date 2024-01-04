package main

import (
	_ "database/sql"
	"fmt"
	_ "time"

	// "bufio"
	// "bytes"
	// "crypto" // .../aes, .../tls, .../cipher
	cm "github.com/JonVojtush/GoLang-Sandbox/customModules"

	// "encoding" // .../xml, .../json
	// "errors"
	// "flag"

	// "io" // .../ioutil
	// "log"
	// "net"  // .../url, .../http
	// "os"   // .../exec
	// "path" // .../filepath
	// "reflect"
	// "strconv"
	// "strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/example/hello/reverse"
)

func fetchDB() string {
	/*db, err := sql.Open("mysql", "root:/hello")
	if err != nil {
		log.Printf("Error: %s\n", *err)
		panic(err)
	}
	connections := 3
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(connections)
	db.SetMaxIdleConns(connections)*/
	name := "Jon"
	return name
}

func helloYou(fullName string) {
	phrase := "Hello " + fullName + "!"
	fmt.Println("Phrase: " + phrase)
	fmt.Println("Upper Case (local pkg): " + cm.HelloYouToUpper(phrase))
	fmt.Println("Reverse (pkg from web): " + reverse.String(phrase))
	// Import XML & TXT file for data display
}

func main() {
	name := fetchDB()
	helloYou(name)
	// fmt.Scanf("h") // keep terminal open
}
