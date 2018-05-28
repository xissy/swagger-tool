package swagger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
	"github.com/xissy/swagger-tool/model"
)

func TestMergePathMaps(t *testing.T) {
	pathMap1 := make(map[string]*model.Path)
	pathMap1["/a"] = nil
	pathMap2 := make(map[string]*model.Path)
	pathMap2["/b"] = nil
	pathMap3 := make(map[string]*model.Path)
	pathMap3["/c"] = nil
	pathMaps := []map[string]*model.Path{pathMap1, pathMap2, pathMap3}

	mergedPathMaps, err := mergePathMaps(pathMaps)
	assert.Nil(t, err)
	assert.NotNil(t, mergedPathMaps)
	assert.Len(t, mergedPathMaps, 3)
}

func TestMergePathMaps_Error_ErrPathAlreadyExists(t *testing.T) {
	pathMap1 := make(map[string]*model.Path)
	pathMap1["sameKey"] = nil
	pathMap2 := make(map[string]*model.Path)
	pathMap2["sameKey"] = nil
	pathMaps := []map[string]*model.Path{pathMap1, pathMap2}

	mergedPathMaps, err := mergePathMaps(pathMaps)
	assert.Nil(t, mergedPathMaps)
	assert.NotNil(t, err)
	assert.EqualError(t, err, errors.Wrap(ErrPathAlreadyExists, "sameKey").Error())
}

func TestMergeDefinitionMaps(t *testing.T) {
	defMap1 := make(map[string]*model.Definition)
	defMap1["a"] = nil
	defMap2 := make(map[string]*model.Definition)
	defMap2["b"] = nil
	defMap3 := make(map[string]*model.Definition)
	defMap3["c"] = nil
	defMaps := []map[string]*model.Definition{defMap1, defMap2, defMap3}

	mergedDefMaps, err := mergeDefinitionMaps(defMaps)
	assert.Nil(t, err)
	assert.NotNil(t, mergedDefMaps)
	assert.Len(t, mergedDefMaps, 3)
}

func TestMergeDefinitionMaps_error_ErrDefinition(t *testing.T) {
	defMap1 := make(map[string]*model.Definition)
	defMap1["sameKey"] = nil
	defMap2 := make(map[string]*model.Definition)
	defMap2["sameKey"] = nil
	defMaps := []map[string]*model.Definition{defMap1, defMap2}

	mergedDefMaps, err := mergeDefinitionMaps(defMaps)
	assert.Nil(t, mergedDefMaps)
	assert.NotNil(t, err)
	assert.EqualError(t, err, errors.Wrap(ErrDefinitionAlreadyExists, "sameKey").Error())
}
