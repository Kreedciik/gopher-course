package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "patient-health/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewHealthServiceClient(conn)
	req := &pb.PatientRequest{PatientId: "1233"}
	stream, err := client.GetPatientVitals(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get patient: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not receive patient data: %v", err)
			break
		}
		fmt.Printf("Patient %s\nBlood pressure = %v\nOxygen level = %v\nTimestamp=%v\n\n", res.GetPatientName(), res.GetHeartRate(), res.GetOxygenLevel(), res.GetTimestamp())
	}
}
