package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name string
	Text string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You can start typing your messages.\n", name)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		msg := Message{Name: name, Text: text}
		var reply string

		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}
