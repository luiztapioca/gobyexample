package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2*r.height + 2*r.width
}

type pessoa struct {
	name string
	sph  float32
	age  int
}

func (p *pessoa) salary() float32 {
	return 30 * 8 * p.sph
}

func main() {
	r := rect{2, 2}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	jonas := pessoa{name: "Teste", age: 30, sph: 8}

	fmt.Println("sslary per month:", jonas.salary())
}
