package auth_pg

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
)

type AuthRepository struct {
	getter *trmsqlx.CtxGetter
	db     *sqlx.DB
}

func New(db *sqlx.DB, getter *trmsqlx.CtxGetter) *AuthRepository {

	if db == nil {
		panic("db is nil")
	}

	if getter == nil {
		panic("getter is nil")
	}

	return &AuthRepository{db: db, getter: getter}
}

func (r *AuthRepository) CreateUser(
	ctx context.Context,
	username domainAuth.AuthUserUsername,
	passwordHash domainAuth.AuthUserPasswordHash,
	role domainAuth.AuthUserRole,
) error {
	sq := `INSERT INTO auth_users (username, password_hash, role) VALUES (?, ?, ?)`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return errors.New("tr is nil")
	}

	_, err := tr.ExecContext(ctx, tr.Rebind(sq), username, passwordHash, role)

	if sqlx_pg.IsDuplicateKeyError(err) {
		return domainAuth.ErrDuplicateUsername
	}

	if err != nil {
		return errors.Wrap(err, "AuthRepository.CreateUser.ExecContext")
	}

	return nil
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username domainAuth.AuthUserUsername) (*domainAuth.AuthUser, error) {
	sq := `SELECT id, username, password_hash, role FROM auth_users WHERE username = ?`

	var userDTO authUserDTO

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("tr is nil")
	}

	err := tr.GetContext(ctx, &userDTO, tr.Rebind(sq), username)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domainAuth.ErrAuthUserNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "AuthRepository.GetUserByUsername.GetContext")
	}

	user, err := fromAuthUserDTO(&userDTO)

	if err != nil {
		return nil, errors.Wrap(err, "AuthRepository.GetUserByUsername.fromAuthUserDTO")
	}

	return user, nil
}

func (r *AuthRepository) ListUsers(ctx context.Context) ([]*domainAuth.AuthUser, error) {
	sq := `SELECT id, username, password_hash, role FROM auth_users`

	var dtos []authUserDTO

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("tr is nil")
	}

	err := tr.SelectContext(ctx, &dtos, tr.Rebind(sq))

	if err != nil {
		return nil, errors.Wrap(err, "AuthRepository.ListUsers.SelectContext")
	}

	authUsers := make([]*domainAuth.AuthUser, len(dtos))

	for i, user := range dtos {
		authUsers[i], err = fromAuthUserDTO(&user)
		if err != nil {
			return nil, errors.Wrap(err, "AuthRepository.ListUsers.fromAuthUserDTO")
		}
	}

	return authUsers, nil
}

var _ domainAuth.AuthRepository = (*AuthRepository)(nil)
