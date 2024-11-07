package main

import (
	"fmt"
)

func main() {

	var text string = "Hello from God!"
	var num int = 42
	var text2 = "Another string"
	num2 := 56

	fmt.Println(text)
	fmt.Printf("The type is %T\n", text)

	fmt.Println(text2)
	fmt.Printf("The type is %T\n", text2)

	fmt.Println(num)
	fmt.Printf("The type is %T\n", num)

	fmt.Println(num2)
	fmt.Printf("The type is %T\n", num2)
}
