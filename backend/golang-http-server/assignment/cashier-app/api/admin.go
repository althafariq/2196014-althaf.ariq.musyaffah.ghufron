package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/repository"
)

type AdminErrorResponse struct {
	Error string `json:"error"`
}

type ProductSales struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}

type AdminResponse struct {
	Sales []repository.Sales `json:"sales"`
}

func (api *API) getDashboard(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	productName := req.URL.Query().Get("product_name")
	var startPeriod time.Time
	var endPeriod time.Time

	startPeriod, err := time.Parse("2006-01-02", req.URL.Query().Get("start_period"))
	endPeriod, err = time.Parse("2006-01-02", req.URL.Query().Get("end_period"))

	// TODO: answer here

	getSalesRequest := repository.GetSalesRequest{
		ProductName: productName,
		StartPeriod: &startPeriod,
		EndPeriod:   &endPeriod,
	}

	if req.URL.Query().Get("start_period") == "" {
		getSalesRequest.StartPeriod = nil
	}
	if req.URL.Query().Get("end_period") == "" {
		getSalesRequest.EndPeriod = nil
	}
	// var getSalesRequest repository.GetSalesRequest

	// if req.URL.Query().Get("start_period") != "" && req.URL.Query().Get("end_period") == "" {
	// 	getSalesRequest = repository.GetSalesRequest{
	// 		StartPeriod: &startPeriod,
	// 		EndPeriod:   nil,
	// 		ProductName: productName,
	// 	}
	// } else if req.URL.Query().Get("end_period") != "" && req.URL.Query().Get("start_period") == "" {
	// 	getSalesRequest = repository.GetSalesRequest{
	// 		StartPeriod: nil,
	// 		EndPeriod:   &endPeriod,
	// 		ProductName: productName,
	// 	}
	// } else if req.URL.Query().Get("start_period") == "" && req.URL.Query().Get("end_period") == "" {
	// 	getSalesRequest = repository.GetSalesRequest{
	// 		StartPeriod: nil,
	// 		EndPeriod:   nil,
	// 		ProductName: productName,
	// 	}
	// } else {
	// 	getSalesRequest = repository.GetSalesRequest{
	// 		StartPeriod: &startPeriod,
	// 		EndPeriod:   &endPeriod,
	// 		ProductName: productName,
	// 	}
	// }

	encoder := json.NewEncoder(w)

	sales, err := api.salesRepo.Get(getSalesRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(AdminResponse{Sales: sales})
	w.WriteHeader(http.StatusOK)
}
