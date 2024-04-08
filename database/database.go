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

var dbConnection DBConnection

const (
	DensityCoefId int    = 1707482375047
	ImbalanceId   int    = 1707482385203
	layout        string = "2006-01-02"
)

/*func ConnectToPostgresDataBase() (*DBConnection, error) {
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

	dbConnection.db = db

	return &DBConnection{db}, nil
}*/

func ConnectToPostgresDataBase() error {
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
		return err
	}

	dbConnection.db = db

	return nil
}

func InsertParametrs(d models.SetManualFuelGas) error {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return cerr
		}
	}

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
	res := dbConnection.db.Exec(queryInsert, strconv.Itoa(d.Id), timestamp, fmt.Sprintf("%f", d.Value), guid, strconv.Itoa(192))
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetData(date time.Time, tag string) ([]models.GetManualFuelGas, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var gas []models.GetManualFuelGas
	dateStart := date.Format(layout)

	queryGetData := `SELECT * FROM "analytics-time-group".get_last_manual_data(?, ?)`
	ans := dbConnection.db.Raw(queryGetData, dateStart, tag).Scan(&gas)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return gas, nil
}

func GetDataHistory(date time.Time, id string) ([]models.UpdateHistory, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var history []models.UpdateHistory
	dateStart := date.Format(layout)

	queryGetHistory := `SELECT * FROM "analytics-time-group".get_manual_data_history(?,?)`
	ans := dbConnection.db.Raw(queryGetHistory, dateStart, id).Scan(&history)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return history, nil
}

func GetDensityCoefficientData(date string) (models.GetDensityCoefficient, error) {
	var res models.GetDensityCoefficient
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return res, cerr
		}
	}

	var history []models.DensityCalculationHistory

	queryGetDensityCoefficient := `SELECT * FROM "analytics-time-group".get_density_coefficient(?)`
	cfans := dbConnection.db.Raw(queryGetDensityCoefficient, date).Scan(&history)
	if cfans.Error != nil {
		return res, cfans.Error
	}

	/*queryGetLastCoefficient := `SELECT * FROM "raw-data".get_day_last_value_by_id_measuring_date(?, ?, ?);`
	cflans := dbConnection.db.Raw(queryGetLastCoefficient, 1703751302145, date, 14).Scan(&res.DensityCoefficient)
	if cflans.Error != nil {
		return res, cflans.Error
	}*/

	queryGetLastCoefficient := `SELECT value, timestamp_insert
    							FROM "raw-data".data
								WHERE timestamp between ?::timestamptz - INTERVAL '14 DAY' and ?
								AND id_measuring = 1703751302145
    							ORDER BY timestamp DESC, id DESC
    							LIMIT 1`
	cflans := dbConnection.db.Raw(queryGetLastCoefficient, date, date).Scan(&res)
	if cflans.Error != nil {
		return res, cflans.Error
	}

	res.CalculationHistory = history

	return res, nil
}

func RecalculateDensityCoefficient(date string, username string) error {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return cerr
		}
	}

	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_density(?, ?)`
	ans := dbConnection.db.Exec(queryRecalculate, date, username)
	if ans.Error != nil {
		return ans.Error
	}
	return nil
}

func GetImbalanceHistory(date string) ([]models.ImbalanceCalculationHistory, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var res []models.ImbalanceCalculationHistory

	queryGetImbalanceHistory := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_history(?)`
	ans := dbConnection.db.Raw(queryGetImbalanceHistory, date).Scan(&res)
	if ans.Error != nil {
		return nil, ans.Error
	}

	return res, nil
}

func GetCalculatedImbalanceDetails(batch string) (models.GetCalculatedImbalanceDetails, error) {
	var res models.GetCalculatedImbalanceDetails
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return res, cerr
		}
	}

	var data models.ImbalanceCalculation
	var nodes []models.Node

	queryGetImbalanceData := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_data(?)`
	dans := dbConnection.db.Raw(queryGetImbalanceData, batch).Scan(&data)
	if dans.Error != nil {
		return res, dans.Error
	}

	queryGetImbalanceNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_data_nodes(?)`
	nans := dbConnection.db.Raw(queryGetImbalanceNodes, batch).Scan(&nodes)
	if nans.Error != nil {
		return res, nans.Error
	}

	res.ImbalanceCalculation = data
	res.Nodes = nodes

	return res, nil
}

