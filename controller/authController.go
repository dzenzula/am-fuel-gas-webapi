package controller

import (
	"main/auth"
	conf "main/configuration"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserInfo
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200 {object} models.AuthUserData
// @Router /api/Authorization/GetCurrentUserInfo [get]
func GetCurrentUserInfo(c *gin.Context) {
	userDB, err := database.ConnectToDataBase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	session := sessions.Default(c)
	sessionID := session.Get("USER_IS_AUTH")
	if sessionID == nil {
		c.JSON(http.StatusBadRequest, "You are not authorized")
		return
	}

	// Получаем данные пользователя из сессии
	displayName := session.Get("USER_NAME").(string)
	login := session.Get("USER_DOMAIN_NAME").(string)
	userType := session.Get("USER_TYPE").(string)
	domain := session.Get("USER_DOMAIN").(string)
	userDomain := domain + "\\" + login

	if conf.GlobalConfig.ServiceId == nil {
		c.JSON(http.StatusBadRequest, "You must complete ServiceID in extension property")
		return
	}

	user, err := userDB.FindUserByUsername(userDomain)
	permissions := userDB.GetMyPermissions(&user)
	authUserData := userDataResponse(login, displayName, 1, "AUTH_USER", domain, permissions)

	// Проверяем наличие необходимых данных
	if displayName == "" || login == "" || userType == "" || domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incomplete user data in session"})
		return
	}

	c.JSON(http.StatusOK, authUserData)
}

// LogIn
// @Tags Authorization
// @Accept json
// @Produce json
// @Param userdata body models.UserData true "Данные пользователя"
// @Success 200 {object} models.AuthUserData
// @Router /api/Authorization/LogInAuthorization [post]
func LogInAuthorization(c *gin.Context) {
	var udata models.UserData
	userDB, err := database.ConnectToDataBase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&udata); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if udata.Domain == nil || udata.Login == nil || udata.Password == nil {
		c.JSON(http.StatusBadRequest, "Missing input parameters!")
		return
	}

	adresult, err := auth.AuthenticateActiveDirectory(*udata.Login, *udata.Password, *udata.Domain)
	if err != nil {
		c.JSON(http.StatusBadRequest, "User not found in Active Directory.")
		return
	}

	if adresult {
		var userDomainName string
		var userFound bool = true
		if *udata.Login != "" && *udata.Domain != "" {
			userDomainName = *udata.Domain + "\\" + *udata.Login
		}

		user, err := userDB.FindUserByUsername(userDomainName)
		if err != nil && user.Id == 0 {
			userFound = false
		}

		set := setAuthUser(user.DisplayName, *udata.Login, 1, "AUTH_USER", *udata.Domain, c)
		if set && userFound {
			err := userDB.UpdateUser(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
		}

		permissions := userDB.GetMyPermissions(&user)
		authUserData := userDataResponse(*udata.Login, user.DisplayName, 1, "AUTH_USER", *udata.Domain, permissions)

		c.JSON(http.StatusOK, authUserData)
	} else {
		c.JSON(http.StatusBadRequest, "Failed to login")
	}
}

// LogOut
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200
// @Router /api/Authorization/LogOutAuthorization [post]
func LogOutAuthorization(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, "User Sign out successfully")
}

func setAuthUser(userName string, domainName string, userIsAuth int, userType string, userDomain string, c *gin.Context) bool {
	session := sessions.Default(c)

	session.Set("USER_NAME", userName)
	session.Set("USER_DOMAIN_NAME", domainName)
	session.Set("USER_IS_AUTH", userIsAuth)
	session.Set("USER_TYPE", userType)
	session.Set("USER_DOMAIN", userDomain)

	session.Save()
	return true
}

func userDataResponse(userName string, userDisplayName string, userIsAuth int, userType string, userDomain string, permissions []models.MyPermission) models.AuthUserData {
	aud := models.AuthUserData{
		DisplayName: userName,
		Domain:      userDomain,
		Name:        userDisplayName,
		AuthStatus:  &userIsAuth,
		Type:        userType,
		Permission:  permissions,
	}

	return aud
}
