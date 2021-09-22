package main

import (
	"encoding/json"
	"fmt"
)

type character struct {
	Alias    string
	Name     string
	IsHero   bool
	Universe string
}

type (
	Animal struct {
		Name  string
		Class string
	}

	Bird struct {
		Animal
		CanFly bool
	}
)

func main() {
	//batman := character{
	//	Alias: "Batman",
	//}
	//
	//fmt.Println(batman)
	//
	b := Bird{}
	b.Animal.Name = "Dog"
	b.Animal.Class = "Mammal"
	b.CanFly = false
	//
	fmt.Println(b)
	//b := Bird{
	//	Animal{ "Dog",  "Mammal"},
	//	false,
	//}
	//

	//bCopy:= b
	//bCopy.Name = "Other name"
	//fmt.Println(b, bCopy)

	user := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ID        int    `json:"id"`
	}{
		FirstName: "David",
	}


	dataJson := `{"first_name":"David","last_name":"Castillo","id":1}`
	_ = json.Unmarshal([]byte(dataJson), &user)

	//fmt.Println(user)

	a := Animal{Name: "animal name"}

	if a == (Animal{}) {
		fmt.Println("enter here")
	}
}
