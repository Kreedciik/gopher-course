package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "heart-rate/proto"
	"log"
	"math/rand"
	"time"
)

func generateHeartRate() int64 {
	randomHeartRate := int64(rand.Intn(100))
	return randomHeartRate
}

func main() {
	fmt.Println("Starting client...")
	conn, err := grpc.NewClient("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHealthServiceClient(conn)
	stream, err := client.MonitorHeartRate(context.Background())
	if err != nil {
		log.Fatalf("MonitorHeartRate = _, %v", err)
	}

	for i := 0; i < 100; i++ {
		heartRate := generateHeartRate()
		err := stream.Send(&pb.HeartRateRequest{HeartRate: heartRate})
		if err != nil {
			log.Fatalf("%v.MonitorHeartRate(_) = _, %v", client, err)
		}
		time.Sleep(2 * time.Second)
	}

}
