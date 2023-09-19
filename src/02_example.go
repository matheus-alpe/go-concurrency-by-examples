package examples

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {	
	free: for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("server is doing a graceful shutdown")
			break free
		}
		// block here until someone is sending a message to the channel
	}

}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Yoe",
		Payload: payload,
	}

	msgch <- msg
}

func gracefulQuitServer(quitch chan struct{}) {
	close(quitch)
}

func ExecuteExample02() {
	fmt.Println("Exec 02")

	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	for i := 10; i > 0; i-- {
		i := i
		go func() {
			time.Sleep(time.Duration(i) * time.Second)
			sendMessageToServer(s.msgch, fmt.Sprintf("Hello from [%d]", i))
			if i == 5 {
				gracefulQuitServer(s.quitch)
			}
		}()
	}

	select {}
}
