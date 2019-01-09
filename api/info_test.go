package api

import (
	"net"
	"runtime"
	"testing"

	"github.com/digitalrebar/provision"
	"github.com/digitalrebar/provision/models"
)

func TestInfo(t *testing.T) {
	test := &crudTest{
		name: "get info",
		expectRes: &models.Info{
			Address:            net.IPv4(127, 0, 0, 1),
			ApiPort:            10011,
			FilePort:           10012,
			BinlPort:           10015,
			DhcpPort:           10014,
			TftpPort:           10013,
			ProvisionerEnabled: true,
			TftpEnabled:        true,
			BinlEnabled:        true,
			DhcpEnabled:        true,
			Stats: []models.Stat{
				{
					Name:  "machines.count",
					Count: 0,
				},
				{
					Name:  "subnets.count",
					Count: 0,
				},
			},
			Arch:    runtime.GOARCH,
			Os:      runtime.GOOS,
			Version: provision.RSVersion,
			Id:      "Fred",
			Features: []string{
				"api-v3",
				"sane-exit-codes",
				"common-blob-size",
				"change-stage-map",
				"job-exit-states",
				"package-repository-handling",
				"profileless-machine",
				"threaded-log-levels",
				"plugin-v2",
				"fsm-runner",
				"plugin-v2-safe-config",
				"workflows",
				"default-workflow",
				"http-range-header",
				"roles",
				"tenants",
				"secure-params",
				"separate-meta-api",
				"slim-objects",
				"secure-param-upgrade",
				"sprig",
				"multiarch",
				"actions-in-task-list",
			},
			License: models.LicenseBundle{Licenses: []models.License{}},
			Scopes: map[string]map[string]struct{}{
				"bootenvs": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"contents": {
					"create": {},
					"delete": {},
					"get":    {},
					"list":   {},
					"update": {},
				},
				"files": {
					"delete": {},
					"get":    {},
					"list":   {},
					"post":   {},
				},
				"info": {
					"get": {},
				},
				"interfaces": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"isos": {
					"delete": {},
					"get":    {},
					"list":   {},
					"post":   {},
				},
				"jobs": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"log":     {},
					"update":  {},
				},
				"leases": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"machines": {
					"action":       {},
					"actions":      {},
					"create":       {},
					"delete":       {},
					"get":          {},
					"getSecure":    {},
					"list":         {},
					"update":       {},
					"updateSecure": {},
				},
				"objects": {
					"list": {},
				},
				"params": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"plugin_providers": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"plugins": {
					"action":       {},
					"actions":      {},
					"create":       {},
					"delete":       {},
					"get":          {},
					"getSecure":    {},
					"list":         {},
					"update":       {},
					"updateSecure": {},
				},
				"preferences": {
					"list": {},
					"post": {},
				},
				"profiles": {
					"action":       {},
					"actions":      {},
					"create":       {},
					"delete":       {},
					"get":          {},
					"getSecure":    {},
					"list":         {},
					"update":       {},
					"updateSecure": {},
				},
				"reservations": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"roles": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"stages": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"subnets": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"tasks": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"tenants": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"templates": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
				"users": {
					"action":   {},
					"actions":  {},
					"create":   {},
					"delete":   {},
					"get":      {},
					"list":     {},
					"password": {},
					"token":    {},
					"update":   {},
				},
				"workflows": {
					"action":  {},
					"actions": {},
					"create":  {},
					"delete":  {},
					"get":     {},
					"list":    {},
					"update":  {},
				},
			},
		},
		expectErr: nil,
		op: func() (interface{}, error) {
			return session.Info()
		},
	}
	test.run(t)

}
