//

// package main

// import "fmt"

// func main() {
// 	var i int = 90
// 	var f float64 = float64(i)
// 	fmt.Printf("%.2f\n", f)
// }

////Converting a Float to an Integer

// package main

// import "fmt"

// func main() {
//     var f float64 = 45.89
//     var i int = int(f)
//     fmt.Printf("%v\n", i)
// }

//Converting an Integer to a String


// package main

// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     var i int = 42
//     var s string = strconv.Itoa(i) // convert int to string
//     fmt.Printf("%q", s)
// }


//Converting a String to an Integer

package main

import (
    "fmt"
    "strconv"
)

func main() {
    var s string = "200"
    i, err := strconv.Atoi(s)
    fmt.Printf("%v, %T \n", i, i)
    fmt.Printf("%v, %T", err, err)
}