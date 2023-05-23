package dependency

import (
	"fmt"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"reflect"
)

// importDependencies Method that imports dependencies from outer source
func (tdc *TDependencyContainer) importDependencies() {
	// for now, we will import dependencies from constant map w/ name : builder func
	tdc.dependencies = dependencyMap

}

// BuildDependencies Method that initializes all dependencies and put it in builder w/ app scope
func (tdc *TDependencyContainer) BuildDependencies() *TDependencyContainer {
	var err error
	for name, builder := range tdc.dependencies {
		err = tdc.builder.Add(di.Def{
			Name:  name,
			Build: builder,
			Scope: di.App,
		})
		if err != nil {
			logrus.Error(err)
			panic(fmt.Errorf("di: can't build dependency w/ name: %s", name))
		}
	}
	return tdc
}

// BuildContainer Makes container's build
func (tdc *TDependencyContainer) BuildContainer() *TDependencyContainer {
	tdc.container = tdc.builder.Build()
	return tdc
}

// Inject dependencies into injectableStruct
func (tdc *TDependencyContainer) Inject(injectableStruct interface{}) {
	injectableValue := reflect.ValueOf(injectableStruct)
	// Checking if there's nil pointer in interface
	if injectableValue.Kind() == reflect.Ptr && injectableValue.IsNil() {
		panic(fmt.Errorf("%s is a nil pointer", injectableValue.Kind().String()))
	}
	// Dereference pointer if it's not nil
	if injectableValue.Kind() == reflect.Ptr {
		injectableValue = injectableValue.Elem()
	}
	// Checking if struct was sent to method
	if injectableValue.Kind() != reflect.Struct {
		panic(fmt.Errorf("injectable object is not a struct but %s", injectableValue.Kind().String()))
	}

	injectableType := injectableValue.Type()
	// Checking all fields of struct
	for i := 0; i < injectableType.NumField(); i++ {
		// Getting tag of current field
		fieldTag := injectableType.Field(i).Tag.Get(TagDI)
		if fieldTag == "" {
			continue
		}
		// Getting dependency object from container
		dependencyObject := tdc.container.Get(fieldTag)
		// Injecting this object into injectable ones
		injectableValue.Field(i).Set(reflect.ValueOf(dependencyObject))
	}
}
