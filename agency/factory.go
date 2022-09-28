package agency

import (
	handlerAPI "github.com/Jiran03/mailku/agency/handler/api"
	repoMySQL "github.com/Jiran03/mailku/agency/repository/mysql"
	service "github.com/Jiran03/mailku/agency/service"
	authMiddleware "github.com/Jiran03/mailku/auth/middlewares"
	"gorm.io/gorm"
)

func NewAgencyFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (agencyHandler handlerAPI.AgencyHandler) {
	agencyRepo := repoMySQL.NewAgencyRepository(db)
	agencyService := service.NewAgencyService(agencyRepo, configJWT)
	agencyHandler = handlerAPI.NewAgencyHandler(agencyService)
	return
}
