package main

import "fmt"
import "crypto/sha256"


var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func popcount(c1, c2 [32]byte) int {
    var count int
    for i := 0; i < 32; i++ {
        count += int(pc[c1[i]^c2[i]])
    }
    return count
}

func reverse(s []int){
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    fmt.Println("exercise 4.1 - count:", popcount(c1, c2))

    // rotate slice
    s := []int{0, 1, 2, 3, 4, 5}
    fmt.Println("exercise 4.2:", s)
    fmt.Println(" s[:2]:", s[:2])
    reverse(s[:2])
    fmt.Println(" s[2:]:", s[2:])
    reverse(s[2:])
    reverse(s)
    fmt.Println("exercise 4.2 - reverse:", s)
    


}
