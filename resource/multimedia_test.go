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

	chapter, ok = ChapterName("lightNovels/3134441396375598/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "chapters", chapter.collection)
	assert.Equal(t, int64(6097286400577570), chapter.id)
	assert.Equal(t, "lightNovels", chapter.Parent.collection)
	assert.Equal(t, int64(3134441396375598), chapter.Parent.id)

	chapterWildcard, ok := ChapterName("lightNovels/-/chapters/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "chapters", chapterWildcard.collection)
	assert.Equal(t, int64(6097286400577570), chapterWildcard.id)
	assert.Equal(t, "lightNovels", chapterWildcard.Parent.collection)
	assert.Equal(t, "-", chapterWildcard.Parent.id)

	chapterFull, ok := ChapterFullName("//multimedia.animeapis.com/lightNovels/3134441396375598/chapters/6097286400577570")
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

	episodeWildcard, ok := EpisodeName("animes/-/episodes/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "episodes", episodeWildcard.collection)
	assert.Equal(t, int64(6097286400577570), episodeWildcard.id)
	assert.Equal(t, "animes", episodeWildcard.Parent.collection)
	assert.Equal(t, "-", episodeWildcard.Parent.id)

	episodeFull, ok := EpisodeFullName("//multimedia.animeapis.com/animes/3134441396375598/episodes/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, episode.String(), episodeFull.String())

	anime, ok := AnimeName("animes/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "animes", anime.collection)
	assert.Equal(t, int64(3134441396375598), anime.id)

	animeFull, ok := AnimeFullName("//multimedia.animeapis.com/animes/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, anime.String(), animeFull.String())

	lightNovel, ok := LightNovelName("lightNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "lightNovels", lightNovel.collection)
	assert.Equal(t, int64(3134441396375598), lightNovel.id)

	lightNovelFull, ok := LightNovelFullName("//multimedia.animeapis.com/lightNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, lightNovel.String(), lightNovelFull.String())

	graphicNovel, ok := GraphicNovelName("graphicNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "graphicNovels", graphicNovel.collection)
	assert.Equal(t, int64(3134441396375598), graphicNovel.id)

	graphicNovelFull, ok := GraphicNovelFullName("//multimedia.animeapis.com/graphicNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, graphicNovel.String(), graphicNovelFull.String())

	visualNovel, ok := VisualNovelName("visualNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "visualNovels", visualNovel.collection)
	assert.Equal(t, int64(3134441396375598), visualNovel.id)

	visualNovelFull, ok := VisualNovelFullName("//multimedia.animeapis.com/visualNovels/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, visualNovel.String(), visualNovelFull.String())
}
