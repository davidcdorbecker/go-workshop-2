package main

import (
	"errors"
	"fmt"
)

var divide func(float64, float64) (float64, error)

func main() {
	g1, _ := sayHello("David", "Castillo", 25)
	fmt.Println(g1)

	fmt.Println(sum(1, 2, 3, 4, 5))

	//IIFE
	func() {
		fmt.Println("hello world!")
	}()

	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, errors.New("Cannot divide by zero")
		}

		return a / b, nil
	}
	//
	res, err := divide(1.2, 0)
	if err != nil {
		// business logic
	}
	fmt.Println(res, err)

	fibo := myFibo()
	fmt.Println(fibo(60))
}

func sayHello(name, lastName string, age int) (string, string) {
	return fmt.Sprintf("Hello %s %s, you are %d years old", name, lastName, age),
		fmt.Sprintf("Hello %s ", name)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0

	for _, value := range values {
		result += value
	}

	return result
}

// 0 1 1 2 3 5 8
// 0 1 2 3 4 5 6
func myFibo() func(pos int) int {
	cache := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	var fibo func(pos int) int
	fibo = func(pos int) int {
		data, ok := cache[pos]
		if ok {
			return data
		}

		fiboNumber := fibo(pos-1) + fibo(pos-2)
		cache[pos] = fiboNumber

		return fiboNumber
	}

	return fibo
}
