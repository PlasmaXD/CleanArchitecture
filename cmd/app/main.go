package main

import (
	"github.com/PlasmaXD/CleanArchitecture/internal/infrastructure/persistence"
	"github.com/PlasmaXD/CleanArchitecture/internal/interface/web"
	"github.com/PlasmaXD/CleanArchitecture/internal/usecase"

	"github.com/gin-gonic/gin"
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
