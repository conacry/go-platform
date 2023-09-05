package initapp

import (
	"github.com/conacry/go-platform/pkg/errors"
)

type DependenciesInitializerBuilder struct {
	dependencyInitChain []DependencyService
	errors              *errors.Errors
}

func NewDependenciesInitializer() *DependenciesInitializerBuilder {
	return &DependenciesInitializerBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *DependenciesInitializerBuilder) DependencyInitChain(dependencyInitChain []DependencyService) *DependenciesInitializerBuilder {
	b.dependencyInitChain = dependencyInitChain
	return b
}

func (b *DependenciesInitializerBuilder) Build() (*DependenciesInitializer, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createInitializer(), nil
}

func (b *DependenciesInitializerBuilder) checkRequiredFields() {
	if b.dependencyInitChain == nil {
		b.errors.AddError(ErrDependencyInitChainIsRequired)
	}
}

func (b *DependenciesInitializerBuilder) createInitializer() *DependenciesInitializer {
	return &DependenciesInitializer{
		dependencyInitChain: b.dependencyInitChain,
	}
}
