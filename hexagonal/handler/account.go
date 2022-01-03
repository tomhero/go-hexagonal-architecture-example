package handler

import (
	"bank/errs"
	"bank/logs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// NOTE : ดูว่าที่ส่งมาเป็น JSON หรือไม่ 🤔
	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("incorrect request header format"))
		return
	}

	request := service.NewAccountRequest{}
	// NOTE : .Decode รับ เป็น pointer
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("incorrect request body format"))
	}

	responseData, err := h.accSrv.NewAcccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated) // NOTE : ตอบ HTTP Status 201 ตาม REST API
	w.Header().Set("content-type", "application/json")
	// NOTE : .Encode รับเป็น pointer เช่นกันแต่ว่า responseData มันเป็น pointer อยู่แล้วหน่ะ
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		logs.Error(err)
	}
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	responseData, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		logs.Error(err)
	}
}
