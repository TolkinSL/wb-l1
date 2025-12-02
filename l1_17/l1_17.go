package main

import "fmt"

func binarySearch(arr []int, num int) int {
    low := 0
		high := len(arr)-1

    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == num {
            return mid
        }

        if arr[mid] < num {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

func main() {
    arr := []int{1, 3, 6, 8, 12, 15}

		fmt.Println("Массив:", arr)
    fmt.Println("Позиция 5:", binarySearch(arr, 5))
		fmt.Println("Позиция 3:", binarySearch(arr, 3))
    fmt.Println("Позиция 12:", binarySearch(arr, 12))
}
