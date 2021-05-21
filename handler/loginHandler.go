package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	"google.golang.org/grpc"

	pb "github.com/aadesh-agarwal8888/User-Service/proto"
)

func Login(res http.ResponseWriter, req *http.Request) {

	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(configuration.User_Service, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewUserServiceClient(conn)
	var loginDetails pb.User

	json.NewDecoder(req.Body).Decode(&loginDetails)

	response, err := client.LoginUser(context.TODO(), &loginDetails)
	if err != nil {
		log.Println(err)
	}

	log.Println(response)

	if response.Valid {
		res.WriteHeader(http.StatusAccepted)
	} else if !response.Valid {
		res.WriteHeader(http.StatusUnauthorized)
	}
}
