/*
 * Yandex Lavka
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type GroupOrders struct {
	GroupOrderId int64 `json:"group_order_id"`

	Orders []OrderDto `json:"orders"`
}
