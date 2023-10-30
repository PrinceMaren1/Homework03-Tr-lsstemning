package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

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

	// File logging adapted from example at https://stackoverflow.com/questions/40443881/how-to-write-log-into-log-files-in-golang
	logFile := "log_" + *id
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

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
	waitc := make(chan struct{})

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				fmt.Println(err)
			}
			updateTime(msg.Time)
			log.Printf("Time: %v. Message received: %v\n", msg.Time, msg.Message)
			fmt.Printf("Time: %v. Message received: %v\n", msg.Time, msg.Message)
		}
	}()

	for {
		var input string
		fmt.Scan(&input)
		if input == "Disconnect" {
			fmt.Print("Shutting down \n")
			close(waitc)
			break
		} else if len([]rune(input)) > 128 {
			fmt.Println("Max length of message is 128 characters")
		} else {
			sendMessage(stream, input)
		}
	}

	stream.CloseSend()
	<-waitc
	return nil
}

func sendMessage(stream gRPC.ServerConnection_SendMessagesClient, message string) {
	time = time + 1
	log.Printf("Sending message to server with time %v", time)
	stream.Send(&gRPC.ClientMessage{ClientId: *id, Message: message, Time: time})
}

func updateTime(receivedTime int64) {
	time = max(receivedTime, time) + 1
}
