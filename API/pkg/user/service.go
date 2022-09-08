package user

import (
	"github.com/MrzBldk/User-API/api/presenter"
	"github.com/MrzBldk/User-API/pkg/entities"
)

type Service interface {
	InsertUser(User *entities.User) (*entities.User, error)
	FetchUser(Id string) (*presenter.User, error)
	FetchUsers() (*[]presenter.User, error)
	UpdateUser(User *entities.User) error
	RemoveUser(Id string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(user)
}

func (s *service) FetchUser(id string) (*presenter.User, error) {
	return s.repository.ReadUser(id)
}

func (s *service) FetchUsers() (*[]presenter.User, error) {
	return s.repository.ReadUsers()
}

func (s *service) UpdateUser(user *entities.User) error {
	return s.repository.UpdateUser(user)
}

func (s *service) RemoveUser(id string) error {
	return s.repository.DeleteUser(id)
}
