package handler

import (
	"encoding/json"
	"hexagonal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintln(w, err)
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customers, err := h.custSrv.GetCustomer(customerID)
	if err != nil {

		// appErr, ok := err.(errs.AppError)
		// if ok {
		// 	w.WriteHeader(appErr.Code)
		// 	fmt.Fprintln(w, appErr.Message)
		// 	return
		// }
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintln(w, err)
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
