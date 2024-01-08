package controller

import "github.com/gin-gonic/gin"

// GetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {

}

// SetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Param userdata body models.SetManualFuelGas true "Данные газ"
// @Success 200
// @Router /api/SetParameters [post]
func SetParameters(c *gin.Context) {

}
