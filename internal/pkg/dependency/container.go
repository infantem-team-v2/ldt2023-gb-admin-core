package dependency

import (
	"fmt"
	"github.com/sarulabs/di"
)

// TDependencyContainer is a Dependency Injection Singleton container which allows automatically inject dependencies in structs
// Typical scenario is a NewDependencyContainer().BuildDependencies().BuildContainer()
type TDependencyContainer struct {
	// builder is a di dependencies builder
	builder *di.Builder
	// container is a di container which will built after builder were fullfilled w/ di
	container di.Container
	// dependencies is a map of fabrics or constructors w/ di
	dependencies map[string]func(ctn di.Container) (interface{}, error)
}

// NewDependencyContainer is a fabric of TDependencyContainer which returns pointer to this structure
func NewDependencyContainer() *TDependencyContainer {
	_builder, err := di.NewBuilder()
	if err != nil {
		panic(fmt.Errorf("can't create di builder: %s", err.Error()))
	}
	if err != nil {
		panic(fmt.Errorf("can't create di builder: can't get config: %s", err.Error()))
	}
	tdc := &TDependencyContainer{
		builder: _builder,
	}
	tdc.importDependencies()
	return tdc
}

// ContainerInstance Returns instance of inner container to work w/ dependencies that was put in this one
func (tdc *TDependencyContainer) ContainerInstance() di.Container {
	if tdc.container != nil {
		return tdc.container
	} else {
		panic("no container was initialized")
	}
}
