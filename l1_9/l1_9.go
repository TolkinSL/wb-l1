package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, n := range arr {
			ch1 <- n
		}
		close(ch1)
	}()

	go func() {
		for n := range ch1 {
			ch2 <- n * 2
		}
		close(ch2)
	}()

	for res := range ch2 {
		fmt.Println(res)
	}
}
