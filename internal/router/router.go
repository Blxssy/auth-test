package router

import (
	"github.com/Blxssy/auth-test/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.POST("/auth/tokens/:id", handlers.GetTokens)
	g.POST("/auth/refresh", handlers.RefreshTokens)
}
