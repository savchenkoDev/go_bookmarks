package repository

import (
	"bookmarks/internal/models"

	"gorm.io/gorm"	
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int64) (models.User, error) {
	var u models.User
	err := r.db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var u models.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Create(u models.UserRequest) (models.User, error) {
	var user models.User

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user = models.User{
		Email: u.Email,
		PasswordHash: string(passwordHash),
	}
	err = r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(id int64, u models.UserRequest) (models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(u).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id int64) error {
	return r.db.Delete(&models.User{}, id).Error
}
