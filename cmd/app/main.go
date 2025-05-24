package main

import (
	"github.com/PlasmaXD/CleanArchitecture/internal/infrastructure/persistence"
	"github.com/PlasmaXD/CleanArchitecture/internal/interface/web"
	"github.com/PlasmaXD/CleanArchitecture/internal/interface/web/handler"
	"github.com/PlasmaXD/CleanArchitecture/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 永続化層（リポジトリ）の初期化
	repo := persistence.NewTodoRepository()

	// 2. ユースケース層の初期化
	uc := usecase.NewTodoUseCase(repo)

	// 3. ハンドラの生成
	todoHandler := handler.NewTodoHandler(uc)

	// 4. Gin エンジン生成＆ルート登録
	r := gin.Default()
	web.RegisterRoutes(r, todoHandler)

	// 5. サーバ起動
	r.Run(":8080")
}
