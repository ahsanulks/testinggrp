package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ahsanulks/testingrpc/client/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var clientService proto.AddServiceClient

func main() {
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	clientService = proto.NewAddServiceClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/add", AddHandler)
	r.HandleFunc("/multiple", MultipleHandler).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]int64)

	json.NewDecoder(r.Body).Decode(&body)
	x, y := body["x"], body["y"]

	req := proto.Request{X: x, Y: y}
	response, err := clientService.Add(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"total": response.Result})
}

func MultipleHandler(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]int64)

	json.NewDecoder(r.Body).Decode(&body)
	x, y := body["x"], body["y"]

	req := proto.Request{X: x, Y: y}
	response, err := clientService.Multiply(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"total": response.Result})
}
