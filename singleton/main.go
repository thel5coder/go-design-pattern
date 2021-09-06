package main

import "fmt"

func main(){
	var vidi = NewPeople("Vidi","Pratama", 8)

	fmt.Println(vidi.firstName+` `+vidi.lastName+` Umur `, vidi.age)
}

type People struct {
	firstName string
	lastName  string
	age       int
}

func NewPeople(firstName,lastName string,age int) People{
	return People{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}