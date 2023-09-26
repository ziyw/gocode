package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ziyw/gocode/go_grpc/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	c := pb.NewNoteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetNote(ctx, &pb.NoteRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get note: %v", err)
	}

	log.Printf("Response :%s\n", r.GetContent())

}
