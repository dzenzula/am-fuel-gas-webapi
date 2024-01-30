package models

import "time"

type GetManualFuelGas struct {
	Id             int
	Name           string
	Description    string
	Value          *float64
	LastUpdateDate *time.Time `gorm:"column:timestamp"`
	Tag            string
}

type SetManualFuelGas struct {
	Id    int
	Value float64
	Tag   string
	Date  time.Time
}
