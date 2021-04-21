package controller

import (
	"backend/model"
	"backend/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func RegisterLogFunc(e *echo.Echo) {
	e.POST("/register", registerHandle)

	e.POST("/login", loginHandle)

	e.POST("/token", refreshTokenHandle, middleware.JWTWithConfig(utils.JwtConfig))
}

// check original token
// if valid, refresh it. else forbidden the request
func refreshTokenHandle(context echo.Context) error {
	isTokenValid := utils.ValidateJwtToken(context)
	if !isTokenValid {
		return context.JSON(http.StatusForbidden, errorReturnMsg(http.StatusForbidden, errors.New("token is invalid")))
	}

	token, err := utils.GenerateJwtToken(utils.GetCurrentUserEmail(context))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}
	return context.JSON(http.StatusOK, logReturnMsg(token))
}

// POST body:
//{
//	"email_address": ""
//	"password": ""
//}
func loginHandle(context echo.Context) error {
	postBodyObj := new(model.User)

	if err := context.Bind(postBodyObj); err != nil {
		return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
	}

	user, err := model.GetUserByEmailAddress(postBodyObj.EmailAddress)
	if err != nil {
		err = errors.New("email-address / password is invalid")
		return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
	}

	isValid := utils.ValidateSecurityHashing(postBodyObj.Password, user.Password)
	if !isValid {
		err = errors.New("email-address / password is invalid")
		return context.JSON(http.StatusForbidden, errorReturnMsg(http.StatusForbidden, err))
	}

	token, err := utils.GenerateJwtToken(user.EmailAddress)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	return context.JSON(http.StatusOK, logReturnMsg(token))
}

// Post Body:
// {
//		"email_address": ""
//	}
func registerHandle(context echo.Context) error {
	postBodyObj := new(struct {
		model.User
		RegisterCode string `json:"register_code"`
	})

	if err := context.Bind(postBodyObj); err != nil {
		return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
	}

	// validate user
	user, err := model.GetUserByEmailAddress(postBodyObj.EmailAddress)
	if err != nil && err != mongo.ErrNoDocuments {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	if err == nil && user.Password != "" {
		reRegisterErr := errors.New("you register again! please login")
		return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, reRegisterErr))
	}

	// validate register code
	//isValid, err := model.IsRegCodeValid(postBodyObj.RegisterCode)
	log.Println(postBodyObj.RegisterCode)

	err = model.InsertUser(postBodyObj.User)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	token, err := utils.GenerateJwtToken(postBodyObj.EmailAddress)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	return context.JSON(http.StatusOK, logReturnMsg(token))
}
