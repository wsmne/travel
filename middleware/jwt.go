package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("wsm_graduation_project")

func GenerateJWT(ID uint) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(12 * time.Hour)

	// 创建 JWT Claims
	claims := &Claims{
		ID: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// 生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 检查 token 是否有效
	if !token.Valid {
		return nil, fmt.Errorf("无效的 token")
	}

	// 检查 token 是否过期
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token 已过期")
	}

	return claims, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 Token"})
			c.Abort()
			return
		}

		//// 按空格分割 "Bearer <token>"
		//parts := strings.Split(authHeader, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token 格式无效"})
		//	c.Abort()
		//	return
		//}

		// 解析 JWT
		claims, err := ParseJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token 无效或已过期"})
			c.Abort()
			return
		}

		// 将解析出的用户名存入上下文
		c.Set("userid", claims.ID)
		c.Next()
	}
}
