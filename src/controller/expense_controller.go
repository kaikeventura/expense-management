package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/expense-management/src/dto"
	"github.com/kaikeventura/expense-management/src/service"
)

var expenseService service.ExpenseService

func ConstructExpenseController(service service.ExpenseService) {
	expenseService = service
}

func CreateExpense(context *gin.Context) {
	var expense dto.Expense
	err := context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var createdExpense, expenseError = expenseService.CreateExpense(expense)

	if expenseError != nil {
		context.JSON(400, gin.H{
			"error": expenseError.Error(),
		})

		return
	}

	context.JSON(201, createdExpense)
}

func CreateExpenseInBatch(context *gin.Context) {
	var expenseBatch dto.ExpenseBatch
	err := context.ShouldBindJSON(&expenseBatch)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var createdExpenseBatch, expenseBatchError = expenseService.CreateExpenseInBatch(expenseBatch)

	if expenseBatchError != nil {
		context.JSON(400, gin.H{
			"error": expenseBatchError.Error(),
		})

		return
	}

	context.JSON(200, createdExpenseBatch)
}

func CreateFixedExpense(context *gin.Context) {
	var fixedExpense dto.FixedExpense
	err := context.ShouldBindJSON(&fixedExpense)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	expenseId := context.Param("expenseId")
	unitExpenseId, err := strconv.ParseUint(expenseId, 10, 16)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Param: " + err.Error(),
		})

		return
	}

	var createdFixedExpense, fixedExpenseError = expenseService.CreateFixedExpense(uint16(unitExpenseId), fixedExpense)

	if fixedExpenseError != nil {
		context.JSON(400, gin.H{
			"error": fixedExpenseError.Error(),
		})

		return
	}

	context.JSON(201, createdFixedExpense)
}

func CreatePurchase(context *gin.Context) {
	var purchase dto.Purchase
	err := context.ShouldBindJSON(&purchase)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	expenseId := context.Param("expenseId")
	unitExpenseId, err := strconv.ParseUint(expenseId, 10, 16)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Param: " + err.Error(),
		})

		return
	}

	var createdPurchase, purchaseError = expenseService.CreatePurchase(uint16(unitExpenseId), purchase)

	if purchaseError != nil {
		context.JSON(400, gin.H{
			"error": purchaseError.Error(),
		})

		return
	}

	context.JSON(201, createdPurchase)
}

func CreateCreditCardPurchase(context *gin.Context) {
	var purchase dto.CreditCardPurchaseRequest

	err := context.ShouldBindJSON(&purchase)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	expenseId := context.Param("expenseId")
	unitExpenseId, err := strconv.ParseUint(expenseId, 10, 16)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Param: " + err.Error(),
		})

		return
	}

	creditCardPurchaseError := expenseService.CreateCreditCardPurchase(uint16(unitExpenseId), purchase)

	if creditCardPurchaseError != nil {
		context.JSON(400, gin.H{
			"error": creditCardPurchaseError.Error(),
		})

		return
	}

	context.Status(204)
}

func GetCurrentExpense(context *gin.Context) {
	expenseId := context.Param("userId")
	unitExpenseId, err := strconv.ParseUint(expenseId, 10, 16)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Param: " + err.Error(),
		})

		return
	}

	currentExpense, creditCardPurchaseError := expenseService.GetCurrentExpenseByUserId(uint8(unitExpenseId))

	if creditCardPurchaseError != nil {
		context.JSON(400, gin.H{
			"error": creditCardPurchaseError.Error(),
		})

		return
	}

	context.JSON(200, currentExpense)
}
