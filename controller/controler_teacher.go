package controller

import (
	"Go/WALIAPP-BACKEND/config"
	"Go/WALIAPP-BACKEND/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveTeacher(c *gin.Context) {
	var RegisterTeacher entity.Teacher
	var err error

	err = c.ShouldBindJSON(&RegisterTeacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("{\"bind_message\": \"%s\" }", err.Error()))
		return
	}

	err = config.DB.Create(RegisterTeacher).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("{create_message: %s }", err.Error()))
		return
	}

	c.JSON(http.StatusAccepted, "{\"message\": \"data has been saved!\"")
}