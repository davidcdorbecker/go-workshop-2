package main

import (
	"fmt"
)

func main() {
	Interfaces()
}

func Interfaces() {
	// Basic interface implementations
	//var w Writer = ConsoleWriter{}
	//w.Write([]byte("hello world!"))
	//
	//// Interfaces mutable
	//i := IntCounter(0)
	//var inc Incrementer = &i
	//for i:= 0; i < 10; i++ {
	//	fmt.Println(inc.Increment())
	//}
}

type Writer interface {
	Write(interface{}) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}