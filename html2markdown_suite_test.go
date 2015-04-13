package html2markdown_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHtml2markdown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html2markdown Suite")
}
