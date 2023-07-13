package handlers

import (
	authDto "accounting-project/dto/auth"
	dto "accounting-project/dto/result"
	"accounting-project/models"
	jwtToken "accounting-project/pkg/jwt"
	"accounting-project/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *Handler {
	return &Handler{UserRepository: UserRepository}
}
func (h *Handler) SingUp(c echo.Context) error {
	request := new(authDto.SignUpRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	//check email not exist in db
	user := h.UserRepository.FindUserByEmail(request.Email)
	if user.Id != 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "This Email already exist"})
	}
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	//Create User And save to db
	user = models.User{
		Email:        request.Email,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Password:     string(hashedPassword),
		IsVerify:     true,
		RegisterTime: time.Now(),
	}
	user, err = h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	//create signUp Response
	response := authDto.SignUpResponse{
		Email: request.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}

func (h *Handler) Login(c echo.Context) error {
	request := new(authDto.LoginRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	//check user exist
	user := h.UserRepository.FindUserByEmail(request.Email)
	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "This Email not exist"})
	}
	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	response := authDto.SignUpResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}

func (h *Handler) Forgot(c echo.Context) error {
	request := new(authDto.ForgotRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	//check user exist
	user := h.UserRepository.FindUserByEmail(request.Email)
	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "This Email not exist"})
	}
	password := "Aa123456789Aa"
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	user.Password = string(hashedPassword)

	response := authDto.SignUpResponse{
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}
