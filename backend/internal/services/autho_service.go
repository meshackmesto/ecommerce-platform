package services

import (
    "errors"
    "strings"

    "ecommerce-platform/internal/models"
    "ecommerce-platform/internal/repositories"
    "ecommerce-platform/internal/utils"
    "github.com/google/uuid"
)

type AuthService struct {
    userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
    return &AuthService{userRepo: userRepo}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
    // Validate input
    if req.Email == "" || req.Password == "" {
        return nil, errors.New("email and password are required")
    }

    // Find user by email
    user, err := s.userRepo.GetByEmail(strings.ToLower(strings.TrimSpace(req.Email)))
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    // Check password
    if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
        return nil, errors.New("invalid email or password")
    }

    // Generate JWT token
    token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role))
    if err != nil {
        return nil, errors.New("failed to generate token")
    }

    return &models.AuthResponse{
        User:  *user,
        Token: token,
    }, nil
}

// Register creates a new user account
func (s *AuthService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
    // Validate input
    if err := s.validateRegisterRequest(req); err != nil {
        return nil, err
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

    // Create user object
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

    // Generate JWT token
    token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role))
    if err != nil {
        return nil, errors.New("account created but failed to generate token")
    }

    return &models.AuthResponse{
        User:  *user,
        Token: token,
    }, nil
}

// GetUserByID retrieves a user by their ID
func (s *AuthService) GetUserByID(userID uuid.UUID) (*models.User, error) {
    return s.userRepo.GetByID(userID)
}

// validateRegisterRequest validates the registration request
func (s *AuthService) validateRegisterRequest(req *models.RegisterRequest) error {
    if req.Email == "" {
        return errors.New("email is required")
    }
    if req.Password == "" {
        return errors.New("password is required")
    }
    if len(req.Password) < 6 {
        return errors.New("password must be at least 6 characters")
    }
    if req.FirstName == "" {
        return errors.New("first name is required")
    }
    if req.LastName == "" {
        return errors.New("last name is required")
    }
    
    // Basic email validation
    if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
        return errors.New("invalid email format")
    }
    
    return nil
}