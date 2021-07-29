package iam

import (
	"context"
	"fmt"
	"io/ioutil"

	admin "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	gapic "github.com/animeshon/pkg/iam/gapic"

	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

// TODO: increase the testing coverage - this package is critical to us.
// TODO: improve error handling - errors should be structured.
// TODO: introduce logging for a better observability.
// TODO: break down the functions even further for a better testability.

type Client struct {
	client gapic.IamClient
}

func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	client, err := gapic.NewIamClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (r *Client) Close() error {
	return r.client.Close()
}

func (r *Client) ApplyFile(ctx context.Context, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return r.Apply(ctx, data)
}

func (r *Client) Apply(ctx context.Context, data []byte) error {
	var service *Service
	if err := yaml.Unmarshal([]byte(data), &service); err != nil {
		return err
	}

	if err := r.ApplyPermissions(ctx, service.Rules.Permissions); err != nil {
		return fmt.Errorf("[%s] %s", service.Service, err)
	}

	if err := r.ApplyRoles(ctx, service.Rules.Roles); err != nil {
		return fmt.Errorf("[%s] %s", service.Service, err)
	}

	return nil
}

func (r *Client) ApplyPermissions(ctx context.Context, rules []*PermissionRule) error {
	for i, rule := range rules {
		if rule.Operation != "RECONCILE" {
			return fmt.Errorf("rules.permissions[%d].operation: expected value to be 'RECONCILE'", i)
		}

		if len(rule.Match) == 0 {
			return fmt.Errorf("rules.permissions[%d].match: no matching rules", i)
		}

		remote, err := r.ListPermissions(ctx, rule.Match)
		if err != nil {
			return fmt.Errorf("rules.permissions[%d].%s", i, err)
		}

		local := []*admin.Permission{}
		for _, permission := range rule.Values {
			local = append(local, permission.ToProtoMessage())
		}

		if err := r.ReconcilePermissions(ctx, local, remote); err != nil {
			return fmt.Errorf("rules.permissions[%d]: failed to reconcile: %s", i, err)
		}
	}

	return nil
}

func (r *Client) ApplyRoles(ctx context.Context, rules []*RoleRule) error {
	for i, rule := range rules {
		if rule.Operation != "RECONCILE" && rule.Operation != "MERGE" {
			return fmt.Errorf("rules.roles[%d].operation: expected value to be 'RECONCILE' or 'MERGE'", i)
		}

		if len(rule.Match) == 0 {
			return fmt.Errorf("rules.roles[%d].match: no matching rules", i)
		}

		remote, err := r.ListRoles(ctx, rule.Match)
		if err != nil {
			return fmt.Errorf("rules.roles[%d].%s", i, err)
		}

		local := []*admin.Role{}
		for _, role := range rule.Values {
			local = append(local, role.ToProtoMessage())
		}

		switch rule.Operation {
		case "RECONCILE":
			if err := r.ReconcileRoles(ctx, local, remote); err != nil {
				return fmt.Errorf("rules.roles[%d]: failed to reconcile: %s", i, err)
			}
		case "MERGE":
			if err := r.MergeRole(ctx, local, remote, rule.Merge); err != nil {
				return fmt.Errorf("rules.roles[%d]: failed to merge: %s", i, err)
			}
		}
	}

	return nil
}
