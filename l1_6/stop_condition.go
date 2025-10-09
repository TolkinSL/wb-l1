package main

import (
	"fmt"
	"time"
)

func byCondition() {
	stop := false

	go func() {
		for !stop {
			fmt.Println(" - Работает")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("\nГорутина завершена по условию")
	}()

	time.Sleep(1 * time.Second)
	stop = true
	time.Sleep(500 * time.Millisecond)
}

func main() {
	fmt.Println("Горутины - Остановка по условию")
	byCondition()
}
