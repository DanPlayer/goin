package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"goin/conf"
	"goin/utils"
	"time"
)

// MakeMemberToken 生产token密钥
func MakeMemberToken(id string, expired time.Time) (tokenStr string, err error) {
	config := conf.GetConf()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"expired": expired.Unix(),
		"issuer":  config.AppName,
	})
	tokenStr, err = token.SignedString([]byte(config.Auth.MemberSecret))
	return
}

// ParseMemberToken 使用jwt验证时使用
func ParseMemberToken(tokenStr, secret string) (id string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = claims["id"].(string)
	}
	return
}

func MakeAdminToken(id string, role string, expired time.Time) (tokenStr string, err error) {
	config := conf.GetConf()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"role":    role,
		"expired": expired.Unix(),
		"issuer":  config.AppName,
	})
	tokenStr, err = token.SignedString([]byte(config.Auth.AdmSecret))
	return
}

func ParseAdminToken(tokenStr, secret string) (id, role string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = claims["id"].(string)
		role = claims["role"].(string)
	}
	return
}

// GenSession ruiz框架的后台session
func GenSession() string {
	timeStr := time.Now().Format("060102150405")
	ruizAdminAccessKey := "SmeZjs212c6uWbpOLKm6AZyfp51Efs-pDHC45nxH"
	return timeStr + "_" + utils.Md5hex(timeStr+ruizAdminAccessKey)
}
