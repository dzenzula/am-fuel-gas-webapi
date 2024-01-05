package database

import (
	"errors"
	c "main/configuration"
	"main/models"
	"time"

	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConnection struct {
	db *gorm.DB
}

func ConnectToDataBase() (*DBConnection, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;schema=auth", c.GlobalConfig.ConStringMsDb.Server, c.GlobalConfig.ConStringMsDb.UserID, c.GlobalConfig.ConStringMsDb.Password, c.GlobalConfig.ConStringMsDb.Database)

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBConnection{db}, nil
}

func (dbc *DBConnection) FindUserByUsername(username string) (models.User, error) {
	var user models.User

	//result := dbc.db.Where("DomainName = ?", username).First(&user)
	result := dbc.db.Raw(c.GlobalConfig.GetACSUser, username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found: %w", result.Error)
		}
		return user, result.Error
	}

	return user, nil
}

func (dbc *DBConnection) UpdateUser(user *models.User) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05.9999999")
	result := dbc.db.Exec(c.GlobalConfig.UpdateACSUser, timeNow, user.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found: %w", result.Error)
		}
		return result.Error
	}

	return nil
}

func (dbc *DBConnection) GetMyPermissions(user *models.User) []models.MyPermission {
	var myPermissions []models.MyPermission

	dbc.db.Raw(c.GlobalConfig.GetPermissions, user.DomainName, c.GlobalConfig.ServiceId).
		Scan(&myPermissions)

	return myPermissions
}
