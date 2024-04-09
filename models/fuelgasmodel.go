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
	DensityCoefficient *float64                    `gorm:"column:value"`
	DateCoefficient    time.Time                   `gorm:"column:timestamp_insert"`
	CalculationHistory []DensityCalculationHistory `gorm:"-"`
}

type DensityCalculationHistory struct {
	StartDate       *time.Time `gorm:"column:start_date"`
	CalculationDate *time.Time `gorm:"column:calculation_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	Error           string     `gorm:"column:error"`
	UserName        string     `gorm:"column:username"`
	SyncMode        string     `gorm:"column:syncmode"`
	Value           *float64   `gorm:"column:dcvalue"`
}

type ImbalanceCalculationHistory struct {
	Id              string     `gorm:"column:calc_batch"`
	StartDate       *time.Time `gorm:"column:start_date"`
	CalculationDate *time.Time `gorm:"column:calculation_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	Error           string     `gorm:"column:error"`
	UserName        string     `gorm:"column:username"`
	SyncMode        string     `gorm:"column:syncmode"`
}

type GetCalculatedImbalanceDetails struct {
	ImbalanceCalculation ImbalanceCalculation
	Nodes                []Node
}

type ImbalanceCalculation struct {
	Id              string     `gorm:"column:calc_batch"`
	CalculationDate *time.Time `gorm:"column:calculation_date"`
	Nitka1Manual    *string    `gorm:"column:nitka1_manual"`
	Nitka2Manual    *string    `gorm:"column:nitka2_manual"`
	Nitka3Manual    *string    `gorm:"column:nitka3_manual"`
	Grp10Manual     *string    `gorm:"column:grp10_manual"`
	Nitka1Auto      *string    `gorm:"column:nitka1_auto"`
	Nitka2Auto      *string    `gorm:"column:nitka2_auto"`
	Nitka3Auto      *string    `gorm:"column:nitka3_auto"`
	Density         *string    `gorm:"column:density"`
	ManualTotal     *string    `gorm:"column:manual_total"`
	AutoTotal       *string    `gorm:"column:auto_total"`
	AggregateTotal  *string    `gorm:"column:aggregate_total"`
	PgRedisTotal    *string    `gorm:"column:pg_redis_total"`
	UserName        *string    `gorm:"column:username"`
}

type Node struct {
	Id                int64  `gorm:"column:id"`
	Description       string `gorm:"column:description"`
	BatchId           string `gorm:"column:batch"`
	Value             string `gorm:"column:node_value"`
	Flag              string `gorm:"column:imbalance_flag"`
	Distributed       string `gorm:"column:distributed"`
	GasRedistribution string `gorm:"column:gas_redistribution"`
	Consumption       string `gorm:"column:consumption"`
	Adjustment        string `gorm:"column:adjustment"`
}

type SetImbalanceFlag struct {
	Id         int
	Flag       string
	Adjustment string
}

type SetAdjustment struct {
	Id    int
	Batch string
	Value string
}

type NodeList struct {
	Id          int
	Description string
	Flag        string `gorm:"column:imbalance_flag"`
}

type CalculationList struct {
	Id          int
	Name        string
	Description string
}

type GetScales struct {
	Id          int64  `gorm:"column:id_measuring"`
	Value       string `gorm:"column:value"`
	Description string `gorm:"column:description"`
}

type UpdateScale struct {
	Id    int64
	Value string
}
