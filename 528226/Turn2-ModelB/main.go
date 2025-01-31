// main.go
package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func counter(id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		println(id, i)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	wg.Add(2)
	go counter(1)
	go counter(2)
	wg.Wait()
}
