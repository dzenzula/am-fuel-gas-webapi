package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"krr-app-gitlab01.europe.mittalco.com/pait/modules/go/authorization"

	c "main/configuration"
	"main/controller"
	"main/docs"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func (p *program) run() {
	startGin()
}

func main() {
	svcConfig := &service.Config{
		Name:        "FuelGasWebApiService",
		DisplayName: "Fuel Gas Web Api Service",
		Description: "",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("Error creating service:", err)
		return
	}

	if err = s.Run(); err != nil {
		fmt.Println("Error starting service:", err)
		return
	}
}

// @BasePath /am-fuel-gas-webapi
func startGin() {
	gin.SetMode(c.GlobalConfig.GinMode)
	docs.SwaggerInfo.Title = "Swagger AmFuelGaz API"
	docs.SwaggerInfo.Description = "This is a sample server AmFuelGaz server."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/am-fuel-gas-webapi/api",
		HttpOnly: true,
		MaxAge:   3600,
		SameSite: http.SameSiteNoneMode,
	})
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	authGroup := r.Group("/api/Authorization")
	{
		authGroup.GET("/GetCurrentUserInfo", authorization.GetCurrentUserInfo)
		authGroup.POST("/LogInAuthorization", authorization.LogInAuthorization)
		authGroup.POST("/LogOutAuthorization", authorization.LogOutAuthorization)
	}

	apiGroup := r.Group("/api")
	apiGroup.Use(authorization.AuthRequired)
	{
		apiGroup.GET("/GetParameters", controller.GetParameters)
		apiGroup.POST("/SetParameters", controller.SetParameters)
	}

	r.Run(c.GlobalConfig.ServerAddress)
}
