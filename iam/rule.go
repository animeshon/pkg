package iam

type Service struct {
	Service string `yaml:"service"`
	Rules   *Rules `yaml:"rules"`
}

type Rules struct {
	Permissions []*PermissionRule `yaml:"permissions"`
	Roles       []*RoleRule       `yaml:"roles"`
}

type Rule struct {
	Operation string       `yaml:"operation"`
	Match     []*RuleMatch `yaml:"match"`
}

type RoleRule struct {
	Rule   `yaml:",inline"`
	Merge  *RoleMerge `yaml:"merge"`
	Values []*Role    `yaml:"values"`
}

type PermissionRule struct {
	Rule   `yaml:",inline"`
	Values []*Permission `yaml:"values"`
}

type RoleMerge struct {
	Title       bool       `yaml:"title"`
	Description bool       `yaml:"description"`
	Permissions *RuleMatch `yaml:"permissions"`
}

type RuleMatch struct {
	Exact  string `yaml:"exact"`
	Prefix string `yaml:"prefix"`
}
