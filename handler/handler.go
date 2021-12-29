package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

// NOTE : Helper function ของการ handle error
// NOTE 2 : ที่ทำเป็น Private เพราะเราจะเรียกใช้แค่ใน package นี้เท่านั้น
func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}
