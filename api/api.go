package api

import (
	"encoding/json"
	"fmt"
	"gobank-backend/helpers"
	"gobank-backend/users"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	// ready body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	// handle login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	// prepare response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// ready body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	// handle registration
	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	// prepare response
	if register["message"] == "all is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/plain")
	// 	w.Write([]byte("Hello World"))
	// })

	router.Post("/login", login)
	router.Post("/register", register)
	log.Fatal(http.ListenAndServe(":3000", router))
	fmt.Println("App is working on port :3000")
}
