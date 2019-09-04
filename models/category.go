package models

import (
	"file-server/util"

	"github.com/jinzhu/gorm"
)

type Category struct {
	Id            int64  `gorm:"primary_key" json:"id"`
	CategoryNo    string `json:"category_no"`
	CategoryTitle string `json:"category_title"`
	CategoryDesc  string `json:"category_desc"`
}

func AddCategory(categoryTitle string, categoryDesc string) error {
	category_no, _ := util.NewUUID()
	categroy := Category{CategoryNo: category_no, CategoryTitle: categoryTitle, CategoryDesc: categoryDesc}
	if err := db.Create(&categroy).Error; err != nil {
		return err
	}
	return nil
}

func ExistCategory(categroyNo string) (bool, error) {
	var category Category
	err := db.Select("categroy_no").Where("category_no = ?", categroyNo).First(&category).Error
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
