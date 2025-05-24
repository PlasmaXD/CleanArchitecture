
# Go Clean Architecture サンプルアプリ (TODO)

以下は、Gin を用いたシンプルな TODO 管理アプリのディレクトリ構成とコード例です。

```
.
├── cmd
│   └── app
│       └── main.go
└── internal
    ├── domain
    │   ├── todo.go
    │   └── todo_repository.go
    ├── usecase
    │   └── todo_usecase.go
    ├── interface
    │   └── web
    │       ├── handler
    │       │   └── todo_handler.go
    │       └── router.go
    └── infrastructure
        └── persistence
            └── todo_repository.go
```

---

// cmd/app/main.go
```go
package main

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/infrastructure/persistence"
    "myapp/internal/interface/web"
    "myapp/internal/usecase"
)

func main() {
    r := gin.Default()

    // インフラ層 (Repository)
    todoRepo := persistence.NewTodoRepository()

    // ユースケース層
    todoUC := usecase.NewTodoUseCase(todoRepo)

    // ハンドラ層
    handler := web.NewTodoHandler(todoUC)

    // ルータ登録
    web.RegisterRoutes(r, handler)

    r.Run(":8080")
}
```

---

// internal/domain/todo.go
```go
package domain

type Todo struct {
    ID    int64  `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
```

// internal/domain/todo_repository.go
```go
package domain

type TodoRepository interface {
    Create(todo *Todo) error
    GetAll() ([]*Todo, error)
}
```

---

// internal/usecase/todo_usecase.go
```go
package usecase

import "myapp/internal/domain"

type TodoUseCase interface {
    CreateTodo(title string) (*domain.Todo, error)
    ListTodos() ([]*domain.Todo, error)
}

type todoUseCase struct {
    repo domain.TodoRepository
}

func NewTodoUseCase(r domain.TodoRepository) TodoUseCase {
    return &todoUseCase{repo: r}
}

func (u *todoUseCase) CreateTodo(title string) (*domain.Todo, error) {
    todo := &domain.Todo{Title: title}
    if err := u.repo.Create(todo); err != nil {
        return nil, err
    }
    return todo, nil
}

func (u *todoUseCase) ListTodos() ([]*domain.Todo, error) {
    return u.repo.GetAll()
}
```

---

// internal/interface/web/handler/todo_handler.go
```go
package handler

import (
    "net/http"
    "myapp/internal/usecase"
    "github.com/gin-gonic/gin"
)

type TodoHandler struct {
    uc usecase.TodoUseCase
}

func NewTodoHandler(uc usecase.TodoUseCase) *TodoHandler {
    return &TodoHandler{uc: uc}
}

func (h *TodoHandler) Create(c *gin.Context) {
    var req struct{ Title string `json:"title"` }
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
```

// internal/interface/web/router.go
```go
package web

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/interface/web/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.TodoHandler) {
    api := r.Group("/api")
    {
        api.POST("/todos", h.Create)
        api.GET("/todos", h.List)
    }
}
```

---

// internal/infrastructure/persistence/todo_repository.go
```go
package persistence

import (
    "errors"
    "myapp/internal/domain"
)

type todoRepository struct {
    store  []*domain.Todo
    nextID int64
}

func NewTodoRepository() domain.TodoRepository {
    return &todoRepository{store: make([]*domain.Todo, 0), nextID: 1}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
    if todo.Title == "" {
        return errors.New("title is empty")
    }
    todo.ID = r.nextID
    r.nextID++
    r.store = append(r.store, todo)
    return nil
}

func (r *todoRepository) GetAll() ([]*domain.Todo, error) {
    return r.store, nil
}
```

---

この構成をベースに、データベース接続やミドルウェアの追加、テストコードの整備などを行ってください。必要に応じてフィードバックをお願いいたします。
