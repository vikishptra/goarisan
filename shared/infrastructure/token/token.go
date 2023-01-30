package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/util"
)

func GenerateToken(user_id vo.UserID) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["uuid"] = util.GenerateID()
	claims["exp"] = time.Now().Add(time.Second * 50).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func RefreshToken(user_id vo.UserID) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["uuid"] = util.GenerateID()
	claims["exp"] = time.Now().Add(time.Second * 50).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorenum.GabisaAksesBro
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}

	return nil
}
func ExtractToken(c *gin.Context) string {
	token := c.Query("access_token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (vo.UserID, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorenum.GabisaAksesBro
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := fmt.Sprintf("%s", claims["user_id"])

		if err != nil {
			return "", err
		}
		return vo.UserID(uid), nil
	}
	return "", nil
}

func ExtractTokenCookie(c *gin.Context) string {
	tokenString, _ := c.Cookie("refresh_token")
	return tokenString
}
func ExtractTokenIDCookies(c *gin.Context) (string, error) {

	tokenString := ExtractTokenCookie(c)
	if tokenString == "" {
		return "", errorenum.GabisaAksesBro
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorenum.GabisaAksesBro
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := fmt.Sprintf("%s", claims["user_id"])

		return uid, nil
	}

	return "", nil
}
