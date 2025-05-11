package accounts_pg

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"

	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type AccountsRepository struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func New(db *sqlx.DB, c *trmsqlx.CtxGetter) *AccountsRepository {

	if db == nil {
		panic("db is nil")
	}

	if c == nil {
		panic("ctxGetter is nil")
	}

	return &AccountsRepository{db: db, getter: c}
}

func (r *AccountsRepository) GetByUserID(ctx context.Context, userID domain.UserID) (*domainAccount.Account, error) {
	sq := `SELECT id, user_id, balance_available, balance_reserved FROM accounts WHERE user_id = ? FOR UPDATE;`

	var accDTO accountDTO

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("AccountsRepository.GetByUserID: tr is nil")
	}

	err := tr.GetContext(ctx, &accDTO, tr.Rebind(sq), userID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domainAccount.ErrAccountNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.GetByUserID.GetContext")
	}

	acc, err := fromAccountDTO(&accDTO)

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.GetByUserID.fromAccountDTO")
	}

	return acc, nil
}

func (r *AccountsRepository) GetByAccountID(ctx context.Context, accountID domainAccount.AccountID) (*domainAccount.Account, error) {
	sq := `SELECT id, user_id, balance_available, balance_reserved FROM accounts WHERE id = ? FOR UPDATE;`

	var accDTO accountDTO

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("AccountsRepository.GetByAccountID: tr is nil")
	}

	err := tr.GetContext(ctx, &accDTO, tr.Rebind(sq), accountID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domainAccount.ErrAccountNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.GetByAccountID.GetContext")
	}

	acc, err := fromAccountDTO(&accDTO)

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.GetByAccountID.fromAccountDTO")
	}

	return acc, nil
}

func (r *AccountsRepository) Save(ctx context.Context, account *domainAccount.Account) (*domainAccount.Account, error) {

	accDTO := toAccountDTO(account)

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return nil, errors.New("AccountsRepository.Save: tr is nil")
	}

	var newDTO *accountDTO
	var err error
	if accDTO.Id == 0 {
		newDTO, err = r.create(ctx, tr, accDTO)
	} else {
		newDTO, err = r.update(ctx, tr, accDTO)
	}

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.Save.create")
	}

	acc, err := fromAccountDTO(newDTO)

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.Save.fromAccountDTO")
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
		return nil, errors.Wrap(err, "AccountsRepository.create.BindNamed")
	}

	var newDTO accountDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)
	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.create.GetContext")
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
		return nil, errors.Wrap(err, "AccountsRepository.update.BindNamed")
	}

	var newDTO accountDTO
	err = tr.GetContext(ctx, &newDTO, sq, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domainAccount.ErrAccountNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "AccountsRepository.update.GetContext")
	}

	return &newDTO, nil
}

var _ domainAccount.AccountRepository = &AccountsRepository{}
