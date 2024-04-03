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
	"main/database"
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

	database.ConnectToPostgresDataBase()
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/am-fuel-gas-webapi/api",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
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
		apiGroup.GET("/GetCalculationsList", controller.GetCalculationsList)

		apiGroup.GET("/GetParameters", controller.GetParameters)
		apiGroup.GET("/GetParameterHistory", controller.GetParameterHistory)
		apiGroup.POST("/SetParameters", controller.SetParameters)

		apiGroup.GET("/GetDensityCoefficientDetails", controller.GetDensityCoefficientDetails)
		apiGroup.POST("/RecalculateDensityCoefficient", controller.RecalculateDensityCoefficient)

		apiGroup.GET("/GetImbalanceHistory", controller.GetImbalanceHistory)
		apiGroup.GET("/GetCalculatedImbalanceDetails", controller.GetCalculatedImbalanceDetails)
		apiGroup.POST("/PrepareImbalanceCalculation", controller.PrepareImbalanceCalculation)
		apiGroup.POST("/CalculateImbalance", controller.CalculateImbalance)
		apiGroup.POST("/RemoveImbalanceCalculation", controller.RemoveImbalanceCalculation)
		apiGroup.GET("/GetNodesList", controller.GetNodesList)

		apiGroup.GET("/GetScales", controller.GetScales)
		apiGroup.POST("/UpdateScale", controller.UpdateScale)
	}

	r.Run(c.GlobalConfig.ServerAddress)
}
