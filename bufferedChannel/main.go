package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	test := make(chan int, 4)
	var times int
	flag.IntVar(&times, "times", 5, "number of times to repeat")
	flag.Parse()
	fmt.Println("Repeat", times, "times")
	
	go func() {
		for i := 1; i <= times; i++ {
			fmt.Println("------>", i, "to channel")
			test <- i
		}
		close(test)
	}()
	for {
		i := <-test
		fmt.Println("Pop", i, "from channel")
		if i == times {
			fmt.Println("Done!")
			return
		}
		time.Sleep(2*time.Second)
	}
}