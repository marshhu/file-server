package models

import (
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID            int64  `gorm:"primary_key" json:"id"`
	CategoryNo    string `json:"category_no"`
	CategoryTitle string `json:"category_title"`
	CategoryDesc  string `json:"category_desc"`
	ParentCategory  string `json:"parent_category"`
	CategoryPath  string `json:"category_path"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func AddCategory(categoryTitle string, categoryDesc string) error {
	uid,_ := uuid.NewV4()
	categoryNo := uid.String()
	category := Category{CategoryNo: categoryNo, CategoryTitle: categoryTitle, CategoryDesc: categoryDesc}
	if err := db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func ExistCategory(categoryNo string) (bool, error) {
	var category Category
	err := db.Select("category_no").Where("category_no = ?", categoryNo).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func DeleteCategory(categoryNo string) error {
	if err := db.Where("category_no = ?", categoryNo).Delete(&Category{}).Error; err != nil {
		return err
	}

	return nil
}

func GetCategory(categoryNo string) (*Category, error) {
	category := Category{}
	err := db.Where("category_no = ?", categoryNo).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategoryList(pageNum int, pageSize int, maps interface{}) ([]Category, error) {
	var (
		categoryList []Category
		err          error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&categoryList).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&categoryList).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categoryList, nil
}
