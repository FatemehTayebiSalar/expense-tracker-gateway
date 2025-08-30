package main

import (
	"log"
	"net/http"

	"github.com/FatemehTayebiSalar/expense-tracker-gateway/routes"
)

func main() {
	http.HandleFunc("/expenses", routes.AddExpenseRoute)
	log.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
