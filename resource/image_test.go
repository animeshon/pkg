package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageAPI(t *testing.T) {
	parent, ok := AlbumParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	album, ok := AlbumName("users/3134441396375598/albums/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "albums", album.collection)
	assert.Equal(t, int64(6097286400577570), album.id)
	assert.Equal(t, "users", album.Parent.collection)
	assert.Equal(t, int64(3134441396375598), album.Parent.id)

	image, ok := ImageName("users/3134441396375598/albums/6097286400577570/images/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, "images", image.collection)
	assert.Equal(t, int64(6097611928899618), image.id)
	assert.Equal(t, "albums", image.Parent.collection)
	assert.Equal(t, int64(6097286400577570), image.Parent.id)
	assert.Equal(t, "users", image.Parent.Parent.collection)
	assert.Equal(t, int64(3134441396375598), image.Parent.Parent.id)

	albumFull, ok := AlbumFullName("//image.animeapis.com/users/3134441396375598/albums/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, album.String(), albumFull.String())

	imageFull, ok := ImageFullName("//image.animeapis.com/users/3134441396375598/albums/6097286400577570/images/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, image.String(), imageFull.String())
}
