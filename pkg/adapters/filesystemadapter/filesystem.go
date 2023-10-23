package filesystemadapter

import (
	"github.com/VictorMarcolino/artifact-manipulator/pkg/core/ports"
	"os"
)

type ArtifactRepository struct {
	BaseDir string
}

var _ ports.ArtifactRepositoryI = &ArtifactRepository{}

func NewTemporaryFileSystemArtifactRepository() (*ArtifactRepository, error) {
	tempDir, err := os.MkdirTemp("", "artifact_repo_")
	if err != nil {
		return nil, err
	}
	return &ArtifactRepository{BaseDir: tempDir}, nil
}

func NewFileSystemArtifactRepository(baseDir string) (*ArtifactRepository, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, err
	}
	return &ArtifactRepository{BaseDir: baseDir}, nil
}
