package routes

import (
	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/sing-up", controllers.SignUp).Methods("POST")
	router.HandleFunc("/sign-in", controllers.SignIn).Methods("POST")
	router.HandleFunc("/sign-out", controllers.SignOut).Methods("POST")
}
