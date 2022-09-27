package middlewares

import (
	"errors"
	"os"
	"time"

	userHandlerAPI "github.com/Jiran03/mailku/user/handler/api"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaim struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (cJWT ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaim{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (cJWT ConfigJWT) CreateToken(userID string, username string) (token string, err error) {
	claims := JWTCustomClaim{
		userID,
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cJWT.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return token, err
	}

	return token, nil
}

func ExtractToken(ctx echo.Context) jwt.MapClaims {
	// paramID = ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	token, _ := jwt.Parse(user.Raw, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	},
	)
	claims := token.Claims.(jwt.MapClaims)

	return claims
}

func UserValidation(userController userHandlerAPI.UserHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			paramID := ctx.Param("id")
			claims := ExtractToken(ctx)
			username := claims["username"].(string)
			validUserID, _, err := userController.GetValidUsername(username)
			if paramID != validUserID {
				return errors.New("anda tidak dapat mengakses endpoint ini")
			}

			if err != nil {
				return err
			}

			return hf(ctx)
		}
	}
}
