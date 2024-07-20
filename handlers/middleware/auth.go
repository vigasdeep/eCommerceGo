package middleware

import (
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = config.JWT_SECRET

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        claims := &handlers.Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
                c.Abort()
                return
            }
            c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
            c.Abort()
            return
        }

        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("email", claims.Email)
        c.Next()
    }
}
