package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := CounterPairs1([]int{6, 90, 234, 11, 39, 700}, 101)
	fmt.Printf("Answer: %d \n", c)
}

func CounterPairs1(l []int, k int) int {
	c := 0
	pairs := map[string]int{}

	for i := 0; i < len(l); i++ {
		if l[i] >= k {
			continue
		}
		for j := i + 1; j < len(l); j++ {
			key := strconv.Itoa(l[i]) + "," + strconv.Itoa(l[j])
			keyInv := strconv.Itoa(l[j]) + "," + strconv.Itoa(l[i])
			_, exists := pairs[key]
			_, existsInv := pairs[keyInv]
			s := l[i] + l[j]
			if s <= k && !(exists || existsInv) {
				pairs[key] = s
				c++
			}
		}
	}

	return c
}
