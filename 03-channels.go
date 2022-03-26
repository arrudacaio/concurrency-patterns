package main

import (
	"fmt"
	"time"

)


func main(){

	ch := make(chan string)

	go func() { // Este fluxo de execução só encerrou depois, e mandou a mensagem pro fluxo da main através do channel

		fmt.Println(time.Now(), "taking a nap")
	
		time.Sleep(2 * time.Second)
	
		ch <- "Hello"
	
	}()

	fmt.Println(time.Now(), "waiting a message")

	messageFromChannel := <- ch 

	fmt.Println(time.Now(), "Received message", messageFromChannel)

}


