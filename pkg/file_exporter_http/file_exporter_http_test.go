package file_exporter_http_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/nktknshn/avito-internship-2022/pkg/file_exporter_http"
	"github.com/stretchr/testify/require"
)

func TestFileExporterHTTP_ExportFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "file_exporter_http_test")

	require.NoError(t, err)

	defer os.RemoveAll(dir)

	cfg := file_exporter_http.Config{
		Folder: dir,
		URL:    "",
	}

	fe, err := file_exporter_http.New(cfg)
	require.NoError(t, err)

	handler := fe.GetHandler()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	filePath, err := fe.ExportFile(context.Background(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	resp, err := http.Get(ts.URL + "/" + filepath.Base(filePath))
	require.NoError(t, err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, "test", string(body))
}

func TestFileExporterHTTP_Cleanup(t *testing.T) {
	dir, err := os.MkdirTemp("", "file_exporter_http_test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	cfg := file_exporter_http.Config{
		Folder: dir,
		TTL:    time.Millisecond * 200,
		URL:    "/data/report_revenue_export",
	}

	fe, err := file_exporter_http.New(cfg)
	require.NoError(t, err)

	filePath, err := fe.ExportFile(context.Background(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	filePath2, err := fe.ExportFile(context.Background(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 120)

	require.NoError(t, fe.Cleanup())

	_, err = os.Stat(filepath.Join(dir, filepath.Base(filePath)))
	require.Error(t, err)

	_, err = os.Stat(filepath.Join(dir, filepath.Base(filePath2)))
	require.NoError(t, err)
}
