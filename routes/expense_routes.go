package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FatemehTayebiSalar/expense-tracker-gateway/models"
)

// AddExpenseRoute
func AddExpenseRoute(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO : Service Layer

	response := map[string]interface{}{
		"message": "Expense created succesfully",
		"expense": expense,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
