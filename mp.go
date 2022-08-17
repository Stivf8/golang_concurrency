package main

import (
	"fmt"
	"time"
)

func doSomethingmp(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomethingmp(d1, c1, 1)
	go doSomethingmp(d2, c2, 2)

	//la multiplexacion nos permite tener diferentes canales y necesitamos saber cual es el canal que ha sido activado, esta pendiente
	//al primero que acabe
	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}
