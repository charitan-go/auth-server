package repository

import (
	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/charitan-go/auth-server/pkg/database"
	"gorm.io/gorm"
	"log"
)

type AuthRepository interface {
	Save(authModel *model.Auth) (*model.Auth, error)

	FindOneByEmail(email string) (*model.Auth, error)
	FindOneByReadableId(readableId string) (*model.Auth, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	db := database.DB
	if db == nil {
		log.Println("db is nil")
	} else {
		log.Println("db is not nil")
	}

	return &authRepositoryImpl{db: database.DB}
}

func (r *authRepositoryImpl) Save(authModel *model.Auth) (*model.Auth, error) {
	result := r.db.Create(authModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return authModel, nil
}

func (r *authRepositoryImpl) FindOneByEmail(email string) (*model.Auth, error) {
	var auth model.Auth

	result := r.db.Where("email = ?", email).First(&auth)
	if result.Error != nil {
		return nil, result.Error
	}

	return &auth, nil
}

// FindOneByReadableId implements AuthRepository.
func (r *authRepositoryImpl) FindOneByReadableId(readableId string) (*model.Auth, error) {
	var auth model.Auth

	result := r.db.Where("readableId = ?", readableId).First(&auth)
	if result.Error != nil {
		return nil, result.Error
	}

	return &auth, nil
}
