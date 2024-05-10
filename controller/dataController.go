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
	logger "krr-app-gitlab01.europe.mittalco.com/pait/modules/go/logging"
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
	logger.Debug("/am-fuel-gas-webapi/api/GetParameters --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetParameters --> Finished with success")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetParameterHistory --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetParameterHistory --> Finished with success")
}

// SetParam
// @Tags Parameters
// @Accept json
// @Produce json
// @Param data body models.SetManualFuelGas true "Данные газ"
// @Success 200
// @Router /api/SetParameters [post]
func SetParameters(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/SetParameters --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/SetParameters --> Finished with success")
}

// GetDensityCoefficientParam
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200 {object} models.GetDensityCoefficient
// @Router /api/GetDensityCoefficientDetails [get]
func GetDensityCoefficientDetails(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetDensityCoefficientDetails --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetDensityCoefficientDetails --> Finished with success")
}

// RecalculateDensityCoefficient
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200
// @Router /api/RecalculateDensityCoefficient [post]
func RecalculateDensityCoefficient(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/RecalculateDensityCoefficient --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/RecalculateDensityCoefficient --> Finished with success")
}

// GetImbalanceHistory
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Success 200 {object} []models.ImbalanceCalculationHistory
// @Router /api/GetImbalanceHistory [get]
func GetImbalanceHistory(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetImbalanceHistory --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetImbalanceHistory --> Finished with success")
}

// GetCalculatedImbalanceDetails
// @Tags Calculations
// @Accept json
// @Produce json
// @Param batch query string true "Id batch расчета"
// @Success 200 {object} models.GetCalculatedImbalanceDetails
// @Router /api/GetCalculatedImbalanceDetails [get]
func GetCalculatedImbalanceDetails(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetCalculatedImbalanceDetails --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetCalculatedImbalanceDetails --> Finished with success")
}

// PrepareImbalanceCalculation
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата расчета"
// @Success 200 {object} string
// @Router /api/PrepareImbalanceCalculation [post]
func PrepareImbalanceCalculation(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/PrepareImbalanceCalculation --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/PrepareImbalanceCalculation --> Finished with success")
}

// CalculateImbalance
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Param batch query string true "Id batch расчета"
// @Param separate query string true "Флаг раздельности расчета"
// @Param data body []models.SetImbalanceFlag true "Данные расчета небаланс"
// @Success 200
// @Router /api/CalculateImbalance [post]
func CalculateImbalance(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/CalculateImbalance --> Called")
	date := c.Query("date")
	batch := c.Query("batch")
	sep := c.Query("separate")
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
	err = database.CalculateImbalance(date, username, string(jsonData), batch, sep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Calculation succsessful")
	logger.Debug("/am-fuel-gas-webapi/api/CalculateImbalance --> Finished with success")
}

// RemoveImbalanceCalculation
// @Tags Calculations
// @Accept json
// @Produce json
// @Param batch query string true "Id batch расчета"
// @Success 200
// @Router /api/RemoveImbalanceCalculation [post]
func RemoveImbalanceCalculation(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/RemoveImbalanceCalculation --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/RemoveImbalanceCalculation --> Finished with success")
}

// GetNodesList
// @Tags Calculations
// @Accept json
// @Produce json
// @Param cloneId query string false "Id batch расчета"
// @Success 200 {object} models.NodeList
// @Router /api/GetNodesList [get]
func GetNodesList(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetNodesList --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetNodesList --> Finished with success")
}

// GetCalculationsList
// @Tags Calculations
// @Accept json
// @Produce json
// @Success 200 {object} models.CalculationList
// @Router /api/GetCalculationsList [get]
func GetCalculationsList(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetCalculationsList --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetCalculationsList --> Finished with success")
}

// GetScales
// @Tags Parameters
// @Accept json
// @Produce json
// @Success 200 {object} models.GetScales
// @Router /api/GetScales [get]
func GetScales(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/GetScales --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/GetScales --> Finished with success")
}

// UpdateScale
// @Tags Parameters
// @Accept json
// @Produce json
// @Param data body models.UpdateScale true "Данные шкалы"
// @Success 200
// @Router /api/UpdateScale [post]
func UpdateScale(c *gin.Context) {
	logger.Debug("/am-fuel-gas-webapi/api/UpdateScale --> Called")
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
	logger.Debug("/am-fuel-gas-webapi/api/UpdateScale --> Finished with success")
}

func isValidDate(c *gin.Context, dateString string) (bool, time.Time) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		logger.Error(err.Error())
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
		logger.Error(fmt.Sprintf("Error: %s", string(checkPermissions)))
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error: %s", string(checkPermissions)))
		return false
	}
	return true
}
