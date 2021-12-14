package user

import "github.com/hyperyuri/webapi-with-go/database/entity"

type IUserService interface {
	GetUsers() (*[]entity.User, error)
	GetUser(userId int64) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) error
}

type ServiceUser struct {
	repository RepositoryUser
}

func NewServiceUser(repository RepositoryUser) IUserService {
	return &ServiceUser{repository: repository}
}

func (s *ServiceUser) GetUsers() (*[]entity.User, error) {
	var users []entity.User
	err := s.repository.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *ServiceUser) GetUser(userId int64) (*entity.User, error) {
	var user entity.User
	err := s.repository.db.Where("id =? ", userId).Find(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *ServiceUser) CreateUser(user *entity.User) (*entity.User, error) {
	err := s.repository.db.Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *ServiceUser) UpdateUser(user *entity.User) (*entity.User, error) {
	err := s.repository.db.Save(&user).Error

	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *ServiceUser) DeleteUser(user *entity.User) error {
	err := s.repository.db.Delete(&user).Error

	if err != nil {
		return err
	}
	return nil
}
