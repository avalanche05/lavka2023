/*
 * Yandex Lavka
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CouriersGroupOrders struct {
	CourierId int64 `json:"courier_id"`

	Orders []GroupOrders `json:"orders"`
}
