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
    "github.com/PlasmaXD/CleanArchitecture/internal/infrastructure/persistence"
    "github.com/PlasmaXD/CleanArchitecture/internal/interface/web"
    "github.com/PlasmaXD/CleanArchitecture/internal/usecase"
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

import (
    "github.com/PlasmaXD/CleanArchitecture/internal/domain"
)

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
    "github.com/PlasmaXD/CleanArchitecture/internal/usecase"
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
    "github.com/PlasmaXD/CleanArchitecture/internal/interface/web/handler"
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
    "github.com/PlasmaXD/CleanArchitecture/internal/domain"
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


以下の手順で起動できます。※Go 1.18 以上がインストールされている前提です。

```bash
# 1. リポジトリをクローン（なければローカルにディレクトリを用意）
git clone https://github.com/yourname/clean-arch-go-app.git
cd clean-arch-go-app

# 2. Go モジュールを初期化／依存取得
go mod tidy

# 3. アプリを「起動」または「ビルド＆実行」
#   a) そのまま起動する場合
go run ./cmd/app

#   b) ビルドして実行ファイルを作る場合
go build -o todo-app ./cmd/app
./todo-app
```

起動に成功するとデフォルトで `:8080` 番号の HTTP サーバが立ち上がります。

---

### 動作確認例

1. **TODO 作成**

```bash
curl -X POST http://localhost:8080/api/todos \
     -H "Content-Type: application/json" \
     -d '{"title":"最初のタスク"}'
```

2. **TODO 一覧取得**

```bash
curl http://localhost:8080/api/todos
```

以上で、Clean Architecture 構成の TODO アプリが動作するはずです。
もしポート番号やログ出力、DB 接続先を変えたい場合は `cmd/app/main.go` や環境変数で調整してください。
