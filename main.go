package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("")
	http.ListenAndServe(":3000", nil)
}

func jiechen(n int) int {
	x := 1
	for i := 2; i <= n; i++ {
		x *= i
	}
	return x
}
