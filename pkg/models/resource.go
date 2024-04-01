package models

import (
	"finbook-server/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Id         uint64 `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	CategoryId uint64 `json:"category_id" gorm:"foreignKey:ResourceCategory.Id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Resource{})
}

func (r *Resource) CreateResource() *Resource {
	db.Save(&r)
	return r
}

func GetAllResources() []Resource {
	var Resources []Resource
	db.Find(&Resources)
	return Resources
}

func GetResourceByID(Id uint64) *Resource {
	var getResource Resource
	if result := db.First(&getResource, Id); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}
	return &getResource
}

func DeleteResource(Id uint64) Resource {
	var resource Resource
	db.Where("id=?", Id).Delete(&resource)
	return resource
}

func UpdateResource(Id uint64, updateResource Resource) Resource {
	resourceDetail := GetResourceByID(Id)
	if updateResource.Title != "" {
		resourceDetail.Title = updateResource.Title
	}
	db.Save(resourceDetail)
	return *resourceDetail
}