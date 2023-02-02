// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class LXCDisk extends pulumi.CustomResource {
    /**
     * Get an existing LXCDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LXCDiskState, opts?: pulumi.CustomResourceOptions): LXCDisk {
        return new LXCDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'proxmoxve:proxmox/lXCDisk:LXCDisk';

    /**
     * Returns true if the given object is an instance of LXCDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LXCDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LXCDisk.__pulumiType;
    }

    public readonly acl!: pulumi.Output<boolean | undefined>;
    public readonly backup!: pulumi.Output<boolean | undefined>;
    public readonly container!: pulumi.Output<string>;
    public readonly mountoptions!: pulumi.Output<outputs.proxmox.LXCDiskMountoptions | undefined>;
    public readonly mp!: pulumi.Output<string>;
    public readonly quota!: pulumi.Output<boolean | undefined>;
    public readonly replicate!: pulumi.Output<boolean | undefined>;
    public readonly shared!: pulumi.Output<boolean | undefined>;
    public readonly size!: pulumi.Output<string>;
    public readonly slot!: pulumi.Output<number>;
    public readonly storage!: pulumi.Output<string>;
    public readonly volume!: pulumi.Output<string>;

    /**
     * Create a LXCDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LXCDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LXCDiskArgs | LXCDiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LXCDiskState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["backup"] = state ? state.backup : undefined;
            resourceInputs["container"] = state ? state.container : undefined;
            resourceInputs["mountoptions"] = state ? state.mountoptions : undefined;
            resourceInputs["mp"] = state ? state.mp : undefined;
            resourceInputs["quota"] = state ? state.quota : undefined;
            resourceInputs["replicate"] = state ? state.replicate : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["slot"] = state ? state.slot : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["volume"] = state ? state.volume : undefined;
        } else {
            const args = argsOrState as LXCDiskArgs | undefined;
            if ((!args || args.container === undefined) && !opts.urn) {
                throw new Error("Missing required property 'container'");
            }
            if ((!args || args.mp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mp'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.slot === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slot'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["backup"] = args ? args.backup : undefined;
            resourceInputs["container"] = args ? args.container : undefined;
            resourceInputs["mountoptions"] = args ? args.mountoptions : undefined;
            resourceInputs["mp"] = args ? args.mp : undefined;
            resourceInputs["quota"] = args ? args.quota : undefined;
            resourceInputs["replicate"] = args ? args.replicate : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["slot"] = args ? args.slot : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["volume"] = args ? args.volume : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LXCDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LXCDisk resources.
 */
export interface LXCDiskState {
    acl?: pulumi.Input<boolean>;
    backup?: pulumi.Input<boolean>;
    container?: pulumi.Input<string>;
    mountoptions?: pulumi.Input<inputs.proxmox.LXCDiskMountoptions>;
    mp?: pulumi.Input<string>;
    quota?: pulumi.Input<boolean>;
    replicate?: pulumi.Input<boolean>;
    shared?: pulumi.Input<boolean>;
    size?: pulumi.Input<string>;
    slot?: pulumi.Input<number>;
    storage?: pulumi.Input<string>;
    volume?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LXCDisk resource.
 */
export interface LXCDiskArgs {
    acl?: pulumi.Input<boolean>;
    backup?: pulumi.Input<boolean>;
    container: pulumi.Input<string>;
    mountoptions?: pulumi.Input<inputs.proxmox.LXCDiskMountoptions>;
    mp: pulumi.Input<string>;
    quota?: pulumi.Input<boolean>;
    replicate?: pulumi.Input<boolean>;
    shared?: pulumi.Input<boolean>;
    size: pulumi.Input<string>;
    slot: pulumi.Input<number>;
    storage: pulumi.Input<string>;
    volume?: pulumi.Input<string>;
}