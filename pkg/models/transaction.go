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

func (t *Transaction) CreateTransaction(tokenString string) (*Transaction, error) {
	verifiedAccount, err := verifyAccount(t.AccountId, tokenString)

	if err != nil {
		return nil, err
	}
	if !verifiedAccount {
		return nil, fmt.Errorf("account is not available")
	}

	db.Save(&t)
	return t, nil
}

func GetAllTransactions(tokenString string) ([]Transaction, error) {
	var Transactions []Transaction
	db.Find(&Transactions)
	return Transactions, nil
}

func GetTransactionByID(Id uint64, tokenString string) (*Transaction, error) {
	var getTransaction Transaction
	verifiedAccount, err := verifyAccount(Id, tokenString)

	if err != nil {
		return nil, err
	}
	if !verifiedAccount {
		return nil, fmt.Errorf("account is not available")
	}

	if result := db.First(&getTransaction, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &getTransaction, nil
}

func DeleteTransaction(Id uint64, tokenString string) (Transaction, error) {
	var transaction Transaction
	verifiedAccount, err := verifyAccount(Id, tokenString)

	if err != nil {
		return Transaction{}, err
	}
	if !verifiedAccount {
		return Transaction{}, fmt.Errorf("account is not available")
	}

	db.Where("id=?", Id).Delete(&transaction)
	return transaction, nil
}

func UpdateTransaction(Id uint64, updateTrx Transaction, tokenString string) (Transaction, error) {
	trxDetail, err := GetTransactionByID(Id, tokenString)
	if err != nil {
		return Transaction{}, nil
	}

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
	return *trxDetail, nil
}
