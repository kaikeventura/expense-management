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

func (service ExpenseService) CreateExpenseInBatch(expenseBatchDto dto.ExpenseBatch) (map[string][]string, error) {
	if len(expenseBatchDto.ReferencesMonth) == 0 || expenseBatchDto.ReferencesMonth == nil {
		return make(map[string][]string), fmt.Errorf("must contain at least one reference date")
	}

	statuses := make(map[string][]string)

	for _, referenceMonth := range expenseBatchDto.ReferencesMonth {
		if dto.ValidateFormatYearMonthDate(referenceMonth) {
			statuses["error"] = append(statuses["error"], fmt.Sprintf("%s => Invalid reference month format", referenceMonth))
			continue
		}

		expenseExisting, err := service.getExpenseByUserIdAndReferenceMonth(expenseBatchDto.UserId, referenceMonth)

		if err != nil {
			statuses["error"] = append(statuses["error"], fmt.Sprintf("%s => %s", referenceMonth, err))
			continue
		}
		if expenseExisting.Id != 0 {
			statuses["error"] = append(statuses["error"], fmt.Sprintf("%s => Reference month already exists", referenceMonth))
			continue
		}

		expenseEntity := entity.Expense{
			UserId:         expenseBatchDto.UserId,
			ReferenceMonth: referenceMonth,
			State:          dto.FUTURE.StateToString(),
			TotalAmount:    0,
		}
		_, err = service.repository.SaveExpense(expenseEntity)

		if err != nil {
			statuses["error"] = append(statuses["error"], fmt.Sprintf("%s => %s", referenceMonth, err))
			continue
		}

		statuses["success"] = append(statuses["success"], referenceMonth)
	}

	return statuses, nil
}

func (service ExpenseService) CreateFixedExpense(expenseId uint16, fixedExpenseDto dto.FixedExpense) (dto.FixedExpense, error) {
	expeseEntity, err := service.getExpenseById(expenseId)

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

	entity.PlusExpenseTotalAmount(&expeseEntity, fixedExpenseDto.Amount)
	err = service.repository.UpdateExpenseTotalAmount(expeseEntity.Id, expeseEntity.TotalAmount)

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

func (service ExpenseService) CreateCreditCardPurchase(expenseId uint16, purchase dto.CreditCardPurchaseRequest) error {
	if purchase.Installments < 1 {
		return fmt.Errorf("installments must be greater than 0")
	}

	expeseEntity, err := service.getExpenseById(expenseId)

	if err != nil {
		return err
	}

	installmentAmount := purchase.Amount / int32(purchase.Installments)
	sequenceNumber := expeseEntity.SequenceNumber

	for installmentNumber := uint8(0); installmentNumber < purchase.Installments; installmentNumber++ {
		expeseEntity, err := service.getExpenseByUserIdSequenceNumber(expeseEntity.UserId, sequenceNumber)

		if err != nil {
			return err
		}

		purchaseEntity := entity.CreditCardPurchase{
			ExpenseId:          expeseEntity.Id,
			Category:           purchase.Category,
			Description:        purchase.Description,
			Amount:             installmentAmount,
			CurrentInstallment: installmentNumber,
			LastInstallment:    purchase.Installments,
		}
		_, err = service.repository.SaveCreditCardPurchase(purchaseEntity)

		if err != nil {
			return err
		}

		entity.PlusExpenseTotalAmount(&expeseEntity, installmentAmount)
		err = service.repository.UpdateExpenseTotalAmount(expeseEntity.Id, expeseEntity.TotalAmount)

		if err != nil {
			return err
		}

		sequenceNumber++
	}

	return nil
}

func (service ExpenseService) GetCurrentExpenseByUserId(userId uint8) (dto.Expense, error) {
	expenseEntity, err := service.getExpenseByUserIdAndStatus(userId, "CURRENT")

	if err != nil {
		return dto.Expense{}, err
	}

	fixedExpensesDto := []dto.FixedExpense{}
	for _, fixedExpense := range expenseEntity.FixedExpenses {
		fixedExpenseDto := dto.FixedExpense{
			Id:          fixedExpense.Id,
			Category:    fixedExpense.Category,
			Description: fixedExpense.Description,
			Amount:      fixedExpense.Amount,
		}
		fixedExpensesDto = append(fixedExpensesDto, fixedExpenseDto)
	}

	purchasesDto := []dto.Purchase{}
	for _, purchase := range expenseEntity.Purchases {
		purchaseDto := dto.Purchase{
			Id:          purchase.Id,
			Category:    purchase.Category,
			Description: purchase.Description,
			Amount:      purchase.Amount,
		}
		purchasesDto = append(purchasesDto, purchaseDto)
	}

	creditCardPurchasesDto := []dto.CreditCardPurchase{}
	for _, creditCardPurchase := range expenseEntity.CreditCardPurchases {
		creditCardPurchaseDto := dto.CreditCardPurchase{
			Id:          creditCardPurchase.Id,
			Category:    creditCardPurchase.Category,
			Description: creditCardPurchase.Description,
			Amount:      creditCardPurchase.Amount,
		}
		creditCardPurchasesDto = append(creditCardPurchasesDto, creditCardPurchaseDto)
	}

	return dto.Expense{
		Id:                  expenseEntity.Id,
		UserId:              expenseEntity.UserId,
		ReferenceMonth:      expenseEntity.ReferenceMonth,
		State:               expenseEntity.State,
		TotalAmount:         expenseEntity.TotalAmount,
		FixedExpenses:       fixedExpensesDto,
		Purchases:           purchasesDto,
		CreditCardPurchases: creditCardPurchasesDto,
	}, nil
}

func (service ExpenseService) getExpenseById(expenseId uint16) (entity.Expense, error) {
	expese, err := service.repository.FindExpenseById(expenseId)

	if err != nil {
		return entity.Expense{}, err
	}

	return expese, nil
}

func (service ExpenseService) getExpenseByUserIdSequenceNumber(userId uint8, sequenceNumber uint16) (entity.Expense, error) {
	expese, err := service.repository.FindExpenseByUserIdSequenceNumber(userId, sequenceNumber)

	if err != nil {
		return entity.Expense{}, err
	}

	return expese, nil
}

func (service ExpenseService) getExpenseByUserIdAndReferenceMonth(userId uint8, referenceMonth string) (entity.Expense, error) {
	expese, err := service.repository.FindExpenseByUserIdAndReferenceMonth(userId, referenceMonth)

	if err != nil {
		return entity.Expense{}, err
	}

	return expese, nil
}

func (service ExpenseService) getExpenseByUserIdAndStatus(userId uint8, status string) (entity.Expense, error) {
	expese, err := service.repository.FindExpenseByUserIdAndStatus(userId, status)

	if err != nil {
		return entity.Expense{}, err
	}

	return expese, nil
}
