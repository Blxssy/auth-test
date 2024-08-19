package router

import (
	"github.com/Blxssy/auth-test/internal/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(g *gin.Engine) {
	g.GET("/", getTokens)
}

func getTokens(c *gin.Context) {
	uid := c.Query("uid")

	accessToken, refreshToken, err := token.GetNewTokens(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
