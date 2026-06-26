/*
Bitwise operators work directly on the binary representation of integer values, manipulating one bit at a time. In Go, these operators are essential for low-level programming, bit masks, flags, and efficient arithmetic (shifts are equivalent to multiplying/dividing by powers of two for non-negative integers). This guide covers the Go bitwise operators with clear binary demonstrations and runnable examples.
Bitwise AND (&)
Bitwise OR (|)
Bitwise XOR (^)
Left shift (<<)
Right shift (>>)
*/

package main

import "fmt"

func main() {
	var x, y int = 12, 25
	z := x | y
	fmt.Println(z) // 29
}
