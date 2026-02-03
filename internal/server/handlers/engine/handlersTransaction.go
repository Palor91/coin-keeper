package handlers

import (
	"coin-keeper/internal/server/http/requests"
	"coin-keeper/internal/server/http/requests/filters"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *HandlersKeeper) HandleCreateTransaction(resp http.ResponseWriter, req *http.Request) {
	transaction := requests.Transaction{}

	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.dbEngine.CreateTransaction(req.Context(), transaction)
}

func (h *HandlersKeeper) HandleReadTransaction(resp http.ResponseWriter, req *http.Request) {

	request := req.PathValue("id")
	if request != "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(request)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.dbEngine.ReadTransaction(req.Context(), id)
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(resp).Encode(user)
	if err != nil {
		fmt.Printf("error while encoding user response: %v\n", err)
	}
}

func (h *HandlersKeeper) HandleUpdateTransaction(resp http.ResponseWriter, req *http.Request) {

	transaction := requests.Transaction{}
	err := json.NewDecoder(req.Body).Decode(&transaction)
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.dbEngine.UpdateTransaction(req.Context(), transaction)
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h *HandlersKeeper) HandleDeleteTransaction(resp http.ResponseWriter, req *http.Request) {

	request := req.PathValue("id")
	if request != "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	idStr := req.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid id"))
		return
	}

	err = h.dbEngine.DeleteTransaction(req.Context(), id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h *HandlersKeeper) GetTrasactionByOption(resp http.ResponseWriter, req *http.Request) {
	// 1. Читаем GET‑параметры
	transactionFilter := filters.NewTransactionFilter(req)
	// 2. Вызываем метод из слоя требований
	result, err := h.dbEngine.GetTransactionByOption(req.Context(), transactionFilter)
	if err != nil {
		// 3. В случае ошибки отдаем 500
		http.Error(resp, err.Error(), 500)
		return
	}

	// 6. Отдаём JSON
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(result)
}
