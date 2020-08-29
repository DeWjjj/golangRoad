package main

import (
	"fmt"
	"time"
)

func main() {
	thisTime, err := time.Parse("2006/01/02 15/04/05", "2020/08/28 15/43/21")
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	NextyearThistime, err := time.Parse("2006/01/02 15/04/05", "2020/08/29 15/43/21")
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	val := NextyearThistime.Sub(thisTime) //24h0m0s
	fmt.Println(val)

}
