package database

import (
	"encoding/json"
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
	var queryInsert string
	var guid string
	parsedDate, errTime := time.Parse("2006-01-02", d.Date)
	if errTime != nil {
		return errTime
	}

	guid, errGuid := generateGUID()
	if errGuid != nil {
		return errGuid
	}

	switch d.Tag {
	case "day":
		queryInsert = "INSERT INTO \"analytics-time-group\".data_day(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
	case "month":
		queryInsert = "INSERT INTO \"analytics-time-group\".data_month(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
		parsedDate = time.Date(parsedDate.Year(), parsedDate.Month(), 1, 0, 0, 0, 0, parsedDate.Location())
	case "year":
		queryInsert = "INSERT INTO \"analytics-time-group\".data_year(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
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
	var tmp []models.GetManualFuelGas
	var gas []models.GetManualFuelGas
	dateStart := date.Format("2006-01-02 15:04:05")
	//dateEnd := date.Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	queryGetData := "SELECT * FROM \"analytics-time-group\".get_manual_data_by_tag_test(?)"
	dbc.db.Raw(queryGetData, dateStart).Scan(&tmp)

	for _, t := range tmp {
		var updateHistory []models.UpdateHistory
		err := json.Unmarshal([]byte(*t.UpdateHistoryJSON), &updateHistory)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		g := t
		if updateHistory[len(updateHistory)-1].Value != nil {
			g.UpdateHistory = updateHistory
		}
		gas = append(gas, g)

	}

	return gas
}

func (dbc *DBConnection) GetDensityCoefficientData(date string) models.GetDensityCoefficient {
	var res models.GetDensityCoefficient
	var history []models.SyncHistory

	queryGetDensityCoefficient := "SELECT * FROM \"analytics-time-group\".get_density_coefficient(?)"
	dbc.db.Raw(queryGetDensityCoefficient, date).Scan(&history)

	queryGetLastCoefficient := `SELECT value FROM "analytics-time-group".data_day 
								WHERE id_measuring = 1703751302145 
								AND "timestamp" >= ? AND "timestamp" < ?::date + INTERVAL '1 DAY'
								ORDER BY id DESC
								LIMIT 1`
	dbc.db.Raw(queryGetLastCoefficient, date, date).Scan(&res.DensityCoefficient)
	res.SyncHistory = history

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
	formattedUUID := fmt.Sprintf("%s", newUUID)

	return formattedUUID, nil
}
