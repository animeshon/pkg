package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIAMAPI(t *testing.T) {
	role, ok := RoleName("roles/iam.test")
	require.True(t, ok)
	assert.Equal(t, "roles", role.collection)
	assert.Equal(t, "iam.test", role.id)

	roleFull, ok := RoleFullName("//iam.animeapis.com/roles/iam.test")
	require.True(t, ok)
	assert.Equal(t, role.String(), roleFull.String())

	permission, ok := PermissionName("permissions/iam.test")
	require.True(t, ok)
	assert.Equal(t, "permissions", permission.collection)
	assert.Equal(t, "iam.test", permission.id)

	permissionFull, ok := PermissionFullName("//iam.animeapis.com/permissions/iam.test")
	require.True(t, ok)
	assert.Equal(t, permission.String(), permissionFull.String())

	serviceAccount, ok := ServiceAccountName("users/3134441396375598/serviceAccounts/system")
	require.True(t, ok)
	assert.Equal(t, "serviceAccounts", serviceAccount.collection)
	assert.Equal(t, "system", serviceAccount.id)
	assert.Equal(t, "users", serviceAccount.Parent.collection)
	assert.Equal(t, int64(3134441396375598), serviceAccount.Parent.id)

	serviceAccountFull, ok := ServiceAccountFullName("//iam.animeapis.com/users/3134441396375598/serviceAccounts/system")
	require.True(t, ok)
	assert.Equal(t, serviceAccount.String(), serviceAccountFull.String())
}
