package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	
	table <- new(Ball) // game on, toss the ball (chan send)
	time.Sleep(1*time.Second)
	c := <- table //game over, grab the ball
	fmt.Println(c.hits)
	
	panic("show me the stacks")
	// chuong trinh ket thuc do ham main ket thuc
	// goroutine van con chay background
	// that's a leak
	// => su dung select trong goroutine de clean up
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100*time.Millisecond)
		table <- ball
	}
}