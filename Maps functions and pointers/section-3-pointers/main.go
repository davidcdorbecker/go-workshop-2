package main

func main() {
	Pointers()
}

func Pointers() {
	//// Checking addresses
	//var a int = 42
	//var b *int
	//
	//b = &a
	//*b = 5
	////
	//fmt.Println(a, *b)

	//// Arithmetic operations
	//a := [3]int{1, 2, 3}
	//b := &a[0]
	//c := &a[1]
	////
	//fmt.Printf("%v %p %p", a, b, c)

	//Pointers in structs
	//type person struct {
	//	name string
	//}

	//var p *person
	//p = &person{"David"}
	////p = new(person)
	//(*p).name = "John"
	//fmt.Println((*p).name)

	////Zero value for section-3-pointers
	//var p *person
	//fmt.Println(p)

	//important note
	//https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac
}
