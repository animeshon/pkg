package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCredentialsAPI(t *testing.T) {
	parent, ok := UserName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	credentials, ok := CredentialsName("users/3134441396375598/credentials/myanimelist-net")
	require.True(t, ok)
	assert.Equal(t, "credentials", credentials.collection)
	assert.Equal(t, "myanimelist-net", credentials.id)
	assert.Equal(t, "users", credentials.Parent.collection)
	assert.Equal(t, int64(3134441396375598), credentials.Parent.id)

	flow, ok := FlowName("users/3134441396375598/flows/myanimelist-net")
	require.True(t, ok)
	assert.Equal(t, "flows", flow.collection)
	assert.Equal(t, "myanimelist-net", flow.id)
	assert.Equal(t, "users", flow.Parent.collection)
	assert.Equal(t, int64(3134441396375598), flow.Parent.id)

	credentialsFull, ok := CredentialsFullName("//credentials.animeapis.com/users/3134441396375598/credentials/myanimelist-net")
	require.True(t, ok)
	assert.Equal(t, credentials.String(), credentialsFull.String())

	flowFull, ok := FlowFullName("//credentials.animeapis.com/users/3134441396375598/flows/myanimelist-net")
	require.True(t, ok)
	assert.Equal(t, flow.String(), flowFull.String())
}
