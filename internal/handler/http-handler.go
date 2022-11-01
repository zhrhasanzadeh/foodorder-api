package handler

import (
	"api-orderfood/internal/model"
	"api-orderfood/pkg/utils"
	"net/http"
	"strconv"
)

type HttpHandler interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	service model.Service
}

func NewHttpHandler(service model.Service) HttpHandler {
	return &httpHandler{
		service: service,
	}
}

func (h httpHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	convertedID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	price := r.FormValue("price")
	convertedPrice, err := strconv.Atoi(price)
	if err != nil {
		panic(err)
	}
	title := r.FormValue("title")
	err = h.service.AddOrder(model.FoodOrder{OrderID: convertedID, Price: convertedPrice, Title: title})
	if err != nil {
		w.Write([]byte("cant add order because of error: " + err.Error()))
	}
	w.Write([]byte("successfully add order"))
}

func (h httpHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	convertedID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	user, err := h.service.GetUser(convertedID)
	if err != nil {
		w.Write([]byte("cant get order because of error: " + err.Error()))
	}
	convertedUser, err := utils.StructToJson(user)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(convertedUser))
}
