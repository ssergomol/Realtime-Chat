package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssergomol/RealtimeChat/pkg/database"
	"github.com/ssergomol/RealtimeChat/pkg/routes"
	"github.com/ssergomol/RealtimeChat/pkg/websocketPool"
)

func main() {
	fmt.Println("Realtime Chat starting...")
	pool := websocketPool.NewPool()
	go pool.Run()

	fmt.Println("Connecting to database...")
	database.Connect()

	router := mux.NewRouter()
	routes.RegisterWsRoute(router, pool)
	routes.RegisterAuthRoutes(router)

	http.Handle("/", router)

	http.ListenAndServe(":9000", nil)
}
