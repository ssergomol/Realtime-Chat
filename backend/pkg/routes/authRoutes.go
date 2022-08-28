package routes

import (
	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/sign-up", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/sign-in", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/sign-out", controllers.SignOut).Methods("POST", "OPTIONS")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET", "OPTIONS")
}
