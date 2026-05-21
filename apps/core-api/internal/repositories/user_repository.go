package repositories

import (
	"context"

	"github.com/amantester/shadowcoreos/apps/core-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
	INSERT INTO users (email, username, password_hash, role)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.DB.QueryRow(
		context.Background(),
		query,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
	SELECT id, email, username, password_hash, role, created_at
	FROM users
	WHERE email = $1
	`

	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
