package service

import (
	"database/sql"
	"job-app/internal/models"
	"job-app/internal/repository"
	"job-app/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)


func RegisterUser(db *sql.DB,user *models.User) error {
	// Implement registration logic here
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil
	}
	user.Password = hashedPassword
	return repository.CreateUser(db, user)
}

func AuthenticateUser(db *sql.DB, login *models.Login) (string, error) {
	user, err := repository.GetUserByUserName(db, login.Username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", err
	}

	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}

func ForgotPassword(db *sql.DB, username string) (string, error) {
	user, err := repository.GetUserByUserName(db, username)
	if err != nil {
		return "", err
	}

	generatedPassword := utils.GenerateFromPassword(6)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)

	if err := repository.UpdateUserPassword(db, user); err != nil {
		return "", err
	}
	return generatedPassword, nil
}