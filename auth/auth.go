package auth

import (
	"crypto/tls"
	c "main/configuration"

	"github.com/go-ldap/ldap"
)

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
