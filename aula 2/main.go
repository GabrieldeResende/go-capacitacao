package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // T1
	canal := make(chan int)

	qtdWorkers := 3

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal)
	}

	// for i := 0; i < 10; i++ {
	// 	canal <- i
	// }
}

// processo = alocar um bloco de memoria
// thread = Acessar o bloco de memoria
// T1 e T2 acessam o mesmo bloco de memoria
// Race Condition = condição de corrida

// Go Routine = canal = go Routine 2
// Input = output
