package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


type PingResponse struct {
	Message string `json:"message"`
}


func PingHandler(w http.ResponseWriter, r *http.Request){
	message := PingResponse{
		Message: "PONG",
	}

	w.Header().Set("COntent-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(message)
}

func main(){

	if  err := godotenv.Load(); err != nil {
		fmt.Printf("Error %d", err)
	}


	http.HandleFunc("/ping", PingHandler)
   port := os.Getenv("PORT")

	defer http.ListenAndServe(port, nil)
	fmt.Printf("Server Running %s", port)
}