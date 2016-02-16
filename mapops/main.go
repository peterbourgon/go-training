package main

func main() {
	m := map[int]int{}
	for i := 1; i <= 100; i++ {
		m[i] = i * i
	}

	var sumSquareEven, sumSquareOdd int
	for k, v := range m {
		if k%2 == 0 {
			sumSquareEven += v
		} else {
			sumSquareOdd += v
		}
	}

	println("sum of squares of even numbers:", sumSquareEven)
	println("sum of squares of odd numbers:", sumSquareOdd)
}
