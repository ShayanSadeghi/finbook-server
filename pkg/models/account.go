package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id            uint64 `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	AccountNumber string `json:"account_number"`
	BankId        uint64 `json:"bank_id" gorm:"foreignKey:Bank.Id"`
	UserID        uint64 `json:"user_id" gorm:"foreignKey:User.Id"`
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

	a.UserID = user.Id

	db.Create(&a)
	return a, nil
}

func GetAllAccounts(tokenString string) ([]Account, error) {

	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	var Accounts []Account
	db.Find(&Accounts, Account{UserID: user.Id})
	return Accounts, nil
}

func GetAccountByID(Id uint64, tokenString string) (*Account, error) {
	var getAccount Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	db.First(&getAccount, Account{Id: Id, UserID: user.Id})

	return &getAccount, nil
}

func DeleteAccount(Id uint64, tokenString string) (Account, error) {
	var account Account

	user, err := verifyToken(tokenString)

	if err != nil {
		return account, err
	}

	db.First(&account, Account{Id: Id, UserID: user.Id}).Delete(&account)
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

	if updateAccount.BankId != 0 {
		accountDetail.BankId = updateAccount.BankId
	}

	if updateAccount.UserID != 0 {
		accountDetail.UserID = updateAccount.UserID
	}

	db.Save(accountDetail)
	return *accountDetail, nil

}
