package apis

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(os.Getenv("JWT_SECRET")),
	}
}

type CustomClaims struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func (j *JWT) CreateJWTTokenByUser(u *model.User) (string, error) {
	if u.Id == 0 {
		return "", fmt.Errorf("invalid user")
	}

	claims := CustomClaims{
		u.Id,
		u.Username,
		u.Email,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().UTC().Add(time.Hour * 24 * 7).Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func JTWMiddware(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("authorization")

	if authorizationHeader == "" {
		c.AbortWithStatus(401)
		return
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	claims, err := NewJWT().ParserToken(token)

	if err != nil {
		c.AbortWithStatus(401)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	j, _ := json.Marshal(claims)

	if claims.UserId == 0 {
		c.AbortWithStatus(401)
		c.Writer.Write(append(j, []byte("invalid token")...))
		return
	}

	c.Set("user_id", claims.UserId)
}
