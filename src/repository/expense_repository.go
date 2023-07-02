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

func (repository ExpenseRepository) FindExpenseByUserIdAndSequenceNumber(userId uint8, sequenceNumber uint16) (entity.Expense, error) {
	var expense entity.Expense
	if err := repository.database.First(&expense, "user_id = ? and sequence_number = ?", userId, sequenceNumber).Error; err != nil {
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

func (repository ExpenseRepository) FindLastExpenseSequenceNumberByUserId(userId uint8) (uint16, error) {
	var expense entity.Expense
	if err := repository.database.Order("created_at DESC").Last(&expense, "user_id = ?", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("Expense does not exists: " + err.Error())
			return 0, nil
		} else {
			fmt.Println("Error occurred:", err)
			return 0, err
		}
	}

	return expense.SequenceNumber, nil
}

func (repository ExpenseRepository) FindExpenseByUserIdAndReferenceMonth(userId uint8, referenceMonth string) (entity.Expense, error) {
	var expense entity.Expense
	if err := repository.database.First(&expense, "user_id = ? and reference_month = ?", userId, referenceMonth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("Expense does not exists: " + err.Error())
			return entity.Expense{}, nil
		} else {
			fmt.Println("Error occurred:", err)
			return entity.Expense{}, err
		}
	}

	return expense, nil
}

func (repository ExpenseRepository) FindExpenseByUserIdAndState(userId uint8, state string) (entity.Expense, error) {
	var expense entity.Expense
	if err := repository.database.Preload("FixedExpenses").Preload("Purchases").Preload("CreditCardPurchases").First(&expense, "expenses.user_id = ? and expenses.state = ?", userId, state).Error; err != nil {
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

func (repository ExpenseRepository) SaveCreditCardPurchase(creditCardPurchase entity.CreditCardPurchase) (entity.CreditCardPurchase, error) {
	err := repository.database.Create(&creditCardPurchase).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return entity.CreditCardPurchase{}, err
	}

	return creditCardPurchase, nil
}
