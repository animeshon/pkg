package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLibraryAPI(t *testing.T) {
	parent, ok := PlaylistParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	playlist, ok := PlaylistName("users/3134441396375598/playlists/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "playlists", playlist.collection)
	assert.Equal(t, int64(6097286400577570), playlist.id)
	assert.Equal(t, "users", playlist.Parent.collection)
	assert.Equal(t, int64(3134441396375598), playlist.Parent.id)

	alias, ok := PlaylistName("users/3134441396375598/playlists/later")
	require.True(t, ok)
	assert.Equal(t, "playlists", alias.collection)
	assert.Equal(t, "later", alias.id)
	assert.Equal(t, "users", alias.Parent.collection)
	assert.Equal(t, int64(3134441396375598), alias.Parent.id)

	item, ok := PlaylistItemName("users/3134441396375598/playlists/6097286400577570/items/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, "items", item.collection)
	assert.Equal(t, int64(6097611928899618), item.id)
	assert.Equal(t, "playlists", item.Parent.collection)
	assert.Equal(t, int64(6097286400577570), item.Parent.id)
	assert.Equal(t, "users", item.Parent.Parent.collection)
	assert.Equal(t, int64(3134441396375598), item.Parent.Parent.id)

	playlistFull, ok := PlaylistFullName("//library.animeapis.com/users/3134441396375598/playlists/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, playlist.String(), playlistFull.String())

	itemFull, ok := PlaylistItemFullName("//library.animeapis.com/users/3134441396375598/playlists/6097286400577570/items/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, item.String(), itemFull.String())
}
