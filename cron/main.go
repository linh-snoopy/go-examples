package main

import (
	"github.com/robfig/cron"
	"fmt"
	"os/signal"
	"os"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("1 * * * * *", func() { fmt.Println("The first second of every seconds", time.Now()) })
	c.Start()
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
	fmt.Println(sig)
}