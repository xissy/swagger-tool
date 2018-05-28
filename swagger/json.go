package swagger

import (
	"encoding/json"

	"github.com/xissy/swagger-tool/model"
)

func ParseJSONFile(jsonFilename string) (*model.Swagger, error) {
	jsonString, err := ReadFile(jsonFilename)
	if err != nil {
		return nil, err
	}

	return parseJSON(jsonString)
}

func parseJSON(jsonString string) (*model.Swagger, error) {
	var swaggerObj model.Swagger

	err := json.Unmarshal([]byte(jsonString), &swaggerObj)
	if err != nil {
		return nil, err
	}

	return &swaggerObj, nil
}

func StringifySwagger(swagger *model.Swagger) (string, error) {
	data, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
