package main

import (
	"BookSpaceServer/auth"
	"BookSpaceServer/booking"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 路由分別對應於不同模塊
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/rooms", booking.GetMeetingRooms).Methods("GET")
	r.HandleFunc("/rooms/{id}/reserve", booking.ReserveRoom).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
