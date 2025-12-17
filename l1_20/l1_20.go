package main

import (
	"fmt"
)

func main() {
	str1 := "snow dog sun"
	str2 := "Hello Го"

	fmt.Println(str1)
	fmt.Println(reverseStr(str1))
	fmt.Println(str2)
	fmt.Println(reverseStr(str2))
}

func reverseStr(s string) string {
	r := []rune(s)
	n := len(r)

	reverseRunes(r, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || r[i] == ' ' {
			reverseRunes(r, start, i-1)
			start = i + 1
		}
	}

	return string(r)
}

func reverseRunes(r []rune, i, j int) {
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}