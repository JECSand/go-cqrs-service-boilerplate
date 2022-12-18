package repositories

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	createUserQuery = `INSERT INTO users (id, email, username, password, root, active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, now(), now()) RETURNING id, email, username, password, root, active, created_at, updated_at`

	updateUserQuery = `UPDATE users p SET 
                      email=COALESCE(NULLIF($1, ''), email), 
                      username=COALESCE(NULLIF($2, ''), username), 
                      updated_at = now()
                      WHERE id=$3
                      RETURNING id, email, username, root, active, created_at, updated_at`

	getUserByIdQuery = `SELECT p.id, p.email, p.username, p.password, p.root, p.active, p.created_at, p.updated_at 
	FROM users p WHERE p.id = $1`

	deleteUserByIdQuery = `DELETE FROM users WHERE id = $1`
)

type userRepository struct {
	log logging.Logger
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewUserRepository(log logging.Logger, cfg *config.Config, db *pgxpool.Pool) *userRepository {
	return &userRepository{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

func (p *userRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepository.CreateUser")
	defer span.Finish()
	var created entities.User
	if err := p.db.QueryRow(ctx, createUserQuery, &user.ID, &user.Email, &user.Username, &user.Password, user.Root, user.Active).Scan(
		&created.ID,
		&created.Email,
		&created.Username,
		&created.Password,
		&created.Root,
		&created.Active,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}
	return &created, nil
}

func (p *userRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepository.UpdateUser")
	defer span.Finish()
	var updated entities.User
	if err := p.db.QueryRow(
		ctx,
		updateUserQuery,
		&user.ID,
		&user.Email,
		&user.Username,
	).Scan(&updated.ID, &updated.Email, &updated.Username, &updated.Root, &updated.Active, &updated.CreatedAt, &updated.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}
	return &updated, nil
}

func (p *userRepository) GetUserById(ctx context.Context, uuid uuid.UUID) (*entities.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepository.GetUserById")
	defer span.Finish()
	var found entities.User
	if err := p.db.QueryRow(ctx, getUserByIdQuery, uuid).Scan(
		&found.ID,
		&found.Email,
		&found.Username,
		&found.Password,
		&found.Root,
		&found.Active,
		&found.CreatedAt,
		&found.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}
	return &found, nil
}

func (p *userRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepository.DeleteUserByID")
	defer span.Finish()
	_, err := p.db.Exec(ctx, deleteUserByIdQuery, id)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}
	return nil
}
