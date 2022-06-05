package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleaseAPI(t *testing.T) {
	parent, ok := ReleaseParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	release, ok := ReleaseName("users/3134441396375598/releases/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "releases", release.collection)
	assert.Equal(t, int64(6097286400577570), release.id)
	assert.Equal(t, "users", release.Parent.collection)
	assert.Equal(t, int64(3134441396375598), release.Parent.id)

	release, ok = ReleaseName("organizations/3134441396375598/releases/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "releases", release.collection)
	assert.Equal(t, int64(6097286400577570), release.id)
	assert.Equal(t, "organizations", release.Parent.collection)
	assert.Equal(t, int64(3134441396375598), release.Parent.id)

	releaseFull, ok := ReleaseFullName("//release.animeapis.com/organizations/3134441396375598/releases/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, release.String(), releaseFull.String())
}
