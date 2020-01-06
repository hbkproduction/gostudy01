package db

import (
	"github.com/hbkproduction/gostudy01/shopping/models"
)

// LoadItem is a function which loads item
func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
