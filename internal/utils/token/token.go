package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const accessTokenDuration = time.Minute * 15
const refreshTokenDuration = time.Hour * 24 * 7

var jwtKey []byte

func InitJWTKey() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func GetNewTokens(uid string, uip string) (string, string, error) {
	accessToken, err := generateToken(uid, uip, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(uid, uip, refreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func RefreshTokens(refreshToken string) (string, string, error) {
	if err := verifyToken(refreshToken); err != nil {
		return "", "", err
	}
	return "1", "", nil
}

func generateToken(uid string, uip string, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["ip"] = uip

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken() {

}

func generateRefreshTokenDB(refreshToken string) (string, error) {
	return refreshToken, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
