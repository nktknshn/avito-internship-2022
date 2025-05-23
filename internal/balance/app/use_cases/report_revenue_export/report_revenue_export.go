package report_revenue_export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/common/file_exporter"
)

type ReportRevenueExportUseCase struct {
	repository report_revenue.ReportRevenueRepository
	exporter   file_exporter.FileExporter
}

func New(
	exporter file_exporter.FileExporter,
	repository report_revenue.ReportRevenueRepository,
) *ReportRevenueExportUseCase {
	if exporter == nil {
		panic("exporter is nil")
	}

	if repository == nil {
		panic("repository is nil")
	}

	return &ReportRevenueExportUseCase{
		exporter:   exporter,
		repository: repository,
	}
}

func (u *ReportRevenueExportUseCase) Handle(ctx context.Context, in In) (Out, error) {

	report, err := u.repository.GetReportRevenueByMonth(ctx, report_revenue.ReportRevenueQuery{
		Year:  in.year,
		Month: in.month,
	})

	if err != nil {
		return Out{}, err
	}

	csvData, err := u.convertToCSV(report)
	if err != nil {
		return Out{}, err
	}

	now := time.Now()
	fileName := fmt.Sprintf("revenue_report_%s.csv", now.Format("2006-01-02_15-04-05"))

	filePath, err := u.exporter.ExportFile(ctx, fileName, strings.NewReader(csvData))

	if err != nil {
		return Out{}, err
	}

	return Out{
		URL: filePath,
	}, nil
}

func (u *ReportRevenueExportUseCase) convertToCSV(report report_revenue.ReportRevenueResponse) (string, error) {
	csvData := &bytes.Buffer{}
	csvWriter := csv.NewWriter(csvData)
	csvWriter.Comma = ';'
	err := csvWriter.Write([]string{"product_id", "product_title", "total_revenue"})
	if err != nil {
		return "", err
	}
	for _, record := range report.Records {
		err := csvWriter.Write([]string{
			strconv.FormatInt(record.ProductID.Value(), 10),
			record.ProductTitle.Value(),
			strconv.FormatInt(record.TotalRevenue, 10),
		})
		if err != nil {
			return "", err
		}
	}
	csvWriter.Flush()
	return csvData.String(), nil
}

func (u *ReportRevenueExportUseCase) GetName() string {
	return use_cases.NameReportRevenueExport
}
