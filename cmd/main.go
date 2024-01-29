package main

import (
	"context"
	"fmt"
	"geoip-service/geoip"
	pb "geoip-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct {
	pb.UnimplementedIPServiceServer
}

func (s *server) GetInfo(ctx context.Context, request *pb.IPRequest) (*pb.IPInfoResponse, error) {
	ip := request.Address
	record := geoip.GetCityRecord(ip)
	return &pb.IPInfoResponse{
		CityName:       record.City.Names["en"],
		CountryName:    record.Country.Names["en"],
		CountryIsoCode: record.Country.IsoCode,
		TimeZone:       record.Location.TimeZone,
		Coordinates:    fmt.Sprintf("%v, %v", record.Location.Latitude, record.Location.Longitude),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterIPServiceServer(s, &server{})
	log.Println("Serving gRPC on connection ")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	conn, err := grpc.Dial("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()
	mux := runtime.NewServeMux()
	err = pb.RegisterIPServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    "localhost:5052",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
}
