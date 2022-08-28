package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/sign-up", controllers.SignUp).Methods("POST")
	router.HandleFunc("/sign-in", controllers.SignIn).Methods("POST")
	router.HandleFunc("/sign-out", controllers.SignOut).Methods("POST")
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
		})
}
