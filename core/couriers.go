package core

import (
	"github.com/lib/pq"
	"log"
	"strconv"
	"time"
	"yandex-team.ru/lavka/models"
)

func CreateCourier(createCourierRequest models.CreateCourierRequest) (models.CreateCouriersResponse, error) {
	var response models.CreateCouriersResponse
	db := GetDB()
	for _, courier := range createCourierRequest.Couriers {
		gormCourier := models.Courier{
			CourierType:  courier.CourierType,
			Regions:      courier.Regions,
			WorkingHours: pq.StringArray(courier.WorkingHours),
		}
		result := db.Create(&gormCourier)
		if result.Error != nil {
			log.Printf("ERROR: %v\n", result.Error)
			return response, result.Error
		}
		response.Couriers = append(response.Couriers, models.CourierDto{
			CourierId:    int64(gormCourier.ID),
			CourierType:  gormCourier.CourierType,
			Regions:      gormCourier.Regions,
			WorkingHours: gormCourier.WorkingHours,
		})

	}
	return response, nil
}

func GetCourier(courierIdStr string) (models.CourierDto, error) {
	var courier models.CourierDto
	courierId, err := strconv.ParseInt(courierIdStr, 10, 64)
	if err != nil {
		return courier, err
	}

	db := GetDB()
	var gormCourier models.Courier
	result := db.First(&gormCourier, courierId) // Replace "1" with the desired ID
	if result.Error != nil {
		return courier, err
	}
	courier.CourierId = int64(gormCourier.ID)
	courier.CourierType = gormCourier.CourierType
	courier.Regions = gormCourier.Regions
	courier.WorkingHours = gormCourier.WorkingHours
	return courier, nil
}

func GetCouriers(offsetStr string, limitStr string) (models.GetCouriersResponse, error) {
	var getCouriersResponse models.GetCouriersResponse
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 1
	}
	getCouriersResponse.Offset = int32(offset)
	getCouriersResponse.Limit = int32(limit)
	var gormCouriers []models.Courier
	db := GetDB()
	res := db.Limit(int(limit)).Offset(int(offset)).Find(&gormCouriers)
	if res.Error != nil {
		return getCouriersResponse, res.Error
	}

	for _, gormCourier := range gormCouriers {
		var courier models.CourierDto
		courier.CourierId = int64(gormCourier.ID)
		courier.CourierType = gormCourier.CourierType
		courier.Regions = gormCourier.Regions
		courier.WorkingHours = gormCourier.WorkingHours

		getCouriersResponse.Couriers = append(getCouriersResponse.Couriers, courier)
	}

	return getCouriersResponse, nil

}

var earnValues map[string]int32 = map[string]int32{
	"FOOT": 2,
	"BIKE": 3,
	"AUTO": 4,
}

var ratingValues map[string]int32 = map[string]int32{
	"FOOT": 3,
	"BIKE": 2,
	"AUTO": 1,
}

func GetMetaInfo(courierIdStr string, startDateStr string, endDateStr string) (models.GetCourierMetaInfoResponse, error) {
	var response models.GetCourierMetaInfoResponse
	courierId, err := strconv.ParseInt(courierIdStr, 10, 64)
	if err != nil {
		return response, err
	}
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return response, err
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return response, err
	}
	db := GetDB()
	var courier models.Courier
	res := db.Find(&courier, courierId)
	if res.Error != nil {
		return response, res.Error
	}

	var orders []models.Order
	res = db.Where("completed_time >= ? AND completed_time < ? AND courier_id = ?", startDate, endDate, courierId).Find(&orders)

	if res.Error != nil {
		return response, res.Error
	}

	var earn int32 = 0
	for _, order := range orders {
		earn += order.Cost
	}
	earn *= earnValues[courier.CourierType]

	ordersCount := int32(len(orders))
	hoursCount := int32(endDate.Sub(startDate).Hours())

	rating := (ordersCount / hoursCount) * ratingValues[courier.CourierType]

	response.CourierId = int64(courier.ID)
	response.CourierType = courier.CourierType
	response.Regions = courier.Regions
	response.WorkingHours = courier.WorkingHours
	response.Earnings = earn
	response.Rating = rating

	return response, nil
}
