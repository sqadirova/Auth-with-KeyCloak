package auth

import (
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
}

type AuthRepository struct {
	connection *gorm.DB
}
