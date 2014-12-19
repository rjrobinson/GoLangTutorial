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

// package main
//
// import "fmt"
//
// func wait() {
// 	//wait around with a forever loop
// 	for {
//
// 	}
// }
//
// func main() {
// 	go wait()
// 	fmt.Println("We didn't wait because it was called as a goroutine!")
// }

// package main
//
// import (
// 	"fmt"
// 	"net"
// )
//
// //notice that in the arguments, the name of
// //the variable comes first, then comes the
// //type of the variable, just like in "var"
// //declarations
//
// func manageClient(conn net.Conn) {
// 	conn.Write([]byte("Hi!"))
// 	conn.Close()
// 	//do something with the client
// }
//
// func main() {
// 	//we are creating a server here that listens
// 	//on port 1337. Notice that, similar to Ruby,
// 	//a method can have two return values (although
// 	//in Ruby, this would be an array instead)
//
// 	listener, err := net.Listen("tcp", ":1337")
// 	for {
// 		//accept a connection
// 		connection, _ := listener.Accept()
// 		go manageClient(connection)
// 	}
// }

// package main
//
// import "fmt"
//
// var eventChannel chan int = make(chan int)
//
// func sayHello() {
// 	fmt.Println("Hello, world!")
//
// 	//pass a message thorugh the eventChannel
// 	//it dosen't matter *what* we actually send across
//
// 	eventChannel <- 1
// }
//
// func main() {
//
// 	//run a goroutine taht says hello
// 	go sayHello()
//
// 	//read the eventChannel
// 	//this call blocks so it waits until sayHello()
// 	//is done
//
// 	<-eventChannel
// }

// package main
//
// import (
// 	"fmt"
// )
//
// var logChannel chan string = make(chan string)
// 
// func loggingLoop() {
// 	for {
// 		//wait for a message
//
// 		msg := <-logChannel
//
// 		//log the message
// 		fmt.Println(msg)
// 	}
// }
//
// func main() {
// 	go loggingLoop()
//
// 	// do some stuff here
// 	logChannel <- "messaged to be logged"
// 	//do some more stuff.
//
// }
