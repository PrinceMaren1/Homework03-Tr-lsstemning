package main

import (
	"context"
	"fmt"

	gRPC "github.com/PrinceMaren1/Homework03-Tr-lsstemning/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var server gRPC.ServerConnectionClient
var ServerConn *grpc.ClientConn

func main(){
	
	fmt.Println("New client")
	fmt.Println("Joining server")
	
	ConnectToServer()
	
	sendMessage()
}


func ConnectToServer(){

	opts := []grpc.DialOption {
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Println("Client: %s: Attemps to dial on port 9000")

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", opts...)
	if err != nil {
		fmt.Println("Failed to Dial : %v", err)
		return
	}

	server = gRPC.NewServerConnectionClient(conn)
	ServerConn = conn
	fmt.Println("The connection is: ", conn.GetState().String()) 

}

func sendMessage(){
	stream, err := server.SendMessages(context.Background())
	if err != nil{
		fmt.Println(err)
		return
	}
	stream.Send(&gRPC.ClientMessage{ClientId: 1, Message: "Første test"})
	for{
		
	}
}