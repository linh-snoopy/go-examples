package main

import "fmt"

type IA interface {
	FB() IB
}

type IB interface {
	Bar() string
}

type A struct {
	b *B
}

func (a *A) FB() IB {
	return a.b
}

type B struct{}

//func (a *A) Bar() string {
//	return "A Bar!"
//}

func (b *B) Bar() string {
	fmt.Println("run Bar")
	return "B Bar!"
}

func main() {
	a := &A{}
	CheckIA(a)
}

func CheckIA(a IA) {
	fmt.Println("run checkIA")
	fmt.Printf("%T", a.FB())
}
