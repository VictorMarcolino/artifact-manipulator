package filesystemadapter_test

import (
	"github.com/VictorMarcolino/artifact-manipulator/pkg/adapters/filesystemadapter"
	"github.com/VictorMarcolino/artifact-manipulator/pkg/core/ports"
	"github.com/VictorMarcolino/artifact-manipulator/pkg/tests/specs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ADAPTERS:FILESYSTEM:ARTIFACT ->", func() {
	var adapter ports.ArtifactRepositoryI
	BeforeEach(func() {
		var err error
		adapter, err = filesystemadapter.NewTemporaryFileSystemArtifactRepository()

		Expect(err).NotTo(HaveOccurred())
	})

	specs.ArtifactRepositorySpec(func() ports.ArtifactRepositoryI {
		return adapter
	})
})
