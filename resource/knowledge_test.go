package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKnowledgeAPI(t *testing.T) {
	parent, ok := ContributionParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	contribution, ok := ContributionName("users/3134441396375598/contributions/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "contributions", contribution.collection)
	assert.Equal(t, int64(6097286400577570), contribution.id)
	assert.Equal(t, "users", contribution.Parent.collection)
	assert.Equal(t, int64(3134441396375598), contribution.Parent.id)

	contributionFull, ok := ContributionFullName("//knowledge.animeapis.com/users/3134441396375598/contributions/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, contribution.String(), contributionFull.String())

}
