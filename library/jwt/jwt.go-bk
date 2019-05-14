// Package jwt This package further simplifies the use of jwt-go
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iphuket/pkt/library/crypto"
)

// NewToken Generating token for users
func NewToken(uuid, sub, ip, aeskey string, sec int64) (string, error) {
	// encrypt user info
	enuuid, err := crypto.EnSting([]byte(uuid), []byte(aeskey))
	if err != nil {
		return "aes error ", err
	}
	ensub, err := crypto.EnSting([]byte(sub), []byte(aeskey))
	if err != nil {
		return "aes error ", err
	}
	enip, err := crypto.EnSting([]byte(ip), []byte(aeskey))
	if err != nil {
		return "aes error ", err
	}
	mip := jwt.MapClaims{"ip": enip}
	fmt.Println("enip:", enip, "string:", mip["ip"])
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": enuuid,
		"sub":  ensub,
		"ip":   enip,
		"exp":  time.Now().Unix() + sec,
		"nbf":  time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(aeskey))
	if err != nil {
		return "Signing error ", err
	}
	return tokenString, nil
}

// CheckIP ip check
func CheckIP(ip, aeskey string, Token *jwt.Token) (bool, error) {
	enip, err := crypto.EnSting([]byte(ip), []byte(aeskey))
	if err != nil {
		return false, err
	}
	claims, ok := Token.Claims.(jwt.MapClaims)
	if ok {
		cip := claims["ip"]
		fmt.Println("enip:", cip, "string:", fmt.Sprint(enip))
		if cip != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New("claims false")
}

// Chcek chcek user login status and Get the UUID. auto chcek ip
func Chcek(ip, aeskey, token string) (string, error) {
	Token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(aeskey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
		bool, err := CheckIP(ip, aeskey, Token)
		if err != nil {
			return "", err
		}
		if bool {
			return claims["uuid"].(string), nil
		}
		return "", errors.New("ip check not pass")
	}
	return "", errors.New("claims exp, iat, nbf Not pass or ok")
}
