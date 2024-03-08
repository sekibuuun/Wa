package routes

import (
	"WanderVive_back/pkg/handlers"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/band", handlers.BandHandler)
	http.HandleFunc("/event", handlers.EventHandler)
	http.HandleFunc("/livehouse", handlers.LivehouseHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
