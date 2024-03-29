package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Id    uint64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Bank{})
}

func (b *Bank) CreateBank() *Bank {
	db.Save(&b)
	return b
}

func GetAllBanks() []Bank {
	var Banks []Bank
	db.Find(&Banks)
	return Banks
}

func GetBankByID(Id uint64) *Bank {
	var getBank Bank
	if result := db.First(&getBank, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getBank
}

func DeleteBank(Id uint64) Bank {
	var bank Bank
	db.Where("id=?", Id).Delete(&bank)
	return bank
}

func UpdateBank(Id uint64, updateBank Bank) Bank {
	bankDetail := GetBankByID(Id)
	if updateBank.Title != "" {
		bankDetail.Title = updateBank.Title
	}
	db.Save(bankDetail)
	return *bankDetail
}
