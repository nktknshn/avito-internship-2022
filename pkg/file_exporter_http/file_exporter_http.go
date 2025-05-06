package file_exporter_http

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	Folder string
	TTL    time.Duration
	URL    string
	Zip    bool
}

// записывает файл в локальную файловую систему
// и предоставляет HTTP-хэндлер для скачивания файла
// очищает файлы по истечении TTL
type FileExporterHTTP struct {
	cfg Config
}

func newRandomPrefix() string {
	return uuid.New().String()
}

func New(cfg Config) (*FileExporterHTTP, error) {

	if cfg.Folder == "" {
		return nil, errors.New("folder is required")
	}

	// проверяем, существует ли папка
	folderInfo, err := os.Stat(cfg.Folder)

	if err != nil {
		return nil, errors.New("can't access folder")
	}

	if !folderInfo.IsDir() {
		return nil, errors.New("folder is not a directory")
	}

	return &FileExporterHTTP{cfg: cfg}, nil
}

// ExportFile экспортирует файл в локальную файловую систему и возвращает URL-адрес файла
func (f *FileExporterHTTP) ExportFile(ctx context.Context, name string, data io.Reader) (string, error) {
	if f.cfg.Zip {
		return f.exportFileZip(ctx, name, data)
	}
	return f.exportFilePlain(ctx, name, data)
}

func (f *FileExporterHTTP) exportFilePlain(ctx context.Context, name string, data io.Reader) (string, error) {
	prefix := newRandomPrefix()
	fileName := prefix + "_" + name
	filePath := filepath.Join(f.cfg.Folder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, data)
	if err != nil {
		return "", err
	}
	return f.cfg.URL + "/" + fileName, nil
}

func (f *FileExporterHTTP) exportFileZip(ctx context.Context, name string, data io.Reader) (string, error) {
	prefix := newRandomPrefix()
	fileName := prefix + "_" + name
	zipFileName := fileName + ".zip"
	zipFilePath := filepath.Join(f.cfg.Folder, zipFileName)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	w, err := zipWriter.Create(fileName)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(w, data)
	if err != nil {
		return "", err
	}

	return f.cfg.URL + "/" + zipFileName, nil
}

// Cleanup очищает файлы по истечении TTL
func (f *FileExporterHTTP) Cleanup() error {
	files, err := filepath.Glob(filepath.Join(f.cfg.Folder, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		s, err := os.Stat(file)
		if err != nil {
			return err
		}
		if s.ModTime().Before(time.Now().Add(-f.cfg.TTL)) {
			err := os.Remove(file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// возвращает HTTP-хэндлер для скачивания файла
func (f *FileExporterHTTP) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		fileName := path.Base(urlPath)
		urlPath = filepath.Join(f.cfg.Folder, fileName)

		file, err := os.Open(urlPath)

		if os.IsNotExist(err) {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(urlPath))
		http.ServeFile(w, r, urlPath)
	})
}
