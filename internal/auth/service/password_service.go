package service

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashPassword, password string) bool
}

type passwordServiceImpl struct {
}

func NewPasswordService() PasswordService {
	return &passwordServiceImpl{}
}

func (s *passwordServiceImpl) HashPassword(password string) (string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(hashedPasswordByte)
	return hashedPassword, nil
}

func (s *passwordServiceImpl) VerifyPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
