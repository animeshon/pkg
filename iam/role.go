package iam

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	"github.com/sirupsen/logrus"
	fieldmask "google.golang.org/protobuf/types/known/fieldmaskpb"

	"google.golang.org/api/iterator"
)

type Role struct {
	Name        string   `yaml:"id"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Permissions []string `yaml:"permissions"`
}

func (r *Role) ToProtoMessage() *admin.Role {
	return &admin.Role{
		Name:                "roles/" + r.Name,
		DisplayName:         r.Title,
		Description:         r.Description,
		IncludedPermissions: r.Permissions,
	}
}

func RoleCompare(a, b *admin.Role) (*fieldmask.FieldMask, bool) {
	var paths []string
	if a.DisplayName != b.DisplayName {
		paths = append(paths, "display_name")
	}
	if a.Description != b.Description {
		paths = append(paths, "description")
	}

	sort.Strings(a.IncludedPermissions)
	sort.Strings(b.IncludedPermissions)

	if !reflect.DeepEqual(a.IncludedPermissions, b.IncludedPermissions) {
		paths = append(paths, "included_permissions")
	}

	if len(paths) == 0 {
		return nil, true
	}

	mask, err := fieldmask.New(a, paths...)
	if err != nil {
		panic(err)
	}

	return mask, false
}

func (r *Client) ListRoles(ctx context.Context, matches []*RuleMatch) ([]*admin.Role, error) {
	roles := []*admin.Role{}
	for k, match := range matches {
		if match.Exact != "" && match.Prefix != "" {
			return nil, fmt.Errorf("match[%d]: 'exact' and 'prefix' cannot both be defined", k)
		}
		if match.Exact == "" && match.Prefix == "" {
			return nil, fmt.Errorf("match[%d]: 'exact' and 'prefix' cannot both be undefined", k)
		}

		if len(match.Exact) != 0 {
			request := &admin.GetRoleRequest{Name: "roles/" + match.Exact}
			role, err := r.client.GetRole(ctx, request)
			if err != nil {
				if IsErrorNotFound(err) {
					continue
				}
				return nil, fmt.Errorf("match[%d]: %s", k, err)
			}

			roles = append(roles, role)
		}

		if len(match.Prefix) != 0 {
			request := &admin.ListRolesRequest{PageSize: 250, Filter: "prefix:" + match.Prefix}
			iter := r.client.ListRoles(ctx, request)

			for {
				role, err := iter.Next()
				if err != nil {
					if errors.Is(err, iterator.Done) {
						break
					}
					return nil, fmt.Errorf("match[%d]: %s", k, err)
				}

				logrus.Infof("fetched role '%s' with %d included permissions", role.Name, len(role.IncludedPermissions))
				roles = append(roles, role)
			}
		}
	}

	return roles, nil
}

func (r *Client) ReconcileRoles(ctx context.Context, local, remote []*admin.Role) error {
	mlocal := make(map[string]*admin.Role)
	for _, role := range local {
		mlocal[role.Name] = role
	}

	mremote := make(map[string]*admin.Role)
	for _, role := range remote {
		mremote[role.Name] = role
	}

	for name, role := range mlocal {
		remoteRole, ok := mremote[name]
		if !ok {
			request := &admin.CreateRoleRequest{Role: role}
			if _, err := r.client.CreateRole(ctx, request); err != nil {
				return err
			}

			continue
		}

		if mask, ok := RoleCompare(role, remoteRole); !ok {
			request := &admin.UpdateRoleRequest{Role: role, UpdateMask: mask}
			if _, err := r.client.UpdateRole(ctx, request); err != nil {
				return err
			}
		}
	}

	for name := range mremote {
		if _, ok := mlocal[name]; ok {
			continue
		}

		request := &admin.DeleteRoleRequest{Name: name}
		if err := r.client.DeleteRole(ctx, request); err != nil {
			return err
		}
	}

	return nil
}

var (
	ErrMergeMultiple   = errors.New("cannot merge role with multiple matches")
	ErrMergeExactlyOne = errors.New("merge requires exactly one value to be defined")
)

func (r *Client) MergeRole(ctx context.Context, local, remote []*admin.Role, merge *RoleMerge) error {
	if len(remote) > 1 {
		return ErrMergeMultiple
	}
	if len(local) != 1 {
		return ErrMergeExactlyOne
	}

	if len(remote) == 0 {
		request := &admin.CreateRoleRequest{Role: local[0]}
		_, err := r.client.CreateRole(ctx, request)
		return err
	}

	permission := merge.Permissions
	if permission.Exact != "" && permission.Prefix != "" {
		return fmt.Errorf("permissions: 'exact' and 'prefix' cannot both be defined")
	}
	if permission.Exact == "" && permission.Prefix == "" {
		return fmt.Errorf("permissions: 'exact' and 'prefix' cannot both be undefined")
	}

	var localP, remoteP, otherP []string
	for _, i := range remote[0].IncludedPermissions {
		if len(permission.Exact) != 0 && i == permission.Exact {
			remoteP = append(remoteP, i)
			continue
		}
		if len(permission.Prefix) != 0 && strings.HasPrefix(i, permission.Prefix) {
			remoteP = append(remoteP, i)
			continue
		}
		otherP = append(otherP, i)
	}

	for _, i := range local[0].IncludedPermissions {
		if len(permission.Exact) != 0 && i == permission.Exact {
			localP = append(localP, i)
		}
		if len(permission.Prefix) != 0 && strings.HasPrefix(i, permission.Prefix) {
			localP = append(localP, i)
		}
	}

	roleA := &admin.Role{
		IncludedPermissions: localP,
	}
	roleB := &admin.Role{
		IncludedPermissions: remoteP,
	}

	if merge.Title {
		roleA.DisplayName = local[0].DisplayName
		roleB.DisplayName = remote[0].DisplayName
	}
	if merge.Description {
		roleA.Description = local[0].Description
		roleB.Description = remote[0].Description
	}

	if mask, ok := RoleCompare(roleA, roleB); !ok {
		remote[0].IncludedPermissions = append(otherP, localP...)

		request := &admin.UpdateRoleRequest{Role: remote[0], UpdateMask: mask}
		if _, err := r.client.UpdateRole(ctx, request); err != nil {
			return err
		}
	}

	return nil
}
