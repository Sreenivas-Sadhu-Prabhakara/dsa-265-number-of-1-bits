package main

import "fmt"

// OnesCount returns how many bits of n are set, clearing the lowest set bit
// per iteration: n & (n-1) — so the loop runs once per 1-bit, not per bit.
func OnesCount(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(OnesCount(0b1011))
	fmt.Println(OnesCount(0xFFFFFFFF))
}
