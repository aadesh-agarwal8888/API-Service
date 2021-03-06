package main

import (
	"log"
	"net/http"

	"github.com/aadesh-agarwal8888/API-Service/configs"
	"github.com/aadesh-agarwal8888/API-Service/handler"
	"github.com/gorilla/mux"
)

func main() {
	configuration, err := configs.LoadConfigurations()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	//User handler
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)           //Login => Post
	router.HandleFunc("/register", handler.RegisterUser).Methods(http.MethodPost) //Register => Post
	router.HandleFunc("/user", handler.GetUserDetails).Methods(http.MethodGet)    //Get User Data => GET

	//ParkingArea handlers
	router.HandleFunc("/home/park", handler.GetParkingAreas).Methods(http.MethodGet) //Fetch Parking Areas => GET

	//Booking History
	router.HandleFunc("/home/park/booking", handler.GetBooking).Methods(http.MethodGet)      //Fetch Last Booking => GET
	router.HandleFunc("/home/park/booking/all", handler.GetBookings).Methods(http.MethodGet) //Fetch all Bookings => GET

	//Payment
	router.HandleFunc("/home/park/payment", handler.GetPayment).Methods(http.MethodGet)       //get payment info => GET
	router.HandleFunc("/home/park/payment/{id}", handler.MakePayment).Methods(http.MethodPut) //Make Payment => Put

	//Booking Slot
	router.HandleFunc("/home/park/{id}", handler.GetFreeSlot).Methods(http.MethodGet) //Get Free Slot => GET
	router.HandleFunc("/home/park/{id}", handler.BookSlot).Methods(http.MethodPost)   //Book the Slot => POST

	server := http.Server{
		Addr:    configuration.Api_Service,
		Handler: router,
	}

	log.Println("Starting API-Service on " + configuration.Api_Service)
	log.Fatal(server.ListenAndServe())

}
