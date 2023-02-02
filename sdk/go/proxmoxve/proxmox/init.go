// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package proxmox

import (
	"fmt"

	"github.com/Doublemine/pulumi-proxmoxve/sdk/v2023/go/proxmoxve"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "proxmoxve:proxmox/lXCContainer:LXCContainer":
		r = &LXCContainer{}
	case "proxmoxve:proxmox/lXCDisk:LXCDisk":
		r = &LXCDisk{}
	case "proxmoxve:proxmox/pool:Pool":
		r = &Pool{}
	case "proxmoxve:proxmox/qemuVM:QemuVM":
		r = &QemuVM{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := proxmoxve.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"proxmoxve",
		"proxmox/lXCContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"proxmoxve",
		"proxmox/lXCDisk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"proxmoxve",
		"proxmox/pool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"proxmoxve",
		"proxmox/qemuVM",
		&module{version},
	)
}
