package main

import (
	"fmt"
	//"log"
	"io"
	"net"

	gRPC "github.com/PrinceMaren1/Homework03-Tr-lsstemning/proto"
	"google.golang.org/grpc"
)


type Server struct{
	gRPC.UnimplementedServerConnectionServer

	clientStreams map[int64] gRPC.ServerConnection_SendMessagesServer

}

func main(){

	fmt.Println("Starting server on port 9000")

	launchServer()

}

func launchServer() {
	
	list, err := net.Listen("tcp", ":9000")
	if err != nil{
		fmt.Println("Failed to listen on port 9000: %v", err)
		return;
	}

	grpcServer := grpc.NewServer()
	
	if err := grpcServer.Serve(list); err != nil{
		fmt.Println("Failed to serve gRPC server over port 9000 %v", err)
	}
}

func (s *Server) SendMessages(msgStream gRPC.ServerConnection_SendMessagesServer) error {

	var id int64;

	for {
		
		msg, err := msgStream.Recv()
		
		fmt.Println("Server recived message from %s: %s", msg.ClientId, msg.Message)
		if err == io.EOF{
			break
		}
		if err != nil {
			return err
		}

		id = msg.ClientId
		clientMessage := msg.Message
		s.clientStreams [id] = msgStream

		fmt.Println("Server recived message from %s: %s", msg.ClientId, msg.Message)

		for key := range s.clientStreams {
			if key != id{
				broadcast := &gRPC.ServerBroadcast{Message: clientMessage, Time: "test"}
				s.clientStreams[key].(gRPC.ServerConnection_SendMessagesServer).Send(broadcast)
			}
		}

	}

	delete(s.clientStreams,id)
	return nil
}