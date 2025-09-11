package utils

import (
	"Golang-Redis-Gin/internal/utils/functions"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenModel struct {
	UserId   int
	UserType string
	Exp      time.Time
}


// HashPassword băm mật khẩu với bcrypt
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost = 10 (đủ an toàn, có thể tăng lên 12-14 nếu cần)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword so sánh password người dùng nhập với password đã hash
func CheckPassword(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
// Extract token
func ExtractToken(c *gin.Context) string {
	if len(c.Request.Header["Authorization"]) > 0 {
		bearerToken := c.Request.Header["Authorization"][0]
		if len(strings.Split(bearerToken, " ")) == 2 {
			return strings.Split(bearerToken, " ")[1]
		}
		return ""
	}
	return ""
}
func InTimeSpan(start, end, check time.Time) bool {
    return check.After(start) && check.Before(end)
}
// Get token info
func GetTokenInfo(c *gin.Context) (tokenModel TokenModel) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		tokenModel.UserId = 0
		tokenModel.UserType = ""
		tokenModel.Exp = functions.CurrentTime()
		return tokenModel
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := claims["user_id"]
	user_type := claims["user_type"]
	expInterface := claims["exp"]
	tokenModel.UserId = int(user_id.(float64))
	tokenModel.UserType = user_type.(string)
	expInt := int64(expInterface.(float64))
	exp := time.Unix(expInt, 0)
	tokenModel.Exp = exp
	return tokenModel
}

// Generate Token String
func GenerateTokenString(user_id int, user_type string) string {
	tokenModel := jwt.New(jwt.SigningMethodHS256)
	claims := tokenModel.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["user_type"] = user_type
	expHoursStr := os.Getenv("EXP_HOURS")
	expHours, err := strconv.Atoi(expHoursStr)
	if err != nil {
    // fallback nếu parse lỗi
    expHours = 1 // mặc định 1h
	}

	claims["exp"] = functions.CurrentTime().
    Add(time.Duration(expHours) * time.Hour).
    Unix()
	token, err := tokenModel.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		functions.ShowLog("GenerateTokenStringError", err.Error())
	}
	return token
}
