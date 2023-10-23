package specs

import (
	"bytes"
	"github.com/VictorMarcolino/artifact-manipulator/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
)

func ArtifactRepositorySpec(repoCreator func() ports.ArtifactRepositoryI) {
	var repo ports.ArtifactRepositoryI
	var testArtifactID string = "some-artifact-id"
	var testArtifactContent = "sample content for artifact"

	BeforeEach(func() {
		repo = repoCreator()
	})

	Describe("PushArtifact", func() {
		It("should push an artifact successfully", func() {
			reader := bytes.NewReader([]byte(testArtifactContent))
			err := repo.PushArtifact(testArtifactID, reader)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("GetArtifact", func() {
		Context("When the artifact exists", func() {
			It("should retrieve the artifact by its ID", func() {
				reader := bytes.NewReader([]byte(testArtifactContent))
				_ = repo.PushArtifact(testArtifactID, reader)

				readCloser, err := repo.GetArtifact(testArtifactID)
				Expect(err).NotTo(HaveOccurred())
				content, _ := io.ReadAll(readCloser)
				Expect(string(content)).To(Equal(testArtifactContent))
			})
		})

		Context("When the artifact does not exist", func() {
			It("should return an error", func() {
				_, err := repo.GetArtifact("nonexistent-id")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("ListArtifacts", func() {
		It("should list all artifacts", func() {
			reader := bytes.NewReader([]byte(testArtifactContent))
			_ = repo.PushArtifact(testArtifactID, reader)

			artifacts, err := repo.ListArtifacts()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(artifacts)).To(Equal(1))
			Expect(artifacts[0]).To(Equal(testArtifactID))
		})
	})

	Describe("DeleteArtifact", func() {
		Context("When the artifact exists", func() {
			It("should remove an artifact by its ID", func() {
				reader := bytes.NewReader([]byte(testArtifactContent))
				_ = repo.PushArtifact(testArtifactID, reader)

				err := repo.DeleteArtifact(testArtifactID)
				Expect(err).NotTo(HaveOccurred())
				_, err = repo.GetArtifact(testArtifactID)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("When the artifact does not exist", func() {
			It("should not return error", func() {
				err := repo.DeleteArtifact("nonexistent-id")
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("CleanArtifactRepository", func() {
		It("should clean all artifacts in the repository", func() {
			reader := bytes.NewReader([]byte(testArtifactContent))
			_ = repo.PushArtifact(testArtifactID, reader)

			err := repo.CleanArtifactRepository()
			Expect(err).NotTo(HaveOccurred())

			artifacts, err := repo.ListArtifacts()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(artifacts)).To(Equal(0))
		})
	})

	Describe("GetFullAddressOfArtifact", func() {
		It("should get the full address of an artifact by its ID", func() {
			reader := bytes.NewReader([]byte(testArtifactContent))
			_ = repo.PushArtifact(testArtifactID, reader)

			address, err := repo.GetFullAddressOfArtifact(testArtifactID)
			Expect(err).NotTo(HaveOccurred())
			Expect(address).NotTo(BeEmpty())
		})
	})
}
