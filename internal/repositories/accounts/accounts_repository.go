package accounts

import (
	"context"
	"database/sql"
	"errors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type AccountsRepository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func (r *AccountsRepository) GetByUserID(ctx context.Context, userID domain.UserID) (*domain.Account, error) {
	sq := `SELECT id, user_id, balance_available, balance_reserved FROM accounts WHERE user_id = ?;`

	var accDTO accountDTO

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	err := tr.GetContext(ctx, &accDTO, r.db.Rebind(sq), userID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	acc, err := fromAccountDTO(&accDTO)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (r *AccountsRepository) Save(ctx context.Context, account *domain.Account) (*domain.Account, error) {

	accDTO, err := toAccountDTO(account)

	if err != nil {
		return nil, err
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	var newDTO *accountDTO

	if accDTO.Id == 0 {
		newDTO, err = r.create(ctx, tr, accDTO)
	} else {
		newDTO, err = r.update(ctx, tr, accDTO)
	}

	if err != nil {
		return nil, err
	}

	acc, err := fromAccountDTO(newDTO)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (r *AccountsRepository) create(ctx context.Context, tr trmsqlx.Tr, accDto *accountDTO) (*accountDTO, error) {
	sq := `
		INSERT INTO accounts 
			(user_id, balance_available, balance_reserved) 
		VALUES 
			(:user_id, :balance_available, :balance_reserved)
		RETURNING *;
		`

	sq, args, err := tr.BindNamed(sq, accDto)

	if err != nil {
		return nil, err
	}

	var newDTO accountDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, err
	}

	return &newDTO, nil
}

func (r *AccountsRepository) update(ctx context.Context, tr trmsqlx.Tr, accDto *accountDTO) (*accountDTO, error) {
	sq := `
		UPDATE accounts 
		SET 
			balance_available = :balance_available, 
			balance_reserved = :balance_reserved 
		WHERE id = :id 
		RETURNING *;
		`

	sq, args, err := tr.BindNamed(sq, accDto)
	if err != nil {
		return nil, err
	}

	var newDTO accountDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, err
	}

	return &newDTO, nil
}

var _ domain.AccountRepository = &AccountsRepository{}
