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

	album, ok := ImageAnnotationName("users/3134441396375598/imageAnnotations/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "imageAnnotations", album.Collection)
	assert.Equal(t, int64(6097286400577570), album.Id)
	assert.Equal(t, "users", album.Parent.Collection)
	assert.Equal(t, int64(3134441396375598), album.Parent.Id)
}
