package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 15; i++ {
		<-tick
		fmt.Println("Запрос", i, "обработан")
	}
}