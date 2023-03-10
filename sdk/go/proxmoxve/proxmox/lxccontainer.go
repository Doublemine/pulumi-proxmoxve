// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package proxmox

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LXCContainer struct {
	pulumi.CustomResourceState

	Arch               pulumi.StringPtrOutput            `pulumi:"arch"`
	Bwlimit            pulumi.IntPtrOutput               `pulumi:"bwlimit"`
	Clone              pulumi.StringPtrOutput            `pulumi:"clone"`
	CloneStorage       pulumi.StringPtrOutput            `pulumi:"cloneStorage"`
	Cmode              pulumi.StringPtrOutput            `pulumi:"cmode"`
	Console            pulumi.BoolPtrOutput              `pulumi:"console"`
	Cores              pulumi.IntPtrOutput               `pulumi:"cores"`
	Cpulimit           pulumi.IntPtrOutput               `pulumi:"cpulimit"`
	Cpuunits           pulumi.IntPtrOutput               `pulumi:"cpuunits"`
	Description        pulumi.StringPtrOutput            `pulumi:"description"`
	Features           LXCContainerFeaturesPtrOutput     `pulumi:"features"`
	Force              pulumi.BoolPtrOutput              `pulumi:"force"`
	Full               pulumi.BoolPtrOutput              `pulumi:"full"`
	Hagroup            pulumi.StringPtrOutput            `pulumi:"hagroup"`
	Hastate            pulumi.StringPtrOutput            `pulumi:"hastate"`
	Hookscript         pulumi.StringPtrOutput            `pulumi:"hookscript"`
	Hostname           pulumi.StringPtrOutput            `pulumi:"hostname"`
	IgnoreUnpackErrors pulumi.BoolPtrOutput              `pulumi:"ignoreUnpackErrors"`
	Lock               pulumi.StringPtrOutput            `pulumi:"lock"`
	Memory             pulumi.IntPtrOutput               `pulumi:"memory"`
	Mountpoints        LXCContainerMountpointArrayOutput `pulumi:"mountpoints"`
	Nameserver         pulumi.StringPtrOutput            `pulumi:"nameserver"`
	Networks           LXCContainerNetworkArrayOutput    `pulumi:"networks"`
	Onboot             pulumi.BoolPtrOutput              `pulumi:"onboot"`
	Ostemplate         pulumi.StringPtrOutput            `pulumi:"ostemplate"`
	Ostype             pulumi.StringOutput               `pulumi:"ostype"`
	Password           pulumi.StringPtrOutput            `pulumi:"password"`
	Pool               pulumi.StringPtrOutput            `pulumi:"pool"`
	Protection         pulumi.BoolPtrOutput              `pulumi:"protection"`
	Restore            pulumi.BoolPtrOutput              `pulumi:"restore"`
	Rootfs             LXCContainerRootfsPtrOutput       `pulumi:"rootfs"`
	Searchdomain       pulumi.StringPtrOutput            `pulumi:"searchdomain"`
	SshPublicKeys      pulumi.StringPtrOutput            `pulumi:"sshPublicKeys"`
	Start              pulumi.BoolPtrOutput              `pulumi:"start"`
	Startup            pulumi.StringPtrOutput            `pulumi:"startup"`
	Swap               pulumi.IntPtrOutput               `pulumi:"swap"`
	Tags               pulumi.StringPtrOutput            `pulumi:"tags"`
	TargetNode         pulumi.StringOutput               `pulumi:"targetNode"`
	Template           pulumi.BoolPtrOutput              `pulumi:"template"`
	Tty                pulumi.IntPtrOutput               `pulumi:"tty"`
	Unique             pulumi.BoolPtrOutput              `pulumi:"unique"`
	Unprivileged       pulumi.BoolPtrOutput              `pulumi:"unprivileged"`
	Unuseds            pulumi.StringArrayOutput          `pulumi:"unuseds"`
	// The VM identifier in proxmox (100-999999999)
	Vmid pulumi.IntOutput `pulumi:"vmid"`
}

