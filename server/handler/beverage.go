package handler

import (
	"errors"
	"io"
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/logger"
	"nipun.io/brew_machine/service"
)

var (
	JsonParseError = errors.New("JsonParseError")
)

type BeverageHandler struct {
	BeverageManager service.BeverageManager
}

func NewBeverageHandler(dependencies *appcontext.Instance) *BeverageHandler {
	return &BeverageHandler{BeverageManager: dependencies.BeverageManager}
}

func (bh *BeverageHandler) AddNewBeverageAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("AddNewBeverageAPI called")

	bodyBytes, _ := io.ReadAll(r.Body)
	beverage := domain.Beverage{}
	err := getBody(r.Context(), bodyBytes, &beverage)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	err = bh.BeverageManager.AddNew(beverage)
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
	}

	WriteResponse(http.StatusOK, beverage, w)
}
