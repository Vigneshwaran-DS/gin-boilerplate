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

// Register 用户注册
func (s *AuthService) Register(username, email, password, fullName string) (*models.User, error) {
	// 检查用户名是否已存在
	existUser, _ := s.userService.GetUserByUsername(username)
	if existUser != nil {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	var user models.User
	err := database.GetDB().Where("email = ?", email).First(&user).Error
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// 创建用户
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

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	// 查找用户
	user, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	// 验证密码
	if !s.userService.VerifyPassword(user, password) {
		return "", nil, errors.New("invalid username or password")
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