// NewLXCContainer registers a new resource with the given unique name, arguments, and options.
func NewLXCContainer(ctx *pulumi.Context,
	name string, args *LXCContainerArgs, opts ...pulumi.ResourceOption) (*LXCContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetNode == nil {
		return nil, errors.New("invalid value for required argument 'TargetNode'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource LXCContainer
	err := ctx.RegisterResource("proxmoxve:proxmox/lXCContainer:LXCContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLXCContainer gets an existing LXCContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLXCContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LXCContainerState, opts ...pulumi.ResourceOption) (*LXCContainer, error) {
	var resource LXCContainer
	err := ctx.ReadResource("proxmoxve:proxmox/lXCContainer:LXCContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LXCContainer resources.
type lxccontainerState struct {
	Arch               *string                  `pulumi:"arch"`
	Bwlimit            *int                     `pulumi:"bwlimit"`
	Clone              *string                  `pulumi:"clone"`
	CloneStorage       *string                  `pulumi:"cloneStorage"`
	Cmode              *string                  `pulumi:"cmode"`
	Console            *bool                    `pulumi:"console"`
	Cores              *int                     `pulumi:"cores"`
	Cpulimit           *int                     `pulumi:"cpulimit"`
	Cpuunits           *int                     `pulumi:"cpuunits"`
	Description        *string                  `pulumi:"description"`
	Features           *LXCContainerFeatures    `pulumi:"features"`
	Force              *bool                    `pulumi:"force"`
	Full               *bool                    `pulumi:"full"`
	Hagroup            *string                  `pulumi:"hagroup"`
	Hastate            *string                  `pulumi:"hastate"`
	Hookscript         *string                  `pulumi:"hookscript"`
	Hostname           *string                  `pulumi:"hostname"`
	IgnoreUnpackErrors *bool                    `pulumi:"ignoreUnpackErrors"`
	Lock               *string                  `pulumi:"lock"`
	Memory             *int                     `pulumi:"memory"`
	Mountpoints        []LXCContainerMountpoint `pulumi:"mountpoints"`
	Nameserver         *string                  `pulumi:"nameserver"`
	Networks           []LXCContainerNetwork    `pulumi:"networks"`
	Onboot             *bool                    `pulumi:"onboot"`
	Ostemplate         *string                  `pulumi:"ostemplate"`
	Ostype             *string                  `pulumi:"ostype"`
	Password           *string                  `pulumi:"password"`
	Pool               *string                  `pulumi:"pool"`
	Protection         *bool                    `pulumi:"protection"`
	Restore            *bool                    `pulumi:"restore"`
	Rootfs             *LXCContainerRootfs      `pulumi:"rootfs"`
	Searchdomain       *string                  `pulumi:"searchdomain"`
	SshPublicKeys      *string                  `pulumi:"sshPublicKeys"`
	Start              *bool                    `pulumi:"start"`
	Startup            *string                  `pulumi:"startup"`
	Swap               *int                     `pulumi:"swap"`
	Tags               *string                  `pulumi:"tags"`
	TargetNode         *string                  `pulumi:"targetNode"`
	Template           *bool                    `pulumi:"template"`
	Tty                *int                     `pulumi:"tty"`
	Unique             *bool                    `pulumi:"unique"`
	Unprivileged       *bool                    `pulumi:"unprivileged"`
	Unuseds            []string                 `pulumi:"unuseds"`
	// The VM identifier in proxmox (100-999999999)
	Vmid *int `pulumi:"vmid"`
}

type LXCContainerState struct {
	Arch               pulumi.StringPtrInput
	Bwlimit            pulumi.IntPtrInput
	Clone              pulumi.StringPtrInput
	CloneStorage       pulumi.StringPtrInput
	Cmode              pulumi.StringPtrInput
	Console            pulumi.BoolPtrInput
	Cores              pulumi.IntPtrInput
	Cpulimit           pulumi.IntPtrInput
	Cpuunits           pulumi.IntPtrInput
	Description        pulumi.StringPtrInput
	Features           LXCContainerFeaturesPtrInput
	Force              pulumi.BoolPtrInput
	Full               pulumi.BoolPtrInput
	Hagroup            pulumi.StringPtrInput
	Hastate            pulumi.StringPtrInput
	Hookscript         pulumi.StringPtrInput
	Hostname           pulumi.StringPtrInput
	IgnoreUnpackErrors pulumi.BoolPtrInput
	Lock               pulumi.StringPtrInput
	Memory             pulumi.IntPtrInput
	Mountpoints        LXCContainerMountpointArrayInput
	Nameserver         pulumi.StringPtrInput
	Networks           LXCContainerNetworkArrayInput
	Onboot             pulumi.BoolPtrInput
	Ostemplate         pulumi.StringPtrInput
	Ostype             pulumi.StringPtrInput
	Password           pulumi.StringPtrInput
	Pool               pulumi.StringPtrInput
	Protection         pulumi.BoolPtrInput
	Restore            pulumi.BoolPtrInput
	Rootfs             LXCContainerRootfsPtrInput
	Searchdomain       pulumi.StringPtrInput
	SshPublicKeys      pulumi.StringPtrInput
	Start              pulumi.BoolPtrInput
	Startup            pulumi.StringPtrInput
	Swap               pulumi.IntPtrInput
	Tags               pulumi.StringPtrInput
	TargetNode         pulumi.StringPtrInput
	Template           pulumi.BoolPtrInput
	Tty                pulumi.IntPtrInput
	Unique             pulumi.BoolPtrInput
	Unprivileged       pulumi.BoolPtrInput
	Unuseds            pulumi.StringArrayInput
	// The VM identifier in proxmox (100-999999999)
	Vmid pulumi.IntPtrInput
}

func (LXCContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*lxccontainerState)(nil)).Elem()
}

type lxccontainerArgs struct {
	Arch               *string                  `pulumi:"arch"`
	Bwlimit            *int                     `pulumi:"bwlimit"`
	Clone              *string                  `pulumi:"clone"`
	CloneStorage       *string                  `pulumi:"cloneStorage"`
	Cmode              *string                  `pulumi:"cmode"`
	Console            *bool                    `pulumi:"console"`
	Cores              *int                     `pulumi:"cores"`
	Cpulimit           *int                     `pulumi:"cpulimit"`
	Cpuunits           *int                     `pulumi:"cpuunits"`
	Description        *string                  `pulumi:"description"`
	Features           *LXCContainerFeatures    `pulumi:"features"`
	Force              *bool                    `pulumi:"force"`
	Full               *bool                    `pulumi:"full"`
	Hagroup            *string                  `pulumi:"hagroup"`
	Hastate            *string                  `pulumi:"hastate"`
	Hookscript         *string                  `pulumi:"hookscript"`
	Hostname           *string                  `pulumi:"hostname"`
	IgnoreUnpackErrors *bool                    `pulumi:"ignoreUnpackErrors"`
	Lock               *string                  `pulumi:"lock"`
	Memory             *int                     `pulumi:"memory"`
	Mountpoints        []LXCContainerMountpoint `pulumi:"mountpoints"`
	Nameserver         *string                  `pulumi:"nameserver"`
	Networks           []LXCContainerNetwork    `pulumi:"networks"`
	Onboot             *bool                    `pulumi:"onboot"`
	Ostemplate         *string                  `pulumi:"ostemplate"`
	Ostype             *string                  `pulumi:"ostype"`
	Password           *string                  `pulumi:"password"`
	Pool               *string                  `pulumi:"pool"`
	Protection         *bool                    `pulumi:"protection"`
	Restore            *bool                    `pulumi:"restore"`
	Rootfs             *LXCContainerRootfs      `pulumi:"rootfs"`
	Searchdomain       *string                  `pulumi:"searchdomain"`
	SshPublicKeys      *string                  `pulumi:"sshPublicKeys"`
	Start              *bool                    `pulumi:"start"`
	Startup            *string                  `pulumi:"startup"`
	Swap               *int                     `pulumi:"swap"`
	Tags               *string                  `pulumi:"tags"`
	TargetNode         string                   `pulumi:"targetNode"`
	Template           *bool                    `pulumi:"template"`
	Tty                *int                     `pulumi:"tty"`
	Unique             *bool                    `pulumi:"unique"`
	Unprivileged       *bool                    `pulumi:"unprivileged"`
	// The VM identifier in proxmox (100-999999999)
	Vmid *int `pulumi:"vmid"`
}

// The set of arguments for constructing a LXCContainer resource.
type LXCContainerArgs struct {
	Arch               pulumi.StringPtrInput
	Bwlimit            pulumi.IntPtrInput
	Clone              pulumi.StringPtrInput
	CloneStorage       pulumi.StringPtrInput
	Cmode              pulumi.StringPtrInput
	Console            pulumi.BoolPtrInput
	Cores              pulumi.IntPtrInput
	Cpulimit           pulumi.IntPtrInput
	Cpuunits           pulumi.IntPtrInput
	Description        pulumi.StringPtrInput
	Features           LXCContainerFeaturesPtrInput
	Force              pulumi.BoolPtrInput
	Full               pulumi.BoolPtrInput
	Hagroup            pulumi.StringPtrInput
	Hastate            pulumi.StringPtrInput
	Hookscript         pulumi.StringPtrInput
	Hostname           pulumi.StringPtrInput
	IgnoreUnpackErrors pulumi.BoolPtrInput
	Lock               pulumi.StringPtrInput
	Memory             pulumi.IntPtrInput
	Mountpoints        LXCContainerMountpointArrayInput
	Nameserver         pulumi.StringPtrInput
	Networks           LXCContainerNetworkArrayInput
	Onboot             pulumi.BoolPtrInput
	Ostemplate         pulumi.StringPtrInput
	Ostype             pulumi.StringPtrInput
	Password           pulumi.StringPtrInput
	Pool               pulumi.StringPtrInput
	Protection         pulumi.BoolPtrInput
	Restore            pulumi.BoolPtrInput
	Rootfs             LXCContainerRootfsPtrInput
	Searchdomain       pulumi.StringPtrInput
	SshPublicKeys      pulumi.StringPtrInput
	Start              pulumi.BoolPtrInput
	Startup            pulumi.StringPtrInput
	Swap               pulumi.IntPtrInput
	Tags               pulumi.StringPtrInput
	TargetNode         pulumi.StringInput
	Template           pulumi.BoolPtrInput
	Tty                pulumi.IntPtrInput
	Unique             pulumi.BoolPtrInput
	Unprivileged       pulumi.BoolPtrInput
	// The VM identifier in proxmox (100-999999999)
	Vmid pulumi.IntPtrInput
}

func (LXCContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lxccontainerArgs)(nil)).Elem()
}

type LXCContainerInput interface {
	pulumi.Input

	ToLXCContainerOutput() LXCContainerOutput
	ToLXCContainerOutputWithContext(ctx context.Context) LXCContainerOutput
}

func (*LXCContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**LXCContainer)(nil)).Elem()
}

