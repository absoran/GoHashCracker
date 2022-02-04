package repository

import (
	"github.com/absoran/goproject/models"
)

type UserRepository interface {
	CreateUser(costumer *models.User) (string, error)
	GetUsers() ([]*models.User, error)
	GetUser(customerId string) (*models.User, error)
	UpdateUser(customerId string, customer *models.User) error
	DeleteUser(id string) error
}
