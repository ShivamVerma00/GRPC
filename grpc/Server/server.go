package main

import (
	pb "grpc/protobuf"
	"grpc/wordcountmain"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":61234"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed...", err)

	}
	server := grpc.NewServer() //server at grpc

	pb.RegisterWordServer(server, &wordcountmain.WordServer{})

	log.Println("Listening... ", listen.Addr())

	if err := server.Serve(listen); err != nil {
		log.Fatal("Failed server... ", err)
	}

}
