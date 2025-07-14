package validator

import (
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/utils"
	"strings"
)

func ValidateUserRegister(user domain.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return utils.NewHttpError(400, "Name, Email, and Password are required")
	}

	if !strings.Contains(user.Email, "@") {
		return utils.NewHttpError(400, "Invalid email format")
	}

	if len(user.Password) < 6 {
		return utils.NewHttpError(400, "Password must be at least 6 characters long")
	}

	if user.Level == "" {
		return utils.NewHttpError(400, "User level is required")
	}

	validLevels := map[string]bool{
		"basic": true,
		"intermediate": true,
		"advanced": true,
	}

	if !validLevels[user.Level] {
		return utils.NewHttpError(400, "Invalid user level, must be one of: basic, intermediate, advanced")
	}
	return nil
}