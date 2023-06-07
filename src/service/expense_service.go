package service

import (
	"fmt"
	"time"

	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/repository"
	"github.com/kaikeventura/expense-management/src/repository/entity"
)

type ExpenseService struct {
	repository repository.ExpenseRepository
}

func ConstructExpenseService(repository repository.ExpenseRepository) ExpenseService {
	return ExpenseService{
		repository: repository,
	}
}

func (service ExpenseService) CreateExpense(expenseDto dto.Expense) (dto.Expense, error) {
	expenseEntity := entity.Expense{
		UserId:         expenseDto.UserId,
		ReferenceMonth: time.Now().Format(dto.YYYYMM),
		State:          dto.CURRENT.StateToString(),
		TotalAmount:    0,
	}
	createdExpense, err := service.repository.SaveExpense(expenseEntity)

	if err != nil {
		return dto.Expense{}, err
	}

	return dto.Expense{
		Id:                  createdExpense.Id,
		UserId:              createdExpense.UserId,
		ReferenceMonth:      createdExpense.ReferenceMonth,
		State:               createdExpense.State,
		TotalAmount:         createdExpense.TotalAmount,
		FixedExpenses:       make([]dto.FixedExpense, 0),
		Purchases:           make([]dto.Purchase, 0),
		CreditCardPurchases: make([]dto.CreditCardPurchase, 0),
	}, nil
}

func (service ExpenseService) CreateFixedExpense(expenseId uint16, fixedExpenseDto dto.FixedExpense) (dto.FixedExpense, error) {
	expeseExists, err := service.repository.ExpenseExists(expenseId)

	if err != nil {
		return dto.FixedExpense{}, err
	}

	if !expeseExists {
		return dto.FixedExpense{}, fmt.Errorf("expense with id %d does not exists", expenseId)
	}

	fixedExpenseEntity := entity.FixedExpense{
		ExpenseId:   expenseId,
		Category:    fixedExpenseDto.Category,
		Description: fixedExpenseDto.Description,
		Amount:      fixedExpenseDto.Amount,
	}
	createdFixedExpense, err := service.repository.SaveFixedExpense(fixedExpenseEntity)

	if err != nil {
		return dto.FixedExpense{}, err
	}

	return dto.FixedExpense{
		Id:          createdFixedExpense.Id,
		Category:    createdFixedExpense.Category,
		Description: createdFixedExpense.Description,
		Amount:      createdFixedExpense.Amount,
	}, nil
}
