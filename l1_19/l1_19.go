package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	fmt.Println(Reverse(str))
}
