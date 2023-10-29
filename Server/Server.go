package main

import (
	"fmt"
	"log"
	"os"

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

var time int64 = 0

func main() {
	fmt.Println("Starting server on port 9000")

	// File logging adapted from example at https://stackoverflow.com/questions/40443881/how-to-write-log-into-log-files-in-golang
	f, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

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

		if err == io.EOF {
			// A client disconnected
			// New event (sending messages to clients), so time is incremented
			time = time + 1
			for key := range s.clientStreams {
				if key != id {
					sendMessage(fmt.Sprintf("Participant %v left Chitty-Chat at Lamport time %v", id, time), key, s.clientStreams[key])
				}
			}
			break
		}
		if err != nil {
			// Unknown error
			fmt.Print(err)
			return err
		}
		id = msg.ClientId
		clientMessage := msg.Message
		s.clientStreams[id] = msgStream
		updateTime(msg.Time) // Update lamport time from msg

		// New event (sending messages to clients), so time is incremented
		// Multiple identical messages sent to different clients are viewed as one event, and are therefore send with identical timestamp
		time = time + 1

		// Time event / messages happens
		if msg.Message == "EstablishConnection" {
			// time-1 because we already incremented time for the new message event, but the connection was received beforehand
			welcomeMsg := fmt.Sprintf("Participant %v joined Chitty-Chat at Lamport time %v", msg.ClientId, time-1)
			for key := range s.clientStreams {
				sendMessage(welcomeMsg, key, s.clientStreams[key])
			}
			continue
		}

		for key := range s.clientStreams {
			if key != id {
				sendMessage(clientMessage, key, s.clientStreams[key])
			}
		}

	}

	delete(s.clientStreams, id)
	return nil
}

func sendMessage(message string, clientId string, stream gRPC.ServerConnection_SendMessagesServer) {
	log.Printf("Sending to message to %v at time %v with content: %v", clientId, time, message)
	broadcast := &gRPC.ServerBroadcast{Message: message, Time: time}
	stream.Send(broadcast)
}

func updateTime(receivedTime int64) {
	time = max(receivedTime, time) + 1
}
