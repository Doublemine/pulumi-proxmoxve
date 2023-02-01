// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the proxmoxve package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'proxmoxve';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * API TokenID e.g. root@pam!mytesttoken
     */
    public readonly pmApiTokenId!: pulumi.Output<string | undefined>;
    /**
     * The secret uuid corresponding to a TokenID
     */
    public readonly pmApiTokenSecret!: pulumi.Output<string | undefined>;
    /**
     * https://host.fqdn:8006/api2/json
     */
    public readonly pmApiUrl!: pulumi.Output<string | undefined>;
    /**
     * Set custom http headers e.g. Key,Value,Key1,Value1
     */
    public readonly pmHttpHeaders!: pulumi.Output<string | undefined>;
    /**
     * Write logs to this specific file
     */
    public readonly pmLogFile!: pulumi.Output<string | undefined>;
    /**
     * OTP 2FA code (if required)
     */
    public readonly pmOtp!: pulumi.Output<string | undefined>;
    /**
     * Password to authenticate into proxmox
     */
    public readonly pmPassword!: pulumi.Output<string | undefined>;
    /**
     * Proxy Server passed to Api client(useful for debugging). Syntax: http://proxy:port
     */
    public readonly pmProxyServer!: pulumi.Output<string | undefined>;
    /**
     * Username e.g. myuser or myuser@pam
     */
    public readonly pmUser!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["pmApiTokenId"] = args ? args.pmApiTokenId : undefined;
            resourceInputs["pmApiTokenSecret"] = args ? args.pmApiTokenSecret : undefined;
            resourceInputs["pmApiUrl"] = args ? args.pmApiUrl : undefined;
            resourceInputs["pmDangerouslyIgnoreUnknownAttributes"] = pulumi.output(args ? args.pmDangerouslyIgnoreUnknownAttributes : undefined).apply(JSON.stringify);
            resourceInputs["pmDebug"] = pulumi.output(args ? args.pmDebug : undefined).apply(JSON.stringify);
            resourceInputs["pmHttpHeaders"] = args ? args.pmHttpHeaders : undefined;
            resourceInputs["pmLogEnable"] = pulumi.output(args ? args.pmLogEnable : undefined).apply(JSON.stringify);
            resourceInputs["pmLogFile"] = args ? args.pmLogFile : undefined;
            resourceInputs["pmLogLevels"] = pulumi.output(args ? args.pmLogLevels : undefined).apply(JSON.stringify);
            resourceInputs["pmOtp"] = args ? args.pmOtp : undefined;
            resourceInputs["pmParallel"] = pulumi.output(args ? args.pmParallel : undefined).apply(JSON.stringify);
            resourceInputs["pmPassword"] = args ? args.pmPassword : undefined;
            resourceInputs["pmProxyServer"] = args ? args.pmProxyServer : undefined;
            resourceInputs["pmTimeout"] = pulumi.output(args ? args.pmTimeout : undefined).apply(JSON.stringify);
            resourceInputs["pmTlsInsecure"] = pulumi.output(args ? args.pmTlsInsecure : undefined).apply(JSON.stringify);
            resourceInputs["pmUser"] = args ? args.pmUser : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * API TokenID e.g. root@pam!mytesttoken
     */
    pmApiTokenId?: pulumi.Input<string>;
    /**
     * The secret uuid corresponding to a TokenID
     */
    pmApiTokenSecret?: pulumi.Input<string>;
    /**
     * https://host.fqdn:8006/api2/json
     */
    pmApiUrl?: pulumi.Input<string>;
    /**
     * By default this provider will exit if an unknown attribute is found. This is to prevent the accidential destruction of
     * VMs or Data when something in the proxmox API has changed/updated and is not confirmed to work with this provider. Set
     * this to true at your own risk. It may allow you to proceed in cases when the provider refuses to work, but be aware of
     * the danger in doing so.
     */
    pmDangerouslyIgnoreUnknownAttributes?: pulumi.Input<boolean>;
    /**
     * Enable or disable the verbose debug output from proxmox api
     */
    pmDebug?: pulumi.Input<boolean>;
    /**
     * Set custom http headers e.g. Key,Value,Key1,Value1
     */
    pmHttpHeaders?: pulumi.Input<string>;
    /**
     * Enable provider logging to get proxmox API logs
     */
    pmLogEnable?: pulumi.Input<boolean>;
    /**
     * Write logs to this specific file
     */
    pmLogFile?: pulumi.Input<string>;
    /**
     * Configure the logging level to display; trace, debug, info, warn, etc
     */
    pmLogLevels?: pulumi.Input<{[key: string]: any}>;
    /**
     * OTP 2FA code (if required)
     */
    pmOtp?: pulumi.Input<string>;
    pmParallel?: pulumi.Input<number>;
    /**
     * Password to authenticate into proxmox
     */
    pmPassword?: pulumi.Input<string>;
    /**
     * Proxy Server passed to Api client(useful for debugging). Syntax: http://proxy:port
     */
    pmProxyServer?: pulumi.Input<string>;
    /**
     * How many seconds to wait for operations for both provider and api-client, default is 20m
     */
    pmTimeout?: pulumi.Input<number>;
    /**
     * By default, every TLS connection is verified to be secure. This option allows terraform to proceed and operate on
     * servers considered insecure. For example if you're connecting to a remote host and you do not have the CA cert that
     * issued the proxmox api url's certificate.
     */
    pmTlsInsecure?: pulumi.Input<boolean>;
    /**
     * Username e.g. myuser or myuser@pam
     */
    pmUser?: pulumi.Input<string>;
}
