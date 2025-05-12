package file_exporter_http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFileExporterHTTP_ExportFile(t *testing.T) {
	dir := t.TempDir()

	cfg := Config{
		Folder: dir,
		URL:    "",
		Zip:    false,
		TTL:    time.Millisecond * 200,
	}

	fe, err := New(cfg)
	require.NoError(t, err)

	handler := fe.GetHandler()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	filePath, err := fe.ExportFile(t.Context(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	resp, err := http.Get(ts.URL + "/" + filepath.Base(filePath))
	require.NoError(t, err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, "test", string(body))
}

func TestFileExporterHTTP_Cleanup(t *testing.T) {
	dir := t.TempDir()

	cfg := Config{
		Folder: dir,
		TTL:    time.Millisecond * 200,
		URL:    "/data/report_revenue_export",
		Zip:    false,
	}

	fe, err := New(cfg)
	require.NoError(t, err)

	filePath, err := fe.ExportFile(t.Context(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	filePath2, err := fe.ExportFile(t.Context(), "test.txt", strings.NewReader("test"))
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 120)

	require.NoError(t, fe.Cleanup())

	_, err = os.Stat(filepath.Join(dir, filepath.Base(filePath)))
	require.Error(t, err)

	_, err = os.Stat(filepath.Join(dir, filepath.Base(filePath2)))
	require.NoError(t, err)
}

func TestFileExporterHTTP_ValidateFilePath(t *testing.T) {
	folder := "/data/report_revenue_export"

	require.True(t, validateFilePath("/data/report_revenue_export/test.txt", folder))
	require.True(t, validateFilePath("/data/report_revenue_export/abc/../test.txt", folder))
	require.False(t, validateFilePath("/data/report_revenue_export/../../test.txt", folder))
}
