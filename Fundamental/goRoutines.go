package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //we use this to notify if there is go routine that is currently running

func say(s string) {
	//we need to notify the wait group if it is done
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	defer wg.Wait()

	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("There")

}
