package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	username := r.Context().Value("username").(string)
	r.ParseForm() 
	_, prodExsist := r.Form["product"]
	if !prodExsist{
		w.WriteHeader(400)
		err := json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Request Product Not Found",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	// TODO: answer here

	var list []model.Product
	var tot float64
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			tot += total
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}
	
	total := tot

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{
		Name: username,
		Cart: list,
		TotalPrice: total,
	} // TODO: replace this
	

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)

}
