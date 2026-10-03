package main

type SetCommand struct {
	key   string
	value int
	done  chan struct{}
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
			cmd.done <- struct{}{}

		case cmd := <-getCh:
			cmd.reply <- state[cmd.key]
		}
	}
}

func main() {
	setCh := make(chan SetCommand)
	getCh := make(chan GetCommand)

	go manager(setCh, getCh)

	done := make(chan struct{})
	setCh <- SetCommand{key: "counter", value: 10, done: done}
	<-done

	reply := make(chan int)
	getCh <- GetCommand{key: "counter", reply: reply}
	fmt.Println("counter =", <-reply)
}
