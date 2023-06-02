package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/kaikeventura/expense-management/src/repository/entity"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	database *gorm.DB
}

func ConstructExpenseRepository(database *gorm.DB) ExpenseRepository {
	return ExpenseRepository{
		database: database,
	}
}

func (repository ExpenseRepository) SaveExpense(expense entity.Expense) (entity.Expense, error) {
	var existingExpense entity.Expense
	if err := repository.database.Where("reference_month = ? AND user_id = ?", expense.ReferenceMonth, expense.UserId).First(&existingExpense).Error; err == nil {
		return entity.Expense{},
			errors.New(fmt.Sprintf("Reference month %s already exists for user_id %d", expense.ReferenceMonth, expense.UserId))
	}

	err := repository.database.Create(&expense).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.Expense{}, err
	}

	return expense, nil
}
