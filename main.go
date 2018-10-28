package main

import (
	"net/http"

	"github.com/serhio83/api-server/controllers"
	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
	// Connect to mongo
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
