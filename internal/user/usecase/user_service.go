package usecase

import (
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/utils"
	"errors"
	"github.com/google/uuid"
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

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.NewHttpError(500, "Failed to hash password")
	}
	user.Password = hashed

	user.ID = uuid.New().String()

	return s.Repo.Create(user)
}

func (s *UserService) Login(email, password string) (*domain.User, error) {
	user, err := s.Repo.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewHttpError(404, "User not found")
		}
		return nil, utils.NewHttpError(500, "Something went wrong")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, utils.NewHttpError(401, "Invalid password")
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.Repo.GetAll()
}
