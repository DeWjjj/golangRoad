package main

var a, b chan int

func main() {
	a = make(chan int)
	b = make(chan int)
	for i := 0; i < 100; i++ {
		a <- i
	}

}
