package step

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollectItems(t *testing.T) {
	tmpDir := t.TempDir()
	createFile(t, filepath.Join(tmpDir, "1.txt"), "1")
	createFile(t, filepath.Join(tmpDir, "a/2.txt"), "2")
	createFile(t, filepath.Join(tmpDir, "a/b/c/3.txt"), "3")
	createFile(t, filepath.Join(tmpDir, "b/4.txt"), "4")
	createFile(t, filepath.Join(tmpDir, "b/d/5"), "5")

	uploader := Uploader{}
	items, err := uploader.collectItems(tmpDir)
	require.NoError(t, err)

	expectedItems := []Item{
		{
			Path:        filepath.Join(tmpDir, "1.txt"),
			Key:         "1.txt",
			ContentType: "text/plain; charset=utf-8",
		},
		{
			Path:        filepath.Join(tmpDir, "a/2.txt"),
			Key:         "a/2.txt",
			ContentType: "text/plain; charset=utf-8",
		},
		{
			Path:        filepath.Join(tmpDir, "a/b/c/3.txt"),
			Key:         "a/b/c/3.txt",
			ContentType: "text/plain; charset=utf-8",
		},
		{
			Path:        filepath.Join(tmpDir, "b/4.txt"),
			Key:         "b/4.txt",
			ContentType: "text/plain; charset=utf-8",
		},
		{
			Path:        filepath.Join(tmpDir, "b/d/5"),
			Key:         "b/d/5",
			ContentType: "application/octet-stream",
		},
	}

	require.ElementsMatch(t, expectedItems, items)
}

func createFile(t *testing.T, path, content string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0o755)
	require.NoError(t, err)

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	require.NoError(t, err)

	defer f.Close()

	_, err = f.Write([]byte(content))
	require.NoError(t, err)
}
