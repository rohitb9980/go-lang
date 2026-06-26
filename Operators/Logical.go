// // AND Operator (&&)

// package main

// import "fmt"

// func main() {
//     var x int = 10
//     fmt.Println((x < 100) && (x < 200)) // Both conditions are true, so the result is true.
//     fmt.Println((x < 300) && (x < 0))     // One condition is false, so the result is false.
// }

// //OR Operator (||)

// package main

// import "fmt"

// func main() {
//     var x int = 10
//     fmt.Println((x < 0) || (x < 200))  // One condition is true, so the result is true.
//     fmt.Println((x < 0) || (x > 200))  // Both conditions are false, so the result is false.
// }

// NOT Operator (!)

package main

import "fmt"

func main() {
	var x, y int = 10, 20
	fmt.Println(!(x > y)) // x > y is false; applying NOT converts it to true.
	fmt.Println(!true)    // Inverts true to false.
	fmt.Println(!false)   // Inverts false to true.
}
