package ports

import "io"

type ArtifactRepositoryI interface {
	PushArtifactI
	GetArtifactI
	DeleteArtifactI
	CleanArtifactRepositoryI
	ListArtifactsI
	GetFullAddressOfArtifactI
}
type GetArtifactI interface {
	GetArtifact(id string) (io.ReadCloser, error)
}
type PushArtifactI interface {
	PushArtifact(id string, c io.Reader) error
}

type DeleteArtifactI interface {
	DeleteArtifact(id string) error
}

type CleanArtifactRepositoryI interface {
	CleanArtifactRepository() error
}

type ListArtifactsI interface {
	ListArtifacts() ([]string, error)
}

type GetFullAddressOfArtifactI interface {
	GetFullAddressOfArtifact(id string) (string, error)
}
