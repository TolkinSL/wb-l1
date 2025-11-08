package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	myMap := make(map[string]struct{})

	for _, v := range strs {
		myMap[v] = struct{}{}
	}

	fmt.Println("Созданное множество слов:")
	for v := range myMap {
		fmt.Println(v)
	}
}
