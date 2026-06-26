package main

import "fmt"

func main() {
	var name string = "ROhit Bondre"
	var age int = 31
	var isMarried bool = true
	var height float64 = 5.9
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Married:", isMarried)
	fmt.Println("Height:", height)
	fmt.Printf("Type of name %T\n", name)
	fmt.Printf("Type of age %T\n", age)
	fmt.Printf("Type of isMarried %T\n", isMarried)
	fmt.Printf("Type of height %T\n", height)
}
