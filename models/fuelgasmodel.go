package models

import "time"

type GetManualFuelGas struct {
	Id                int
	Name              string
	Description       string
	Value             *float64
	Timestamp         *time.Time
	Tag               string
	LastUpdateDate    *time.Time      `gorm:"column:timestamp_insert"`
	UpdateHistoryJSON *string         `gorm:"column:update_history" json:"-"`
	UpdateHistory     []UpdateHistory `gorm:"-"`
}

type UpdateHistory struct {
	TimestampInsert time.Time `json:"TimestampInsert"`
	Value           *float64  `json:"Value,string"`
}

type SetManualFuelGas struct {
	Id    int
	Value float64
	Tag   string
	Date  string
}

type GetDensityCoefficient struct {
	DensityCoefficient *float64 `gorm:"column:value"`
	CalculationHistory []CalculationHistory
}

type CalculationHistory struct {
	StartDate *time.Time `gorm:"column:startdate"`
	EndDate   *time.Time `gorm:"column:enddate"`
	Error     string     `gorm:"column:error"`
	UserName  string     `gorm:"column:username"`
	SyncMode  string     `gorm:"column:syncmode"`
	Value     *float64   `gorm:"column:dcvalue"`
}

type CalculationList struct {
	Id          int
	Name        string
	Description string
}
