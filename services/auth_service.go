package services

import (
	"errors"

	"gin-boilerplate/database"
	"gin-boilerplate/models"
	"gin-boilerplate/utils"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: NewUserService(),
	}
}

// Register user registration
func (s *AuthService) Register(username, email, password, fullName string) (*models.User, error) {
	// Check if username already exists
	existUser, _ := s.userService.GetUserByUsername(username)
	if existUser != nil {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	var user models.User
	err := database.GetDB().Where("email = ?", email).First(&user).Error
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// Create user
	newUser := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		FullName: fullName,
	}

	if err := s.userService.CreateUser(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Login user login
func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	// Find user
	user, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	// Verify password
	if !s.userService.VerifyPassword(user, password) {
		return "", nil, errors.New("invalid username or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
