package iam

import (
	"context"
	"testing"

	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	"github.com/animeshon/pkg/iam/gapictest"
	"github.com/googleapis/gax-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEverythingIsMissing(t *testing.T) {
	gapic := &gapictest.IamClient{}

	roleIterator := &gapictest.RoleIterator{}
	roleIterator.On("Next").Return(nil, iterator.Done)

	permissionIterator := &gapictest.PermissionIterator{}
	permissionIterator.On("Next").Return(nil, iterator.Done)

	notFound := status.New(codes.NotFound, "").Err()

	roles := map[string]bool{
		"roles/iam.viewer": false,
		"roles/iam.admin":  false,
		"roles/viewer":     false,
	}

	fnCreateRole := func(ctx context.Context, req *admin.CreateRoleRequest, opts ...gax.CallOption) *admin.Role {
		seen, ok := roles[req.Role.Name]
		assert.True(t, ok, req.Role.Name)
		assert.False(t, seen, req.Role.Name)

		roles[req.Role.Name] = true
		return req.Role
	}

	permissions := map[string]bool{
		"permissions/iam.roles.get":          false,
		"permissions/iam.roles.list":         false,
		"permissions/iam.permissions.update": false,
	}

	fnCreatePermission := func(ctx context.Context, req *admin.CreatePermissionRequest, opts ...gax.CallOption) *admin.Permission {
		seen, ok := permissions[req.Permission.Name]
		assert.True(t, ok, req.Permission.Name)
		assert.False(t, seen, req.Permission.Name)

		permissions[req.Permission.Name] = true
		return req.Permission
	}

	gapic.On("GetRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, notFound)
	gapic.On("ListRoles", mock.Anything, mock.Anything, mock.Anything).Return(roleIterator)
	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(fnCreateRole, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeleteRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	gapic.On("GetPermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, notFound)
	gapic.On("ListPermissions", mock.Anything, mock.Anything, mock.Anything).Return(permissionIterator)
	gapic.On("CreatePermission", mock.Anything, mock.Anything, mock.Anything).Return(fnCreatePermission, nil)
	gapic.On("UpdatePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeletePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	client := &Client{client: gapic}

	assert.Nil(t, client.ApplyFile(context.TODO(), "testdata/iam.yaml"))

	gapic.AssertNumberOfCalls(t, "GetRole", 2)
	gapic.AssertNumberOfCalls(t, "ListRoles", 1)
	gapic.AssertNumberOfCalls(t, "CreateRole", 3)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 0)
	gapic.AssertNumberOfCalls(t, "DeleteRole", 0)

	gapic.AssertNumberOfCalls(t, "GetPermission", 1)
	gapic.AssertNumberOfCalls(t, "ListPermissions", 1)
	gapic.AssertNumberOfCalls(t, "CreatePermission", 3)
	gapic.AssertNumberOfCalls(t, "UpdatePermission", 0)
	gapic.AssertNumberOfCalls(t, "DeletePermission", 0)
}

func TestNothingChanged(t *testing.T) {
	gapic := &gapictest.IamClient{}

	roleIteratorErr := error(nil)
	roles := []*admin.Role{
		{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

			IncludedPermissions: []string{
				"iam.roles.get",
				"iam.roles.list",
			},
		},
	}
	fnNextRole := func() *admin.Role {
		if len(roles) == 0 {
			roleIteratorErr = iterator.Done
			return nil
		}
		role := roles[0]
		roles = roles[1:]
		return role
	}
	fnNextRoleErr := func() error {
		return roleIteratorErr
	}

	permissionIteratorErr := error(nil)
	permissions := []*admin.Permission{
		{
			Name:        "permissions/iam.roles.get",
			DisplayName: "IAM Role Get",
			Description: "This is the IAM Role Get permission.",
		},
		{
			Name:        "permissions/iam.roles.list",
			DisplayName: "IAM Role List",
			Description: "This is the IAM Role List permission.",
		},
	}
	fnNextPermission := func() *admin.Permission {
		if len(permissions) == 0 {
			permissionIteratorErr = iterator.Done
			return nil
		}
		role := permissions[0]
		permissions = permissions[1:]
		return role
	}
	fnNextPermissionErr := func() error {
		return permissionIteratorErr
	}

	roleIterator := &gapictest.RoleIterator{}
	roleIterator.On("Next").Return(fnNextRole, fnNextRoleErr)

	permissionIterator := &gapictest.PermissionIterator{}
	permissionIterator.On("Next").Return(fnNextPermission, fnNextPermissionErr)

	fnGetRole := func(ctx context.Context, req *admin.GetRoleRequest, opts ...gax.CallOption) *admin.Role {
		m := map[string]*admin.Role{
			"roles/iam.admin": {
				Name:        "roles/iam.admin",
				DisplayName: "IAM Admin",
				Description: "This is the IAM Admin role.",

				IncludedPermissions: []string{
					"iam.permissions.update",
				},
			},
			"roles/viewer": {
				Name:        "roles/viewer",
				DisplayName: "Viewer",
				Description: "Read access to all resources.",

				IncludedPermissions: []string{
					"iam.roles.get",
					"iam.roles.list",
					"image.folders.get",
					"image.folders.list",
					"hub.repositories.get",
					"hub.repositories.list",
				},
			},
		}

		role, ok := m[req.Name]
		require.True(t, ok, req.Name)

		return role
	}

	fnGetPermission := func(ctx context.Context, req *admin.GetPermissionRequest, opts ...gax.CallOption) *admin.Permission {
		m := map[string]*admin.Permission{
			"permissions/iam.permissions.update": {
				Name:        "permissions/iam.permissions.update",
				DisplayName: "IAM Permission Update",
				Description: "This is the IAM Permission Update permission.",
			},
		}

		role, ok := m[req.Name]
		require.True(t, ok, req.Name)

		return role
	}

	gapic.On("GetRole", mock.Anything, mock.Anything, mock.Anything).Return(fnGetRole, nil)
	gapic.On("ListRoles", mock.Anything, mock.Anything, mock.Anything).Return(roleIterator)
	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeleteRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	gapic.On("GetPermission", mock.Anything, mock.Anything, mock.Anything).Return(fnGetPermission, nil)
	gapic.On("ListPermissions", mock.Anything, mock.Anything, mock.Anything).Return(permissionIterator)
	gapic.On("CreatePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdatePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeletePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	client := &Client{client: gapic}

	assert.Nil(t, client.ApplyFile(context.TODO(), "testdata/iam.yaml"))

	gapic.AssertNumberOfCalls(t, "GetRole", 2)
	gapic.AssertNumberOfCalls(t, "ListRoles", 1)
	gapic.AssertNumberOfCalls(t, "CreateRole", 0)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 0)
	gapic.AssertNumberOfCalls(t, "DeleteRole", 0)

	gapic.AssertNumberOfCalls(t, "GetPermission", 1)
	gapic.AssertNumberOfCalls(t, "ListPermissions", 1)
	gapic.AssertNumberOfCalls(t, "CreatePermission", 0)
	gapic.AssertNumberOfCalls(t, "UpdatePermission", 0)
	gapic.AssertNumberOfCalls(t, "DeletePermission", 0)
}

func TestEverythingChanged(t *testing.T) {
	gapic := &gapictest.IamClient{}

	roleIteratorErr := error(nil)
	roles := []*admin.Role{
		{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

			IncludedPermissions: []string{},
		}, {
			Name: "roles/iam.deleteme",
		},
	}
	fnNextRole := func() *admin.Role {
		if len(roles) == 0 {
			roleIteratorErr = iterator.Done
			return nil
		}
		role := roles[0]
		roles = roles[1:]
		return role
	}
	fnNextRoleErr := func() error {
		return roleIteratorErr
	}

	permissionIteratorErr := error(nil)
	permissions := []*admin.Permission{
		{
			Name:        "permissions/iam.roles.get",
			DisplayName: "IAM Role Get",
			Description: "-",
		},
		{
			Name:        "permissions/iam.roles.list",
			DisplayName: "IAM Role List",
			Description: "-",
		}, {
			Name: "permissions/iam.delete.me",
		},
	}
	fnNextPermission := func() *admin.Permission {
		if len(permissions) == 0 {
			permissionIteratorErr = iterator.Done
			return nil
		}
		role := permissions[0]
		permissions = permissions[1:]
		return role
	}
	fnNextPermissionErr := func() error {
		return permissionIteratorErr
	}

	roleIterator := &gapictest.RoleIterator{}
	roleIterator.On("Next").Return(fnNextRole, fnNextRoleErr)

	permissionIterator := &gapictest.PermissionIterator{}
	permissionIterator.On("Next").Return(fnNextPermission, fnNextPermissionErr)

	fnGetRole := func(ctx context.Context, req *admin.GetRoleRequest, opts ...gax.CallOption) *admin.Role {
		m := map[string]*admin.Role{
			"roles/iam.admin": {
				Name:        "roles/iam.admin",
				DisplayName: "IAM Admin",
				Description: "This is the IAM Admin role.",

				IncludedPermissions: []string{},
			},
			"roles/viewer": {
				Name:        "roles/viewer",
				DisplayName: "Viewer",
				Description: "Read access to all resources.",

				IncludedPermissions: []string{
					"image.folders.get",
					"image.folders.list",
					"hub.repositories.get",
					"hub.repositories.list",
				},
			},
		}

		role, ok := m[req.Name]
		require.True(t, ok, req.Name)

		return role
	}

	fnGetPermission := func(ctx context.Context, req *admin.GetPermissionRequest, opts ...gax.CallOption) *admin.Permission {
		m := map[string]*admin.Permission{
			"permissions/iam.permissions.update": {
				Name:        "permissions/iam.permissions.update",
				DisplayName: "-",
				Description: "This is the IAM Permission Update permission.",
			},
		}

		role, ok := m[req.Name]
		require.True(t, ok, req.Name)

		return role
	}

	gapic.On("GetRole", mock.Anything, mock.Anything, mock.Anything).Return(fnGetRole, nil)
	gapic.On("ListRoles", mock.Anything, mock.Anything, mock.Anything).Return(roleIterator)
	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeleteRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	gapic.On("GetPermission", mock.Anything, mock.Anything, mock.Anything).Return(fnGetPermission, nil)
	gapic.On("ListPermissions", mock.Anything, mock.Anything, mock.Anything).Return(permissionIterator)
	gapic.On("CreatePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdatePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("DeletePermission", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	client := &Client{client: gapic}

	assert.Nil(t, client.ApplyFile(context.TODO(), "testdata/iam.yaml"))

	gapic.AssertNumberOfCalls(t, "GetRole", 2)
	gapic.AssertNumberOfCalls(t, "ListRoles", 1)
	gapic.AssertNumberOfCalls(t, "CreateRole", 0)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 3)
	gapic.AssertNumberOfCalls(t, "DeleteRole", 1)

	gapic.AssertNumberOfCalls(t, "GetPermission", 1)
	gapic.AssertNumberOfCalls(t, "ListPermissions", 1)
	gapic.AssertNumberOfCalls(t, "CreatePermission", 0)
	gapic.AssertNumberOfCalls(t, "UpdatePermission", 3)
	gapic.AssertNumberOfCalls(t, "DeletePermission", 1)
}
