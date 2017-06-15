package main

import (
	"fmt"
	"sync"
	"time"
)

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func say(s string, len int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say("world", 5, &wg)
	wg.Add(1)
	say("hello", 2, &wg)
	fmt.Println("Done!")
	wg.Wait()
}
