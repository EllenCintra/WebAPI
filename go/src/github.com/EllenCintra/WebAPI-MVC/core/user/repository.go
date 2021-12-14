package user

import (
	"gorm.io/gorm"
)

/*
type IUserRepository interface {
	GetUsers() (*[]entity.User, error)
	GetUser(userId int64) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
}
*/
type RepositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) RepositoryUser {
	return RepositoryUser{db: db}
}

/*
func (r *RepositoryUser) GetUsers() (*[]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error

	return &users, err
}

func (r *RepositoryUser) GetUser(userId int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id =? ", userId).Find(&user).Error
	return &user, err
}

func (r *RepositoryUser) CreateUser(user *entity.User) (*entity.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *RepositoryUser) UpdateUser(user *entity.User) (*entity.User, error) {
	return user, r.db.Save(&user).Error
}

func (r *RepositoryUser) DeleteUser(user *entity.User) error {
	return r.db.Delete(&user).Error
}
*/
