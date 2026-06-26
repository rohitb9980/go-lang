package main

import "fmt"

// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanf("%s", &name)
// 	fmt.Println("Hey there,", name)
// }

// func main() {
// 	var name string
// 	var isMuggle bool
// 	fmt.Print("Enter your name & are you a muggle: ")
// 	fmt.Scanf("%s %t", &name, &isMuggle)
// 	fmt.Println(name, isMuggle)
// }

func main() {
	var a string
	var b int
	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count:", count)
	fmt.Println("error:", err)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
