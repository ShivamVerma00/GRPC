package main

import (
	"bufio"
	"context"
	"fmt"
	pb "grpc/protobuf"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const address = "localhost:61234"

func main() {
	fmt.Println("Input String :")

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid...")
	}

	//connecting to grpc server
	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("Dial Failed...")
	}

	defer connection.Close()

	//Creating client
	cl := pb.NewWordClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Result
	response, err := cl.Word_Count(ctx, &pb.Request{Text: s})
	if err != nil {
		log.Fatal("Error: \n", err)
	}

	//log.Println("WordCount:")
	for index, value := range response.Word_Count_ {
		log.Printf("Word: %s \t	Count: %d", value.Word, value.Count)
		if index == 9 {
			break
		}
	}
}
