package shopping

import (
	"github.com/hbkproduction/gostudy01/shopping/db"
)

// PriceCheck is a function which returns
// 1) the price of an item and
// 2) the boolean value whether item exists
func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}