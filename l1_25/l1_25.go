package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	<-timer.C
}

func main() {
	myTime := 2
	fmt.Println("Страт")
	fmt.Printf("Ожидание %d секунды\n", myTime)
	mySleep(time.Duration(myTime) * time.Second)
	fmt.Println("Завершение")
}
