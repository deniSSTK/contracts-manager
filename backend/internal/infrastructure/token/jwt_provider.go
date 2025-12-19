package token

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/config"
	"contracts-manager/internal/infrastructure/logger"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
)

type JWTProvider struct {
	secretKey string
	log       *logger.Logger
}

func NewJWTProvider(cfg *config.Config, log *logger.Logger) *JWTProvider {
	return &JWTProvider{secretKey: cfg.JWTSecret, log: log}
}

func (p *JWTProvider) generateToken(userID uuid.UUID, duration time.Duration, tokenType string) (string, int64, error) {
	expiration := time.Now().Add(duration).Unix()

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    expiration,
		"type":   tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(p.secretKey))
	if err != nil {
		p.log.Errorf(ErrFailedToSignToken(tokenType), err)
		return "", 0, err
	}

	return signedToken, expiration, nil
}

func (p *JWTProvider) GenerateAccessToken(userID uuid.UUID) (auth.AuthResponse, error) {
	token, exp, err := p.generateToken(userID, 15*time.Minute, "refresh")
	if err != nil {
		return auth.AuthResponse{}, err
	}

	return auth.AuthResponse{
		AccessToken: token,
		Exp:         exp,
	}, nil
}

func (p *JWTProvider) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	token, _, err := p.generateToken(userID, 7*24*time.Hour, "refresh")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (p *JWTProvider) ValidateToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(p.secretKey), nil
	})

	if err != nil {
		p.log.Errorf(ErrFailedToParseToken, err)
		return false
	}

	return token.Valid
}

func (p *JWTProvider) ParseUserID(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(p.secretKey), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, ErrFailedToParseToken
	}

	return auth.ParseUserIDFromJWTClaims(token.Claims)
}

func (p *JWTProvider) RefreshAccessToken(token string) (auth.AuthResponse, error) {
	userID, err := p.ParseUserID(token)
	if err != nil {
		return auth.AuthResponse{}, err
	}

	return p.GenerateAccessToken(userID)
}
