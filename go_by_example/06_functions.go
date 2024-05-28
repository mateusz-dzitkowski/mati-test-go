package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func times(a int, b int) int {
	output := 0
	for range b {
		output = plus(output, a)
	}
	return output
}

func power(a int, b int) int {
	output := 1
	for range b {
		output = times(output, a)
	}
	return output
}

func mod(a int, b int) (int, int) {
	n := 1
	for {
		if b*n >= a {
			return n - 1, b*n - a
		}
		n += 1
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = plus(total, num)
	}
	return total
}

func prod(nums ...int) int {
	total := 1
	for _, num := range nums {
		total = times(total, num)
	}
	return total
}

func factorial(n int) int {
	nums := make([]int, n)
	for i := range n {
		nums[i] = i + 1
	}
	return prod(nums...)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(plus(3, 2))
	fmt.Println(times(3, 2))
	fmt.Println(power(3, 2))
	fmt.Println(mod(5, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(prod(1, 2, 5, 2, 2))
	fmt.Println(factorial(5))
	fmt.Println(fib(50))
}
