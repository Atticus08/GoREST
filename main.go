package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Atticus08/GoREST/middleware"
	"github.com/justinas/alice"
)

// User - User object
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	requestHandler := http.HandlerFunc(processRequest)
	/*
		Middlware chaining without Alice
	*/
	// http.Handle("/user",
	// 	middleware.VerifyContentType(
	// 		middleware.SetServerTime(requestHandler)))

	/*
		Middleware chaining with Alice
	*/
	middleChain := alice.New(middleware.VerifyContentType, middleware.SetServerTime).Then(requestHandler)
	http.Handle("/user", middleChain)
	http.ListenAndServe(":8080", nil)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	var user User
	switch restMethod := r.Method; restMethod {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		error := decoder.Decode(&user)
		if error != nil {
			panic(error)
		}
		defer r.Body.Close()
		log.Printf("Here's the user I see from request: %s %s %d\n",
			user.Name, user.Username, user.Age)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - User Created"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}
