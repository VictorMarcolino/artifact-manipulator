package filesystemadapter

import (
	"io"
	"os"
	"path/filepath"
)

func (fs *ArtifactRepository) GetArtifact(id string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(fs.BaseDir, id))
}

func (fs *ArtifactRepository) PushArtifact(id string, c io.Reader) error {
	filePath := filepath.Join(fs.BaseDir, id)
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Set file permissions to r-xr-xr-x
	if err = os.Chmod(filePath, 0o755); err != nil {
		return err
	}
	_, err = io.Copy(f, c)
	return err
}

func (fs *ArtifactRepository) DeleteArtifact(id string) error {
	return os.RemoveAll(filepath.Join(fs.BaseDir, id))
}

func (fs *ArtifactRepository) CleanArtifactRepository() error {
	return os.RemoveAll(fs.BaseDir)
}

func (fs *ArtifactRepository) ListArtifacts() ([]string, error) {
	if _, err := os.Stat(fs.BaseDir); os.IsNotExist(err) {
		return []string{}, nil
	}
	var fileNames []string
	err := filepath.Walk(fs.BaseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			rel, pErr := filepath.Rel(fs.BaseDir, path)
			if pErr != nil {
				return pErr
			}
			fileNames = append(fileNames, rel)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

func (fs *ArtifactRepository) GetFullAddressOfArtifact(id string) (string, error) {
	fullPath := filepath.Join(fs.BaseDir, id)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", err
	}

	return fullPath, nil
}
