package controller

import (
	"fmt"
	conf "main/configuration"
	"main/database"
	"main/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"krr-app-gitlab01.europe.mittalco.com/pait/modules/go/authorization"
)

// GetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {
	date := c.Query("date")
	
	isValid, truncatedTime := isValidDate(date)
	if !isValid {
		c.JSON(http.StatusBadRequest, "Дата не корректна.")
		return
	}

	permissions := []string{conf.GlobalConfig.Permissions.Show, conf.GlobalConfig.Permissions.Edit}
	authorization.Init(c)

	checkPermissions := authorization.CheckAnyPermission(permissions)
	if checkPermissions != authorization.Ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error code: %d", checkPermissions))
		return
	}

	db, err := database.ConnectToPostgresDataBase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gas := db.GetData(truncatedTime)
	db.Close()
	c.JSON(http.StatusOK, gas)
}

// SetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Param data body models.SetManualFuelGas true "Данные газ"
// @Success 200
// @Router /api/SetParameters [post]
func SetParameters(c *gin.Context) {
	var data models.SetManualFuelGas
	permissions := []string{conf.GlobalConfig.Permissions.Show, conf.GlobalConfig.Permissions.Edit}

	authorization.Init(c)
	checkPermissions := authorization.CheckAnyPermission(permissions)
	if checkPermissions != authorization.Ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error code: %d", checkPermissions))
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.ConnectToPostgresDataBase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.InsertParametrs(data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db.Close()
	c.JSON(http.StatusOK, "Insert successful")
}

func isValidDate(dateString string) (bool, time.Time) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return false, time.Time{}
	}

	// Возвращаем дату с временем 00:00:00
	return true, parsedTime.Truncate(24 * time.Hour)
}