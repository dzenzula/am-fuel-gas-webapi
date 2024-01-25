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
	timeNow := time.Now().Format("2006-01-02 15:04:05.999")

	queryInsert := "INSERT INTO \"raw-data\".data(id_measuring, \"timestamp\", value, quality, batch_id) VALUES (?, ?, ?, ?, ?)"

	dbc.db.Exec(queryInsert, strconv.Itoa(d.Id), timeNow, fmt.Sprintf("%f", d.Value), strconv.Itoa(192), nil)
}

func (dbc *DBConnection) GetData() []models.GetManualFuelGas {
	var gas []models.GetManualFuelGas
	var ids []int

	queryGetIds := "SELECT  * FROM \"raw-data\".get_id_measuring_by_tags('AmFuelGas', 'NatGas', 'manual')"
	queryGetData := `SELECT d.id_measuring, ms.name, ms.description, d.value, d.\"timestamp\" FROM \"raw-data\".data AS d
					JOIN \"raw-data\".measurings AS ms ON ms.id = d.id_measuring WHERE d.id_measuring IN ?`
	
	dbc.db.Raw(queryGetIds).Scan(&ids)
	dbc.db.Raw(queryGetData, ids).Scan(&gas)

	return gas
}
