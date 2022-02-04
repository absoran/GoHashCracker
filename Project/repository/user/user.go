package user

import (
	"github.com/absoran/goproject/models"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/postgres"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) CreateUser(user *models.User) (string, error) {
	if err := r.db.Create(user).Error; err != nil {
		return "error on creating user", err
	}
	return user.ID, nil
}

func (r *Repo) GetUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repo) GetUser(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) UpdateUser(id string, user *models.User) error {
	if err := r.db.Model(&models.User{}).Where("id = ?", id).Update(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteUser(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
