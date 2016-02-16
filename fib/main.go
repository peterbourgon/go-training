package main

func fibonacci(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		println(i)
	}
}
