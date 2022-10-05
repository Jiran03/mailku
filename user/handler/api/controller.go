package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/mailku/user/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (uh UserHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := uh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := uh.service.InsertData(toDomain(req))

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

func (uh UserHandler) Login(ctx echo.Context) error {
	var req RequestLoginJSON
	ctx.Bind(&req)
	err := uh.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	username := req.Username
	password := req.Password
	token, userRes, err := uh.service.CreateToken(username, password)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"token":   token,
		"data":    userObj,
	})
}

func (uh UserHandler) GetAll(ctx echo.Context) error {
	userRes, err := uh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := []ResponseJSON{}

	for _, value := range userRes {
		userObj = append(userObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	userRes, err := uh.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id := ctx.Param("id")
	userRes, err := uh.service.UpdateData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetByUsername(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	userRes, err := uh.service.GetByUsername(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := uh.service.DeleteData(id)
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

func (uh UserHandler) GetValidUsername(username string) (id string, role string, err error) {
	userObj, err := uh.service.GetByUsername(username)
	if err != nil {
		return userObj.IDX, userObj.Role, err
	}

	return userObj.IDX, userObj.Role, nil
}
