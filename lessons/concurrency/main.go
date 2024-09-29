package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	//sends data to the channel
	doneChan <- true
}

func main() {
	dones := make([]chan bool, 4)
	// done := make(chan bool)
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])

	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("How...are...you..?", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope you are ok!", dones[3])
	
	//<-done //data is coming out let it flow into the void
	for _, done := range dones {
		<-done
	}
}
