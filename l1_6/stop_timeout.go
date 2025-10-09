package main

import (
	"fmt"
	"time"
)

func byTimeout() {
	timeout := time.After(1 * time.Second)

	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("\nГорутина завершена таймаут")
				return
			default:
				fmt.Println(" - Работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
}

func main() {
	fmt.Println("Горутины - Остановка по таймауту")
	byTimeout()
}
