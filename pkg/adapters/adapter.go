package adapters

import "github.com/VictorMarcolino/artifact-manipulator/pkg/adapters/filesystemadapter"

var (
	NewTemporaryFileSystemArtifactRepository = filesystemadapter.NewTemporaryFileSystemArtifactRepository
	NewFileSystemArtifactRepository          = filesystemadapter.NewFileSystemArtifactRepository
)
