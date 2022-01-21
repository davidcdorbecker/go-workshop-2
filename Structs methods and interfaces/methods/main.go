package main

import "fmt"

func main() {
	Methods()
}

type Person struct {
	name string
	age  int
}

type Names []string

func (n *Names) AddName(names ...string) {
	*n = append(*n, names...)
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func NewPerson (name string, age int) Person {
	return Person{name, age}
}

func Methods() {
	//p := Person{"David", 25}
	//fmt.Println(p.GetName())

	var names Names
	names.AddName("David", "Alvaro", "Cesar")

	fmt.Println(names)
}
