package wordcountmain

import (
	"context"
	pb "grpc/protobuf"
	"log"
	"sort"
	"strings"
)

type WordServer struct {
	pb.UnimplementedWordServer //Service
}

//function
func (server *WordServer) Word_Count(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Print("Working... ")
	str := in.Text //input string

	count := make(map[string]int) //map

	for _, word := range strings.Fields(str) {
		count[word]++ // counting  words
	}

	words := make([]string, 0, len(count))
	for i := range count {
		words = append(words, i) //appending words and their frequency
	}

	var WordCPb []*pb.Word_Count // slice of Word_Count type
	for k, v := range count {
		WordCPb = append(WordCPb, &pb.Word_Count{Word: k, Count: uint64(v)}) //appending words and their frequency in slice
	}

	sort.Slice(WordCPb, func(i, j int) bool {
		return WordCPb[i].Count > WordCPb[j].Count //sorting in Non-Increasing order
	})

	return &pb.Response{Word_Count_: WordCPb}, nil
}
