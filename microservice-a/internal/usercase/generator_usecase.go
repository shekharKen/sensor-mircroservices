package usecase

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "microservice-a/proto" // your generated proto package
	"microservice-a/internal/data"
)

type GeneratorUsecase struct {
	SensorType string
	Frequency  time.Duration
	stopChan   chan bool

	// new optional gRPC client (nil if not connected)
	client pb.SensorServiceClient
}

func NewGeneratorUsecase(sensorType string, freq time.Duration) *GeneratorUsecase {
	return &GeneratorUsecase{
		SensorType: sensorType,
		Frequency:  freq,
		stopChan:   make(chan bool),
	}
}

func (g *GeneratorUsecase) ConnectGRPC(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	g.client = pb.NewSensorServiceClient(conn)
	log.Printf("Connected to Microservice B at %s", addr)
	return nil
}

func (g *GeneratorUsecase) Start() {
	ticker := time.NewTicker(g.Frequency)
	go func() {
		for {
			select {
			case <-ticker.C:
				sensorData := data.RandomSensorData(g.SensorType)

				log.Printf("Generated: %+v\n", sensorData)

				if g.client != nil {
					ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					defer cancel()
					_, err := g.client.SendSensorData(ctx, &pb.SensorData{
						SensorValue: sensorData.SensorValue,
						SensorType:  sensorData.SensorType,
						ID1:         sensorData.ID1,
						ID2:         int32(sensorData.ID2),
						Timestamp:   sensorData.Timestamp.Format(time.RFC3339),
					})
					if err != nil {
						log.Printf("gRPC send failed: %v", err)
					} else {
						log.Println("Data sent to Microservice B via gRPC")
					}
				}
			case <-g.stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}

func (g *GeneratorUsecase) UpdateFrequency(freq time.Duration) {
	g.stopChan <- true
	g.Frequency = freq
	g.stopChan = make(chan bool)
	g.Start()
}