func (i *LXCContainer) ToLXCContainerOutput() LXCContainerOutput {
	return i.ToLXCContainerOutputWithContext(context.Background())
}

func (i *LXCContainer) ToLXCContainerOutputWithContext(ctx context.Context) LXCContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCContainerOutput)
}

// LXCContainerArrayInput is an input type that accepts LXCContainerArray and LXCContainerArrayOutput values.
// You can construct a concrete instance of `LXCContainerArrayInput` via:
//
//	LXCContainerArray{ LXCContainerArgs{...} }
type LXCContainerArrayInput interface {
	pulumi.Input

	ToLXCContainerArrayOutput() LXCContainerArrayOutput
	ToLXCContainerArrayOutputWithContext(context.Context) LXCContainerArrayOutput
}

type LXCContainerArray []LXCContainerInput

func (LXCContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LXCContainer)(nil)).Elem()
}

func (i LXCContainerArray) ToLXCContainerArrayOutput() LXCContainerArrayOutput {
	return i.ToLXCContainerArrayOutputWithContext(context.Background())
}

func (i LXCContainerArray) ToLXCContainerArrayOutputWithContext(ctx context.Context) LXCContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCContainerArrayOutput)
}

// LXCContainerMapInput is an input type that accepts LXCContainerMap and LXCContainerMapOutput values.
// You can construct a concrete instance of `LXCContainerMapInput` via:
//
//	LXCContainerMap{ "key": LXCContainerArgs{...} }
type LXCContainerMapInput interface {
	pulumi.Input

	ToLXCContainerMapOutput() LXCContainerMapOutput
	ToLXCContainerMapOutputWithContext(context.Context) LXCContainerMapOutput
}

