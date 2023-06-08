package service

import (
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
	expeseEntity, err := service.getExpenseById(expenseId)

	if err != nil {
		return dto.FixedExpense{}, err
	}

	entity.PlusExpenseTotalAmount(&expeseEntity, fixedExpenseDto.Amount)
	err = service.repository.UpdateExpenseTotalAmount(expeseEntity.Id, expeseEntity.TotalAmount)

	if err != nil {
		return dto.FixedExpense{}, err
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

func (service ExpenseService) CreatePurchase(expenseId uint16, purchase dto.Purchase) (dto.Purchase, error) {
	expeseEntity, err := service.getExpenseById(expenseId)

	if err != nil {
		return dto.Purchase{}, err
	}

	purchaseEntity := entity.Purchase{
		ExpenseId:   expenseId,
		Category:    purchase.Category,
		Description: purchase.Description,
		Amount:      purchase.Amount,
	}
	createdPurchase, err := service.repository.SavePurchase(purchaseEntity)

	if err != nil {
		return dto.Purchase{}, err
	}

	entity.PlusExpenseTotalAmount(&expeseEntity, purchase.Amount)
	err = service.repository.UpdateExpenseTotalAmount(expeseEntity.Id, expeseEntity.TotalAmount)

	if err != nil {
		return dto.Purchase{}, err
	}

	return dto.Purchase{
		Id:          createdPurchase.Id,
		Category:    createdPurchase.Category,
		Description: createdPurchase.Description,
		Amount:      createdPurchase.Amount,
	}, nil
}

func (service ExpenseService) getExpenseById(expenseId uint16) (entity.Expense, error) {
	expese, err := service.repository.FindExpenseById(expenseId)

	if err != nil {
		return entity.Expense{}, err
	}

	return expese, nil
}
