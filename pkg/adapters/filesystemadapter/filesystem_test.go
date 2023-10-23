package filesystemadapter_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsecaseSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ADAPTERS:FILESYSTEM")
}
