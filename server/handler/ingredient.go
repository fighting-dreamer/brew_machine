package handler

import (
	"io"
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/logger"
	"nipun.io/brew_machine/service"
)

type IngredientHandler struct {
	IngredientManager service.IngredientManager
}

func NewIngredientHandler(dependencies *appcontext.Instance) *IngredientHandler {
	return &IngredientHandler{
		IngredientManager: dependencies.IngredientManager,
	}
}

func (ih *IngredientHandler) AddNewIngredientAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("AddNewIngredientAPI called")

	bodyBytes, _ := io.ReadAll(r.Body)
	ingredient := domain.Ingredient{}
	err := getBody(r.Context(), bodyBytes, &ingredient)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	err = ih.IngredientManager.AddNew(ingredient)

	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
	}

	WriteResponse(http.StatusOK, ingredient, w)
}

func (ih *IngredientHandler) RefillIngredientAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("RefillIngredientAPI called")

	bodyBytes, _ := io.ReadAll(r.Body)
	refillReq := domain.RefillIngredientRequest{}
	err := getBody(r.Context(), bodyBytes, &refillReq)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, []string{JsonParseError.Error()}, w)
		return
	}

	err = ih.IngredientManager.UpdateQuantity(refillReq.Name, refillReq.Quantity)

	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
	}

	WriteResponse(http.StatusOK, refillReq, w)
}

func (ih *IngredientHandler) IsIngredientAvailableAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("IsIngredientAvailableAPI called")

	ingredientName := r.URL.Query().Get("Name")

	isAvailable, err := ih.IngredientManager.IsAvailable(ingredientName)

	if err != nil {
		WriteErrorResponse(http.StatusInternalServerError, []string{err.Error()}, w)
	}

	WriteResponse(http.StatusOK, map[string]bool{"available": isAvailable}, w)
}
