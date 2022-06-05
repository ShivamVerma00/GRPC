package wordcountmain

import (
	"context"
	pb "grpc/protobuf"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterWordServer(s, &WordServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server Failed...: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestWordCount(t *testing.T) {

	//Dial a connection to grpc Server
	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer connection.Close()

	//Create new Client
	c := pb.NewWordClient(connection)

	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer connection.Close()

	//Result
	response, err := c.Word_Count(ctx, &pb.Request{Text: "Demo text for testing grpc "})
	if err != nil {
		t.Fatal("Could not count word: \n", err)
	}
	t.Log("WordCount:\n")

	for index, value := range response.Word_Count_ {
		t.Logf("Word: %s \t	Count: %d", value.Word, value.Count)
		if index == 9 {
			break
		}
	}
}
