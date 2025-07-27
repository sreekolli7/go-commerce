package infrastructure

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sreekolli7/go-commerce/user-service/internal/domain"
	"github.com/sreekolli7/go-commerce/user-service/internal/usecase"
)

type PostgresUserRepo struct {
	DB *sqlx.DB
}

func NewPostgresUserRepo(db *sqlx.DB) usecase.UserRepository {
	return &PostgresUserRepo{DB: db}
}

func (r *PostgresUserRepo) Create(user *domain.User) error {
	return r.DB.QueryRowx(
		"INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
		user.Email, user.PasswordHash, user.FirstName, user.LastName,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *PostgresUserRepo) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepo) GetByID(id int64) (*domain.User, error) {
	var user domain.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepo) Update(user *domain.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.DB.Exec(
		"UPDATE users SET email = $1, password_hash = $2, first_name = $3, last_name = $4, updated_at = $5 WHERE id = $6",
		user.Email, user.PasswordHash, user.FirstName, user.LastName, user.UpdatedAt, user.ID,
	)
	return err
}

func (r *PostgresUserRepo) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
