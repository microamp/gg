package main

import (
	"fmt"
	"log"
        "os/exec"
)

func send(ch chan<- string, command string, args string) {
	cmd := exec.Command(command, args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n != 0 {
			msg := string(buf[:n])
			fmt.Println("message sent:", msg)
			ch <- msg
		}
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func receive(ch <-chan string) {
	/* receives messages from channel, writes to db */
	for {
		msg, ok := <-ch
		if !ok {
			log.Fatal("aborted: failed to receive messages")
			break
		}
		fmt.Println("message received:", msg)
	}
}

func main() {
	fmt.Println("gg starts")

	ch := make(chan string)

	command := "journalctl"
	args := "-f"

	go send(ch, command, args)
	go receive(ch)

	var input string
	fmt.Scanln(&input)
}
