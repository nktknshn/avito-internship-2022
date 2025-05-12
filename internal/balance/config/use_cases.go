package config

import "time"

type ConfigUseCases struct {
	ReportRevenueExport *ConfigReportRevenueExport `yaml:"report_revenue_export" env-required:"true"`
}

func (c *ConfigUseCases) GetReportRevenueExport() *ConfigReportRevenueExport {
	return c.ReportRevenueExport
}

type ConfigReportRevenueExport struct {
	Folder string        `yaml:"folder" env-required:"true" env:"REPORT_REVENUE_EXPORT_FOLDER"`
	TTL    time.Duration `yaml:"ttl"    env-required:"true" env:"REPORT_REVENUE_EXPORT_TTL"`
	URL    string        `yaml:"url"    env-required:"true" env:"REPORT_REVENUE_EXPORT_URL"`
	Zip    bool          `yaml:"zip"    env-required:"true" env:"REPORT_REVENUE_EXPORT_ZIP"`
}

func (c *ConfigReportRevenueExport) GetFolder() string {
	return c.Folder
}

func (c *ConfigReportRevenueExport) GetTTL() time.Duration {
	return c.TTL
}

func (c *ConfigReportRevenueExport) GetURL() string {
	return c.URL
}

func (c *ConfigReportRevenueExport) GetZip() bool {
	return c.Zip
}
