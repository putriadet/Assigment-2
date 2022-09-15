package repositories

import (
	structs "assigment2/struct"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return OrderRepo{
		db: db,
	}
}

func (or *OrderRepo) GetOrders() ([]structs.Order, error) {
	var orders []structs.Order
	err := or.db.Model(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepo) GetOrderID(id int) (structs.Order, error) {
	var order structs.Order
	err := or.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *OrderRepo) Create(order structs.Order, item structs.Item) (structs.Order, structs.Item, error) {
	fmt.Println(order)
	err := r.db.Create(&order).Error
	if err != nil {
		return order, item, err
	}
	return order, item, nil
}

func (r *OrderRepo) Update(order structs.Order) (structs.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *OrderRepo) Delete(order structs.Order) (structs.Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
