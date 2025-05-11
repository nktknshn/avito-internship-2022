package postgres_test

import (
	"testing"
	"time"

	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestPostgres(t *testing.T) {
	s := &Suite{}
	s.NeedsPostgres = true
	suite.Run(t, s)
}

type Suite struct {
	testing_pg.TestSuitePg
}

func (s *Suite) TearDownTest() {
	s.ExecSqlMust(`DROP TABLE IF EXISTS test_time;`)
}

func (s *Suite) TestEqaultDates() {
	t0 := time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local)
	t1 := t0.UTC()

	require.True(s.T(), t0.Equal(t1))
	require.NotEqual(s.T(), t0.Day(), t1.Day())
}

func (s *Suite) TestTimeSaveAndGet_TIMESTAMPTZ() {
	// Если время сохраняется в TIMESTAMPTZ, то оно сохраняется в UTC, конвертируя часовой пояс.
	_, err := s.ExecSql(`
		CREATE TABLE test_time_tz (
			created_at TIMESTAMPTZ NOT NULL
		);
	`)
	s.Require().NoError(err)

	loc, err := time.LoadLocation("Europe/Moscow")
	s.Require().NoError(err)
	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, loc)

	_, err = s.Conn.NamedExec(`
		INSERT INTO test_time_tz (created_at) VALUES (:created_at)
	`, map[string]any{
		"created_at": t0,
	})
	s.Require().NoError(err)

	rows, err := s.ExecSql(`
		SELECT * FROM test_time_tz;
	`)
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))

	t := rows.Rows[0]["created_at"].(time.Time)

	// Проверяем, что время сохранилось корректно и вернулось в UTC
	s.Require().Equal(time.UTC, t.Location())
	s.Require().Equal(t0, t.In(loc))
}

func (s *Suite) TestTimeSaveAndGet_TIMESTAMP() {
	// Если время сохраняется в TIMESTAMP, то оно сохраняется в часовом поясе в UTC, а таймзона отбрасывается.
	_, err := s.ExecSql(`
		CREATE TABLE test_time (
			created_at TIMESTAMP NOT NULL
		);
	`)
	s.Require().NoError(err)

	loc, err := time.LoadLocation("Europe/Moscow")
	s.Require().NoError(err)
	t0 := time.Date(2024, 1, 1, 0, 0, 0, 0, loc)

	_, err = s.Conn.NamedExec(`
		INSERT INTO test_time (created_at) VALUES (:created_at)
	`, map[string]any{
		"created_at": t0,
	})
	s.Require().NoError(err)

	rows, err := s.ExecSql(`
		SELECT * FROM test_time;
	`)
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))

	t := rows.Rows[0]["created_at"].(time.Time)

	// Время сохранилось в UTC, отбросив часовой пояс, и вернулось в UTC
	s.Require().Equal(time.UTC, t.Location())
	s.Require().Equal(t0.Day(), t.Day())
	s.Require().Equal(t0.Month(), t.Month())
	s.Require().Equal(t0.Year(), t.Year())
	s.Require().Equal(t0.Hour(), t.Hour())
	s.Require().Equal(t0.Minute(), t.Minute())
	s.Require().Equal(t0.Second(), t.Second())
	s.Require().Equal(t0.Nanosecond(), t.Nanosecond())

	s.Require().NotEqual(t0, t.In(loc))
}
