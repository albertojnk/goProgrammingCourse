package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("HELLOL PORRA")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Alberto"}
	saySomething(&p)

	// does not work -> saySomething(p)
	// this does work -> p.speak()

	p.speak()

}
