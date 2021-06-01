package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	pb "github.com/aadesh-agarwal8888/API-Service/proto"
	"github.com/aadesh-agarwal8888/API-Service/types"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func ConnectToPaymentService() (pb.PaymentServiceClient, error) {

	configuration, err := configs.LoadConfigurations()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(configuration.Payment_Service, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewPaymentServiceClient(conn)

	return client, nil
}

func GetPayment(res http.ResponseWriter, req *http.Request) {
	var user *types.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, "Id not found", http.StatusBadGateway)
		return
	}

	paymentClient, err := ConnectToPaymentService()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		return
	}

	request := &pb.GetPaymentRequest{
		Id: user.ID,
	}

	response, err := paymentClient.GetPayment(context.TODO(), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(response)
}

func MakePayment(res http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]

	paymentClient, err := ConnectToPaymentService()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		return
	}

	request := &pb.MakePaymentRequest{
		Id: id,
	}

	response, err := paymentClient.MakePayment(context.TODO(), request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(response)
}
