package handlers

import (
	authDto "accounting-project/dto/auth"
	"accounting-project/dto/otp"
	dto "accounting-project/dto/result"
	userDao "accounting-project/dto/user"
	"accounting-project/models"
	jwtToken "accounting-project/pkg/jwt"
	"accounting-project/pkg/redis"
	"accounting-project/repositories"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var path_file = "http://localhost:3535/uploads/"

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository: UserRepository}
}
func (h *handlerUser) SignUp(c echo.Context) error {
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

	//check phoneNumber not exist in db
	user := h.UserRepository.FindUserByPhoneNumber(request.PhoneNumber)
	if user.Id != 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "This PhoneNumber already exist"})
	}
	code, err := sendSMS(request.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, dto.ErrorResult{Code: http.StatusServiceUnavailable, Message: err.Error()})
	}
	//add code to request model and convert to json
	request.Code = code
	redisValue, err := json.Marshal(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}
	//set Phone number and code in redis
	var ctx = context.Background()
	err = redis.Rdb.Set(ctx, request.PhoneNumber, string(redisValue), time.Minute+10*time.Second).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: request.PhoneNumber})
}

func (h *handlerUser) Verify(c echo.Context) error {
	request := new(authDto.VerifyRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	//get origin code with Phone number in redis
	ctx := context.Background()
	redisValue, err := redis.Rdb.Get(ctx, request.PhoneNumber).Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	//convert redis value to signUp value and get code
	var signUpRequest authDto.SignUpRequest
	err = json.Unmarshal([]byte(redisValue), &signUpRequest)
	//check input code with origin code
	if signUpRequest.Code != request.Code {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Code not valid"})
	}
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	//Create User And save to db
	user := models.User{
		FirstName:    signUpRequest.FirstName,
		LastName:     signUpRequest.LastName,
		PhoneNumber:  request.PhoneNumber,
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

	user.PhotoURL = path_file + user.PhotoURL
	response := authDto.SignUpResponse{
		User:  user,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}
func (h *handlerUser) Login(c echo.Context) error {
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
	user := h.UserRepository.FindUserByPhoneNumber(request.PhoneNumber)
	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "The phone number or password is wrong"})
	}
	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "The phone number or password is wrong"})
	}
	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	user.PhotoURL = path_file + user.PhotoURL
	response := authDto.SignUpResponse{
		User:  user,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}

func (h *handlerUser) Forgot(c echo.Context) error {
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
	user := h.UserRepository.FindUserByPhoneNumber(request.PhoneNumber)
	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "This Phone number not exist"})
	}
	password, err := sendSMSForPassword(request.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, dto.ErrorResult{Code: http.StatusServiceUnavailable, Message: err.Error()})
	}

	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	user = h.UserRepository.ChangePassword(user.Id, string(hashedPassword))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: ""})

}

func sendSMS(phoneNumber string) (string, error) {
	client := &http.Client{}
	url := os.Getenv("KAVENEGAR_URL")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", errors.New("errored when create request for sms server")
	}
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	code := strconv.Itoa(rand.Intn(999999-100000) + 10000)
	q := req.URL.Query()
	q.Add("receptor", phoneNumber)
	q.Add("token", code)
	q.Add("template", "accounting")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("errored when sending request to the server")
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var response otp.Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", errors.New("errored when converting json to response model")
	}
	if response.Return.Status == http.StatusOK {
		return code, nil
	}
	return "", errors.New(response.Entries[0].Message)
}

func (h *handlerUser) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user := h.UserRepository.CheckAuth(int(userId))

	user.PhotoURL = path_file + user.PhotoURL
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) EditProfileStatus(c echo.Context) error {
	request := new(userDao.EditProfileStatusRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user := h.UserRepository.EditProfileStatus(int(userId), request.IsPublic)

	user.PhotoURL = path_file + user.PhotoURL
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) Resend(c echo.Context) error {
	request := new(authDto.ResendRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	//get origin code with Phone number in redis
	ctx := context.Background()
	redisValue, err := redis.Rdb.Get(ctx, request.PhoneNumber).Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	//convert redis value to signUp value and get code
	var signUpRequest authDto.SignUpRequest
	err = json.Unmarshal([]byte(redisValue), &signUpRequest)
	//send sms to client
	code, err := sendSMS(request.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, dto.ErrorResult{Code: http.StatusServiceUnavailable, Message: err.Error()})
	}
	//add code to request model and convert to json
	signUpRequest.Code = code
	newRedisValue, err := json.Marshal(signUpRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}
	//set Phone number and code in redis
	err = redis.Rdb.Set(ctx, request.PhoneNumber, string(newRedisValue), time.Minute+2*time.Second).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: signUpRequest.PhoneNumber})

}
func (h *handlerUser) ChangePassword(c echo.Context) error {
	request := new(userDao.ChangePasswordRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user := h.UserRepository.CheckAuth(int(userId))

	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "The phone number or password is wrong"})
	}
	if request.NewPassword != request.ConfirmNewPassword {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "The new password doesn't match the confirm password"})

	}
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	user = h.UserRepository.ChangePassword(int(userId), string(hashedPassword))

	user.PhotoURL = path_file + user.PhotoURL
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func sendSMSForPassword(phoneNumber string) (string, error) {
	client := &http.Client{}
	url := os.Getenv("KAVENEGAR_URL")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", errors.New("errored when create request for sms server")
	}
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	password := strconv.Itoa(rand.Intn(999999999-100000000) + 10000000)
	q := req.URL.Query()
	q.Add("receptor", phoneNumber)
	q.Add("token", password)
	q.Add("template", "accounting")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("errored when sending request to the server")
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var response otp.Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", errors.New("errored when converting json to response model")
	}
	if response.Return.Status == http.StatusOK {
		return password, nil
	}
	return "", errors.New(response.Entries[0].Message)
}

func (h *handlerUser) EditProfile(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	request := new(userDao.EditProfileRequest)
	request.FirstName = c.FormValue("first_name")
	request.LastName = c.FormValue("last_name")
	request.Email = c.FormValue("email")
	request.Country = c.FormValue("country")
	request.State = c.FormValue("state")
	request.City = c.FormValue("city")
	request.ZipCode = c.FormValue("zip_code")
	request.Address = c.FormValue("address")
	request.About = c.FormValue("about")

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user := h.UserRepository.CheckAuth(int(userId))

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Country = request.Country
	user.State = request.State
	user.City = request.City
	user.ZipCode = request.ZipCode
	user.Address = request.Address
	user.About = request.About
	user.PhotoURL = dataFile
	user = h.UserRepository.ChangeProfile(int(userId), user)
	user.PhotoURL = path_file + user.PhotoURL
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})

}

func (h *handlerUser) EditSocialLinks(c echo.Context) error {
	request := new(userDao.EditSocialLinksRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user := h.UserRepository.CheckAuth(int(userId))

	user.FacebookLink = request.FacebookLink
	user.InstagramLink = request.InstagramLink
	user.LinkedinLink = request.LinkedinLink
	user.TwitterLink = request.TwitterLink

	user = h.UserRepository.ChangeProfile(int(userId), user)

	user.PhotoURL = path_file + user.PhotoURL
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})

}
