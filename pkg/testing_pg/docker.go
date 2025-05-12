package testing_pg

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	"github.com/nktknshn/avito-internship-2022/pkg/sqlx_pg"
)

type DatabaseTestDockerConfig struct {
	Repository     string
	Tag            string
	DockerPassword string
	Timezone       string
	RestartPolicy  docker.RestartPolicy
	ResourceExpire uint
}

type RunningPostgres struct {
	Container *dockertest.Resource
}

func (rp *RunningPostgres) GetPort(name string) string {
	if rp == nil {
		return ""
	}

	return rp.Container.GetPort(name)
}

type DockerDatabase struct {
	dockerConfig DatabaseTestDockerConfig

	runningDocker *RunningPostgres
	pool          *dockertest.Pool
}

//nolint:mnd // дефолтный конфиг
var defaultDockerDatabaseConfig = DatabaseTestDockerConfig{
	Repository:     "postgres",
	Tag:            "latest",
	DockerPassword: "who-cares-about-the-test-password",
	RestartPolicy:  docker.RestartPolicy{Name: "no"},
	ResourceExpire: 120,
	Timezone:       "UTC",
	// Timezone: "Europe/Moscow",
}

func NewDockerDatabase(dockerConfig DatabaseTestDockerConfig) *DockerDatabase {
	return &DockerDatabase{
		dockerConfig: dockerConfig,
	}
}

func (dt *DockerDatabase) IsRunning() bool {
	return dt.runningDocker != nil
}

func (dt *DockerDatabase) GetRunningPort() (string, error) {
	if dt.runningDocker == nil {
		return "", errors.New("postgres docker is not running")
	}

	return dt.runningDocker.GetPort("5432/tcp"), nil
}

func (dt *DockerDatabase) GetDockerConfig() DatabaseTestDockerConfig {
	return dt.dockerConfig
}

func (dt *DockerDatabase) RunPostgresDocker(ctx context.Context) error {

	pool, err := dockertest.NewPool("")

	dt.pool = pool

	if err != nil {
		return err
	}

	if err = pool.Client.PingWithContext(ctx); err != nil {
		return err
	}

	if dt.dockerConfig.DockerPassword == "" {
		dt.dockerConfig.DockerPassword = defaultDockerDatabaseConfig.DockerPassword
	}

	if dt.dockerConfig.Repository == "" {
		dt.dockerConfig.Repository = defaultDockerDatabaseConfig.Repository
	}

	if dt.dockerConfig.Tag == "" {
		dt.dockerConfig.Tag = defaultDockerDatabaseConfig.Tag
	}

	if dt.dockerConfig.RestartPolicy == (docker.RestartPolicy{}) {
		dt.dockerConfig.RestartPolicy = defaultDockerDatabaseConfig.RestartPolicy
	}

	if dt.dockerConfig.ResourceExpire == 0 {
		dt.dockerConfig.ResourceExpire = defaultDockerDatabaseConfig.ResourceExpire
	}

	if dt.dockerConfig.Timezone == "" {
		dt.dockerConfig.Timezone = defaultDockerDatabaseConfig.Timezone
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: dt.dockerConfig.Repository,
		Tag:        dt.dockerConfig.Tag,
		Env: []string{
			"POSTGRES_PASSWORD=" + dt.dockerConfig.DockerPassword,
			"TZ=" + dt.dockerConfig.Timezone,
			"PGTZ=" + dt.dockerConfig.Timezone,
			"timezone=" + dt.dockerConfig.Timezone,
		},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = dt.dockerConfig.RestartPolicy
	})

	if err != nil {
		return err
	}

	if err = resource.Expire(dt.dockerConfig.ResourceExpire); err != nil {
		return err
	}

	dt.runningDocker = &RunningPostgres{
		Container: resource,
	}

	return nil
}

const (
	maxOpenConnections = 100
	maxIdleConnections = 10
)

func (dt *DockerDatabase) Connect(ctx context.Context, migrationsDir string) (*sqlx.DB, error) {

	cfg := &config{
		Addr:                  "localhost:" + dt.runningDocker.GetPort("5432/tcp"),
		User:                  "postgres",
		Password:              dt.dockerConfig.DockerPassword,
		Database:              "postgres",
		Schema:                "public",
		MaxIdleConnections:    maxOpenConnections,
		MaxOpenConnections:    maxIdleConnections,
		ConnectionMaxLifetime: time.Hour,
		UpMigrations:          true,
		ReturnUTC:             true,
	}

	var conn *sqlx.DB

	err := dt.pool.Retry(func() error {
		var err error
		conn, err = sqlx_pg.Connect(ctx, cfg)
		if err != nil {
			slog.Error("postgres.NewPostgresAdapter(ctx, cfg)", "error", err)
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if conn == nil {
		return nil, errors.New("postgres connection is nil")
	}

	if cfg.UpMigrations && migrationsDir != "" {
		if _, statErr := os.Stat(migrationsDir); statErr != nil {
			return nil, errors.New("migrations directory does not exist or not accessible")
		}
		err = sqlx_pg.Migrate(ctx, conn.DB, migrationsDir)
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}

func (dt *DockerDatabase) StopPostgresDocker() error {
	if dt.runningDocker == nil {
		return nil
	}
	container := dt.runningDocker.Container
	dt.runningDocker = nil
	return dt.pool.Purge(container)
}
