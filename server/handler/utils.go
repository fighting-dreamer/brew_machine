package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"nipun.io/brew_machine/logger"
)

func getBody(ctx context.Context, reqBodyByte []byte, body interface{}) error {
	err := json.Unmarshal(reqBodyByte, body)
	if err != nil {
		status := http.StatusBadRequest
		logMessage := "Failed parsing request payload for returning %d, err: %v"
		logger.Logger.Error().Msg(fmt.Sprintf(logMessage, status, err))
	}
	return err
}

type errorResponse struct {
	Errors []string `json:"errors"`
}

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	if response == nil {
		response = struct{}{}
	}
	body, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Error().Err(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(body)
}

func WriteErrorResponse(status int, errorMessages []string, rw http.ResponseWriter) {
	if errorMessages == nil {
		errorMessages = []string{}
	}
	errorResponse := errorResponse{Errors: errorMessages}
	respBytes, err := json.Marshal(errorResponse)
	if err != nil {
		logger.Logger.Error().Err(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
