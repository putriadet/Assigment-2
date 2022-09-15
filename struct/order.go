package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Order_id      int
	Customer_name string
	Ordered_at    time.Time
	Items         []Item
}

type Item struct {
	gorm.Model
	Item_id     int
	Item_code   int
	Description string
	Quantity    int
	Order_id    int
}
