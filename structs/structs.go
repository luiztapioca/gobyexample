package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Anna", age: 30})
	fmt.Println(person{name: "Ricahrd"})
	fmt.Println(person{age: 20})
	fmt.Println(&person{name: "John"})

	fmt.Println(newPerson("Cesar"))

	s := person{"Giancarlo", 20}
	sp := &s

	fmt.Println(sp.name)
	fmt.Println(sp.age)

	sp.age = 40
	fmt.Println(sp.age)
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
