package controllers

import (
	structs "assigment2/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.Orderepository.GetOrders()
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var result gin.H
	var order structs.Order
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	customer_name := c.PostForm("Customer_Name")
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)
	order, err := idb.Orderepository.GetOrderID(int(order.ID))
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newOrder.Customer_name = customer_name
	order, err = idb.Orderepository.Update(order)
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "succesfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	order, err := idb.Orderepository.GetOrderID(int(order.ID))
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	order, err = idb.Orderepository.Delete(order)
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data succesfully deleted",
		}
	}
	c.JSON(http.StatusOK, result)
}
