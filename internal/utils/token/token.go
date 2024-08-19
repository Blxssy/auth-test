package token

import (
	"os"
	"time"
)

const accessTokenDuration = time.Minute * 15
const refreshTokenDuration = time.Hour * 24 * 7

var jwtKey []byte

func InitJWTKey() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func GetNewTokens(uid string) (string, string, error) {
	return "", "", nil
}

func RefreshTokens() {

}
