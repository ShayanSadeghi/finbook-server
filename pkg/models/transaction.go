package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id          uint64  `json:"id" gorm:"primaryKey"`
	Amount      float64 `json:"amount"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ResourceId  uint64  `json:"resource_id" gorm:"foreignKey:Resource.Id"`
	AccountId   uint64  `json:"account_id" gorm:"foreignKey:Account.Id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}

func (t *Transaction) CreateTransaction() *Transaction {
	db.Save(&t)
	return t
}

func GetAllTransactions() []Transaction {
	var Transactions []Transaction
	db.Find(&Transactions)
	return Transactions
}

func GetTransactionByID(Id uint64) *Transaction {
	var getTransaction Transaction
	if result := db.First(&getTransaction, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getTransaction
}

func DeleteTransaction(Id uint64) Transaction {
	var transaction Transaction
	db.Where("id=?", Id).Delete(&transaction)
	return transaction
}

func UpdateTransaction(Id uint64, updateTrx Transaction) Transaction {
	trxDetail := GetTransactionByID(Id)
	if updateTrx.Title != "" {
		trxDetail.Title = updateTrx.Title
	}

	if updateTrx.Amount != 0 {
		trxDetail.Amount = updateTrx.Amount
	}

	if updateTrx.Description != "" {
		trxDetail.Description = updateTrx.Description
	}

	if updateTrx.ResourceId != 0 {
		trxDetail.ResourceId = updateTrx.ResourceId
	}

	if updateTrx.AccountId != 0 {
		trxDetail.AccountId = updateTrx.AccountId
	}

	db.Save(trxDetail)
	return *trxDetail
}
