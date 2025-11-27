package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := []int{}
	great := []int{}

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			great = append(great, v)
		}
	}

	result := append(quickSort(less), pivot)
	result = append(result, quickSort(great)...)
	return result
}

func main() {
	arr := []int{3, 5, 2, 4, 1, 7}
	fmt.Println("Первоначальный массив, pivot первый элемент 3")
	fmt.Println(arr)

	sorted := quickSort(arr)
	fmt.Println("Отсортированный массив")
	fmt.Println(sorted)
}
