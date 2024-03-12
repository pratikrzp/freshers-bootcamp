package main

import (
	"time"

	"github.com/pratikrzp/freshers-bootcamp/examples"
)

func main() {
	examples.ArraySlice()
	examples.MapPractice()
	examples.TestMethodInGo()
	examples.InterfaceExample()
	go examples.SayHello("Pratik")
	time.Sleep(1000)
}
