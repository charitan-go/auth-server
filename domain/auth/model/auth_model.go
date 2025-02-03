package model

import (
	"time"

	"github.com/charitan-go/auth-server/domain/auth/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	id                uuid.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ReadableId        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid()" json:"readable_id"`
	Email             string       `gorm:"type:varchar(255);not null;unique" json:"email"`
	HashedPassword    string       `gorm:"type:varchar(255);not null" json:"hashed_password"`
	Role              dto.RoleEnum `gorm:"type:varchar(20);not null" json:"role"`
	ProfileReadableId uuid.UUID    `gorm:"type:uuid;not null" json:"profile_readable_id"`
	createdAt         time.Time
	updatedAt         time.Time
}

func (a *Auth) BeforeCreate(db *gorm.DB) (err error) {
	if a.id == uuid.Nil {
		a.id = uuid.New()
	}

	if a.ReadableId == uuid.Nil {
		a.ReadableId = uuid.New()
	}

	return nil
}

func NewAuth(req dto.RegisterDonorRequestDto, hashedPassword string, role dto.RoleEnum, profileReableId uuid.UUID) *Auth {
	return &Auth{
		Email:             req.Email,
		HashedPassword:    hashedPassword,
		Role:              role,
		ProfileReadableId: profileReableId,
	}
}
