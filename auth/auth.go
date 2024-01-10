package auth

import (
	"crypto/tls"
	c "main/configuration"
	"main/database"
	"main/models"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap"
)

type PermissionInfo int

const (
	NotAuthorized PermissionInfo = 2
	NoCredentials PermissionInfo = 3
	Ok            PermissionInfo = 1
	Error         PermissionInfo = -1
)

var (
	login       string
	permissions []models.MyPermission
	isAuth      bool = false
)

func Init(s *gin.Context) {
	session := sessions.Default(s)

	if userDomainName, ok := session.Get("USER_DOMAIN_NAME").(string); ok {
		domain := session.Get("USER_DOMAIN").(string)
		login = userDomainName
		loginDomain := domain + "\\" + login
		mssql, _ := database.ConnectToMSDataBase()
		permissions = mssql.GetMyPermissions(loginDomain)
		isAuth = true
	} else {
		s.JSON(http.StatusBadRequest, "You are not authorized")
		return
	}
}

func AuthenticateActiveDirectory(username, password, domain string) (bool, error) {
	l, err := ldap.DialTLS("tcp", c.GlobalConfig.AdAddress, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return false, err
	}
	defer l.Close()

	usernamedomain := domain + "\\" + username
	err = l.Bind(usernamedomain, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckAnyPermission(perms []string) PermissionInfo {
	if !isAuth {
		return NotAuthorized
	}

	if len(perms) == 0 {
		return NoCredentials
	}

	for _, p := range perms {
		p = strings.ToLower(strings.TrimSpace(p))
		if hasValidPermission(p) {
			return Ok
		}
	}

	return NoCredentials
}

func hasValidPermission(p string) bool {
	for _, perm := range permissions {
		if perm.Name == p && perm.Permission == 1 {
			return true
		}
	}
	return false
}
