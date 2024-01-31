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
	Value           *string   `json:"Value"`
}

type SetManualFuelGas struct {
	Id    int
	Value float64
	Tag   string
	Date  string
}
