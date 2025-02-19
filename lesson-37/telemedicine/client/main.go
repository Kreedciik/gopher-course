package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "telemedicine/proto"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTelemedicineServiceClient(conn)
	stream, err := client.Consult(context.Background())
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("error receiving from stream: %v", err)
			}
			fmt.Printf("From: %s\nMessage: %s\nTimestamp: %s\n", res.GetUsername(), res.GetMessage(), res.GetTimestamp())
			fmt.Println("Enter message to doctor: ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter message to doctor: ")
	for scanner.Scan() {
		msg := scanner.Text()
		if err := stream.Send(&pb.ConsultationMessage{
			Username:  "Ilfat",
			Message:   msg,
			Timestamp: time.Now().String(),
		}); err != nil {
			log.Fatalf("error sending message to stream: %v", err)
		}
	}

}
