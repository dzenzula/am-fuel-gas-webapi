package auth

import (
	"crypto/tls"
	"fmt"
	c "main/configuration"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap"
)

var login string

func Init(s *gin.Context) {
	session := sessions.Default(s)
	login = session.Get("USER_DOMAIN_NAME").(string)
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

func CheckPermission() {
	fmt.Println(login)
}
