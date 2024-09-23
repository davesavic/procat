package seeders

import (
	"embed"
)

type accessControlList struct {
	Roles []role `yaml:"roles"`
}

type role struct {
	Name        string   `yaml:"name"`
	Permissions []string `yaml:"permissions"`
}

type AccessControlSeeder struct {
	EmbededFiles embed.FS
}

func (acs AccessControlSeeder) Run() error {
	return nil
}

func (acs AccessControlSeeder) Name() string {
	return "acl"
}
