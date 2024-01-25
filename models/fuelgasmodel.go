package models

import "time"

type GetManualFuelGas struct {
	Id             int `gorm:"column:id_measuring"`
	Name           string
	Description    string
	Value          float64
	LastUpdateDate time.Time `gorm:"column:timestamp"`
}

type SetManualFuelGas struct {
	Id    int
	Value float64
}
