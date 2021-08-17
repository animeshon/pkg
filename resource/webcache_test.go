package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebCacheAPI(t *testing.T) {
	parent, revision, ok := CacheName("caches/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "caches", parent.Collection)
	assert.Empty(t, revision)
	assert.Equal(t, int64(3134441396375598), parent.Id)

	parent, revision, ok = CacheName("caches/3134441396375598@abcd1234")
	require.True(t, ok)
	assert.Equal(t, "caches", parent.Collection)
	assert.Equal(t, "abcd1234", revision)
	assert.Equal(t, int64(3134441396375598), parent.Id)
}
