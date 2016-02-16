package main

func main() {
	for i := 0; i < 3; i++ {
		func() {
			println("hello", i)
			defer println("goodbye", i)
		}()
	}
}
