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

type CompleteOrder struct {
	CourierId int64 `json:"courier_id"`

	OrderId int64 `json:"order_id"`

	CompleteTime time.Time `json:"complete_time"`
}
