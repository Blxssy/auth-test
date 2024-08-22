package handlers

import (
	"github.com/Blxssy/auth-test/internal/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTokens(c *gin.Context) {
	uid := c.Param("id")
	clientIP := c.ClientIP()

	accessToken, refreshToken, err := token.GetNewTokens(uid, clientIP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshTokens(c *gin.Context) {
	uid := c.Query("uid")

	refreshToken := c.Query("refresh_token")

	accessToken, refreshToken, err := token.RefreshTokens(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
