package main

import (
	"fmt"
	"strings"
)

// Main function Tutorial 1 //
// This is the first tutorial of the GO language. //
func mainTutorial1() {
	fmt.Println("Hello, GO!!")
	inputStuffSlice := []string{"My", "name", "is", "Matt", "McQuaid"}
	var s, name, sep string

	// For loop //
	for i := 0; i < len(inputStuffSlice); i++ {
		fmt.Println(i)
		s += sep + inputStuffSlice[i]
		sep = " "
	}

	fmt.Println(s)

	// For loop with range //
	for _, arg := range inputStuffSlice{
		name += strings.Join([] string{arg}, " ")
	}
	
	fmt.Println(name)


}