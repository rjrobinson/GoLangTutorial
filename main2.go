// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Println("Hello world! . Again... we printed it again. ")
//
// 	// in ruby
// 	// puts "Hello World"
// 	// or
// 	// p "Hello World"
// }

package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
