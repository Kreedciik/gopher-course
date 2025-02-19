package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "telemedicine/proto"
	"time"
)

type TelemedicineServer struct {
	pb.UnimplementedTelemedicineServiceServer
}

func (t *TelemedicineServer) Consult(stream grpc.BidiStreamingServer[pb.ConsultationMessage, pb.ConsultationMessage]) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("%v.ConsultationMessage(_) = _, %v", stream, err)
		}
		fmt.Printf("From: %s\nMessage: %s\nTimestamp: %s", msg.GetUsername(), msg.GetMessage(), msg.GetTimestamp())

		response := &pb.ConsultationMessage{
			Username:  "Doctor D",
			Message:   msg.GetMessage(),
			Timestamp: time.Now().String(),
		}

		if err := stream.Send(response); err != nil {
			log.Fatalf("%v.ConsultationMessage(_) = _, %v", stream, err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTelemedicineServiceServer(s, &TelemedicineServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
