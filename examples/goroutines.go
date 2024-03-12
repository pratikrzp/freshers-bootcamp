package examples

import "fmt"

func SayHello(name string) {
	if len(name) == 0 {
		fmt.Println("Hello Stranger!!")
		return
	}
	fmt.Printf("Hello %v.\n", name)
}
