package main

import (
	"fmt"
	"time"

)


func main(){

	ch := make(chan int, 2) // Canal possui 2 de espaço

	go func() { 

		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "enviando")
			ch <- i 
			fmt.Println(time.Now(), i, "enviado")
		}

		// XXX: Pode haver casos onde essa mensagem não será completada
		// E isso será resolvido nos proximos exemplos


		fmt.Println(time.Now(), "tudo feito com sucesso.")
	
	}()

	// Colocamos esse fluxo para dormir, no intuito de dar tempo para o fluxo da
	// função acima terminar de ser executado

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "esperando as mensagens")

	// messageFromChannel := <- ch 

	fmt.Println(time.Now(), "Mensagem recebida na main", <- ch )
	fmt.Println(time.Now(), "Mensagem recebida na main", <- ch )
	fmt.Println(time.Now(), "Mensagem recebida na main", <- ch )

	fmt.Println(time.Now(), "FIM")


}


