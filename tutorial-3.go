package main

import (
	"fmt"
	"math"	
)

const value string = "132456778990"

func main() {
	//mainTutorial3()

	fmt.Printf("  %s\n", comma(value))
	fmt.Printf("commaNoRecursion::  %s\n", commaNoRecursion(value))
	
}
// Main function Tutorial 3 //
// This is the third tutorial of the GO language. //
func mainTutorial3() {
	for x := 0;	x < 8; x++ {
		fmt.Printf("x = %d e^x = %.5f\n", x, math.Exp(float64(x)))
	}

}


func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNoRecursion(s string) string {
	
	result := []byte{}
    
    // Iterate backwards to add commas
    for i := len(s) - 1; i >= 0; i-- {

        result = append(result, s[i])

        if (len(s)-i)%3 == 0 && i != 0 {
            result = append(result, ',')
        } 
    }
    
    // Reverse the result slice
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result)
}
