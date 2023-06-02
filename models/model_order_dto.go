/*
 * Yandex Lavka
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type OrderDto struct {
	OrderId int64 `json:"order_id" gorm:"primaryKey"`

	Weight float32 `json:"weight"`

	Regions int32 `json:"regions"`

	DeliveryHours []string `json:"delivery_hours" gorm:"type:text[]"`

	// Стоимость доставки заказа
	Cost int32 `json:"cost"`

	CompletedTime time.Time `json:"completed_time,omitempty"`
}
