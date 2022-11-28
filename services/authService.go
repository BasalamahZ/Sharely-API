package services

import (
	"sharely/models"
	"sharely/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) (models.User, error)
	Login(loginRequest models.LoginRequest) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

// Register implements AuthService
func (as *authService) Register(user models.User) (models.User, error) {
	requestRegister := models.User{}
	requestRegister.Fullname = user.Fullname
	requestRegister.Email = user.Email
	requestRegister.PhoneNumber = user.PhoneNumber
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return user, err
	}
	requestRegister.Password = string(hash)
	newUser, err := as.authRepository.Create(requestRegister)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// Login implements AuthService
func (as *authService) Login(loginRequest models.LoginRequest) (models.User, error) {
	user, err := as.authRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewAuthService(authRepository *repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: *authRepository,
	}
}
