package main

import (
	"net/http"
	"strconv"

	"github.com/ahsanulks/testingrpc/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/add/{x}/{y}", AddHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x, _ := strconv.ParseInt(vars["x"], 10, 64)
	y, _ := strconv.ParseInt(vars["y"], 10, 64)

	req := proto.Request{X: x, Y: y}
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
