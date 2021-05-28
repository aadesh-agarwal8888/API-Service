package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	pb "github.com/aadesh-agarwal8888/API-Service/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func GetFreeSlot(res http.ResponseWriter, req *http.Request) {
	client, err := ConnectToSlotService()
	if err != nil {
		log.Println(err)
		return
	}

	parkingId := mux.Vars(req)["id"]

	freeSlotRequest := &pb.FreeSlotRequest{
		Id: parkingId,
	}

	freeSlotResponse, err := client.GetFreeSlot(context.TODO(), freeSlotRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusOK)
		return
	}

	json.NewEncoder(res).Encode(freeSlotResponse)

}

func ConnectToSlotService() (pb.SlotServiceClient, error) {
	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn, err := grpc.Dial(configuration.Slot_Service, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := pb.NewSlotServiceClient(conn)

	return client, nil
}
