package routes

import (
	"fmt"
	"net/http"
)

// AddExpenseRoute
func AddExpenseRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Hello from Post /expenses")
}
