package main

import "fmt"

type SetCommand struct {
	key   string
	value int
	done  chan bool
}

type GetCommand struct {
	key   string
	reply chan int
}

func manager(setCh chan SetCommand, getCh chan GetCommand) {
	state := make(map[string]int)

	for {
		select {
		case cmd := <-setCh:
			state[cmd.key] = cmd.value
			cmd.done <- true

		case cmd := <-getCh:
			cmd.reply <- state[cmd.key]
		}
	}
}

func main() {
	setCh := make(chan SetCommand)
	getCh := make(chan GetCommand)

	go manager(setCh, getCh)

	// Установить значение
	done := make(chan bool)
	setCh <- SetCommand{key: "counter", value: 10, done: done}
	<-done

	// Получить значение
	reply := make(chan int)
	getCh <- GetCommand{key: "counter", reply: reply}
	fmt.Println("counter =", <-reply)
}