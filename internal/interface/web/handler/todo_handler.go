package handler

import (
	"myapp/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	uc usecase.TodoUseCase
}

func NewTodoHandler(uc usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := h.uc.CreateTodo(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) List(c *gin.Context) {
	todos, err := h.uc.ListTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
