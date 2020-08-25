package main

import (
	"fmt"
)

type allCar interface {
	run()
}

type ferrari struct {
	models string
}

func (f ferrari) run() {
	fmt.Printf("%v is a ferrari.", f.models)

}

type bentley struct {
	models string
}

func (b bentley) run() {
	fmt.Printf("%v is a bentley.", b.models)
}

func whatCar(a allCar) {
	a.run()
}

func main() {
	var car1 = ferrari{
		models: "Portofino",
	}
	whatCar(car1)
}
