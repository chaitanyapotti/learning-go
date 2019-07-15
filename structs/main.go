package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateNameByPointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	chai := person{firstName: "Chaitanya", lastName: "Potti"}
	alex := person{"Alexie", "Hopper", contactInfo{"alex@hoo.com", 245322}}
	fmt.Println(chai)
	fmt.Println(alex)

	var joyce person
	joyce.firstName = "Joyce"
	joyce.lastName = "Byers"
	fmt.Printf("%+v", joyce)
	fmt.Println(joyce.toString())
	joyce.updateName("Will")

	joycePointer := &joyce
	joycePointer.updateNameByPointer("Nathan")

	joyce.updateName("Joyce")

	joyce.print()

}

// slices, maps, channels, pointers, functions are all reference types in go
// go makes a copy of this ds but the copy still points to the same underlying ds
// it means, they reference to a different underlying data structure

// struct, string are value types
