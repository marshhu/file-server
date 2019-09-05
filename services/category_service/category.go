package category_service

import "file-server/models"

type Category struct {
	Id            int64
	CategoryNo    string
	CategoryTitle string
	CategoryDesc  string

	PageNum  int
	PageSize int
}

func (c *Category) Add() error {
	return models.AddCategory(c.CategoryTitle, c.CategoryTitle)
}
