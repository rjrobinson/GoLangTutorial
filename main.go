// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Println("Hello, World!")
// }

// package main
//
// import "fmt"
//
// func main() {
// 	//the basic for loop
// 	for i := 1; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main
//
// import "fmt"
//
// func main() {
// 	// defines the var a
// 	a := 5
// 	fmt.Println(a)
//
// 	// sets a different value to a
// 	a = 10
// 	fmt.Println(a)
//
// 	//another way to define a variable
// 	var b int
// 	b = 15
// 	fmt.Println(b)
// }

// package main
//
// import "fmt"
//
// func main() {
// 	/// this creates a slice of integers with len of 15
// 	mySlice := make([]int, 15)
// }

package main

import "fmt"

func wait() {
	//wait around with a forever loop
	for {

	}
}

func main() {
	go wait()
	fmt.Println("We didn't wait because it was called as a goroutine!")
}
