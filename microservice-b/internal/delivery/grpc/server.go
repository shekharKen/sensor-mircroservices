package grpc

import (
	"context"
	"log"
	"net"
	"time"

	"microservice-b/internal/domain"
	"microservice-b/internal/usecase"
	pb "microservice-b/proto"

	"google.golang.org/grpc"
)

type SensorServer struct {
	pb.UnimplementedSensorServiceServer
	usecase *usecase.SensorUsecase
}

func NewSensorServer(uc *usecase.SensorUsecase) *SensorServer {
	return &SensorServer{usecase: uc}
}

func (s *SensorServer) SendSensorData(ctx context.Context, req *pb.SensorData) (*pb.Ack, error) {
	t, _ := time.Parse(time.RFC3339, req.Timestamp)
	data := domain.SensorData{
		SensorValue: req.SensorValue,
		SensorType:  req.SensorType,
		ID1:         req.ID1,
		ID2:         int(req.ID2),
		Timestamp:   t,
	}

	if err := s.usecase.Save(data); err != nil {
		return &pb.Ack{Message: "Failed"}, err
	}

	log.Printf("Received: %+v\n", data)
	return &pb.Ack{Message: "Success"}, nil
}

func StartGRPCServer(uc *usecase.SensorUsecase) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSensorServiceServer(s, NewSensorServer(uc))

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
