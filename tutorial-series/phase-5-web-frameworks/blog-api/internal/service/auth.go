package service

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"blog-api/utils/jwt"
	"errors"
)

// AuthService 认证服务层
type AuthService struct {
	userRepo *repository.UserRepository
}

// NewAuthService 创建 AuthService 实例
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register 用户注册
func (s *AuthService) Register(username, password, email string) error {
	// 1. 检查用户名是否已存在
	_, err := s.userRepo.FindByUsername(username)
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 2. 检查邮箱是否已存在
	_, err = s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("邮箱已被注册")
	}

	// 3. 创建用户（密码会在 BeforeCreate Hook 中自动加密）
	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	return s.userRepo.Create(user)
}

// Login 用户登录
// 返回：Token 字符串和错误
func (s *AuthService) Login(username, password string) (string, error) {
	// 1. 查找用户
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 2. 验证密码
	if !user.CheckPassword(password) {
		return "", errors.New("用户名或密码错误")
	}

	// 3. 生成 JWT Token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.New("生成 Token 失败")
	}

	return token, nil
}
