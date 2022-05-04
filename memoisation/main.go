package main

import "fmt"

func main() {
	cache := make(map[int64]int64)
	for i := int64(1); i < 100; i++ {

		fmt.Println(fib(i, cache))
	}
}

func fib(n int64, cache map[int64]int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if cache[n] == 0 {
		cache[n] = fib(n-1, cache) + fib(n-2, cache)
	}
	return cache[n]
}
