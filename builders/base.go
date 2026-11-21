package builders

import "github.com/et891/go-api-autotests/fake"

type Builder struct {
	generator *fake.Fake
}

func New(generator *fake.Fake) *Builder {
	return &Builder{generator: generator}
}