func PrepareImbalanceCalculation(date string, username string) (string, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return "", cerr
		}
	}

	var res string

	queryCancel := `CALL "analytics-time-group".del_natural_gas_imbalance_empty_calculation(?)`
	cans := dbConnection.db.Exec(queryCancel, username)
	if cans.Error != nil {
		return "", cans.Error
	}

	queryCreateCalc := `CALL "analytics-time-group".ins_day_natural_gas_empty_imbalance(?, ?);`
	ans := dbConnection.db.Raw(queryCreateCalc, date, username).Scan(&res)
	if ans.Error != nil {
		return "", ans.Error
	}

	return res, nil
}

func CalculateImbalance(date string, username string, setData string, batch string) error {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return cerr
		}
	}

	queryRecalculate := `CALL "analytics-time-group".ins_calculate_day_natural_gas_imbalance_main(?, ?, ?, ?)`
	ans := dbConnection.db.Exec(queryRecalculate, date, username, setData, batch)
	if ans.Error != nil {
		return ans.Error
	}

	return nil
}

func RemoveImbalanceCalculation(username string, batch string) error {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return cerr
		}
	}

	queryCancel := `CALL "analytics-time-group".del_natural_gas_imbalance_empty_calculation(?, ?)`
	cans := dbConnection.db.Exec(queryCancel, username, batch)
	if cans.Error != nil {
		return cans.Error
	}

	return nil
}

func GetNodesList(batch string) ([]models.NodeList, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var res []models.NodeList
	var ans *gorm.DB

	if batch != "" {
		queryGetNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_nodes_flag(?)`
		ans = dbConnection.db.Raw(queryGetNodes, batch).Scan(&res)
	} else {
		queryGetNodes := `SELECT * FROM "analytics-time-group".get_imbalance_calculation_nodes_flag()`
		ans = dbConnection.db.Raw(queryGetNodes).Scan(&res)
	}
	if ans.Error != nil {
		return nil, ans.Error
	}

	return res, nil
}

func GetCalculationsList() ([]models.CalculationList, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var res []models.CalculationList
	arr := []int{DensityCoefId, ImbalanceId}
	queryGetCalculationsList := `SELECT id, name, description FROM "raw-data".measurings WHERE id IN (?)`
	if err := dbConnection.db.Raw(queryGetCalculationsList, arr).Scan(&res); err.Error != nil {
		return nil, err.Error
	}

	return res, nil
}

func GetScales() ([]models.GetScales, error) {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return nil, cerr
		}
	}

	var res []models.GetScales
	queryGetScales := `SELECT id_measuring, value, description FROM "raw-data".measuring_scales
					   WHERE id_measuring IN (
							SELECT "raw-data".get_id_measuring_by_tags('AmFuelGas', 'NatGas')
					   )
					   ORDER BY description ASC`
	if db := dbConnection.db.Raw(queryGetScales).Scan(&res); db.Error != nil {
		return nil, db.Error
	}

	return res, nil
}

func UpdateScale(scale models.UpdateScale) error {
	sqldb, _ := dbConnection.db.DB()
	if err := sqldb.Ping(); err != nil {
		if cerr := ConnectToPostgresDataBase(); cerr != nil {
			return cerr
		}
	}

	queryUpdateScale := `UPDATE "raw-data".measuring_scales SET value=? WHERE id_measuring=?`
	if err := dbConnection.db.Exec(queryUpdateScale, scale.Value, scale.Id); err.Error != nil {
		return err.Error
	}

	return nil
}

func Close() {
	sqldb, _ := dbConnection.db.DB()
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
