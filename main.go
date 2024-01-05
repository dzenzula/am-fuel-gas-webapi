package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	c "main/configuration"
	"main/controller"
	"main/docs"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger AmFuelGaz API"
	docs.SwaggerInfo.Description = "This is a sample server AmFuelGaz server."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 3600,
	})
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	work := r.Group("/api")
	{
		work.GET("/GetParameters", controller.GetParameters)
	}

	auth := r.Group("/api/Authorization")
	{
		auth.GET("/GetCurrentUserInfo", controller.GetCurrentUserInfo)
		auth.POST("/LogInAuthorization", controller.LogInAuthorization)
		auth.POST("/LogOutAuthorization", controller.LogOutAuthorization)
	}

	r.Run(c.GlobalConfig.Port)
}
