package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	pb "github.com/aadesh-agarwal8888/API-Service/proto"
	"google.golang.org/grpc"
)

func GetParkingAreas(res http.ResponseWriter, req *http.Request) {

	client, err := ConnectToParkingService()
	if err != nil {
		log.Println(err)
		http.Error(res, "Cannot Connect to Parking Service", http.StatusInternalServerError)
		return
	}

	response, err := client.GetParkingAreas(context.Background(), &pb.Request{})
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(response)
}

func ConnectToParkingService() (pb.ParkingServiceClient, error) {
	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn, err := grpc.Dial(configuration.Parking_Service, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := pb.NewParkingServiceClient(conn)

	return client, nil
}
