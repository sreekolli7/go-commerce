package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sreekolli7/go-commerce/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	GetByID(id int64) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int64) error
}

type UserUsecase struct {
	Repo UserRepository
}

func (u *UserUsecase) Register(req *domain.RegisterRequest) (*domain.User, error) {
	// Check if user already exists
	existingUser, _ := u.Repo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &domain.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
	}

	// Save to database
	err = u.Repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) Login(req *domain.LoginRequest) (string, error) {
	// Find user by email
	user, err := u.Repo.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	})

	tokenString, err := token.SignedString([]byte("your-secret-key")) // In production, use environment variable
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *UserUsecase) GetByID(id int64) (*domain.User, error) {
	return u.Repo.GetByID(id)
}
