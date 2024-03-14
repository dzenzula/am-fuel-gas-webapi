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
// @Success 200 {object} models.GetManualFuelGas
// @Router /api/GetParameters [get]
func GetParameters(c *gin.Context) {
	date := c.Query("date")

	isValid, truncatedTime := isValidDate(c, date)
	if !isValid {
		return
	}

	permissions := []string{conf.GlobalConfig.Permissions.Show}
	if !checkPermissions(c, permissions) {
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gas := db.GetData(truncatedTime)
	db.Close()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	history := db.GetDataHistory(truncatedTime, idMeasuring)

	db.Close()
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isValid := isValidValue(data.Value)
	if !isValid {
		c.JSON(http.StatusBadRequest, "Введите корректное значение")
		return
	}

	db, err := connectToDatabase()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data = db.GetDensityCoefficientData(date)
	db.Close()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	username := authorization.ReturnDomainUser()
	db.RecalculateDensityCoefficient(date, username)
	db.Close()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data = db.GetImbalanceHistory(date)
	db.Close()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data = db.GetCalculatedImbalanceDetails(batch)
	db.Close()
	c.JSON(http.StatusOK, data)
}

// CalculateImbalance
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Param data body []models.SetImbalanceFlag true "Данные расчета небаланс"
// @Success 200
// @Router /api/CalculateImbalance [post]
func CalculateImbalance(c *gin.Context) {
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	username := authorization.ReturnDomainUser()
	db.CalculateImbalance(date, username, string(jsonData))
	db.Close()
	c.JSON(http.StatusOK, "Calculation succsessful")
}

// SetAdjustment
// @Tags Calculations
// @Accept json
// @Produce json
// @Param date query string true "Дата получения параметров"
// @Param data body []models.SetAdjustment true "Данные корректировки"
// @Success 200
// @Router /api/SetAdjustment [post]
func SetAdjustment(c *gin.Context) {
	date := c.Query("date")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	isValid, _ := isValidDate(c, date)
	if !isValid {
		return
	}

	if !checkPermissions(c, permissions) {
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db.SetAdjustment(date, string(jsonData))
	db.Close()
	c.JSON(http.StatusOK, "Calculation succsessful")
}

// GetNodesList
// @Tags Calculations
// @Accept json
// @Produce json
// @Param batch query string false "Id batch расчета"
// @Success 200 {object} models.NodeList
// @Router /api/GetNodesList [get]
func GetNodesList(c *gin.Context) {
	batch := c.Query("batch")
	permissions := []string{conf.GlobalConfig.Permissions.Calculate}

	if !checkPermissions(c, permissions) {
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data := db.GetNodesList(batch)
	db.Close()
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

	db, err := connectToDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data := db.GetCalculationsList()
	db.Close()
	c.JSON(http.StatusOK, data)
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

func connectToDatabase() (*database.DBConnection, error) {
	db, err := database.ConnectToPostgresDataBase()
	if err != nil {
		return nil, err
	}
	return db, nil
}
