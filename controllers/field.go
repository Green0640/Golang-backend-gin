package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFields(c *gin.Context) {
	var field []Field
	db.Table("field").Scan(&field)
	c.JSON(http.StatusOK, field)
}

func GetField(c *gin.Context) {
	id := c.Param("id")
	var field Field
	db.Table("field").Where("id = ?", id).Scan(&field)
	c.JSON(http.StatusOK, field)
}

func AddField(c *gin.Context) {
	var field Field
	c.Bind(&field)
	db.Table("field").Create(&field)
	c.JSON(http.StatusOK, field)
}

func UpdateField(c *gin.Context) {
	id := c.Param("id")
	var field Field
	c.Bind(&field)
	db.Table("field").Model(&field).Where("id = ?", id).Updates(&field)
}

func DeleteField(c *gin.Context) {
	var field Field
	id := c.Param("id")
	db.Table("field").Where("id = ?", id).Scan(&field)
	db.Table("field").Delete(&User{}, id)
	c.JSON(http.StatusOK, "Delete OK!")
}

type Field struct {
	ID            int    `gorm:"column:id"`
	FieldName     string `json:"fieldName" gorm:"column:fieldName"`
	CompanyName   string `json:"companyName" gorm:"column:companyName"`
	RealTimeField string `json:"realTimeField" gorm:"column:realTimeFieldName"`
	Status        string `json:"status" gorm:"column:status"`
}
