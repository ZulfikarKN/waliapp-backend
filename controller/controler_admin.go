package controller

import (
	"Go/WALIAPP-BACKEND/config"
	"Go/WALIAPP-BACKEND/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AdminController interface {
	SaveAdmin(db *gorm.DB, c *gin.Context)
	GetAdminData(loginVm entity.LoginViewModel) (*entity.Admin, error)
}

func SaveAdmin(c *gin.Context) {
	var registerAdmin entity.Admin
	var err error
	var message string

	err = c.ShouldBindJSON(&registerAdmin)
	if err != nil {
		message, err = errorHandling(err.Error())
		c.JSON(http.StatusInternalServerError, message)
	}

	password, err := registerAdmin.EncryptPassword()
	if err != nil {
		message, err = errorHandling(err.Error())
		c.JSON(http.StatusInternalServerError, message)
	}

	registerAdmin.Password = password

	err = config.DB.Create(&registerAdmin).Error
	if err != nil {
		message, err = errorHandling(err.Error())
		c.JSON(http.StatusInternalServerError, message)
	}

	afterReqVM := entity.AdminViewModel{
		ID: registerAdmin.ID,
		FullName: fmt.Sprintf("%s %s", registerAdmin.FirstName, registerAdmin.LastName),
		Email: registerAdmin.Email,
	}

	c.JSON(http.StatusCreated, afterReqVM)
}

func errorHandling(errStr string) (string, error) {
	return fmt.Sprintf("{\"error\": %s}", errStr), nil
}