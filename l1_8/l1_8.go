package main

import (
	"fmt"
	"os"
	"strconv"
)

func setBit(n int64, i uint, v int) int64 {
	if v == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Укажите параметр: go run l1_8.go число номер_бита 0_или_1")
		return
	}

	n, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	i, err2 := strconv.Atoi(os.Args[2])
	v, err3 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Ошибка: пример правильного вызова → go run l1_8.go 5 1 0")
		return
	}

	fmt.Printf("Начальное число: %d (в двоичном виде %08b)\n", n, n)
	// n = setBit(n, uint(i), v) // в Go индекс 0 - самый первый бит
	n = setBit(n, uint(i-1), v)
	fmt.Printf("Результат: %d (в двоичном виде %08b)\n", n, n)
}
