package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebCacheAPI(t *testing.T) {
	cache, revision, ok := CacheName("caches/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "caches", cache.Collection)
	assert.Empty(t, revision)
	assert.Equal(t, int64(3134441396375598), cache.Id)

	cache, revision, ok = CacheName("caches/3134441396375598@abcd1234")
	require.True(t, ok)
	assert.Equal(t, "caches", cache.Collection)
	assert.Equal(t, "abcd1234", revision)
	assert.Equal(t, int64(3134441396375598), cache.Id)

	cacheFull, _, ok := CacheNameFullName("//webcache.animeapis.com/caches/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, cache.String(), cacheFull.String())
}
