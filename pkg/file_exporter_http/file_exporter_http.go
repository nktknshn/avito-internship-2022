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
	"strings"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	Folder string
	TTL    time.Duration
	URL    string
	Zip    bool
}

func newRandomPrefix() string {
	return uuid.New().String()
}

type FileExporterHTTP struct {
	cfg  Config
	stop chan struct{}
}

// New создает новый экспортер файлов,
// который записывает файл в локальную файловую систему
// и предоставляет HTTP-хэндлер для скачивания файла
// очищает файлы по истечении TTL
func New(cfg Config) (*FileExporterHTTP, error) {

	if cfg.Folder == "" {
		return nil, errors.New("folder is required")
	}

	err := os.MkdirAll(cfg.Folder, 0755)
	if err != nil {
		return nil, err
	}

	return &FileExporterHTTP{cfg: cfg, stop: make(chan struct{})}, nil
}

// ExportFile экспортирует файл в локальную файловую систему и возвращает URL-адрес файла
func (f *FileExporterHTTP) ExportFile(ctx context.Context, name string, data io.Reader) (string, error) {
	if f.cfg.Zip {
		return f.exportFileZip(ctx, name, data)
	}
	return f.exportFilePlain(ctx, name, data)
}

func (f *FileExporterHTTP) exportFilePlain(_ context.Context, name string, data io.Reader) (string, error) {
	prefix := newRandomPrefix()
	fileName := prefix + "_" + name
	filePath := path.Join(f.cfg.Folder, fileName)

	if !validateFilePath(filePath, f.cfg.Folder) {
		return "", errors.New("invalid file path")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, data)
	if err != nil {
		return "", err
	}

	return path.Clean(f.cfg.URL + "/" + fileName), nil
}

func (f *FileExporterHTTP) exportFileZip(_ context.Context, name string, data io.Reader) (string, error) {
	prefix := newRandomPrefix()
	fileName := prefix + "_" + name
	zipFileName := fileName + ".zip"
	zipFilePath := path.Join(f.cfg.Folder, zipFileName)

	if !validateFilePath(zipFilePath, f.cfg.Folder) {
		return "", errors.New("invalid file path")
	}

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

	return path.Clean(f.cfg.URL + "/" + zipFileName), nil
}

func (f *FileExporterHTTP) Stop() {
	close(f.stop)
}

// Cleanup очищает файлы по истечении TTL
func (f *FileExporterHTTP) Cleanup() error {
	files, err := filepath.Glob(path.Join(f.cfg.Folder, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		s, statErr := os.Stat(file)
		if statErr != nil {
			return statErr
		}
		if s.ModTime().Before(time.Now().Add(-f.cfg.TTL)) {
			removeErr := os.Remove(file)
			if removeErr != nil {
				return removeErr
			}
		}
	}

	return nil
}

func (f *FileExporterHTTP) CleanupWorker(ctx context.Context) {
	go func() {
		for {
			select {
			case <-f.stop:
			case <-ctx.Done():
				// logger.Info("File exporter cleanup goroutine finished")
				return
			case <-time.After(f.cfg.TTL):
				// logger.Info("Running file exporter cleanup")
				f.Cleanup()
			}
		}
	}()
}

// GetHandler возвращает HTTP-хэндлер для скачивания файла
func (f *FileExporterHTTP) GetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		fileName := path.Base(urlPath)
		urlPath = path.Join(f.cfg.Folder, fileName)

		if !validateFilePath(urlPath, f.cfg.Folder) {
			http.Error(w, "invalid file path", http.StatusBadRequest)
			return
		}

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
		w.Header().Set("Content-Disposition", "attachment; filename="+path.Base(urlPath))
		http.ServeFile(w, r, urlPath)
	})
}

// validateFilePath проверяет, является ли путь до файла вложенным в папку folder
func validateFilePath(filePath string, folder string) bool {
	cleanPath := path.Clean(filePath)
	return strings.HasPrefix(cleanPath, folder)
}
