package controller

import (
	"fmt"
	"io"
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
// @Param tag query string false "Тэг типы: 'day', 'month' or empty"
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {
	date := c.Query("date")
	tag := c.Query("tag")

	isValid, truncatedTime := isValidDate(c, date)
	if !isValid {
		return
	}

	permissions := []string{conf.GlobalConfig.Permissions.Show}
	if !checkPermissions(c, permissions) {
		return
	}

	gas, err := database.GetData(truncatedTime, tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gas)
}

// GetParamHistory
// @Tags Parameters
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Param id query int true "Id параметра"
// @Success 200 {object} models.UpdateHistory
// @Router /api/GetParameterHistory [get]
func GetParameterHistory(c *gin.Context) {
	idMeasuring := c.Query("id")
	date := c.Query("date")

	isValid, truncatedTime := isValidDate(c, date)
	if !isValid {
		return
	}

	permissions := []string{conf.GlobalConfig.Permissions.Show}
	if !checkPermissions(c, permissions) {
		return
	}

	history, err := database.GetDataHistory(truncatedTime, idMeasuring)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, history)
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
	permissions := []string{conf.GlobalConfig.Permissions.Edit}

	if !checkPermissions(c, permissions) {
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	isValid := isValidValue(data.Value)
	if !isValid {
		c.JSON(http.StatusBadRequest, "Введите корректное значение")
		return
	}

	if err := database.InsertParametrs(data); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Insert successful")
}

// GetDensityCoefficientParam
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200 {object} models.GetDensityCoefficient
// @Router /api/GetDensityCoefficientDetails [get]
func GetDensityCoefficientDetails(c *gin.Context) {
	var data models.GetDensityCoefficient
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Show}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetDensityCoefficientData(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// RecalculateDensityCoefficient
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200
// @Router /api/RecalculateDensityCoefficient [post]
func RecalculateDensityCoefficient(c *gin.Context) {
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	username := authorization.ReturnDomainUser()
	err := database.RecalculateDensityCoefficient(date, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Calculation succsessful")
}

// GetImbalanceHistory
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200 {object} []models.ImbalanceCalculationHistory
// @Router /api/GetImbalanceHistory [get]
func GetImbalanceHistory(c *gin.Context) {
	var data []models.ImbalanceCalculationHistory
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Show}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetImbalanceHistory(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetCalculatedImbalanceDetails
// @Tags Calculations
// @Accept json
// @Produce json
// @Param batch query string true "Id batch расчета"
// @Success 200 {object} models.GetCalculatedImbalanceDetails
// @Router /api/GetCalculatedImbalanceDetails [get]
func GetCalculatedImbalanceDetails(c *gin.Context) {
	var data models.GetCalculatedImbalanceDetails
	batch := c.Query("batch")
	permissions := []string{conf.GlobalConfig.Permissions.Show}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetCalculatedImbalanceDetails(batch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// PrepareImbalanceCalculation
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата расчета"
// @Success 200 {object} string
// @Router /api/PrepareImbalanceCalculation [post]
func PrepareImbalanceCalculation(c *gin.Context) {
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	username := authorization.ReturnDomainUser()
	batch, err := database.PrepareImbalanceCalculation(date, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, batch)
}

// CalculateImbalance
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Param batch query string true "Id batch расчета"
// @Param data body []models.SetImbalanceFlag true "Данные расчета небаланс"
// @Success 200
// @Router /api/CalculateImbalance [post]
func CalculateImbalance(c *gin.Context) {
	date := c.Query("date")
	batch := c.Query("batch")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	username := authorization.ReturnDomainUser()
	err = database.CalculateImbalance(date, username, string(jsonData), batch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Calculation succsessful")
}

// RemoveImbalanceCalculation
// @Tags Calculations
// @Accept json
// @Produce json
// @Param batch query string true "Id batch расчета"
// @Success 200
// @Router /api/RemoveImbalanceCalculation [post]
func RemoveImbalanceCalculation(c *gin.Context) {
	batch := c.Query("batch")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	if !checkPermissions(c, permissions) {
		return
	}

	username := authorization.ReturnDomainUser()
	err := database.RemoveImbalanceCalculation(username, batch)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}

// GetNodesList
// @Tags Calculations
// @Accept json
// @Produce json
// @Param cloneId query string false "Id batch расчета"
// @Success 200 {object} models.NodeList
// @Router /api/GetNodesList [get]
func GetNodesList(c *gin.Context) {
	batch := c.Query("cloneId")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetNodesList(batch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetCalculationsList
// @Tags Calculations
// @Accept json
// @Produce json
// @Success 200 {object} models.CalculationList
// @Router /api/GetCalculationsList [get]
func GetCalculationsList(c *gin.Context) {
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetCalculationsList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetScales
// @Tags Parameters
// @Accept json
// @Produce json
// @Success 200 {object} models.GetScales
// @Router /api/GetScales [get]
func GetScales(c *gin.Context) {
	var data []models.GetScales
	permissions := []string{conf.GlobalConfig.Permissions.Show}

	if !checkPermissions(c, permissions) {
		return
	}

	data, err := database.GetScales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpdateScale
// @Tags Parameters
// @Accept json
// @Produce json
// @Param data body models.UpdateScale true "Данные шкалы"
// @Success 200
// @Router /api/UpdateScale [post]
func UpdateScale(c *gin.Context) {
	var data models.UpdateScale
	permissions := []string{conf.GlobalConfig.Permissions.EditScales}

	if !checkPermissions(c, permissions) {
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.UpdateScale(data); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Update successful")
}

func isValidDate(c *gin.Context, dateString string) (bool, time.Time) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Дата не корректна.")
		return false, time.Time{}
	}

	// Возвращаем дату с временем 00:00:00
	return true, parsedTime.Truncate(24 * time.Hour)
}

func isValidValue(value float64) bool {
	res := value >= 0.0001
	return res
}

func checkPermissions(c *gin.Context, permissions []string) bool {
	authorization.Init(c)
	checkPermissions := authorization.CheckAnyPermission(permissions)
	if checkPermissions != authorization.Ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error code: %s", string(checkPermissions)))
		return false
	}
	return true
}
