package services

import (
	"errors"
	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/repositories"
	"ecommerce-platform/internal/utils"
	"strings"

	"github.com/google/uuid"
)

// AuthService now holds the jwtSecret
type AuthService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

// NewAuthService now correctly accepts the jwtSecret
func NewAuthService(userRepo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// ... (validation logic is the same)
	user, err := s.userRepo.GetByEmail(strings.ToLower(strings.TrimSpace(req.Email)))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token using the stored secret
	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role), s.jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	user.PasswordHash = "" // Don't send hash to client
	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

// Register creates a new user account
func (s *AuthService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Validate input (logic from your original file)
	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return nil, errors.New("all fields are required")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Check if email already exists
	exists, err := s.userRepo.EmailExists(strings.ToLower(strings.TrimSpace(req.Email)))
	if err != nil {
		return nil, errors.New("failed to check email availability")
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// THIS IS THE MISSING PART: Create the user object
	user := &models.User{
		Email:        strings.ToLower(strings.TrimSpace(req.Email)),
		PasswordHash: hashedPassword,
		FirstName:    strings.TrimSpace(req.FirstName),
		LastName:     strings.TrimSpace(req.LastName),
		Role:         models.RoleCustomer, // Default role is customer
	}
	if req.Phone != "" {
		phone := strings.TrimSpace(req.Phone)
		user.Phone = &phone
	}

	// Save user to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate JWT token using the stored secret
	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role), s.jwtSecret)
	if err != nil {
		return nil, errors.New("account created but failed to generate token")
	}

	user.PasswordHash = "" // Don't send hash to client
	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

// GetUserByID retrieves a user by their ID
func (s *AuthService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	return s.userRepo.GetByID(userID)
}

// ... (your validation logic for registration)