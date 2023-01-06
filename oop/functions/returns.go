package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func getMultiples(x int) (double, triple, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(4, 10, 30, 45))
	double, triple, quad := getMultiples(35)
	fmt.Println(double, triple, quad)
}
