package examples

import "fmt"

type Hello struct {
	name string
}

type Age int

func (a Age) PrintAge() {
	fmt.Printf("User age is %v\n", a)
}

func (h Hello) methodExample() {
	fmt.Printf(h.name)
}

func TestMethodInGo() {
	d := Hello{"Pratik"}
	d.methodExample()

	userAge := Age(28)
	userAge.PrintAge()
}
