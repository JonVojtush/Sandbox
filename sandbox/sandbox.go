package main

import (
	"fmt"
	"reflect"
)

var (
	testInt     int8  = 42
	inputValue  *int8 = &testInt
	outputValue string
)

func EmptyInt(intPtr interface{}) string {
	if intPtr == nil {
		return ""
	} else {
		value := reflect.ValueOf(intPtr)
		retValue := fmt.Sprintf("%d", value)
		return retValue
	}
}

func main() {
	outputValue = EmptyInt(*inputValue)
}
