package handlers

import (
	dto "accounting-project/dto/result"
	"accounting-project/repositories"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var pathFileBanks = "http://localhost:3535/uploads/banks/"

type handlerBank struct {
	BankRepository repositories.BankRepository
}

func HandlerBank(BankRepository repositories.BankRepository) *handlerBank {
	return &handlerBank{BankRepository}
}

func (h handlerBank) GetAllBanks(c echo.Context) error {

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	banks := h.BankRepository.GetAllBankWithUserId(int(userId))

	for i, bank := range banks {
		banks[i].Icon = pathFileBanks + bank.Icon
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: banks})

}
