package controllers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = database.CreateDatabase()

func GetAllUsers(c *gin.Context) {
	var user []User
	db.Table("user").Scan(&user)
	c.JSON(http.StatusOK, user)
}

type User struct {
	Id              int
	CompanyId       int
	Name            string
	HashedPassword  string
	Phone           string
	NotifyType      string
	LineNotifyToken string
}
