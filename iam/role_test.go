package iam

import (
	"context"
	"testing"

	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	"github.com/animeshon/pkg/iam/gapictest"
	"github.com/googleapis/gax-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMergeRoleRequireCreate(t *testing.T) {
	gapic := &gapictest.IamClient{}
	client := &Client{client: gapic}

	local := []*admin.Role{
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

	fnCreateRole := func(ctx context.Context, req *admin.CreateRoleRequest, opts ...gax.CallOption) *admin.Role {
		_, ok := RoleCompare(req.Role, local[0])
		assert.True(t, ok)

		return nil
	}

	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(fnCreateRole, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	merge := &RoleMerge{Permissions: &RuleMatch{Prefix: "iam."}}
	assert.Nil(t, client.MergeRole(context.TODO(), local, nil, merge))

	gapic.AssertNumberOfCalls(t, "CreateRole", 1)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 0)
}

func TestMergeRoleRequireUpdate(t *testing.T) {
	gapic := &gapictest.IamClient{}
	client := &Client{client: gapic}

	local := []*admin.Role{
		{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

			IncludedPermissions: []string{
				"iam.roles.get",
				"iam.roles.list",
				"filter.wrong.get",
				"filter.wrong.list",
			},
		},
	}

	remote := []*admin.Role{
		{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

			IncludedPermissions: []string{
				"iam.roles.get",
				"iam.roles.list",
				"iam.permissions.get",
				"iam.permissions.list",
				"image.folders.get",
				"image.folders.list",
				"hub.repositories.get",
				"hub.repositories.list",
			},
		},
	}

	fnUpdateRole := func(ctx context.Context, req *admin.UpdateRoleRequest, opts ...gax.CallOption) *admin.Role {
		expected := &admin.Role{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

			IncludedPermissions: []string{
				"iam.roles.get",
				"iam.roles.list",
				"image.folders.get",
				"image.folders.list",
				"hub.repositories.get",
				"hub.repositories.list",
			},
		}

		_, ok := RoleCompare(req.Role, expected)
		assert.True(t, ok)

		return nil
	}

	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(fnUpdateRole, nil)

	merge := &RoleMerge{Permissions: &RuleMatch{Prefix: "iam."}}
	assert.Nil(t, client.MergeRole(context.TODO(), local, remote, merge))

	gapic.AssertNumberOfCalls(t, "CreateRole", 0)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 1)
}

func TestMergeRoleNoChanges(t *testing.T) {
	gapic := &gapictest.IamClient{}
	client := &Client{client: gapic}

	local := []*admin.Role{
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

	remote := []*admin.Role{
		{
			Name:        "roles/iam.viewer",
			DisplayName: "IAM Viewer",
			Description: "This is the IAM Viewer role.",

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

	gapic.On("CreateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	gapic.On("UpdateRole", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	merge := &RoleMerge{Permissions: &RuleMatch{Prefix: "iam."}}
	assert.Nil(t, client.MergeRole(context.TODO(), local, remote, merge))

	gapic.AssertNumberOfCalls(t, "CreateRole", 0)
	gapic.AssertNumberOfCalls(t, "UpdateRole", 0)
}
