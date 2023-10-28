package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	gRPC "github.com/PrinceMaren1/Homework03-Tr-lsstemning/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var time int64 = 0
var server gRPC.ServerConnectionClient
var ServerConn *grpc.ClientConn
var id = flag.String("id", "", "client name")

func main() {
	flag.Parse()
	fmt.Println("New client")
	fmt.Println("Joining server")


	ConnectToServer()
	defer ServerConn.Close()
	sendMessages()
}

func ConnectToServer() {

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Printf("Client %v: Attemps to dial on port 9000\n", *id)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", opts...)
	if err != nil {
		fmt.Printf("Failed to Dial : %v\n", err)
		return
	}

	server = gRPC.NewServerConnectionClient(conn)
	ServerConn = conn
	fmt.Println("The connection is: ", conn.GetState().String())
}

func sendMessages() error {
	stream, _ := server.SendMessages(context.Background())
	// ...
	sendMessage(stream, "EstablishConnection")

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				// stream end from server
				// should we stop execution of sendMessage() here?
				fmt.Print("Shuting down")
				return
			}

			if err != nil {
				fmt.Println(err)
			}
			updateTime(msg.Time)
			fmt.Printf("Time: %v. Message: %v\n", msg.Time, msg.Message)
		}
	}()

	for {
		var input string
		fmt.Scan(&input)
		if(input == "Disconnect"){
			stream.CloseSend()
			break
		}else{
			sendMessage(stream, input)
		}
	}
	return nil
}

func sendMessage(stream gRPC.ServerConnection_SendMessagesClient, message string) {
	time = time + 1
	stream.Send(&gRPC.ClientMessage{ClientId: *id, Message: message})
}

func updateTime(receivedTime int64) {
	time = max(receivedTime, time) + 1
}
