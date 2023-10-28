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

var time int = 0

func main() {
	fmt.Println("Starting server on port 9000")
	launchServer()
}

func launchServer() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("Failed to listen on port 9000: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()

	server := &Server{
		clientStreams: make(map[string]gRPC.ServerConnection_SendMessagesServer),
	}

	gRPC.RegisterServerConnectionServer(grpcServer, server)

	if err := grpcServer.Serve(list); err != nil {
		fmt.Printf("Failed to serve gRPC server over port 9000 %v\n", err)
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

		if msg.Message == "EstablishConnection" {
			welcomeMsg := fmt.Sprintf("Participant %v joined Chitty-Chat at Lamport time %v", msg.ClientId, "L")
			for key := range s.clientStreams {
				broadcast := &gRPC.ServerBroadcast{Message: welcomeMsg, Time: "test"}
				s.clientStreams[key].Send(broadcast)
			}
			continue
		}

		fmt.Printf("Server recived message from client %v: %v\n", msg.ClientId, msg.Message)

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

func updateTime(receivedTime int) {
	time = max(receivedTime, time) + 1
}
