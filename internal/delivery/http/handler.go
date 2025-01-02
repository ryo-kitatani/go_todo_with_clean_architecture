package http

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/domain/models"
	"todo-api/internal/usecase"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	useCase usecase.TodoUseCase
}

func NewTodoHandler(useCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		useCase: useCase,
	}
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.useCase.GetTodos()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := h.useCase.GetTodo(params["id"])
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdTodo, err := h.useCase.CreateTodo(todo)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, createdTodo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedTodo, err := h.useCase.UpdateTodo(params["id"], todo)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, updatedTodo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := h.useCase.DeleteTodo(params["id"]); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
