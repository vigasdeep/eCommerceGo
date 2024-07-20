package middleware

import (
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"ecommerce-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = config.JWT_SECRET

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        
        authHeader := c.GetHeader("Authorization")
        
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or improperly formatted"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

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
        var user models.User
        if err := config.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        // Store the user model in Gin context
        c.Set("userID", user.ID)
        c.Set("email", claims.Email)
        c.Next()
    }
}
