package file_exporter

import (
	"context"
	"io"
)

// FileExporter это интерфейс для экспорта файлов. Возвращает информацию о том, как получить файл (например,URL)
type FileExporter interface {
	ExportFile(ctx context.Context, name string, data io.Reader) (string, error)
}
