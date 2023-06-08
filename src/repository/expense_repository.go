package repository

import (
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
			fmt.Errorf("reference month %s already exists for user_id %d", expense.ReferenceMonth, expense.UserId)
	}

	err := repository.database.Create(&expense).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.Expense{}, err
	}

	return expense, nil
}

func (repository ExpenseRepository) FindExpenseById(expenseId uint16) (entity.Expense, error) {
	var expense entity.Expense
	if err := repository.database.First(&expense, expenseId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("Expense does not exists: " + err.Error())
			return entity.Expense{}, err
		} else {
			fmt.Println("Error occurred:", err)
			return entity.Expense{}, err
		}
	}

	return expense, nil
}

func (repository ExpenseRepository) UpdateExpenseTotalAmount(expenseId uint16, newTotalAmount int32) error {
	var expense entity.Expense
	return repository.database.Model(&expense).Where("id", expenseId).UpdateColumn("total_amount", newTotalAmount).Error
}

func (repository ExpenseRepository) SaveFixedExpense(fixedExpense entity.FixedExpense) (entity.FixedExpense, error) {
	err := repository.database.Create(&fixedExpense).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.FixedExpense{}, err
	}

	return fixedExpense, nil
}

func (repository ExpenseRepository) SavePurchase(purchase entity.Purchase) (entity.Purchase, error) {
	err := repository.database.Create(&purchase).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.Purchase{}, err
	}

	return purchase, nil
}
