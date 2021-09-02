package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentityAPI(t *testing.T) {
	user, ok := UserName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", user.Collection)
	assert.Equal(t, int64(3134441396375598), user.Id)

	userFull, ok := UserNameFullName("//identity.animeapis.com/users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, user.String(), userFull.String())
}
