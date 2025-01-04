package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func mainTut3a() {
	counts := make(map[string]int)
	// files := os.Args[1:]
	// if (len(files) == 0) {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		file, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(file, counts)
	// 		file.Close()
	// 	}
	// }
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}