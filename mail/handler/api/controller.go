package handlerAPI

import (
	"net/http"

	"github.com/Jiran03/mailku/mail/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MailHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewMailHandler(service domain.Service) MailHandler {
	return MailHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (mh MailHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := mh.validation.Struct(req)
	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	receipt, _ := ctx.FormFile("receipt")
	src, _ := receipt.Open()
	defer src.Close()

	req.Receipt = src
	_, err := mh.service.InsertData(toDomain(req))
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

func (mh MailHandler) GetAll(ctx echo.Context) error {
	mailRes, err := mh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	mailObj := []ResponseJSON{}
	for _, value := range mailRes {
		mailObj = append(mailObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    mailObj,
	})
}

func (mh MailHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	mailRes, err := mh.service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	mailObj := fromDomain(mailRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    mailObj,
	})
}

func (mh MailHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	receipt, _ := ctx.FormFile("receipt")
	src, _ := receipt.Open()
	defer src.Close()

	req.Receipt = src
	mailRes, err := mh.service.UpdateData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	mailObj := fromDomain(mailRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    mailObj,
	})
}

func (mh MailHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := mh.service.DeleteData(id)
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
