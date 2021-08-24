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
}
