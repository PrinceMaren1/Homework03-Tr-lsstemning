package main

import (
	"fmt"
	//"log"
	"io"
	"net"

	gRPC "github.com/PrinceMaren1/Homework03-Tr-lsstemning/proto"
	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedServerConnectionServer
	clientStreams map[string]gRPC.ServerConnection_SendMessagesServer
}

func main() {

	fmt.Println("Starting server on port 9000")

	launchServer()

}

func launchServer() {
	fmt.Println("Start")
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("Failed to listen on port 9000: %v", err)
		return
	}
	fmt.Println("Listen")
	grpcServer := grpc.NewServer()

	server := &Server{
		clientStreams: make(map[string]gRPC.ServerConnection_SendMessagesServer),
	}

	gRPC.RegisterServerConnectionServer(grpcServer, server)

	fmt.Println("Server")
	if err := grpcServer.Serve(list); err != nil {
		fmt.Printf("Failed to serve gRPC server over port 9000 %v", err)
	}

	fmt.Println("Server started")
}

func (s *Server) SendMessages(msgStream gRPC.ServerConnection_SendMessagesServer) error {

	var id string

	for {

		msg, err := msgStream.Recv()

		//fmt.Println("Server recived message from %s: %s", msg.ClientId, msg.Message)
		if err == io.EOF {
			fmt.Print("break")
			break
		}
		if err != nil {
			fmt.Print(err)
			return err
		}
		id = msg.ClientId
		clientMessage := msg.Message
		s.clientStreams[id] = msgStream

		fmt.Printf("Server recived message from client %v: %v", msg.ClientId, msg.Message)

		for key := range s.clientStreams {
			if key != id {
				broadcast := &gRPC.ServerBroadcast{Message: clientMessage, Time: "test"}
				s.clientStreams[key].Send(broadcast)
			}
		}

	}

	delete(s.clientStreams, id)
	return nil
}
