package database

import (
	c "main/configuration"
	"main/models"
	"strconv"
	"time"

	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	db *gorm.DB
}

type CalculationIds int

const (
	DensityCoefId int    = 1707482375047
	ImbalanceId   int    = 1707482385203
	layout        string = "2006-01-02"
)

func ConnectToPostgresDataBase() (*DBConnection, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.GlobalConfig.ConStringPgDb.Host,
		c.GlobalConfig.ConStringPgDb.User,
		c.GlobalConfig.ConStringPgDb.Password,
		c.GlobalConfig.ConStringPgDb.DBName,
		c.GlobalConfig.ConStringPgDb.Port,
		c.GlobalConfig.ConStringPgDb.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBConnection{db}, nil
}

func (dbc *DBConnection) InsertParametrs(d models.SetManualFuelGas) error {
	var queryInsert string = "INSERT INTO \"raw-data\".data(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
	var guid string
	parsedDate, errTime := time.Parse(layout, d.Date)
	if errTime != nil {
		return errTime
	}

	guid, errGuid := generateGUID()
	if errGuid != nil {
		return errGuid
	}

	switch d.Tag {
	case "month":
		parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), 1, 0, 0, 0, 0, parsedDate.Location())
	case "year":
		parsedDate = time.Date(parsedDate.Year(), time.January, 1, 0, 0, 0, 0, parsedDate.Location())
	}

	timestamp := parsedDate.Format("2006-01-02 15:04:05.999")
	res := dbc.db.Exec(queryInsert, strconv.Itoa(d.Id), timestamp, fmt.Sprintf("%f", d.Value), guid, strconv.Itoa(192))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (dbc *DBConnection) GetData(date time.Time) []models.GetManualFuelGas {
	var gas []models.GetManualFuelGas
	dateStart := date.Format(layout)

	queryGetData := "SELECT * FROM \"analytics-time-group\".get_last_manual_data(?)"
	dbc.db.Raw(queryGetData, dateStart).Scan(&gas)

	return gas
}

func (dbc *DBConnection) GetDataHistory(date time.Time, id string) []models.UpdateHistory {
	var history []models.UpdateHistory
	dateStart := date.Format(layout)

	queryGetHistory := "SELECT * FROM \"analytics-time-group\".get_manual_data_history(?,?)"
	dbc.db.Raw(queryGetHistory, dateStart, id).Scan(&history)

	return history
}

func (dbc *DBConnection) GetDensityCoefficientData(date string) models.GetDensityCoefficient {
	var res models.GetDensityCoefficient
	var history []models.DensityCalculationHistory

	queryGetDensityCoefficient := "SELECT * FROM \"analytics-time-group\".get_density_coefficient(?)"
	dbc.db.Raw(queryGetDensityCoefficient, date).Scan(&history)

	queryGetLastCoefficient := `SELECT value FROM "raw-data".data 
								WHERE id_measuring = 1703751302145 
								AND "timestamp" >= ? AND "timestamp" < ?::timestamptz + INTERVAL '1 DAY'
								ORDER BY id DESC
								LIMIT 1`
	dbc.db.Raw(queryGetLastCoefficient, date, date).Scan(&res.DensityCoefficient)
	res.CalculationHistory = history

	return res
}

func (dbc *DBConnection) RecalculateDensityCoefficient(date string, username string) {
	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_density_or_imbalance(?, ?, ?)`
	dbc.db.Exec(queryRecalculate, date, username, DensityCoefId)
}

func (dbc *DBConnection) GetImbalanceHistory(date string) []models.ImbalanceCalculationHistory {
	var res []models.ImbalanceCalculationHistory

	queryGetImbalanceHistory := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_history(?)`
	dbc.db.Raw(queryGetImbalanceHistory, date).Scan(&res)

	return res
}

func (dbc *DBConnection) GetCalculatedImbalanceDetails(batch string) models.GetCalculatedImbalanceDetails {
	var res models.GetCalculatedImbalanceDetails

	queryGetImbalanceDetails := `SELECT FROM * FROM "analytics-time-group".get_imbalance_calculation_data(?)`
	dbc.db.Raw(queryGetImbalanceDetails, batch).Scan(&res)

	return res
}

func (dbc *DBConnection) RecalculateImbalance(date string, username string, setData string) {
	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_density_or_imbalance(?, ?, ?, ?)`
	dbc.db.Exec(queryRecalculate, date, username, ImbalanceId, setData)
}

func (dbc *DBConnection) GetNodesList() []models.NodeList {
	var res []models.NodeList

	queryGetNodes := `SELECT id, description FROM "raw-data".measurings
						WHERE id IN (
							SELECT "raw-data".get_id_measuring_by_tags('AmFuelGas', 'NatGas', 'input', 'day')
						)
						AND id NOT IN (1703746063311, 1703746063312, 1703746063313, 1707468553, 1703751302145)`
	dbc.db.Raw(queryGetNodes).Scan(&res)

	return res
}

func (dbc *DBConnection) GetCalculationsList() []models.CalculationList {
	var res []models.CalculationList
	arr := []int{DensityCoefId, ImbalanceId}
	queryGetCalculationsList := "SELECT id, name, description FROM \"raw-data\".measurings WHERE id IN (?)"
	if err := dbc.db.Raw(queryGetCalculationsList, arr).Scan(&res); err.Error != nil {

	}

	return res
}

func (dbc *DBConnection) Close() {
	sqldb, _ := dbc.db.DB()
	sqldb.Close()
}

func generateGUID() (string, error) {
	// Generate a new random UUID
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// Format the UUID as a string in the specified format
	formattedUUID := newUUID.String()

	return formattedUUID, nil
}
