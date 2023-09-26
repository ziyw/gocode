package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ziyw/gocode/go_grpc/note"
	"google.golang.org/grpc"
)

type note struct {
	ID      int32
	Title   string
	Content string
}

var notes = map[int32]note{
	1: note{ID: 1, Title: "FirstEntry", Content: "First note"},
	2: note{ID: 2, Title: "Second", Content: "Second Content"},
}

type server struct {
	pb.UnimplementedNoteServiceServer
}

func (s *server) GetNote(ctx context.Context, in *pb.NoteRequest) (*pb.NoteResponse, error) {
	n := notes[in.GetId()]
	return &pb.NoteResponse{Title: n.Title, Content: n.Content}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen to: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNoteServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
