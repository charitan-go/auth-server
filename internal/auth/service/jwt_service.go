package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	SignToken(authModel *model.Auth) (string, error)
	UpdatePrivateKey(privateKeyStr string) error
}

type jwtServiceImpl struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey

	jwtExpirationDuration time.Duration
	jwtIssuer             string
}

type JwtClaims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewJwtService() JwtService {
	jwtService := &jwtServiceImpl{}

	// Read config
	jwtService.readConfig()

	return jwtService
}

func (s *jwtServiceImpl) readConfig() {
	// Read JWT_EXPIRATION_DURATION
	jwtExpirationDurationStr := os.Getenv("JWT_EXPIRATION_DURATION")
	if jwtExpirationDurationStr == "" {
		log.Fatalln("Error in reading JWT_EXPIRATION_DURATION")
	}

	if jwtExpirationDuration, err := time.ParseDuration(jwtExpirationDurationStr); err != nil {
		log.Fatalln("Error in parsing jwt expiration duration")
	} else {
		s.jwtExpirationDuration = jwtExpirationDuration
	}

	// Read JWT_ISSUER
	s.jwtIssuer = os.Getenv("JWT_ISSUER")
	if s.jwtIssuer == "" {
		s.jwtIssuer = "charitan-go"
	}
}

func (s *jwtServiceImpl) SignToken(authModel *model.Auth) (string, error) {
	// Read expiration time
	expirationTime := time.Now().Add(s.jwtExpirationDuration)

	// Get claims
	claims := &JwtClaims{
		Sub:  authModel.ReadableId.String(),
		Role: string(authModel.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.jwtIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *jwtServiceImpl) UpdatePrivateKey(keyStr string) error {
	// Decode the PEM block from the string.
	block, _ := pem.Decode([]byte(keyStr))
	if block == nil {
		return errors.New("failed to decode PEM block containing the key")
	}

	var rsaKey *rsa.PrivateKey
	var err error

	switch block.Type {
	case "RSA PRIVATE KEY":
		// Parse the key in PKCS#1 format.
		rsaKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return err
		}
	case "PRIVATE KEY":
		// Parse the key in PKCS#8 format.
		var key interface{}
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return err
		}
		var ok bool
		rsaKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return errors.New("parsed key is not an RSA private key")
		}
	default:
		return fmt.Errorf("unsupported key type %q", block.Type)
	}

	s.privateKey = rsaKey

	return nil

}
