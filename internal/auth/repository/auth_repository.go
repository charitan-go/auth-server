package repository

import (
	"fmt"

	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/charitan-go/auth-server/pkg/database"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOneByEmail(email string) (*model.Auth, error)
	Save(authModel *model.Auth) (*model.Auth, error)
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

func (r *authRepositoryImpl) FindOneByEmail(email string) (*model.Auth, error) {
	var auth model.Auth

	result := r.db.Where("email = ?", email).First(&auth)
	if result.Error != nil {
		return nil, result.Error
	}

	return &auth, nil
}

func (r *authRepositoryImpl) Save(authModel *model.Auth) (*model.Auth, error) {
	result := r.db.Create(authModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return authModel, nil
}
