package main

import "fmt"

func TwoSum(a []int, target int) [2]int {
	var result [2]int
	m := make(map[int]int)

	for i := 0; i < len(a); i++ {

		value, ok := m[target - a[i]]

		if ok {
			// Match found
			result[0] = value
			result[1] = i

			break
		} else {
			m[a[i]] = i
		}
	}

	return result
}

func main() {
	// Example

	a := []int{2, 7, 11, 15}
	result := TwoSum(a, 9)

	fmt.Println("Given input: ")
	fmt.Println(a)
	fmt.Println("TwoSum returns: ")
	fmt.Println(result)
}
