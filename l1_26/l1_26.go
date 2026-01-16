package main

import (
	"fmt"
	"strings"
)

func isUniq(str string) bool {
	var lower_str string
	lower_str = strings.ToLower(str)
	// lower_str = str // для тестирования ошибки

	runes := []rune(lower_str)
	r_map := make(map[rune]struct{})

	for _, r := range runes {
		_, ok := r_map[r]

		if !ok {
			r_map[r] = struct{}{}
		} else {
			return false
		}
	}

	return true
}

func main() {
	var input_str string
	var result bool

	fmt.Scanln(&input_str)
	result = isUniq(input_str)

	if result {
		fmt.Printf("Строка: %s содержит уникальные символы\n", input_str)
	} else {
		fmt.Printf("Строка: %s не содержит уникальные символы\n", input_str)
	}
}
