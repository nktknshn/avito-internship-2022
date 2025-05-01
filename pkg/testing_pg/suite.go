package testing_pg

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type TestSuitePg struct {
	suite.Suite

	MigrationsDir string

	DT   *DockerDatabase
	Conn *sqlx.DB

	NeedsPostgres bool
}

func (suite *TestSuitePg) SetPostgresMigrationsDir(dir string) {
	suite.NeedsPostgres = true
	suite.MigrationsDir = dir
}

// GetPostgresAdapter returns the adapter
func (suite *TestSuitePg) GetPostgresAdapter() *sqlx.DB {
	return suite.Conn
}

// GetDT returns the docker database
func (suite *TestSuitePg) GetDT() *DockerDatabase {
	return suite.DT
}

func (suite *TestSuitePg) Context() context.Context {
	return suite.T().Context()
}

// Logf(format string, args ...any)
func (suite *TestSuitePg) Logf(format string, args ...any) {
	suite.T().Logf(format, args...)
}

func (suite *TestSuitePg) Log(args ...any) {
	suite.T().Log(args...)
}

func (s *TestSuitePg) ExecSqlMust(sql string) *ResultRows {
	rows, err := s.ExecSql(sql)
	s.Require().NoError(err)
	return rows
}

func (s *TestSuitePg) ExecSqlExpectRowsLen(sql string, expectedRowsLen int) {
	rows, err := s.ExecSql(sql)
	s.Require().NoError(err)
	s.Require().Equal(expectedRowsLen, len(rows.Rows))
}

func (s *TestSuitePg) ExecSql(sql string) (*ResultRows, error) {
	rows := []map[string]any{}

	err := sqlx_pg.NamedSelectMapScan(s.Context(), s.Conn, &rows, sql, map[string]any{})

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return &ResultRows{}, nil
	}

	headers := make([]string, 0, len(rows[0]))
	for key := range rows[0] {
		headers = append(headers, key)
	}

	return &ResultRows{
		Headers: headers,
		Rows:    rows,
	}, nil
}

func (suite *TestSuitePg) SetupSuite() {
	suite.T().Log("Setting up docker")

	if suite.NeedsPostgres {
		suite.SetupPostgres()
	}

}

func (suite *TestSuitePg) SetupPostgres() {
	suite.T().Log("Setting up postgres")

	ctx := suite.Context()

	db := NewDockerDatabase(defaultDockerDatabaseConfig)

	if err := db.RunPostgresDocker(ctx); err != nil {
		suite.T().Fatal(errors.Wrap(err, "failed to run postgres docker"))
	}

	port, err := db.GetRunningPort()

	if err != nil {
		suite.T().Fatal(errors.Wrap(err, "failed to get running port"))
	}

	suite.T().Logf("Running port: %s", port)

	if err != nil {
		suite.T().Fatal(errors.Wrap(err, "failed to get running port"))
	}

	adapter, err := db.Connect(ctx, suite.MigrationsDir)
	if err != nil {
		suite.T().Fatal(errors.Wrap(err, "failed to connect to docker database"))
	}

	suite.DT = db
	suite.Conn = adapter

}

func (suite *TestSuitePg) TearDownSuite() {
	suite.T().Log("Stopping docker")

	if suite.DT != nil {
		if err := suite.DT.StopPostgresDocker(); err != nil {
			suite.FailNow(err.Error())
		}
	}
}

func (suite *TestSuitePg) TearDownTest() {
	suite.T().Log("Cleaning database")
}
