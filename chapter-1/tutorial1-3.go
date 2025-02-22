package main

import (
	"bufio"
	"fmt"
	"os"
)

func maintut3() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Enter Text: ")
		input.Scan()
		text := input.Text()
		if len(text) != 0 {
			counts[text]++
		} else {
			break
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

// func main() {
// 	// To create dynamic array
// 	arr := make([]string, 0)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("Enter Text: ")
// 		// Scans a line from Stdin(Console)
// 		scanner.Scan()
// 		// Holds the string that scanned
// 		text := scanner.Text()
// 		if len(text) != 0 {
// 			fmt.Println(text)
// 			arr = append(arr, text)
// 		} else {
// 			break
// 		}

// 	}
// 	// Use collected inputs
// 	fmt.Println(arr)
// }
