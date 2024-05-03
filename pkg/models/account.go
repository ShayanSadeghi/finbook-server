package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID            uint64 `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	AccountNumber string `json:"account_number"`
	BankID        uint64 `json:"bank_id"`
	UserID        uint64 `json:"user_id"`
	Bank          Bank   `gorm:"references:ID"`
	User          User   `gorm:"references:ID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Account{})
}

func (a *Account) CreateAccount(tokenString string) (*Account, error) {
	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	a.UserID = user.ID

	db.Create(&a)
	return a, nil
}

func GetAllAccounts(tokenString string) ([]Account, error) {

	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	var Accounts []Account
	db.Preload("Bank").Find(&Accounts, Account{UserID: user.ID})
	return Accounts, nil
}

func GetAccountByID(Id uint64, tokenString string) (*Account, error) {
	var getAccount Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	db.First(&getAccount, Account{ID: Id, UserID: user.ID})

	return &getAccount, nil
}

func DeleteAccount(Id uint64, tokenString string) (Account, error) {
	var account Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return account, err
	}

	db.First(&account, Account{ID: Id, UserID: user.ID}).Delete(&account)
	if &account == nil {
		return account, fmt.Errorf("account not found")
	}
	return account, nil
}

func UpdateAccount(Id uint64, updateAccount Account, tokenString string) (Account, error) {
	accountDetail, err := GetAccountByID(Id, tokenString)

	if err != nil {
		return Account{}, err
	}

	if updateAccount.Title != "" {
		accountDetail.Title = updateAccount.Title
	}

	if updateAccount.AccountNumber != "" {
		accountDetail.AccountNumber = updateAccount.AccountNumber
	}

	if updateAccount.BankID != 0 {
		accountDetail.BankID = updateAccount.BankID
	}

	if updateAccount.UserID != 0 {
		accountDetail.UserID = updateAccount.UserID
	}

	db.Save(accountDetail)
	return *accountDetail, nil

}
