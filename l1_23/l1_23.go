package main

import "fmt"

func remove(sl []string, i int) []string {
	if i >= len(sl) {
		return sl
	}

	copy(sl[i:], sl[i+1:])
	sl[len(sl)-1] = ""

	return sl[:len(sl)-1]
}

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g"}

	fmt.Println(arr)
	arr = remove(arr, 1)
	fmt.Println(arr)
	arr = remove(arr, 3)
	fmt.Println(arr)
}
