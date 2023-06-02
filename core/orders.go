package core

import (
	"github.com/lib/pq"
	"log"
	"strconv"
	"yandex-team.ru/lavka/models"
)

func CreateOrder(createOrderRequest models.CreateOrderRequest) ([]models.OrderDto, error) {
	var orders []models.OrderDto
	db := GetDB()
	for _, order := range createOrderRequest.Orders {
		gormOrder := models.Order{
			Weight:        order.Weight,
			Regions:       order.Regions,
			DeliveryHours: pq.StringArray(order.DeliveryHours),
			Cost:          order.Cost,
		}
		result := db.Create(&gormOrder)
		if result.Error != nil {
			log.Printf("ERROR: %v\n", result.Error)
			return orders, result.Error
		}
		orders = append(orders, models.OrderDto{
			OrderId:       int64(gormOrder.ID),
			Weight:        gormOrder.Weight,
			Regions:       gormOrder.Regions,
			DeliveryHours: gormOrder.DeliveryHours,
			Cost:          gormOrder.Cost,
		})

	}
	return orders, nil
}

func GetOrder(orderIdStr string) (models.OrderDto, error) {
	var order models.OrderDto
	orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
	if err != nil {
		return models.OrderDto{}, err
	}

	db := GetDB()
	var gormOrder models.Order
	result := db.First(&gormOrder, orderId) // Replace "1" with the desired ID
	if result.Error != nil {
		return models.OrderDto{}, err
	}
	order.OrderId = int64(gormOrder.ID)
	order.Weight = gormOrder.Weight
	order.Regions = gormOrder.Regions
	order.DeliveryHours = gormOrder.DeliveryHours
	order.Cost = gormOrder.Cost
	if gormOrder.CompletedTime != nil {
		order.CompletedTime = *gormOrder.CompletedTime
	}
	return order, nil
}

func GetOrders(offsetStr string, limitStr string) ([]models.OrderDto, error) {
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 1
	}

	var gormOrders []models.Order
	var orders []models.OrderDto
	db := GetDB()
	res := db.Limit(int(limit)).Offset(int(offset)).Find(&gormOrders)
	if res.Error != nil {
		return orders, res.Error
	}

	for _, gormOrder := range gormOrders {
		var order models.OrderDto
		order.OrderId = int64(gormOrder.ID)
		order.Weight = gormOrder.Weight
		order.DeliveryHours = gormOrder.DeliveryHours
		order.Regions = gormOrder.Regions
		if gormOrder.CompletedTime != nil {
			order.CompletedTime = *gormOrder.CompletedTime
		}

		orders = append(orders, order)
	}

	return orders, nil

}

func CompleteOrder(completeOrderRequest models.CompleteOrderRequestDto) ([]models.OrderDto, error) {
	var orders []models.OrderDto
	db := GetDB()
	for _, completeOrder := range completeOrderRequest.CompleteInfo {
		var order models.Order
		res := db.First(&order, completeOrder.OrderId)
		if res.Error != nil {
			return orders, res.Error
		}
		order.CompletedTime = &completeOrder.CompleteTime
		order.CourierId = completeOrder.CourierId
		res = db.Save(order)
		if res.Error != nil {
			return orders, res.Error
		}

		var responseOrder models.OrderDto
		responseOrder.OrderId = int64(order.ID)
		responseOrder.Weight = order.Weight
		responseOrder.DeliveryHours = order.DeliveryHours
		responseOrder.Regions = order.Regions
		responseOrder.Cost = order.Cost
		responseOrder.CompletedTime = *order.CompletedTime

		orders = append(orders, responseOrder)
	}

	return orders, nil
}
