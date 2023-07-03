package chat

import (
	"log"

	context "golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received Message Body From Client: %s", message.Body)
	return &Message{Body: "Hello From The Server!"}, nil
}
