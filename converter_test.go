package html2markdown_test

import (
	"github.com/dgodd/html2markdown"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converter", func() {
	It("Converts <b>", func() {
		html := "Hi <b>Bob</b>, from Jane"
		Expect(html2markdown.Convert(html)).To(Equal("Hi **Bob**, from Jane"))
	})

	It("Handles headings", func() {
		html := `<BS><h2>Subheading 1</h2>
<h3>Third-level <em>heading</em></h3>
<h4>Fourth-level</h4>
<h1>heading #2</h1>
<h6>Six hashes</h6>
<h6>Six '#'es</h6>`

		markdown := `## Subheading 1
### Third-level *heading*
#### Fourth-level
# heading #2
###### Six hashes
###### Six '#'es`

		Expect(html2markdown.Convert(html)).To(Equal(markdown))
	})

	It("Converts hr tags", func() {
		html := `Some Stuff
<hr>
Other Stuff`

		markdown := `Some Stuff
***
Other Stuff`

		Expect(html2markdown.Convert(html)).To(Equal(markdown))
	})
})
