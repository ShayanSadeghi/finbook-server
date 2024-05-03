package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID          uint64   `json:"id" gorm:"primaryKey"`
	Amount      float64  `json:"amount"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ResourceID  uint64   `json:"resource_id"`
	AccountID   uint64   `json:"account_id"`
	Resource    Resource `gorm:"references:ID"`
	Account     Account  `gorm:"references:ID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}

func (t *Transaction) CreateTransaction(tokenString string) (*Transaction, error) {
	verifiedAccount, err := verifyAccount(t.AccountID, tokenString)

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
	var Accounts []Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	db.Find(&Accounts, Account{UserID: user.ID})
	//extract account ids
	var accountIDs []uint64
	for _, acc := range Accounts {
		accountIDs = append(accountIDs, acc.ID)
	}
	db.Preload("Resource").Preload("Account").Preload("Account.Bank").Where("account_id IN ?", accountIDs).Find(&Transactions)
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

	if result := db.Preload("Resource").Preload("Account").First(&getTransaction, Id); result.Error != nil {
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

	if updateTrx.ResourceID != 0 {
		trxDetail.ResourceID = updateTrx.ResourceID
	}

	if updateTrx.AccountID != 0 {
		trxDetail.AccountID = updateTrx.AccountID
	}

	db.Save(trxDetail)
	return *trxDetail, nil
}
