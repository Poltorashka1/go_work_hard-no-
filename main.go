package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "12"
	strConv, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Before conv %v - %T\n", x, x)
	fmt.Printf("After conv %v - %T\n", strConv, strConv)
	fmt.Println()
	y := 24
	intConv := strconv.Itoa(y)
	fmt.Printf("Before conv %v - %T\n", y, y)
	fmt.Printf("After conv %v - %T\n", intConv, intConv)

}
