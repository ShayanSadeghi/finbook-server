package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type ResourceCategory struct {
	gorm.Model
	ID      uint64 `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	IsInput bool   `json:"is_input"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&ResourceCategory{})
}

func (rc *ResourceCategory) CreateResourceCategory() *ResourceCategory {
	db.Save(&rc)
	return rc
}

func GetAllResourceCategories() []ResourceCategory {
	var ResourceCat []ResourceCategory
	db.Find(&ResourceCat)
	return ResourceCat
}

func GetResourceCategoryByID(Id uint64) *ResourceCategory {
	var getResourceCat ResourceCategory
	if result := db.First(&getResourceCat, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getResourceCat
}

func DeleteResourceCategory(Id uint64) ResourceCategory {
	var resourceCat ResourceCategory
	db.Where("id=?", Id).Delete(&resourceCat)
	return resourceCat
}

func UpdateResourceCategory(Id uint64, updateResourceCategory ResourceCategory) ResourceCategory {
	ResourceCat := GetResourceCategoryByID(Id)
	if updateResourceCategory.Title != "" {
		ResourceCat.Title = updateResourceCategory.Title
	}

	ResourceCat.IsInput = updateResourceCategory.IsInput

	db.Save(ResourceCat)
	return *ResourceCat
}
