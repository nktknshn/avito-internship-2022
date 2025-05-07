package app_impl

import (
	"context"
	"net/http"

	"github.com/avito-tech/go-transaction-manager/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	ucReportRevenue "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	ucReportTransactions "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainTransactions "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/common/file_exporter"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
	"github.com/nktknshn/avito-internship-2022/internal/common/password_hasher"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
	"github.com/nktknshn/avito-internship-2022/pkg/file_exporter_http"
	"github.com/nktknshn/avito-internship-2022/pkg/metrics_prometheus"
	"github.com/nktknshn/avito-internship-2022/pkg/password_hasher_argon"
	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
	"github.com/nktknshn/avito-internship-2022/pkg/token_generator_jwt"
)

type TransactionRepository interface {
	domainTransactions.TransactionRepository
	ucReportTransactions.ReportTransactionsRepository
	ucReportRevenue.ReportRevenueRepository
}

type Repositories struct {
	AuthRepository         domainAuth.AuthRepository
	AccountsRepository     domainAccount.AccountRepository
	TransactionsRepository TransactionRepository
}

type AppDeps struct {
	PasswordHasher password_hasher.HasherVerifier
	TokenGenerator token_generator.TokenGeneratorVerifier[domainAuth.AuthUserTokenClaims]
	Repositories   Repositories
	Trm            *manager.Manager
	FileExporter   FileExporterWithHandler
	MetricsClient  MetricsWithHandler
	Logger         logging.Logger
}

type MetricsWithHandler interface {
	metrics.Metrics
	GetHandler() http.Handler
}

type FileExporterWithHandler interface {
	file_exporter.FileExporter
	GetHandler() http.Handler
	CleanupWorker(ctx context.Context)
}

func NewDeps(ctx context.Context, cfg *config.Config) (*AppDeps, error) {

	db, err := sqlx_pg.Connect(ctx, cfg.GetPostgres())
	if err != nil {
		return nil, err
	}

	if cfg.GetPostgres().GetMigrationsDir() != "" {
		err = sqlx_pg.Migrate(ctx, db.DB, cfg.GetPostgres().GetMigrationsDir())
		if err != nil {
			return nil, err
		}
	}

	trmFactory := trmsqlx.NewFactory(db, sql.NewSavePoint())
	trm, err := manager.New(trmFactory)
	if err != nil {
		return nil, err
	}

	exporter, err := file_exporter_http.New(
		file_exporter_http.Config{
			Folder: cfg.GetUseCases().GetReportRevenueExport().GetFolder(),
			TTL:    cfg.GetUseCases().GetReportRevenueExport().GetTTL(),
			URL:    cfg.GetUseCases().GetReportRevenueExport().GetURL(),
			Zip:    cfg.GetUseCases().GetReportRevenueExport().GetZip(),
		},
	)

	if err != nil {
		return nil, err
	}

	metricsClient, err := metrics_prometheus.NewMetricsPrometheus("app_balance")
	if err != nil {
		return nil, err
	}

	logger := logging.NewSlog()

	return &AppDeps{
		PasswordHasher: password_hasher_argon.New(),
		TokenGenerator: token_generator_jwt.NewTokenGeneratorJWTVerifier[domainAuth.AuthUserTokenClaims](
			[]byte(cfg.GetJWT().GetSecret()),
			cfg.GetJWT().GetTTL(),
		),
		Repositories: Repositories{
			AuthRepository:         auth_pg.New(db, trmsqlx.DefaultCtxGetter),
			AccountsRepository:     accounts_pg.New(db, trmsqlx.DefaultCtxGetter),
			TransactionsRepository: transactions_pg.New(db, trmsqlx.DefaultCtxGetter),
		},
		Trm:           trm,
		FileExporter:  exporter,
		MetricsClient: metricsClient,
		Logger:        logger,
	}, nil
}
