package main

import (
	"encoding/json"
	"net/http"
	"strings"
)


type CreateUserRequest struct {
	Name string `json:"name"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	users []User
	lastId int
)

func WriteJSON(w http.ResponseWriter, status int, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string){
	WriteJSON(w, status, map[string]string{"error": msg})
}


func GetUserhandler(w http.ResponseWriter, r *http.Request){
	if r.Method  != http.MethodGet{
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	
	WriteJSON(w, http.StatusOK, users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method  != http.MethodPost{
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}

	lastId++
	user := User{
		ID: lastId,
		Name: req.Name,
	}
	users = append(users, user)

	WriteJSON(w, http.StatusCreated, user)
}

func main(){
	http.HandleFunc("/users", CreateUserHandler)
	http.HandleFunc("/userget", GetUserhandler)
	http.ListenAndServe(":8080", nil)
}