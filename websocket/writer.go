package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

func (s *Socket) SendMessage(msg []byte) {
	s.out <- msg
}

func (s *Socket) makeWriter() {
	ws := s.ws
	out := s.out
	go func () {
		for {
			msg, ok := <-out
			if ! ok {
				// stop writer
				return
			}
			err := ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Writer error:", err)
				return
			}
		}
	} ()
}

