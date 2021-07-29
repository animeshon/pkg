package gapic

import (
	"context"
	"io"

	gapic "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"

	"github.com/googleapis/gax-go/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type RoleIterator interface {
	PageInfo() *iterator.PageInfo
	Next() (*admin.Role, error)
}

type PermissionIterator interface {
	PageInfo() *iterator.PageInfo
	Next() (*admin.Permission, error)
}

type IamClient interface {
	io.Closer

	GetRole(context.Context, *admin.GetRoleRequest, ...gax.CallOption) (*admin.Role, error)
	ListRoles(context.Context, *admin.ListRolesRequest, ...gax.CallOption) RoleIterator
	CreateRole(context.Context, *admin.CreateRoleRequest, ...gax.CallOption) (*admin.Role, error)
	UpdateRole(context.Context, *admin.UpdateRoleRequest, ...gax.CallOption) (*admin.Role, error)
	DeleteRole(context.Context, *admin.DeleteRoleRequest, ...gax.CallOption) error
	GetPermission(context.Context, *admin.GetPermissionRequest, ...gax.CallOption) (*admin.Permission, error)
	ListPermissions(context.Context, *admin.ListPermissionsRequest, ...gax.CallOption) PermissionIterator
	CreatePermission(context.Context, *admin.CreatePermissionRequest, ...gax.CallOption) (*admin.Permission, error)
	UpdatePermission(context.Context, *admin.UpdatePermissionRequest, ...gax.CallOption) (*admin.Permission, error)
	DeletePermission(context.Context, *admin.DeletePermissionRequest, ...gax.CallOption) error
}

var _ IamClient = &IamClientImpl{}

type IamClientImpl struct {
	internalClient *gapic.IamClient
}

func NewIamClient(ctx context.Context, opts ...option.ClientOption) (*IamClientImpl, error) {
	client, err := gapic.NewIamClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &IamClientImpl{internalClient: client}, nil
}

func (c *IamClientImpl) Close() error {
	return c.internalClient.Close()
}

func (c *IamClientImpl) GetRole(ctx context.Context, req *admin.GetRoleRequest, opts ...gax.CallOption) (*admin.Role, error) {
	return c.internalClient.GetRole(ctx, req, opts...)
}

func (c *IamClientImpl) ListRoles(ctx context.Context, req *admin.ListRolesRequest, opts ...gax.CallOption) RoleIterator {
	return c.internalClient.ListRoles(ctx, req, opts...)
}
func (c *IamClientImpl) CreateRole(ctx context.Context, req *admin.CreateRoleRequest, opts ...gax.CallOption) (*admin.Role, error) {
	logrus.Infof("the following role will be created: %s", req.GetRole().GetName())
	return c.internalClient.CreateRole(ctx, req, opts...)
}
func (c *IamClientImpl) UpdateRole(ctx context.Context, req *admin.UpdateRoleRequest, opts ...gax.CallOption) (*admin.Role, error) {
	logrus.Infof("the following role will be updated: %s", req.GetRole().GetName())
	return c.internalClient.UpdateRole(ctx, req, opts...)
}
func (c *IamClientImpl) DeleteRole(ctx context.Context, req *admin.DeleteRoleRequest, opts ...gax.CallOption) error {
	logrus.Errorf("the following role must be deleted manually: %s", req.GetName())
	return nil
}
func (c *IamClientImpl) GetPermission(ctx context.Context, req *admin.GetPermissionRequest, opts ...gax.CallOption) (*admin.Permission, error) {
	return c.internalClient.GetPermission(ctx, req, opts...)
}
func (c *IamClientImpl) ListPermissions(ctx context.Context, req *admin.ListPermissionsRequest, opts ...gax.CallOption) PermissionIterator {
	return c.internalClient.ListPermissions(ctx, req, opts...)
}
func (c *IamClientImpl) CreatePermission(ctx context.Context, req *admin.CreatePermissionRequest, opts ...gax.CallOption) (*admin.Permission, error) {
	logrus.Infof("the following permission will be created: %s", req.GetPermission().GetName())
	return c.internalClient.CreatePermission(ctx, req, opts...)
}
func (c *IamClientImpl) UpdatePermission(ctx context.Context, req *admin.UpdatePermissionRequest, opts ...gax.CallOption) (*admin.Permission, error) {
	logrus.Infof("the following permission will be updated: %s", req.GetPermission().GetName())
	return c.internalClient.UpdatePermission(ctx, req, opts...)
}
func (c *IamClientImpl) DeletePermission(ctx context.Context, req *admin.DeletePermissionRequest, opts ...gax.CallOption) error {
	logrus.Infof("the following permission will be deleted: %s", req.GetName())
	return c.internalClient.DeletePermission(ctx, req, opts...)
}
