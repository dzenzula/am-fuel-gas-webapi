package models

import "time"

type GetManualFuelGas struct {
	Id             int
	Name           string
	Description    string
	Value          *float64
	Timestamp      *time.Time
	Tag            string
	LastUpdateDate *time.Time `gorm:"column:timestamp_insert"`
	//UpdateHistoryJSON *string         `gorm:"column:update_history" json:"-"`
	//UpdateHistory     []UpdateHistory `gorm:"-"`
}

type UpdateHistory struct {
	TimestampInsert time.Time `gorm:"column:timestamp_insert"`
	Value           *float64
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
	StartDate       *time.Time `gorm:"column:start_date"`
	CalculationDate *time.Time `gorm:"column:calculation_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	Error           string     `gorm:"column:error"`
	UserName        string     `gorm:"column:username"`
	SyncMode        string     `gorm:"column:syncmode"`
	Value           *float64   `gorm:"column:dcvalue"`
}

type GetImbalanceDetails struct {
	Id              string     `gorm:"column:calc_batch"`
	StartDate       *time.Time `gorm:"column:start_date"`
	CalculationDate *time.Time `gorm:"column:calculation_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	ManualTotal     *string    `gorm:"column:manual_total"`
	AutoTotal       *string    `gorm:"column:auto_total"`
	AggregateTotal  *string    `gorm:"column:aggregate_total"`
	PgRedisTotal    *string    `gorm:"column:pg_redis_total"`
	Error           *string    `gorm:"column:error"`
	Username        *string    `gorm:"column:username"`
	SyncMode        *string    `gorm:"column:syncmode"`
	NodesString     string     `gorm:"column:nodes" swaggerignore:"true" json:"-"`
	Nodes           []Node     `gorm:"-"`
}

type Node struct {
	MeasuringId       int64 `json:",string"`
	Value             string
	Flag              string
	Consumption       string
	GasRedistribution string
	Distributed       string
}

type SetImbalanceFlag struct {
	MeasuringId int
	Flag        string
}

type CalculationList struct {
	Id          int
	Name        string
	Description string
}
