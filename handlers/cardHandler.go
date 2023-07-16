package handlers

import (
	cardDto "accounting-project/dto/card"
	dto "accounting-project/dto/result"
	"accounting-project/models"
	"accounting-project/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handlerCard struct {
	CardRepository repositories.CardRepository
}

func HandlerCard(CardRepository repositories.CardRepository) *handlerCard {
	return &handlerCard{CardRepository}
}

func (h *handlerCard) CreateCard(c echo.Context) error {
	request := new(cardDto.CreateCardRequest)
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

	card := models.Card{
		Name:       request.Name,
		CardNumber: request.CardNumber,
		UserId:     int64(userId),
		BankId:     request.BankId,
		Balance:    request.Balance,
	}
	card, err = h.CardRepository.CreateCard(card)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: card})
}
func (h *handlerCard) GetAllCards(c echo.Context) error {

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	cards := h.CardRepository.GetAllCards(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cards})

}
func (h *handlerCard) DeleteCard(c echo.Context) error {
	request := new(cardDto.DeleteCardRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	err = h.CardRepository.DeleteCard(request.CardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	cards := h.CardRepository.GetAllCards(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cards})

}

func (h *handlerCard) EditCard(c echo.Context) error {
	request := new(cardDto.EditCardRequest)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	card := h.CardRepository.GetCardById(request.CardId)

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	card.CardNumber = request.CardNumber
	card.BankId = request.BankId
	card.Balance = request.Balance
	card.Name = request.Name

	card, err = h.CardRepository.EditCard(int(card.Id), card)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	cards := h.CardRepository.GetAllCards(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cards})

}
