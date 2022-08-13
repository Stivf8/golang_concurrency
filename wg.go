package main

import (
	"fmt"
	"sync"
	"time"
)

/*
creates a wait group that gets incremented by on on every iteration, which
will as well run the doSmth() func that will substract 1 from the wait group
after it finishes. wg.Wait() at the end guarantees that it'll wait for the wg (counter)
to be 0.
*/

func doSmth(u int, wg *sync.WaitGroup) {
	//se ejecuta al final de la func doSmth, restando 1 al waitGroup
	defer wg.Done()

	fmt.Printf("Started at #%d\n", u)
	time.Sleep(time.Second * 2)
	fmt.Println("Ended...")
}

func main2() {
	//isntanciamos un waitGroup
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//Cada que ejecute una goRoutine incrementa el grupo en 1, para que sea esperado
		wg.Add(1)
		go doSmth(i, &wg)
	}
	//espera que el wg este en 0 para finalizar la ejecucion
	wg.Wait()
}
