package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_auth_login": "error",
				"message":           "Akses gagal, Tidak ada Token",
			})
			return
		}

		Token := strings.TrimPrefix(authHeader, "Bearer ")

		setClaims, err := ValidasiJWT(Token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_auth_login": "error",
				"message":           err.Error(),
			})
			return
		}

		ctx.Set("username_sekarang", setClaims.Username)
		ctx.Set("IdUser_sekarang", setClaims.UserID)

		ctx.Next()
	}
}
