package auth

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/config"
	"birthdays/internal/pkg/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthManager struct {
	secret   []byte
	lifetime time.Duration
}

type jwtClaim struct {
	ID       uint64
	Username string
	jwt.RegisteredClaims
}

func NewManager(config *config.JWTConfig) *AuthManager {
	return &AuthManager{secret: []byte(config.Secret), lifetime: config.Lifetime}
}

func (am *AuthManager) GenerateToken(user *dto.TokenInfo) (string, error) {
	expiresAt := time.Now().Add(am.lifetime)
	claims := &jwtClaim{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(am.secret)
}

func (am *AuthManager) ValidateToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&jwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(am.secret), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(*jwtClaim)
	if !ok {
		return apperrors.ErrCouldNotParseClaims
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		return apperrors.ErrTokenExpired
	}
	if claims.IssuedAt.After(time.Now().Local()) {
		return apperrors.ErrInvalidIssuedTime
	}

	return nil
}
