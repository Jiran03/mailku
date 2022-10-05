package handlerAPI

import (
	"net/http"

	"github.com/Jiran03/mailku/agency/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AgencyHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewAgencyHandler(service domain.Service) AgencyHandler {
	return AgencyHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (ah AgencyHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := ah.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := ah.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}

func (ah AgencyHandler) GetAll(ctx echo.Context) error {
	agencyRes, err := ah.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	agencyObj := []ResponseJSON{}
	for _, value := range agencyRes {
		agencyObj = append(agencyObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    agencyObj,
	})
}

func (ah AgencyHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	agencyRes, err := ah.service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	agencyObj := fromDomain(agencyRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    agencyObj,
	})
}

func (ah AgencyHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	agencyRes, err := ah.service.UpdateData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	agencyObj := fromDomain(agencyRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    agencyObj,
	})
}

func (ah AgencyHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := ah.service.DeleteData(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}
