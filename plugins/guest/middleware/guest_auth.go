package middleware

import (
	"net/http"
	"strings"

	guestutils "gin-boilerplate/plugins/guest/utils"
	"github.com/gin-gonic/gin"
)

// GuestAuth JWT authentication middleware for guest users
func GuestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization from Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Verify Bearer Token format
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Authorization header format must be Bearer {token}",
			})
			c.Abort()
			return
		}

		// Parse guest token using guest-specific JWT parser
		claims, err := guestutils.ParseGuestToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Invalid or expired guest token",
			})
			c.Abort()
			return
		}

		// Token is validated as guest token by ParseGuestToken
		// No need for additional verification

		// Save guest information to context
		c.Set("guest_id", claims.GuestID)
		c.Set("guest_uid", claims.GuestUID)

		c.Next()
	}
}
