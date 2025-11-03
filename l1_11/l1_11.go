package main

import "fmt"

func intersection(a []int, b []int) []int {
	seen := make(map[int]bool)
	for _, v := range a {
		seen[v] = true
	}

	var result []int
	for _, v := range b {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	fmt.Println(intersection(a, b))
}
