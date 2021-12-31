package handler

import (
	"bank/logs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
อันนี้คือ Handler หรือ Presentation view ซึ่งไฟล์นี้ก็คือ Adapter
แล้ว Handler จะต้องไปเรียกใช้ตัว Service อีกที
*/

// NOTE : อยาลืมตั้งเป็น Private
type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		// NOTE : ตรงนี้มีการดึง error จากที่ปั้นมาจาก Business Layer เรียบร้อยแล้ว
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		logs.Error(err)
	}
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// NOTE : ดึงค่าจาก path variable (customerID) แล้วแปลงเป็น int ทันที
	customerId, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	customer, err := h.custSrv.GetCustomer(customerId)
	if err != nil {
		// NOTE : Service มีการ return เป็น AppError ออกมาเลยต้อง cast type ดูก่อน
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		logs.Error(err)
	}
}
