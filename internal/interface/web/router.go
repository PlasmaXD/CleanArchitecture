package web

import (
	"github.com/PlasmaXD/CleanArchitecture/internal/interface/web/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handler.TodoHandler) {
	api := r.Group("/api")
	{
		api.POST("/todos", h.Create)
		api.GET("/todos", h.List)
	}
}
