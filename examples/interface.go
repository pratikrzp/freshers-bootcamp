package examples

import "fmt"

type Food interface {
	Prepare()
	Serve()
}

type Pizza struct {
	pizzaType string
}
type Salad struct {
	vegan bool
}
type Rasmalai struct {
	hasAlmond bool
}

func (p Pizza) Prepare() {
	fmt.Printf("prepare pizza with user's type: %v\n", p.pizzaType)
}
func (p Pizza) Serve() {
	fmt.Println("server it hot with some coke")
}

func (s Salad) Prepare() {
	fmt.Printf("chop some veggies with some pepper and it is vegan: %v\n", s.vegan)
}
func (s Salad) Serve() {
	fmt.Println("serve it in a bowl")
}

func (r Rasmalai) Prepare() {
	fmt.Printf("Yumm! It contains almond: %v\n", r.hasAlmond)
}
func (r Rasmalai) Serve() {
	fmt.Println("serve it with some dry fruits")
}

func orderFood(food Food) {
	food.Prepare()
	food.Serve()
}

func InterfaceExample() {
	pizza := Pizza{pizzaType: "Extra cheese"}
	orderFood(pizza)

	rasmalai := Rasmalai{hasAlmond: false}
	orderFood(rasmalai)

	salad := Salad{vegan: true}
	orderFood(salad)
}
