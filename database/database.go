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
	var queryInsert string = `INSERT INTO "raw-data".data(id_measuring, "timestamp", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)`
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

func (dbc *DBConnection) GetData(date time.Time, tag string) ([]models.GetManualFuelGas, error) {
	var gas []models.GetManualFuelGas
	dateStart := date.Format(layout)

	queryGetData := `SELECT * FROM "analytics-time-group".get_last_manual_data(?)`
	ans := dbc.db.Raw(queryGetData, dateStart, tag).Scan(&gas)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return gas, nil
}

func (dbc *DBConnection) GetDataHistory(date time.Time, id string) ([]models.UpdateHistory, error) {
	var history []models.UpdateHistory
	dateStart := date.Format(layout)

	queryGetHistory := `SELECT * FROM "analytics-time-group".get_manual_data_history(?,?)`
	ans := dbc.db.Raw(queryGetHistory, dateStart, id).Scan(&history)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return history, nil
}

func (dbc *DBConnection) GetDensityCoefficientData(date string) (models.GetDensityCoefficient, error) {
	var res models.GetDensityCoefficient
	var history []models.DensityCalculationHistory

	queryGetDensityCoefficient := `SELECT * FROM "analytics-time-group".get_density_coefficient(?)`
	cfans := dbc.db.Raw(queryGetDensityCoefficient, date).Scan(&history)
	if cfans.Error != nil {
		return res, cfans.Error
	}

	queryGetLastCoefficient := `SELECT value FROM "raw-data".data 
								WHERE id_measuring = 1703751302145 
								AND "timestamp" >= ? AND "timestamp" < ?::timestamptz + INTERVAL '1 DAY'
								ORDER BY id DESC
								LIMIT 1`
	cflans := dbc.db.Raw(queryGetLastCoefficient, date, date).Scan(&res.DensityCoefficient)
	if cflans.Error != nil {
		return res, cflans.Error
	}

	res.CalculationHistory = history

	return res, nil
}

func (dbc *DBConnection) RecalculateDensityCoefficient(date string, username string) error {
	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_density_or_imbalance(?, ?, ?)`
	ans := dbc.db.Exec(queryRecalculate, date, username, DensityCoefId)
	if ans.Error != nil {
		return ans.Error
	}
	return nil
}

func (dbc *DBConnection) GetImbalanceHistory(date string) ([]models.ImbalanceCalculationHistory, error) {
	var res []models.ImbalanceCalculationHistory

	queryGetImbalanceHistory := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_history(?)`
	ans := dbc.db.Raw(queryGetImbalanceHistory, date).Scan(&res)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return res, nil
}

func (dbc *DBConnection) GetCalculatedImbalanceDetails(batch string) (models.GetCalculatedImbalanceDetails, error) {
	var res models.GetCalculatedImbalanceDetails
	var data models.ImbalanceCalculation
	var nodes []models.Node

	queryGetImbalanceData := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_data(?)`
	dans := dbc.db.Raw(queryGetImbalanceData, batch).Scan(&data)
	if dans.Error != nil {
		return res, dans.Error
	}

	queryGetImbalanceNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_data_nodes(?)`
	nans := dbc.db.Raw(queryGetImbalanceNodes, batch).Scan(&nodes)
	if nans.Error != nil {
		return res, nans.Error
	}

	res.ImbalanceCalculation = data
	res.Nodes = nodes

	return res, nil
}

func (dbc *DBConnection) CalculateImbalance(date string, username string, setData string) error {
	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_density_or_imbalance(?, ?, ?, ?)`
	ans := dbc.db.Exec(queryRecalculate, date, username, ImbalanceId, setData)
	if ans.Error != nil {
		return ans.Error
	}

	return nil
}

func (dbc *DBConnection) SetAdjustment(date string, setData string) error {
	queryAdjustment := `CALL "analytics-time-group".set_update_imbalance_calculation_adjustment(?, ?)`
	ans := dbc.db.Exec(queryAdjustment, date, setData)
	if ans.Error != nil {
		return ans.Error
	}

	return nil
}

func (dbc *DBConnection) GetNodesList(batch string) ([]models.NodeList, error) {
	var res []models.NodeList
	var ans *gorm.DB

	if batch != "" {
		queryGetNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_nodes_flag(?)`
		ans = dbc.db.Raw(queryGetNodes, batch).Scan(&res)
	} else {
		queryGetNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_nodes_flag()`
		ans = dbc.db.Raw(queryGetNodes).Scan(&res)
	}
	if ans.Error != nil {
		return nil, ans.Error
	}

	return res, nil
}

func (dbc *DBConnection) GetCalculationsList() ([]models.CalculationList, error) {
	var res []models.CalculationList
	arr := []int{DensityCoefId, ImbalanceId}
	queryGetCalculationsList := `SELECT id, name, description FROM "raw-data".measurings WHERE id IN (?)`
	if err := dbc.db.Raw(queryGetCalculationsList, arr).Scan(&res); err.Error != nil {
		return nil, err.Error
	}

	return res, nil
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
