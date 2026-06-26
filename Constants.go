/*There are two main types of constants in Go:
Untyped Constants: These are flexible and do not have a fixed data type unless explicitly defined.
Typed Constants: These require a specific data type, offering stronger type checking.
*/

// package main

// import "fmt"

// func main() {
// 	const name = "Hermione Granger"
// 	fmt.Printf("%v: %T \n", name, name)
// }


//Use Case: Calculating the Area of a Circle
package main
import "fmt"

const PI float64 = 3.14 // global constant

func main() {
    var radius float64 = 5.0
    var area float64
    area = PI * radius * radius
    fmt.Println("Area of Circle is :", area)
}