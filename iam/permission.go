package iam

import (
	"context"
	"errors"
	"fmt"

	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	fieldmask "google.golang.org/protobuf/types/known/fieldmaskpb"

	"google.golang.org/api/iterator"
)

type Permission struct {
	Name        string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func (p *Permission) ToProtoMessage() *admin.Permission {
	return &admin.Permission{
		Name:        "permissions/" + p.Name,
		DisplayName: p.Title,
		Description: p.Description,
	}
}

func PermissionCompare(a, b *admin.Permission) (*fieldmask.FieldMask, bool) {
	var paths []string
	if a.DisplayName != b.DisplayName {
		paths = append(paths, "display_name")
	}
	if a.Description != b.Description {
		paths = append(paths, "description")
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

func (r *Client) ListPermissions(ctx context.Context, matches []*RuleMatch) ([]*admin.Permission, error) {
	permissions := []*admin.Permission{}
	for k, match := range matches {
		if match.Exact != "" && match.Prefix != "" {
			return nil, fmt.Errorf("match[%d]: 'exact' and 'prefix' cannot both be defined", k)
		}
		if match.Exact == "" && match.Prefix == "" {
			return nil, fmt.Errorf("match[%d]: 'exact' and 'prefix' cannot both be undefined", k)
		}

		if len(match.Exact) != 0 {
			request := &admin.GetPermissionRequest{Name: "permissions/" + match.Exact}
			permission, err := r.client.GetPermission(ctx, request)
			if err != nil {
				if IsErrorNotFound(err) {
					continue
				}
				return nil, fmt.Errorf("match[%d]: %s", k, err)
			}

			permissions = append(permissions, permission)
		}

		if len(match.Prefix) != 0 {
			request := &admin.ListPermissionsRequest{PageSize: 250, Filter: "prefix:" + match.Prefix}
			iter := r.client.ListPermissions(ctx, request)

			for {
				permission, err := iter.Next()
				if err != nil {
					if errors.Is(err, iterator.Done) {
						break
					}
					return nil, fmt.Errorf("match[%d]: %s", k, err)
				}

				permissions = append(permissions, permission)
			}
		}
	}

	return permissions, nil
}

func (r *Client) ReconcilePermissions(ctx context.Context, local, remote []*admin.Permission) error {
	mlocal := make(map[string]*admin.Permission)
	for _, permission := range local {
		mlocal[permission.Name] = permission
	}

	mremote := make(map[string]*admin.Permission)
	for _, permission := range remote {
		mremote[permission.Name] = permission
	}

	for name, permission := range mlocal {
		remotePermission, ok := mremote[name]
		if !ok {
			request := &admin.CreatePermissionRequest{Permission: permission}
			if _, err := r.client.CreatePermission(ctx, request); err != nil {
				return err
			}

			continue
		}

		if mask, ok := PermissionCompare(permission, remotePermission); !ok {
			request := &admin.UpdatePermissionRequest{Permission: permission, UpdateMask: mask}
			if _, err := r.client.UpdatePermission(ctx, request); err != nil {
				return err
			}
		}
	}

	for name := range mremote {
		if _, ok := mlocal[name]; ok {
			continue
		}

		request := &admin.DeletePermissionRequest{Name: name}
		if err := r.client.DeletePermission(ctx, request); err != nil {
			return err
		}
	}

	return nil
}
