package randSequence_test

import (
	"fmt"

	. "github.com/nicday/randSequence"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RandSequence", func() {
	generators := map[string]*RandGenerator{
		"default": New(),
		"numeric": &RandGenerator{
			Characters: []rune("1234567890"),
		},
		"special": &RandGenerator{
			Characters: []rune("!@#$%^&*()"),
		},
	}

	for name := range generators {
		generator := generators[name]

		Context(fmt.Sprintf("with the %s RandGenerator", name), func() {
			Describe("#New", func() {
				It("generates a sequence of a given length", func() {
					for l := range []int{0, 1, 2, 3, 10, 100} {
						s := generator.Generate(l)

						Expect(len(s)).To(Equal(l))
					}
				})

				It("generates random sequences", func() {
					s1 := generator.Generate(5)
					s2 := generator.Generate(5)

					Expect(s1).NotTo(Equal(s2))
				})

				It("only generates from the source characters", func() {
					source := generator.Characters
					sourceLen := len(source)

					s := generator.Generate(sourceLen * 20)

					for _, c := range s {
						Expect(source).To(ContainElement(c))
					}
				})
			})
		})
	}
})
