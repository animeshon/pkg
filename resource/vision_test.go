package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVisionAPI(t *testing.T) {
	parent, ok := ImageAnnotationParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	annotation, ok := ImageAnnotationName("users/3134441396375598/imageAnnotations/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "imageAnnotations", annotation.collection)
	assert.Equal(t, int64(6097286400577570), annotation.id)
	assert.Equal(t, "users", annotation.Parent.collection)
	assert.Equal(t, int64(3134441396375598), annotation.Parent.id)

	annotationFull, ok := ImageAnnotationFullName("//vision.animeapis.com/users/3134441396375598/imageAnnotations/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, annotation.String(), annotationFull.String())

	analysis, ok := ImageAnalysisName("users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/123456789")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.collection)
	assert.Equal(t, int64(123456789), analysis.id)
	assert.Equal(t, "images", analysis.Parent.collection)
	assert.Equal(t, int64(6097611928899618), analysis.Parent.id)
	assert.Equal(t, "albums", analysis.Parent.Parent.collection)
	assert.Equal(t, int64(6097286400577570), analysis.Parent.Parent.id)
	assert.Equal(t, "users", analysis.Parent.Parent.Parent.collection)
	assert.Equal(t, int64(3134441396375598), analysis.Parent.Parent.Parent.id)

	analysis, ok = ImageAnalysisName("users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/latest")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.collection)
	assert.Equal(t, int64(-1), analysis.id)

	analysis, ok = ImageAnalysisFullName("//vision.animeapis.com/users/3134441396375598/albums/6097286400577570/images/6097611928899618/analyses/latest")
	require.True(t, ok)
	assert.Equal(t, "analyses", analysis.collection)
	assert.Equal(t, int64(-1), analysis.id)
}
