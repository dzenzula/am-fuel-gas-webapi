package controller

import (
	"fmt"
	"main/auth"
	conf "main/configuration"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {
	permissions := []string{conf.GlobalConfig.Permissions.Show, conf.GlobalConfig.Permissions.Edit}
	auth.Init(c)
	checkPermissions := auth.CheckAnyPermission(permissions)
	if checkPermissions != auth.Ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error code: %d", checkPermissions))
	}

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
	auth.Init(c)
	checkPermissions := auth.CheckAnyPermission(permissions)
	if checkPermissions != auth.Ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error code: %d", checkPermissions))
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
}
