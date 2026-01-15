package handlers

import (
	"coin-keeper/internal/server/handlers/requirements"
	"coin-keeper/internal/server/http/requests"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type HandlersKeeper struct {
	dbEngine requirements.RequiredDatabase
}

func NewHandlersKeeper(dbEngine requirements.RequiredDatabase) *HandlersKeeper {
	return &HandlersKeeper{
		dbEngine: dbEngine,
	}
}

func (h *HandlersKeeper) HandleCreateUser(resp http.ResponseWriter, req *http.Request) {
	// Создаем пустую структуру User (она отображает тело ожидаемого запроса)
	user := requests.User{}

	// Тут однострочник. Создаем новый декодировщик JSON на основе тела запроса
	// (NewDecoder это по сути конструктор объекта декодировщика) и сразу вызываем
	// метод Decode, которому нужна ссылка на объект в который он будет распаршивать тело
	err := json.NewDecoder(req.Body).Decode(&user)
	// Разумеется проверяем на ошибку
	if err != nil {
		// Если произошло говно, то у ошибки вызываем метод Error который вернет строку
		// описывающую ошибку и результат сразу преобразуем в байты, чтобы записать в тело ответа
		resp.Write([]byte(err.Error()))
		// После записи тела выставляем 400 код
		resp.WriteHeader(http.StatusBadRequest)
		// Ну и вызываем return так как дальнейшее выполнение функции не имеет смысла
		return
	}

	// У нашего объекта с хендлерами вызываем метод на создание пользователя
	err = h.dbEngine.CreateUser(req.Context(), user)
	// Ну и тоже проверяем на ошибку
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	// В случае успеха (если проверка выше не прошла) отправляем 201 код (объект создан)
	resp.WriteHeader(http.StatusCreated)
}

func (h *HandlersKeeper) HandleReadUser(resp http.ResponseWriter, req *http.Request) {

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

	user, err := h.dbEngine.ReadUser(req.Context(), id)
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

func (h *HandlersKeeper) HandleUpdateUser(resp http.ResponseWriter, req *http.Request) {

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

	err = h.dbEngine.UpdateUser(req.Context(), id)
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h *HandlersKeeper) HandleDeleteUser(resp http.ResponseWriter, req *http.Request) {

	request := req.PathValue("id")
	if request != "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	// не уверен нужно ли удалить по айди а не оп имени

	idStr := req.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("invalid id"))
		return
	}

	err = h.dbEngine.DeleteUser(req.Context(), id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.WriteHeader(http.StatusOK)
}
