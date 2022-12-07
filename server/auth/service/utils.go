package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/config"
	"gitlab.cs.washington.edu/assafv/fs3/server/shared/jwtutils"
)

func GetKeyFromUsername(username string) string {
	usernameHash := sha256.Sum256([]byte(username))
	usernameHashEncoded := hex.EncodeToString(usernameHash[:])
	return usernameHashEncoded
}

func GetToken(username string) (string, error) {
	key, err := config.GetJwtPrivateKey()
	if err != nil {
		return "", errors.New("internal error: could not load key")
	}
	claims := &jwtutils.Fs3JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "fs3-authservice",
			Audience:  jwt.ClaimStrings{"fs3-primary"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(key)
}
