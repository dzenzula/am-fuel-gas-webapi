package models

import (
	"time"
)

type User struct {
	Id          int    `gorm:"primaryKey"`
	DomainName  string `gorm:"size:256"`
	DisplayName string `gorm:"size:256"`
	Mail        string `gorm:"size:256"`
	Disable     bool   `gorm:"type:bit"`
	LastAuth    time.Time
	IsDeleted   *int
	UserRoles   []UserRole
	CreatedBy   string `gorm:"size:256"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	UpdatedBy   string
}

func (User) TableName() string {
	return "Users"
}

type UserRole struct {
	Id        int
	UserId    int
	User      *User
	RoleId    int
	Role      *Role
	IsDeleted *int
	CreatedBy string `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	UpdatedBy string `gorm:"size:256"`
}

type Role struct {
	Id              int
	Name            string `gorm:"size:256"`
	Description     string `gorm:"size:256"`
	ServiceId       *int
	IsDeleted       *int
	Service         *Service
	RolePermissions []RolePermission
	UserRoles       []UserRole
	CreatedBy       string `gorm:"size:256"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	UpdatedBy       string `gorm:"size:256"`
}

type RolePermission struct {
	Id          int
	Permission  int `gorm:"size:256"`
	OperationId int
	Operation   *Operation
	RoleId      int
	Role        *Role
	IsDeleted   *int
	CreatedBy   string `gorm:"size:256"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	UpdatedBy   string `gorm:"size:256"`
}

type Operation struct {
	Id              int
	Name            string `gorm:"size:256"`
	Description     string `gorm:"size:256"`
	ServiceId       int
	IsDeleted       *int
	Service         *Service
	RolePermissions []RolePermission
	CreatedBy       string `gorm:"size:256"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	UpdatedBy       string `gorm:"size:256"`
}

type Service struct {
	Id                 int
	Name               string `gorm:"size:256"`
	Description        string `gorm:"size:256"`
	IsDeleted          *int
	Roles              []Role
	SyncAuthExtSystems []SyncAuthExtSystem
	Operations         []Operation
	CreatedBy          string `gorm:"size:256"`
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	UpdatedBy          string `gorm:"size:256"`
}

type SyncAuthExtSystem struct {
	Id           int
	DatabaseName string `gorm:"size:256"`
	ServerName   string `gorm:"size:256"`
	ServiceId    int
	Service      *Service
	LastSync     *time.Time
	Disable      int
	Duration     *int
	IsDeleted    *int
	Error        string `gorm:"size:4000"`
	ForceSync    bool
	CreatedBy    string `gorm:"size:256"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	UpdatedBy    string `gorm:"size:256"`
}

type UserData struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
	Domain   *string `json:"domain"`
}
