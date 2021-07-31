package handler

import (
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/logger"
	"nipun.io/brew_machine/service"
)

type LockHandler struct {
	TransactionLockManager service.TransactionLockManager
}

func NewLockHandler(dependencies *appcontext.Instance) *LockHandler {
	return &LockHandler{TransactionLockManager: dependencies.TransactionLockManager}
}

func (lh *LockHandler) GetCurrentLocksStateAPI(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug().Msg("GetCurrentLocksStateAPI called")
	lh.TransactionLockManager.GetCurrentLockState()
	WriteResponse(http.StatusOK, "", w)
}
