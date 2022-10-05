package routes

import (
	"os"
	"strconv"

	"github.com/Jiran03/mailku/agency"
	"github.com/Jiran03/mailku/auth/middlewares"
	"github.com/Jiran03/mailku/config"
	"github.com/Jiran03/mailku/mail"
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
	mail := mail.NewMailFactory(db, configJWT)
	agency := agency.NewAgencyFactory(db, configJWT)
	e := echo.New()
	middlewares.LogMiddleware(e)
	v1 := e.Group("/v1")
	v1.POST("/su-adminmailku/register", user.Create)
	v1.POST("/login", user.Login)

	adminG := v1.Group("/admin")
	adminG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.RoleValidation(os.Getenv("ROLE_ADMIN"), user))
	adminToUserG := adminG.Group("/user")
	adminToUserG.GET("", user.GetAll)
	adminToUserG.GET("/:id", user.GetByID)
	adminToUserG.PUT("/:id", user.Update)
	adminToUserG.DELETE("/:id", user.Delete)
	adminToMailG := adminG.Group("/mail")
	adminToMailG.POST("", mail.Create)
	adminToMailG.PUT("/:id", mail.Update)
	adminToMailG.DELETE("/:id", mail.Delete)

	userG := v1.Group("/user")
	userG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	userG.GET("/:id", user.GetByID)
	userG.PUT("/:id", user.Update, middlewares.UserValidation(user))
	userG.DELETE("/:id", user.Delete, middlewares.UserValidation(user))

	agencyG := v1.Group("/agency")
	agencyG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	agencyG.POST("", agency.Create)
	agencyG.GET("", agency.GetAll)
	agencyG.GET("/:id", agency.GetByID)
	agencyG.PUT("/:id", agency.Update)
	agencyG.DELETE("/:id", agency.Delete)

	mailG := v1.Group("/mail")
	mailG.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	mailG.GET("", mail.GetAll)
	mailG.GET("/:id", mail.GetByID)

	return e
}
