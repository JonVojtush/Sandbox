package main

import (
    "flag"
    "fmt"
)

func greet(name string) string {
    return fmt.Sprintf("Hello %s", name)
}

func main() {
    namePtr := flag.String("name", "World", "Name of the client to greet")

    flag.Parse()

    greeting := greet(*namePtr)
    fmt.Println(greeting)
}
