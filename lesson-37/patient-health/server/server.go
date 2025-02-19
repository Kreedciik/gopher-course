package main

import (
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	pb "patient-health/proto"
	"time"
)

type HealthServer struct {
	pb.UnimplementedHealthServiceServer
}

var patients = map[string]string{
	"1233": "Adam",
	"1244": "Sarah",
}

func generateData(id string) *pb.PatientVitalsResponse {
	heartRate := rand.Intn(100)
	oxygenLevel := rand.Intn(100)
	return &pb.PatientVitalsResponse{
		HeartRate:   int32(heartRate),
		OxygenLevel: int32(oxygenLevel),
		Timestamp:   time.Now().String(),
		PatientName: patients[id],
	}
}
func (h *HealthServer) GetPatientVitals(req *pb.PatientRequest, stream pb.HealthService_GetPatientVitalsServer) error {
	id := req.GetPatientId()
	for i := 0; i < 10; i++ {
		res := generateData(id)
		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to send response: %v", err)
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, &HealthServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
