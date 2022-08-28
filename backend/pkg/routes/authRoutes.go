package routes

import (
	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/controllers"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/sign-up", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/sign-in", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/sign-out", controllers.SignOut).Methods("POST", "OPTIONS")

	// router.HandleFunc("sign-up", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	w.Header().Set("Access-Control-Allow-Methods", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "*")
	// }).Methods("OPTIONS")

	// router.HandleFunc("sign-in", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	w.Header().Set("Access-Control-Allow-Methods", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "*")
	// 	w.Header().Set("Access-Control-Allow-Credentials", "true")
	// 	w.Header().Set("Access-Control-Max-Age", "86400")
	// }).Methods("OPTIONS")

	// router.Methods("OPTIONS").HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 		w.Header().Set("Access-Control-Allow-Methods", "*")
	// 		w.Header().Set("Access-Control-Allow-Headers", "*")
	// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
	// 		w.Header().Set("Access-Control-Max-Age", "86400")
	// 	})
	// router.HandleFunc("sign-in", func(w http.ResponseWriter, r *http.Request) {

	// }).Methods("OPTIONS")
}
