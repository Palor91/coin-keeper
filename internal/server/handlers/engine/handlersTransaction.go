package handlers

import (
	"coin-keeper/internal/server/http/requests"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	q := req.URL.Query()

	userID := q.Get("user_id")
	minAmount := q.Get("min_amount")
	maxAmount := q.Get("max_amount")
	dateFrom := q.Get("date_from")
	dateTo := q.Get("date_to")

	// 2. Готовим динамические фильтры
	filters := []string{}
	args := []interface{}{}
	i := 1

	if userID != "" {
		filters = append(filters, fmt.Sprintf("user_id = $%d", i))
		args = append(args, userID)
		i++
	}

	if minAmount != "" {
		filters = append(filters, fmt.Sprintf("amount >= $%d", i))
		args = append(args, minAmount)
		i++
	}

	if maxAmount != "" {
		filters = append(filters, fmt.Sprintf("amount <= $%d", i))
		args = append(args, maxAmount)
		i++
	}

	if dateFrom != "" {
		filters = append(filters, fmt.Sprintf("date >= $%d", i))
		args = append(args, dateFrom)
		i++
	}

	if dateTo != "" {
		filters = append(filters, fmt.Sprintf("date <= $%d", i))
		args = append(args, dateTo)
		i++
	}

	// 3. Собираем SQL‑запрос
	query := `
        SELECT id, user_id, description, amount, date
        FROM transactions
    `
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	// 4. Выполняем запрос
	rows, err := h.db.Query(query, args...)
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}
	defer rows.Close()

	// 5. Сканируем результат
	var result []Transactions

	for rows.Next() {
		var t Transactions
		if err := rows.Scan(&t.ID, &t.User_id, &t.Description, &t.Amount, &t.Date); err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		result = append(result, t)
	}

	// 6. Отдаём JSON
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(result)
}
