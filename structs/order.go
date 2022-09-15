package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

type Item struct {
	gorm.Model
	ItemID      int
	ItemCode    int
	Description string
	Quantity    int
	OrderID     int
}
