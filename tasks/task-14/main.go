package main

import (
	"fmt"
	"reflect"
)

func identifyType(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case bool:
		fmt.Println("bool:", val)
	default:
		// check if a type of i is chan
		if reflect.TypeOf(i).Kind() == reflect.Chan {
			fmt.Println("chan:", reflect.TypeOf(i))
		} else {
			fmt.Println("unknown type:", val)
		}
	}
}

func main() {
	cInt := make(chan int)
	cBool := make(chan bool)

	identifyType(1)       // int
	identifyType("hello") // string
	identifyType(true)    // bool
	identifyType(cInt)    // chan
	identifyType(cBool)   // chan
	identifyType(12.3)    // unknown type
}
