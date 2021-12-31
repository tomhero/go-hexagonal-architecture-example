package handler

import (
	"bank/service"
	"net/http"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
}
