package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	"google.golang.org/grpc"

	pb "github.com/aadesh-agarwal8888/API-Service/proto"
)

func Login(res http.ResponseWriter, req *http.Request) {
	client, err := ConnecctToUserService()
	if err != nil {
		log.Println(err)
		http.Error(res, "Server Error", http.StatusNotFound)
		return
	}

	var loginDetails pb.LoginDetails

	json.NewDecoder(req.Body).Decode(&loginDetails)

	response, err := client.LoginUser(context.TODO(), &loginDetails)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	json.NewEncoder(res).Encode(response)
}

func GetUserDetails(res http.ResponseWriter, req *http.Request) {

	client, err := ConnecctToUserService()
	if err != nil {
		log.Println(err)
		http.Error(res, "Server Error", http.StatusNotFound)
		return
	}

	id := req.URL.Query().Get("id")

	userDetails := &pb.UserDetails{
		Id: id,
	}

	userData, err := client.GetUserData(context.Background(), userDetails)
	if err != nil {
		http.Error(res, "No User Found", http.StatusOK)
		return
	}

	json.NewEncoder(res).Encode(userData)
}

func RegisterUser(res http.ResponseWriter, req *http.Request) {

	var userRegistrationData *pb.UserRegistrationData
	json.NewDecoder(req.Body).Decode(&userRegistrationData)

	client, err := ConnecctToUserService()
	if err != nil {
		log.Println(err)
		http.Error(res, "Server Error", http.StatusNotFound)
		return
	}

	registrationResponse, err := client.RegisterUser(context.TODO(), userRegistrationData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(res).Encode(registrationResponse)

}

func ConnecctToUserService() (pb.UserServiceClient, error) {
	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn, err := grpc.Dial(configuration.User_Service, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)

	return client, nil
}
