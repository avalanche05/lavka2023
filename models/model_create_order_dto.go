/*
 * Yandex Lavka
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CreateOrderDto struct {
	Weight float32 `json:"weight"`

	Regions int32 `json:"regions"`

	DeliveryHours []string `json:"delivery_hours"`

	Cost int32 `json:"cost"`
}
