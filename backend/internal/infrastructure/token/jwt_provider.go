package token

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/config"
	"contracts-manager/internal/infrastructure/logger"
	"contracts-manager/internal/utils"
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

func (p *JWTProvider) generateToken(userID uuid.UUID, age int, tokenType string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Duration(age) * time.Second).Unix(),
		"type":   tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(p.secretKey))

	if err != nil {
		p.log.Errorf(ErrFailedToSignToken(tokenType), err)
		return "", err
	}

	return signedToken, nil
}

func (p *JWTProvider) GenerateAccessToken(userId uuid.UUID) (string, error) {
	return p.generateToken(userId, 15*60, "access")
}

func (p *JWTProvider) GenerateRefreshToken(userId uuid.UUID) (string, error) {
	return p.generateToken(userId, utils.Week, "refresh")
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

func (p *JWTProvider) RefreshAccessToken(token string) (string, error) {
	userID, err := p.ParseUserID(token)
	if err != nil {
		return "", err
	}

	return p.GenerateAccessToken(userID)
}
