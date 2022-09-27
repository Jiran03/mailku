package routes

import (
	"os"
	"strconv"

	"github.com/Jiran03/mailku/auth/middlewares"
	"github.com/Jiran03/mailku/config"
	"github.com/Jiran03/mailku/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	config.Init()
	db := config.DBInit()
	config.DBMigrate(db)
	expDuration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: expDuration,
	}

	user := user.NewUserFactory(db, configJWT)
	e := echo.New()
	middlewares.LogMiddleware(e)
	v1 := e.Group("/v1")
	v1.POST("/register", user.Create)
	v1.POST("/login", user.Login)

	userG := v1.Group("/user")
	userG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	userG.GET("", user.GetAll)
	userG.GET("/:id", user.GetByID)
	userG.PUT("/:id", user.Update, middlewares.UserValidation(user))
	userG.DELETE("/:id", user.Delete, middlewares.UserValidation(user))

	return e
}
