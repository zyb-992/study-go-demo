package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const serverSecretKey = "zhengyb"

type TestClaim struct {
	Uid      string `json:"uid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken() (token string, err error) {
	tc := TestClaim{
		"uid-1",
		"zhengyb-1",
		"email-1",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "zhengyb",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)
	token, err = tokenClaims.SignedString([]byte(serverSecretKey))
	if err != nil {
		log.Printf("sign token error:%w", err)
		return "", nil
	}

	return token, nil
}

func ParseToken(token string) *TestClaim {
	var claim TestClaim
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(serverSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &claim, keyFunc)
	if err != nil {
		// FIXME
	}

	if parsedClaim, ok := jwtToken.Claims.(TestClaim); ok && jwtToken.Valid {
		return &TestClaim{
			parsedClaim.Uid,
			parsedClaim.UserName,
			parsedClaim.Email,
			parsedClaim.StandardClaims,
		}
	}

	return nil
}
