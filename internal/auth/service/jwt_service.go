package service

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	SignToken(authModel *model.Auth) (string, error)
}

type jwtServiceImpl struct {
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	jwtExpirationDuration time.Duration
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

	// Gen key pair
	jwtService.generateRSAKeyPair(2048)

	return jwtService
}

func (s *jwtServiceImpl) readConfig() {
	jwtExpirationDurationStr := os.Getenv("JWT_EXPIRATION_DURATION")
	if jwtExpirationDurationStr == "" {
		log.Fatalln("Error in reading JWT_EXPIRATION_DURATION")
	}

	if jwtExpirationDuration, err := time.ParseDuration(jwtExpirationDurationStr); err != nil {
		log.Fatalln("Error in parsing jwt expiration duration")
	} else {
		s.jwtExpirationDuration = jwtExpirationDuration
	}
}

func (s *jwtServiceImpl) generateRSAKeyPair(bits int) {
	// TODO: Impl
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	s.privateKey = privateKey
	s.publicKey = &(privateKey.PublicKey)
}

func (s *jwtServiceImpl) SignToken(authModel *model.Auth) (string, error) {
	// TODO: Impl
	expirationTime := time.Now().Add(os.Getenv("JWT_EXPIRATION"))
	return "", nil

}
