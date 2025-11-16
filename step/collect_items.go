package step

import (
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Path        string
	Key         string
	ContentType string
}

func (u Uploader) collectItems(path string) ([]Item, error) {
	var items []Item

	err := filepath.WalkDir(path, func(pathOnDisk string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			// S3 does not require creating directories; they are implied by keys.
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			u.logger.Warnf("Skipping non-regular file: %s\n", pathOnDisk)
			return nil
		}

		relativePath, err := filepath.Rel(path, pathOnDisk)
		if err != nil {
			return err
		}
		if relativePath == "." {
			relativePath = d.Name()
		}
		key := filepath.ToSlash(relativePath)

		contentType := mime.TypeByExtension(strings.ToLower(filepath.Ext(pathOnDisk)))
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		items = append(items, Item{
			Path:        pathOnDisk,
			Key:         key,
			ContentType: contentType,
		})

		return nil
	})
	return items, err
}
