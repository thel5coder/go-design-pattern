package main

import "fmt"

func main() {
	builder := NewPeopleBuilder()
	vidi := builder.SetFirstName("vidi").SetLastName("pandu").SetAge(9).Build()

	fmt.Println(vidi.firstName+` `+ vidi.lastName+` umur `,vidi.age)
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

type IBuilder interface {
	SetFirstName(firstName string) IBuilder

	FirstName() string

	SetLastName(lastName string) IBuilder

	LastName() string

	SetAge(age int) IBuilder

	Age() int

	Build() People
}

type PeopleBuilder struct {
	firstName string
	lastName  string
	age       int
}

func NewPeopleBuilder() PeopleBuilder {
	return PeopleBuilder{}
}

func (p *PeopleBuilder) SetFirstName(firstName string) IBuilder {
	p.firstName = firstName

	return p
}

func (p *PeopleBuilder) FirstName() string {
	return p.firstName
}

func (p *PeopleBuilder) SetLastName(lastName string) IBuilder {
	p.lastName = lastName

	return p
}

func (p *PeopleBuilder) LastName() string {
	return p.lastName
}

func (p *PeopleBuilder) SetAge(age int) IBuilder {
	p.age = age

	return p
}

func (p *PeopleBuilder) Age() int {
	return p.age
}


func (p *PeopleBuilder) Build() People {
	return NewPeople(p.firstName,p.lastName,p.age)
}