type LXCContainerMap map[string]LXCContainerInput

func (LXCContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LXCContainer)(nil)).Elem()
}

func (i LXCContainerMap) ToLXCContainerMapOutput() LXCContainerMapOutput {
	return i.ToLXCContainerMapOutputWithContext(context.Background())
}

func (i LXCContainerMap) ToLXCContainerMapOutputWithContext(ctx context.Context) LXCContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCContainerMapOutput)
}

type LXCContainerOutput struct{ *pulumi.OutputState }

func (LXCContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LXCContainer)(nil)).Elem()
}

func (o LXCContainerOutput) ToLXCContainerOutput() LXCContainerOutput {
	return o
}

func (o LXCContainerOutput) ToLXCContainerOutputWithContext(ctx context.Context) LXCContainerOutput {
	return o
}

func (o LXCContainerOutput) Arch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Arch }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Bwlimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Bwlimit }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Clone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Clone }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) CloneStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.CloneStorage }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Cmode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Cmode }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Console() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Console }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Cores }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Cpulimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Cpulimit }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Cpuunits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Cpuunits }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Features() LXCContainerFeaturesPtrOutput {
	return o.ApplyT(func(v *LXCContainer) LXCContainerFeaturesPtrOutput { return v.Features }).(LXCContainerFeaturesPtrOutput)
}

