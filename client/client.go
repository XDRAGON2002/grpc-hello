package main

import (
	"context"
	"log"

	"example.com/proj/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could Not Connect: %s", err.Error())
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	resp, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From The Client!"})
	if err != nil {
		log.Fatalf("Could Not Say Hello: %s", err.Error())
	}
	log.Printf("Response From Server: %s", resp.Body)
}