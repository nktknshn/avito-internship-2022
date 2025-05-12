package testing_pg

import (
	"context"

	"github.com/pkg/errors"

	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"

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

func (s *TestSuitePg) SetPostgresMigrationsDir(dir string) {
	s.NeedsPostgres = true
	s.MigrationsDir = dir
}

// GetPostgresAdapter returns the adapter
func (s *TestSuitePg) GetPostgresAdapter() *sqlx.DB {
	return s.Conn
}

// GetDT returns the docker database
func (s *TestSuitePg) GetDT() *DockerDatabase {
	return s.DT
}

func (s *TestSuitePg) Context() context.Context {
	return s.T().Context()
}

// Logf logs a formatted message
func (s *TestSuitePg) Logf(format string, args ...any) {
	s.T().Logf(format, args...)
}

func (s *TestSuitePg) Log(args ...any) {
	s.T().Log(args...)
}

func (s *TestSuitePg) ExecSQLMust(sql string) *ResultRows {
	rows, err := s.ExecSQL(sql)
	s.Require().NoError(err)
	return rows
}

func (s *TestSuitePg) ExecSQLExpectRowsLen(sql string, expectedRowsLen int) {
	rows, err := s.ExecSQL(sql)
	s.Require().NoError(err)
	s.Require().Len(rows.Rows, expectedRowsLen)
}

func (s *TestSuitePg) ExecSQL(sql string) (*ResultRows, error) {
	var result ResultRows

	rows := []map[string]any{}

	err := sqlx_pg.NamedSelectMapScan(s.Context(), s.Conn, &rows, sql, map[string]any{})

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return &result, nil
	}

	headers := make([]string, 0, len(rows[0]))
	for key := range rows[0] {
		headers = append(headers, key)
	}

	result.Headers = headers
	result.Rows = rows

	return &result, nil
}

func (s *TestSuitePg) SetupSuite() {
	s.T().Log("Setting up docker")

	if s.NeedsPostgres {
		s.SetupPostgres()
	}

}

func (s *TestSuitePg) SetupPostgres() {
	s.T().Log("Setting up postgres")

	ctx := s.Context()

	db := NewDockerDatabase(defaultDockerDatabaseConfig)

	if err := db.RunPostgresDocker(ctx); err != nil {
		s.T().Fatal(errors.Wrap(err, "failed to run postgres docker"))
	}

	port, err := db.GetRunningPort()

	if err != nil {
		s.T().Fatal(errors.Wrap(err, "failed to get running port"))
	}

	s.T().Logf("Running port: %s", port)

	if err != nil {
		s.T().Fatal(errors.Wrap(err, "failed to get running port"))
	}

	adapter, err := db.Connect(ctx, s.MigrationsDir)
	if err != nil {
		s.T().Fatal(errors.Wrap(err, "failed to connect to docker database"))
	}

	s.DT = db
	s.Conn = adapter

}

func (s *TestSuitePg) TearDownSuite() {
	s.T().Log("Stopping docker")

	if s.DT != nil {
		if err := s.DT.StopPostgresDocker(); err != nil {
			s.FailNow(err.Error())
		}
	}
}

func (s *TestSuitePg) TearDownTest() {
	s.T().Log("Cleaning database")
}
