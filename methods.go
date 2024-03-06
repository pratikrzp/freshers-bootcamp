package main

import "fmt"

type Hello struct {
	name string
}

func (h Hello) methodExample() {
	fmt.Println(h.name)
}

func TestMethodInGo() {
	d := Hello{"Pratik"}

	d.methodExample()
}
