package deps

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

var (
	_config *config.Config
	_app    *app_impl.Application
	_cli    *cli.CliAdapter
)

func GetApplication(ctx context.Context) (*app_impl.Application, error) {
	if _app != nil {
		return _app, nil
	}
	cfg, err := GetConfig(ctx)
	if err != nil {
		return nil, err
	}
	app, _, err := app_impl.NewApplication(ctx, cfg)
	if err != nil {
		return nil, err
	}
	_app = app
	return _app, nil
}

func GetConfig(_ context.Context) (*config.Config, error) {
	if _config != nil {
		return _config, nil
	}
	cfgPath := root.GetConfigPath()
	cfg, err := config.LoadConfigFromFile(cfgPath)
	if err != nil {
		return nil, err
	}
	_config = cfg
	return cfg, nil
}

func GetCliAdapter(ctx context.Context) (*cli.CliAdapter, error) {
	if _cli != nil {
		return _cli, nil
	}
	app, err := GetApplication(ctx)
	if err != nil {
		return nil, err
	}
	cliAdapter := cli.NewCliAdapter(app.GetApp())
	_cli = cliAdapter
	return cliAdapter, nil
}
