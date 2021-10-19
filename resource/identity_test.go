package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentityAPI(t *testing.T) {
	user, ok := UserName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", user.collection)
	assert.Equal(t, int64(3134441396375598), user.id)

	userFull, ok := UserFullName("//identity.animeapis.com/users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, user.String(), userFull.String())

	group, ok := GroupName("groups/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "groups", group.collection)
	assert.Equal(t, int64(3134441396375598), group.id)

	groupFull, ok := GroupFullName("//identity.animeapis.com/groups/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, group.String(), groupFull.String())
}
