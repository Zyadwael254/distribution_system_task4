package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// هيكل الرسالة
type Message struct {
	Name string
	Text string
}

// هيكل السيرفر
type ChatServer struct {
	messages []Message
	mu       sync.Mutex
}

// السيرفر يستقبل الرسائل فقط
func (cs *ChatServer) SendMessage(msg Message, reply *string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	// حفظ الرسالة في القائمة
	cs.messages = append(cs.messages, msg)

	// عرض الرسالة على شاشة السيرفر
	fmt.Printf("%s: %s\n", msg.Name, msg.Text)

	*reply = "Message received"
	return nil
}

func main() {
	chatServer := new(ChatServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
