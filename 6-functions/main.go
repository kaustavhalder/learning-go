// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("this is the main function")

// 	// Function Expression
// 	f := func() {
// 		fmt.Println("this is an function expression")
// 	}
// 	f()
// 	foo()
// 	bar()
// 	b2 := bar2()
// 	fmt.Println(b2())
// 	ii := []int{2, 5, 5, 5}
// 	su := sum(ii...)
// 	fmt.Println(su)

// }

// // Returning a string
// func foo() string {
// 	return "hello world"
// }

// // Returning a function
// // need to mention the return type of the inner function as well
// func bar() func() int {
// 	return func() int {
// 		return 451
// 	}
// }

// func bar2() func() string {
// 	return func() string {
// 		fmt.Println("Alibaba")
// 		return "Alibaba"
// 	}
// }

// // Callback
// // Passing a function as a parameter eg: JavaScript

// func sum(x ...int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }



package main
import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var text string
    count := 0
    fmt.Scanln(&text)
    for i := 0; i < len(text); i ++ {
        count += 1
    }
    fmt.Println(count)
}