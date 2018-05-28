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
	key := "key"
	pathMap1 := make(map[string]*model.Path)
	pathMap1[key] = nil
	pathMap2 := make(map[string]*model.Path)
	pathMap2[key] = nil
	pathMaps := []map[string]*model.Path{pathMap1, pathMap2}

	mergedPathMaps, err := mergePathMaps(pathMaps)
	assert.Nil(t, mergedPathMaps)
	assert.NotNil(t, err)
	assert.EqualError(t, err, errors.Wrap(ErrPathAlreadyExists, key).Error())
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

func TestMergeDefinitionMaps_Error_ErrDefinition(t *testing.T) {
	key := "key"
	defMap1 := make(map[string]*model.Definition)
	defMap1[key] = nil
	defMap2 := make(map[string]*model.Definition)
	defMap2[key] = nil
	defMaps := []map[string]*model.Definition{defMap1, defMap2}

	mergedDefMaps, err := mergeDefinitionMaps(defMaps)
	assert.Nil(t, mergedDefMaps)
	assert.NotNil(t, err)
	assert.EqualError(t, err, errors.Wrap(ErrDefinitionAlreadyExists, key).Error())
}
