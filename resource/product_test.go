package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductAPI(t *testing.T) {
	parent, ok := ProductChapterParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	chapter, ok := ProductChapterName("users/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "chapters", chapter.collection)
	assert.Equal(t, int64(6097286400577570), chapter.id)
	assert.Equal(t, "users", chapter.Parent.collection)
	assert.Equal(t, int64(3134441396375598), chapter.Parent.id)

	chapter, ok = ProductChapterName("organizations/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "chapters", chapter.collection)
	assert.Equal(t, int64(6097286400577570), chapter.id)
	assert.Equal(t, "organizations", chapter.Parent.collection)
	assert.Equal(t, int64(3134441396375598), chapter.Parent.id)

	chapterFull, ok := ProductChapterFullName("//product.animeapis.com/organizations/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, chapter.String(), chapterFull.String())
}
