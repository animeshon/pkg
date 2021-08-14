package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageAPI(t *testing.T) {
	parent, ok := AlbumParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.Collection)
	assert.Equal(t, int64(3134441396375598), parent.Id)

	album, ok := AlbumName("users/3134441396375598/albums/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "albums", album.Collection)
	assert.Equal(t, int64(6097286400577570), album.Id)
	assert.Equal(t, "users", album.Parent.Collection)
	assert.Equal(t, int64(3134441396375598), album.Parent.Id)

	image, ok := ImageName("users/3134441396375598/albums/6097286400577570/images/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, "images", image.Collection)
	assert.Equal(t, int64(6097611928899618), image.Id)
	assert.Equal(t, "albums", image.Parent.Collection)
	assert.Equal(t, int64(6097286400577570), image.Parent.Id)
	assert.Equal(t, "users", image.Parent.Parent.Collection)
	assert.Equal(t, int64(3134441396375598), image.Parent.Parent.Id)
}
