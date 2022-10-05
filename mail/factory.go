package mail

import (
	authMiddleware "github.com/Jiran03/mailku/auth/middlewares"
	handlerAPI "github.com/Jiran03/mailku/mail/handler/api"
	repoMySQL "github.com/Jiran03/mailku/mail/repository/mysql"
	service "github.com/Jiran03/mailku/mail/service"
	"gorm.io/gorm"
)

func NewMailFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (mailHandler handlerAPI.MailHandler) {
	mailRepo := repoMySQL.NewMailRepository(db)
	mailService := service.NewMailService(mailRepo, configJWT)
	mailHandler = handlerAPI.NewMailHandler(mailService)
	return
}
