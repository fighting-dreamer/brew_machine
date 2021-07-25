package handler

import (
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/logger"
	"nipun.io/brew_machine/service"
	"strconv"
)

type DispenserHandler struct {
	DispenserService service.DispenserService
}

func NewDispenserHandler(dependencies *appcontext.Instance) *DispenserHandler {
	return &DispenserHandler{
		DispenserService: dependencies.DispenserService,
	}
}

func (dh *DispenserHandler) MakeBeverageAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("MakeBeverageAPI called")

	beverageName := r.URL.Query().Get("Name")
	outletNumber, err := strconv.Atoi(r.URL.Query().Get("Outlet"))
	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
		return
	}

	beverage, err := dh.DispenserService.MakeBeverage(beverageName, outletNumber)

	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
		return
	}

	WriteResponse(http.StatusOK, beverage, w)
}
