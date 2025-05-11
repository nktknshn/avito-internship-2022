package lagging

import (
	"context"
	"math/rand"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type LaggingDeps struct {
}

func NewLaggingDeps(ctx context.Context, cfg *config.Config) (*app_impl.AppDeps, func(), error) {

	deps, closer, err := app_impl.NewDeps(ctx, cfg)
	if err != nil {
		return nil, nil, err
	}

	return &app_impl.AppDeps{
		PasswordHasher: deps.PasswordHasher,
		TokenGenerator: deps.TokenGenerator,
		Repositories: app_impl.Repositories{
			AccountsRepository: &laggingAccountsRepository{
				repo: deps.Repositories.AccountsRepository,
			},
			TransactionsRepository: deps.Repositories.TransactionsRepository,
			AuthRepository:         deps.Repositories.AuthRepository,
		},
		Trm:           deps.Trm,
		FileExporter:  deps.FileExporter,
		MetricsClient: deps.MetricsClient,
		Logger:        deps.Logger,
	}, closer, nil
}

type laggingAccountsRepository struct {
	repo domainAccount.AccountRepository
}

func (r *laggingAccountsRepository) GetByUserID(ctx context.Context, userID domain.UserID) (*domainAccount.Account, error) {

	// if rand.Intn(10) > 8 {
	// 	time.Sleep(60 * time.Second)
	// }

	if rand.Intn(11) > 9 {
		panic("lagging")
	}

	return r.repo.GetByUserID(ctx, userID)
}

func (r *laggingAccountsRepository) GetByAccountID(ctx context.Context, accountID domainAccount.AccountID) (*domainAccount.Account, error) {
	// if rand.Intn(10) > 8 {
	// 	time.Sleep(60 * time.Second)
	// }

	if rand.Intn(11) > 9 {
		panic("lagging")
	}

	return r.repo.GetByAccountID(ctx, accountID)
}

func (r *laggingAccountsRepository) Save(ctx context.Context, account *domainAccount.Account) (*domainAccount.Account, error) {
	// if rand.Intn(10) > 8 {
	// 	time.Sleep(60 * time.Second)
	// }

	if rand.Intn(11) > 9 {
		panic("lagging")
	}

	return r.repo.Save(ctx, account)
}
