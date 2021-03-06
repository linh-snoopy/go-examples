package main

import (
    "fmt"
    "os"
    "os/signal"
    "time" // or "runtime"
)

func cleanup() {
    fmt.Println("cleanup")
}

func main() {
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        cleanup()
        os.Exit(1)
    }()

    for {
        fmt.Println("sleeping...")
        time.Sleep(5 * time.Second) // or runtime.Gosched() or similar per @misterbee
    }
}