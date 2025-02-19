package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "heart-rate/proto"
	"io"
	"log"
	"net"
)

type HeartRateServer struct {
	pb.UnimplementedHealthServiceServer
}

func (h *HeartRateServer) MonitorHeartRate(stream pb.HealthService_MonitorHeartRateServer) error {
	messageCount := 0
	var sum int64 = 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed.")
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive a message from client: %v", err)
		}
		messageCount++
		if messageCount%10 == 0 {
			fmt.Printf("AverageHeartRate: %.2f\n", float32((sum)/10))
			sum = 0
		} else {
			sum += msg.GetHeartRate()
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, &HeartRateServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
