package usecase

import (
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/utils"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	Repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(user *domain.User) error {
	existingUser, err := s.Repo.FindByEmail(user.Email)
	if existingUser.ID != "" {
		return utils.NewHttpError(400, "Email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewHttpError(500, "Something went wrong")
	}

	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.Repo.GetAll()
}
