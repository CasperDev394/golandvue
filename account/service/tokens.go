package service

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/CasperDev394/golandvue/account/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type IDTokenCustomClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(u *model.User, key *rsa.PrivateKey, exp int64) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + exp

	claims := IDTokenCustomClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		log.Println("Failed to sign id token string")
		return "", err
	}

	return ss, nil
}

type RefreshToken struct {
	SS         string
	ID         string
	ExpirensIn time.Duration
}

type RefreshTokenCastomClaims struct {
	UID uuid.UUID `json:"uid"`
	jwt.StandardClaims
}

func generateRefreshToken(uid uuid.UUID, key string, exp int64) (*RefreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(time.Duration(exp) * time.Second)
	tokenID, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate refresh token ID")
		return nil, err
	}

	claims := RefreshTokenCastomClaims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("Failed to sign refresh token string")
		return nil, err
	}

	return &RefreshToken{
		SS:         ss,
		ID:         tokenID.String(),
		ExpirensIn: tokenExp.Sub(currentTime),
	}, nil
}
