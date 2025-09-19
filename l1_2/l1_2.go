package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 8}
	results := make(chan string, len(numbers))

	for _, num := range numbers {
		go func(n int) {
			results <- fmt.Sprintf("%d^2 = %d", n, n*n)
		}(num)
	}

	fmt.Println("Конкурентное возведение в квадрат")
	fmt.Println(numbers)

	for range numbers {
		fmt.Println(<-results)
	}

}
