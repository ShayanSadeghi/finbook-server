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

func (a *Account) CreateAccount() *Account {
	db.Create(&a)
	return a
}

func GetAllAccounts() []Account {
	var Accounts []Account
	db.Find(&Accounts)
	return Accounts
}

func GetAccountByID(Id uint64) *Account {
	var getAccount Account
	if result := db.First(&getAccount, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getAccount
}

func DeleteAccount(Id uint64) Account {
	var account Account
	db.Where("id=?", Id).Delete(&account)
	return account
}

func UpdateAccount(Id uint64, updateAccount Account) Account {
	accountDetail := GetAccountByID(Id)

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
	return *accountDetail

}
