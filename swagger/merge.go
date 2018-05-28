package swagger

import (
	"github.com/pkg/errors"

	"github.com/xissy/swagger-tool/model"
)

var (
	ErrEmptySlice              = errors.New("the slice is empty")
	ErrPathAlreadyExists       = errors.New("the path already exists")
	ErrDefinitionAlreadyExists = errors.New("the definition already exists")
)

func MergeSwaggerFiles(inputFilenames []string, swaggerHeader *model.Swagger) (string, error) {
	var swaggers []*model.Swagger

	for _, inputFilename := range inputFilenames {
		swagger, err := ParseJSONFile(inputFilename)
		if err != nil {
			return "", err
		}
		swaggers = append(swaggers, swagger)
	}

	swagger, err := mergeSwaggers(swaggers, swaggerHeader)
	if err != nil {
		return "", err
	}

	return StringifySwagger(swagger)
}

func mergeSwaggers(swaggers []*model.Swagger, swaggerHeader *model.Swagger) (*model.Swagger, error) {
	switch len(swaggers) {
	case 0:
		return nil, ErrEmptySlice
	case 1:
		return swaggers[0], nil
	}

	if swaggerHeader == nil {
		swaggerHeader = NewSwaggerHeader("API", "v1")
	}

	var pathMaps []map[string]*model.Path
	var definitionMaps []map[string]*model.Definition

	for _, swagger := range swaggers {
		pathMaps = append(pathMaps, swagger.Paths)
		definitionMaps = append(definitionMaps, swagger.Definitions)
	}

	mergedPaths, err := mergePaths(pathMaps)
	if err != nil {
		return nil, err
	}
	mergedDefinitions, err := mergeDefinitions(definitionMaps)
	if err != nil {
		return nil, err
	}

	mergesSwagger := &model.Swagger{
		Swagger:     swaggerHeader.Swagger,
		Info:        swaggerHeader.Info,
		Schemes:     swaggerHeader.Schemes,
		Consumes:    swaggerHeader.Consumes,
		Produces:    swaggerHeader.Produces,
		Paths:       mergedPaths,
		Definitions: mergedDefinitions,
	}

	return mergesSwagger, nil
}

func NewSwaggerHeader(title, version string) *model.Swagger {
	return &model.Swagger{
		Swagger: "2.0",
		Info: map[string]string{
			"title":   title,
			"version": version,
		},
		Schemes: []string{
			"http",
			"https",
		},
		Consumes: []string{
			"application/json",
		},
		Produces: []string{
			"application/json",
		},
	}
}

func mergePaths(pathMaps []map[string]*model.Path) (map[string]*model.Path, error) {
	mergedPaths := make(map[string]*model.Path)

	for _, pathMap := range pathMaps {
		for path, value := range pathMap {
			_, ok := mergedPaths[path]
			if ok {
				return nil, errors.Wrap(ErrPathAlreadyExists, path)
			}
			mergedPaths[path] = value
		}
	}

	return mergedPaths, nil
}

func mergeDefinitions(definitionMaps []map[string]*model.Definition) (map[string]*model.Definition, error) {
	mergedDefinitions := make(map[string]*model.Definition)

	for _, definitionMap := range definitionMaps {
		for definition, value := range definitionMap {
			_, ok := mergedDefinitions[definition]
			if ok {
				return nil, errors.Wrap(ErrDefinitionAlreadyExists, definition)
			}
			mergedDefinitions[definition] = value
		}
	}

	return mergedDefinitions, nil
}
