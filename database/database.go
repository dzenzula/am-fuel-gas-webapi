package database

import (
	c "main/configuration"
	"main/models"
	"strconv"
	"time"

	"fmt"

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

func (dbc *DBConnection) InsertParametrs(d models.SetManualFuelGas) {
	timestamp := d.Date.Format("2006-01-02 15:04:05.999")
	var queryInsert string

	if d.Tag == "day" {
		queryInsert = "INSERT INTO \"analytics-time-group\".data_day(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
	} else if d.Tag == "month" {
		queryInsert = "INSERT INTO \"analytics-time-group\".data_month(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
	} else if d.Tag == "year" {
		queryInsert = "INSERT INTO \"analytics-time-group\".data_year(id_measuring, \"timestamp\", value, batch_id, quality) VALUES (?, ?, ?, ?, ?)"
	}

	dbc.db.Exec(queryInsert, strconv.Itoa(d.Id), timestamp, fmt.Sprintf("%f", d.Value), nil, strconv.Itoa(192))
}

func (dbc *DBConnection) GetData(date time.Time) []models.GetManualFuelGas {
	var gas []models.GetManualFuelGas
	var temp []models.GetManualFuelGas
	var periods = []string{"day", "month", "year"}
	dateStart := date.Format("2006-01-02 15:04:05")
	dateEnd := date.Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	for _, p := range periods {
		queryGetData := "SELECT * FROM \"analytics-time-group\".get_manual_data_by_tag(?, ?, ?)"
		dbc.db.Raw(queryGetData, dateStart, dateEnd, p).Scan(&temp)
		for _, t := range temp {
			ct := t
			ct.Tag = p
			gas = append(gas, ct)
		}
		temp = []models.GetManualFuelGas{}
	}

	return gas
}

func (dbc *DBConnection) Close() {
	sqldb, _ := dbc.db.DB()
	sqldb.Close()
}
