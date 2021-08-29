package jwt

import (
	UserModels "api-go/app/models/user"
	"api-go/lib/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte(config.AppConfig.JwtSecret)
var UserTokenKey = "UToken"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}

type Claims struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
	jwt.StandardClaims
}

// GenerateToken 生成 token
func GenerateToken(user *UserModels.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := Claims{
		user.ID,
		user.Name,
		user.Password,
		user.RoleId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppConfig.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// UserInfo 当前用户登录信息
func UserInfo(c *gin.Context) *User {
	token := c.Request.Header.Get("token")
	user := &User{}
	if token == "" {
		return user
	} else {
		claims, err := ParseToken(token)
		if err != nil {
			return user
		} else {
			user.ID = claims.ID
			user.Name = claims.Name
			user.RoleId = claims.RoleId
			return user
		}
	}
}

// UserId 当前用户 ID
func UserId(c *gin.Context) int64 {
	return UserInfo(c).ID
}
