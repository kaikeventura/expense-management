package migration

import (
	"log"

	"github.com/kaikeventura/expense-management/src/repository/entity"
	"gorm.io/gorm"
)

func ExecuteMigrations(database *gorm.DB) {
	// trucateTables(database)
	err := database.AutoMigrate(entity.User{}, entity.Expense{}, entity.FixedExpense{}, entity.Purchase{}, entity.CreditCardPurchase{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

func trucateTables(database *gorm.DB) {
	database.Exec("truncate table purchases")
	database.Exec("truncate table credit_card_purchases")
	database.Exec("truncate table expenses")
	database.Exec("truncate table fixed_expenses")
	database.Exec("truncate table users")
}
