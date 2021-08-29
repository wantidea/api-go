package jwt

import (
	SystemModels "api-go/app/models/system"
	UploadServices "api-go/app/services/upload"
	"api-go/lib/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtAdminSecret = []byte(config.AppConfig.JwtAdminSecret)
var AdminTokenKey = "AToken"

type Admin struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id"`
	AvatarUrl string `json:"avatar_url"`
}

type AdminClaims struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id"`
	AvatarUrl string `json:"avatar_url"`
	jwt.StandardClaims
}

// GenerateAdminToken 生成 token
func GenerateAdminToken(admin *SystemModels.Admin) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := AdminClaims{
		admin.ID,
		admin.Name,
		admin.Password,
		admin.RoleId,
		UploadServices.ImageUrl(admin.AvatarId),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppConfig.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	adminToken, err := tokenClaims.SignedString(jwtAdminSecret)

	return adminToken, err
}

// ParseAdminToken 解析 token
func ParseAdminToken(adminToken string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(adminToken, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtAdminSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// AdminInfo 当前管理员登录信息
func AdminInfo(c *gin.Context) *Admin {
	adminToken := c.Request.Header.Get(AdminTokenKey)
	admin := &Admin{}
	if adminToken == "" {
		return admin
	} else {
		claims, err := ParseAdminToken(adminToken)
		if err != nil {
			return admin
		} else {
			admin.ID = claims.ID
			admin.Name = claims.Name
			admin.RoleId = claims.RoleId
			admin.AvatarUrl = claims.AvatarUrl
			return admin
		}
	}
}

// AdminId 当前用户 ID
func AdminId(c *gin.Context) int64 {
	return AdminInfo(c).ID
}
