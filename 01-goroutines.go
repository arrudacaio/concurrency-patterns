package main

import "fmt"

func main(){
	go hello() // Não printa nada pois nesse codigo não há nada falando: ei função main, voce pode esperar até que a função hello termine de executar? 

}

func hello() {
	fmt.Println("Startandoooo tudo")
}


