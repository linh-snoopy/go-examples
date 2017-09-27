package main

import (
	"strings"
	"os/exec"
	"fmt"
)

func main() {
	fmt.Println("Starting ping")
	cmd := exec.Command("ping", "www.google.com")
	out, err := cmd.CombinedOutput()
	fmt.Println("1111111111111111")
	if err != nil {
		fmt.Println("AAAAAAAAAAAAAAA")
	}
	fmt.Println(string(out))
	if strings.Contains(string(out), "Destination Host Unreachable") {
	    fmt.Println("TANGO DOWN")
	} else {
	    fmt.Println("IT'S ALIVEEE")
	}
}