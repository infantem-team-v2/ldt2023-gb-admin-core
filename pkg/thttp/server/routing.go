package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"io"
	"log"

	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

const (
	wrapLabel = "can't open file with routes on %s"
)

var (
	ErrEmptyFile = fmt.Errorf("there's no data in file")
)

// MapRoutes to r fiber.Router via reflection by passing handler struct which has functions that returns fiber.Handler
func MapRoutes(h IHandler) {
	hVal := reflect.ValueOf(h)
	hType := hVal.Type()
	hMapping, err := ParseRoutes(fmt.Sprintf("/http/%s.yaml", h.GetPrefix()))
	if err != nil {
		panic(err)
	}
	for i := 0; i < hType.NumMethod(); i++ {
		fHandler := hType.Method(i)
		if hMapping[fHandler.Name] == nil {
			continue
		}
		fStruct := hMapping[fHandler.Name]
		fRes := fHandler.Func
		//if !ok {
		//	panic(fmt.Errorf("can't map method %s. it doesn't implement fiber.Handler", fHandler.Name))
		//}
		f := fRes.Call([]reflect.Value{hVal})
		h.GetRouter().Add(fStruct.HttpMethod, fStruct.Route, f[0].Interface().(fiber.Handler))
	}
}

// ParseRoutes for mapping handlers on fiber.App
func ParseRoutes(path string) (routeMapping map[string]*RouteMapping, err error) {
	workDirection, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(wrapLabel, path))
	}
	log.Printf("Parsing routes from: %s", workDirection+path)
	fMap, err := os.Open(workDirection + path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(wrapLabel, path))
	}
	bMap, err := io.ReadAll(fMap)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(wrapLabel, path))
	}
	if len(bMap) == 0 {
		return nil, errors.Wrap(ErrEmptyFile, fmt.Sprintf(wrapLabel, path))
	}
	log.Printf("Yaml from file:\n %s", bMap)
	routeMapping = make(map[string]*RouteMapping)
	if err := yaml.Unmarshal(bMap, &routeMapping); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(wrapLabel, path))
	}
	return routeMapping, nil
}
