/*
 * Yandex Lavka
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controllers

import (
	"net/http"
	"yandex-team.ru/lavka/core"
	"yandex-team.ru/lavka/models"

	"github.com/gin-gonic/gin"
)

// CouriersAssignments - Список распределенных заказов
func CouriersAssignments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateCourier -
func CreateCourier(c *gin.Context) {
	var createCourierRequest models.CreateCourierRequest
	err := c.BindJSON(&createCourierRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	response, err := core.CreateCourier(createCourierRequest)
	c.JSON(http.StatusOK, response)
}

// GetCourierById -
func GetCourierById(c *gin.Context) {
	response, err := core.GetCourier(c.Param("courier_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetCourierMetaInfo -
func GetCourierMetaInfo(c *gin.Context) {
	response, err := core.GetMetaInfo(c.Param("courier_id"), c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetCouriers -
func GetCouriers(c *gin.Context) {
	response, err := core.GetCouriers(c.Query("offset"), c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, response)
}
