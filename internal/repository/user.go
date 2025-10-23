package repository

import (
	"auctionhouse/internal/models"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (repository *UserRepository) CreateUser(email, password string) (*models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	result, err := repository.DB.Exec("INSERT INTO users (email, password, created_at) VALUES (?, ?, ?)", email, hashedPassword, time.Now())
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.User{
		ID:        int(id),
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}, nil
}
func (repository *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	row := repository.DB.QueryRow("SELECT id, email, password, created_at FROM users WHERE email = ?", email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := ur.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !CheckPassword(user.Password, password) {
		return nil, nil
	}
	user.Password = "" // Clear password before returning
	return user, nil
}
