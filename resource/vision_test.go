package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVisionAPI(t *testing.T) {
	parent, ok := ImageAnnotationParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.Collection)
	assert.Equal(t, int64(3134441396375598), parent.Id)

	annotation, ok := ImageAnnotationName("users/3134441396375598/imageAnnotations/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "imageAnnotations", annotation.Collection)
	assert.Equal(t, int64(6097286400577570), annotation.Id)
	assert.Equal(t, "users", annotation.Parent.Collection)
	assert.Equal(t, int64(3134441396375598), annotation.Parent.Id)

	annotationFull, ok := ImageAnnotationFullName("//vision.animeapis.com/users/3134441396375598/imageAnnotations/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, annotation.String(), annotationFull.String())

	analysis, ok := ImageAnalysisName("users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/123456789")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.Collection)
	assert.Equal(t, int64(123456789), analysis.Id)
	assert.Equal(t, "images", analysis.Parent.Collection)
	assert.Equal(t, int64(6097611928899618), analysis.Parent.Id)
	assert.Equal(t, "albums", analysis.Parent.Parent.Collection)
	assert.Equal(t, int64(6097286400577570), analysis.Parent.Parent.Id)
	assert.Equal(t, "users", analysis.Parent.Parent.Parent.Collection)
	assert.Equal(t, int64(3134441396375598), analysis.Parent.Parent.Parent.Id)

	analysis, ok = ImageAnalysisName("users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/latest")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.Collection)
	assert.Equal(t, int64(-1), analysis.Id)

	analysis, ok = ImageAnalysisFullName("//vision.animeapis.com/users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/latest")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.Collection)
	assert.Equal(t, int64(-1), analysis.Id)
}
