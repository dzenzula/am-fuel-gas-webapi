package controller

import (
	"fmt"
	conf "main/configuration"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"krr-app-gitlab01.europe.mittalco.com/pait/modules/go/authorization"
)

// GetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {
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

	gas := db.GetData()
	c.JSON(http.StatusOK, gas)
}

// SetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Param userdata body models.SetManualFuelGas true "Данные газ"
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

	db.InsertParametrs(data)
	c.JSON(http.StatusOK, "Insert successful")
}
