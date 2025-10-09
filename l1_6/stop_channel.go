package main

import (
	"fmt"
	"time"
)

func byChannel() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("\nГорутина завершена через канал")
				return
			default:
				fmt.Println(" - Работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(stop)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	fmt.Println("Горутины - Остановка через канал")
	byChannel()
}
