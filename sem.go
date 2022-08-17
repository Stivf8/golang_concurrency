package main

import (
	"fmt"
	"sync"
	"time"
)

//trabaja como semaforo, espera que el canal sea liberado con la ejecucion de la linea 29
// c := [][]
// c := [goRoutine1][goRoutine2]
// c:= [goRoutine3][goRoutine2]

func mainsem() {
	//define canal de 5 espacios
	c := make(chan int, 5)
	//define waitGroup
	var wg sync.WaitGroup
	//for para enviar rutinas 10 veces
	for i := 0; i < 10; i++ {
		//enviamos dato al canal
		c <- 1
		//incrementamos el waitGroup
		wg.Add(1)
		//ejecutamos funcion, enviando indice, la copia del waitGroup, y el canal
		go doSomething(i, &wg, c)
	}
	//ejecutamos .Wait para que este espere hasta que el waitGroup este en 0 para finalizar ejecucion
	wg.Wait()
}

//funcion recibimos variable i int, waitGroup, canal
func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	//al final se ejecuta el defer wg.Done para decrementar el waitGroup
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
	//confirmamos al canal c
	<-c
}
