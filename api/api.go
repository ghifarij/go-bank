package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ghifarij/go-bank/helpers"
	"github.com/ghifarij/go-bank/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on Port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
