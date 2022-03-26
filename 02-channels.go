package main

import "fmt"

// This program panics because there is no goroutine
// outside of `main` interacting with `channel`

// fatal error: all goroutines are asleep - deadlock!


func main(){

	var channel chan int // Declare variable channel

	channel = make(chan int) // Creates a channel


	output := <- channel // Output has channel value

	fmt.Println("Printando channel ", output)




}