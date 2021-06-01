package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	pb "github.com/aadesh-agarwal8888/API-Service/proto"
	"github.com/aadesh-agarwal8888/API-Service/types"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func BookSlot(res http.ResponseWriter, req *http.Request) {

	var user *types.User
	json.NewDecoder(req.Body).Decode(&user)

	client, err := ConnectToBookingService()
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	slotId := mux.Vars(req)["id"]

	bookingRequest := &pb.BookingRequest{
		SlotId:     slotId,
		CustomerId: user.ID,
	}

	bookingResponse, err := client.BookSlot(context.TODO(), bookingRequest)
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(bookingResponse)

}

func GetBookings(res http.ResponseWriter, req *http.Request) {
	var user *types.User
	json.NewDecoder(req.Body).Decode(&user)

	client, err := ConnectToBookingService()
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	bookingRequest := &pb.GetBookingRequest{
		Id: user.ID,
	}

	response, err := client.GetBookings(context.TODO(), bookingRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(response)
}

func GetBooking(res http.ResponseWriter, req *http.Request) {
	var user *types.User
	json.NewDecoder(req.Body).Decode(&user)

	client, err := ConnectToBookingService()
	if err != nil {
		log.Println(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	bookingRequest := &pb.GetBookingRequest{
		Id: user.ID,
	}

	response, err := client.GetBooking(context.TODO(), bookingRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(response)
}

func ConnectToBookingService() (pb.BookingServiceClient, error) {

	configuration, err := configs.LoadConfigurations()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(configuration.Booking_Service, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewBookingServiceClient(conn)

	return client, nil
}
