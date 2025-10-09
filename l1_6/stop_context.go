package main

import (
	"context"
	"fmt"
	"time"
)

func byContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("\nГорутина завершена через context")
				return
			default:
				fmt.Println(" - Работает")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func main() {
	fmt.Println("Горутины - Остановка через context")
	byContext()
}