func (o LXCContainerOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Full() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Full }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Hagroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Hagroup }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Hastate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Hastate }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Hookscript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Hookscript }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) IgnoreUnpackErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.IgnoreUnpackErrors }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Lock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Lock }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Memory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Memory }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Mountpoints() LXCContainerMountpointArrayOutput {
	return o.ApplyT(func(v *LXCContainer) LXCContainerMountpointArrayOutput { return v.Mountpoints }).(LXCContainerMountpointArrayOutput)
}

func (o LXCContainerOutput) Nameserver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Nameserver }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Networks() LXCContainerNetworkArrayOutput {
	return o.ApplyT(func(v *LXCContainer) LXCContainerNetworkArrayOutput { return v.Networks }).(LXCContainerNetworkArrayOutput)
}

func (o LXCContainerOutput) Onboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Onboot }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Ostemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Ostemplate }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Ostype() pulumi.StringOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringOutput { return v.Ostype }).(pulumi.StringOutput)
}

func (o LXCContainerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Pool }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Protection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Protection }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Restore() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Restore }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Rootfs() LXCContainerRootfsPtrOutput {
	return o.ApplyT(func(v *LXCContainer) LXCContainerRootfsPtrOutput { return v.Rootfs }).(LXCContainerRootfsPtrOutput)
}

func (o LXCContainerOutput) Searchdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Searchdomain }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) SshPublicKeys() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.SshPublicKeys }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Start() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Start }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Startup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Startup }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) Swap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Swap }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Tags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringPtrOutput { return v.Tags }).(pulumi.StringPtrOutput)
}

func (o LXCContainerOutput) TargetNode() pulumi.StringOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringOutput { return v.TargetNode }).(pulumi.StringOutput)
}

func (o LXCContainerOutput) Template() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Template }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Tty() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntPtrOutput { return v.Tty }).(pulumi.IntPtrOutput)
}

func (o LXCContainerOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Unique }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Unprivileged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.BoolPtrOutput { return v.Unprivileged }).(pulumi.BoolPtrOutput)
}

func (o LXCContainerOutput) Unuseds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.StringArrayOutput { return v.Unuseds }).(pulumi.StringArrayOutput)
}

// The VM identifier in proxmox (100-999999999)
func (o LXCContainerOutput) Vmid() pulumi.IntOutput {
	return o.ApplyT(func(v *LXCContainer) pulumi.IntOutput { return v.Vmid }).(pulumi.IntOutput)
}

type LXCContainerArrayOutput struct{ *pulumi.OutputState }

func (LXCContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LXCContainer)(nil)).Elem()
}

func (o LXCContainerArrayOutput) ToLXCContainerArrayOutput() LXCContainerArrayOutput {
	return o
}

func (o LXCContainerArrayOutput) ToLXCContainerArrayOutputWithContext(ctx context.Context) LXCContainerArrayOutput {
	return o
}

func (o LXCContainerArrayOutput) Index(i pulumi.IntInput) LXCContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LXCContainer {
		return vs[0].([]*LXCContainer)[vs[1].(int)]
	}).(LXCContainerOutput)
}

type LXCContainerMapOutput struct{ *pulumi.OutputState }

func (LXCContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LXCContainer)(nil)).Elem()
}

func (o LXCContainerMapOutput) ToLXCContainerMapOutput() LXCContainerMapOutput {
	return o
}

func (o LXCContainerMapOutput) ToLXCContainerMapOutputWithContext(ctx context.Context) LXCContainerMapOutput {
	return o
}

func (o LXCContainerMapOutput) MapIndex(k pulumi.StringInput) LXCContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LXCContainer {
		return vs[0].(map[string]*LXCContainer)[vs[1].(string)]
	}).(LXCContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LXCContainerInput)(nil)).Elem(), &LXCContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LXCContainerArrayInput)(nil)).Elem(), LXCContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LXCContainerMapInput)(nil)).Elem(), LXCContainerMap{})
	pulumi.RegisterOutputType(LXCContainerOutput{})
	pulumi.RegisterOutputType(LXCContainerArrayOutput{})
	pulumi.RegisterOutputType(LXCContainerMapOutput{})
}
