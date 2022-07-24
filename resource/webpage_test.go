package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebPageAPI(t *testing.T) {
	site, ok := SiteName("sites/animeshon-com")
	require.True(t, ok)
	assert.Equal(t, "sites", site.collection)
	assert.Equal(t, string("animeshon-com"), site.id)

	page, ok := PageName("sites/animeshon-com/pages/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "pages", page.collection)
	assert.Equal(t, int64(6097286400577570), page.id)
	assert.Equal(t, "sites", page.Parent.collection)
	assert.Equal(t, string("animeshon-com"), page.Parent.id)

	siteFull, ok := SiteFullName("//webpage.animeapis.com/sites/animeshon-com")
	require.True(t, ok)
	assert.Equal(t, site.String(), siteFull.String())

	pageFull, ok := PageFullName("//webpage.animeapis.com/sites/animeshon-com/pages/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, page.String(), pageFull.String())
}
