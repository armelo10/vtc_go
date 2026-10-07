package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/domain/auth"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEmailExists = errors.New("email already exists")

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (r *AuthRepository) CreatePassenger(ctx context.Context, email, phone, passwordHash, firstName, lastName string) (auth.User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return auth.User{}, err
	}
	defer tx.Rollback(ctx)

	var user auth.User
	err = tx.QueryRow(ctx, "INSERT INTO users (email, phone, password_hash, role) VALUES ($1, $2, $3, 'PASSENGER') RETURNING id, email, COALESCE(phone, ''), role", email, phone, passwordHash).Scan(&user.ID, &user.Email, &user.Phone, &user.Role)
	if err != nil {
		if isUniqueViolation(err) {
			return auth.User{}, ErrEmailExists
		}
		return auth.User{}, err
	}

	if _, err = tx.Exec(ctx, "INSERT INTO passengers (user_id, first_name, last_name) VALUES ($1, $2, $3)", user.ID, firstName, lastName); err != nil {
		return auth.User{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return auth.User{}, err
	}
	user.FirstName, user.LastName = firstName, lastName
	return user, nil
}

func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (string, auth.User, error) {
	var hash string
	var user auth.User
	err := r.pool.QueryRow(ctx, "SELECT u.password_hash, u.id, u.email, COALESCE(u.phone, ''), u.role, COALESCE(p.first_name, ''), COALESCE(p.last_name, '') FROM users u LEFT JOIN passengers p ON p.user_id = u.id WHERE lower(u.email) = lower($1)", email).Scan(&hash, &user.ID, &user.Email, &user.Phone, &user.Role, &user.FirstName, &user.LastName)
	return hash, user, err
}

func (r *AuthRepository) CreateSession(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO sessions (user_id, token_hash, expires_at) VALUES ($1, $2, $3)", userID, tokenHash, expiresAt)
	return err
}

func (r *AuthRepository) FindSession(ctx context.Context, tokenHash string) (auth.User, time.Time, error) {
	var user auth.User
	var expiresAt time.Time
	err := r.pool.QueryRow(ctx, "SELECT u.id, u.email, COALESCE(u.phone, ''), u.role, COALESCE(p.first_name, ''), COALESCE(p.last_name, ''), s.expires_at FROM sessions s JOIN users u ON u.id = s.user_id LEFT JOIN passengers p ON p.user_id = u.id WHERE s.token_hash = $1 AND s.revoked_at IS NULL", tokenHash).Scan(&user.ID, &user.Email, &user.Phone, &user.Role, &user.FirstName, &user.LastName, &expiresAt)
	return user, expiresAt, err
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
