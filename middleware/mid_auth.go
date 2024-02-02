package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("X-API-Key")
		config := &oauth2.Config{}
		token := &oauth2.Token{AccessToken: accessToken}
		client := config.Client(context.Background(), token)

		response, err := client.Get("https://www.googleapis.com/oauth2/v3/tokeninfo")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 400,
				"msg":  "Connection error to Google",
			})

			c.Abort()
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Invalid token",
			})

			c.Abort()
			return
		}

		var tokenInfo map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&tokenInfo); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Forged token",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
