package services

import (
	"errors"

	"github.com/FatemehTayebiSalar/expense-tracker-gateway/models"
)

// ValidateExpense checks if the given expense is valid
func ValidateExpense(expense models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if expense.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

// AddExpense handles creating an expense(later with DB  or API call)
// Right now, it just validates and returns the expense
func AddExpense(expense models.Expense) (models.Expense, error) {
	if err := ValidateExpense(expense); err != nil {
		return models.Expense{}, err
	}

	//TODO: save expense to DB
	return expense, nil
}
