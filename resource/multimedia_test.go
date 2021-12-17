package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultimediaAPI(t *testing.T) {
	parent, ok := ChapterParentName("graphicNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "graphicNovels", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	chapter, ok := ChapterName("graphicNovels/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "chapters", chapter.collection)
	assert.Equal(t, int64(6097286400577570), chapter.id)
	assert.Equal(t, "graphicNovels", chapter.Parent.collection)
	assert.Equal(t, int64(3134441396375598), chapter.Parent.id)

	chapterFull, ok := ChapterFullName("//multimedia.animeapis.com/graphicNovels/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, chapter.String(), chapterFull.String())

	parent, ok = EpisodeParentName("animes/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "animes", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	episode, ok := EpisodeName("animes/3134441396375598/episodes/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "episodes", episode.collection)
	assert.Equal(t, int64(6097286400577570), episode.id)
	assert.Equal(t, "animes", episode.Parent.collection)
	assert.Equal(t, int64(3134441396375598), episode.Parent.id)

	episodeFull, ok := EpisodeFullName("//multimedia.animeapis.com/animes/3134441396375598/episodes/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, episode.String(), episodeFull.String())
}
