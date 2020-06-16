// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random..RandomId` generates random numbers that are intended to be
 * used as unique identifiers for other resources.
 *
 * This resource *does* use a cryptographic random number generator in order
 * to minimize the chance of collisions, making the results of this resource
 * when a 16-byte identifier is requested of equivalent uniqueness to a
 * type-4 UUID.
 *
 * This resource can be used in conjunction with resources that have
 * the `createBeforeDestroy` lifecycle flag set to avoid conflicts with
 * unique names during the brief period where both the old and new resources
 * exist concurrently.
 * ## Example Usage
 *
 * The following example shows how to generate a unique name for an AWS EC2
 * instance that changes each time a new AMI id is selected.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * const serverRandomId = new random.RandomId("server", {
 *     byteLength: 8,
 *     keepers: {
 *         // Generate a new id each time we switch to a new AMI id
 *         ami_id: var_ami_id,
 *     },
 * });
 * const serverInstance = new aws.ec2.Instance("server", {
 *     ami: serverRandomId.keepers.apply(keepers => keepers.amiId),
 *     tags: {
 *         Name: pulumi.interpolate`web-server ${serverRandomId.hex}`,
 *     },
 * });
 * ```
 */
export class RandomId extends pulumi.CustomResource {
    /**
     * Get an existing RandomId resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomIdState, opts?: pulumi.CustomResourceOptions): RandomId {
        return new RandomId(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomId:RandomId';

    /**
     * Returns true if the given object is an instance of RandomId.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomId {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomId.__pulumiType;
    }

    /**
     * @deprecated Use b64_url for old behavior, or b64_std for standard base64 encoding
     */
    public /*out*/ readonly b64!: pulumi.Output<string>;
    /**
     * The generated id presented in base64 without additional transformations.
     */
    public /*out*/ readonly b64Std!: pulumi.Output<string>;
    /**
     * The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
     */
    public /*out*/ readonly b64Url!: pulumi.Output<string>;
    /**
     * The number of random bytes to produce. The
     * minimum value is 1, which produces eight bits of randomness.
     */
    public readonly byteLength!: pulumi.Output<number>;
    /**
     * The generated id presented in non-padded decimal digits.
     */
    public /*out*/ readonly dec!: pulumi.Output<string>;
    /**
     * The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
     */
    public /*out*/ readonly hex!: pulumi.Output<string>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Arbitrary string to prefix the output value with. This
     * string is supplied as-is, meaning it is not guaranteed to be URL-safe or
     * base64 encoded.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;

    /**
     * Create a RandomId resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomIdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomIdArgs | RandomIdState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RandomIdState | undefined;
            inputs["b64"] = state ? state.b64 : undefined;
            inputs["b64Std"] = state ? state.b64Std : undefined;
            inputs["b64Url"] = state ? state.b64Url : undefined;
            inputs["byteLength"] = state ? state.byteLength : undefined;
            inputs["dec"] = state ? state.dec : undefined;
            inputs["hex"] = state ? state.hex : undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
        } else {
            const args = argsOrState as RandomIdArgs | undefined;
            if (!args || args.byteLength === undefined) {
                throw new Error("Missing required property 'byteLength'");
            }
            inputs["byteLength"] = args ? args.byteLength : undefined;
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["b64"] = undefined /*out*/;
            inputs["b64Std"] = undefined /*out*/;
            inputs["b64Url"] = undefined /*out*/;
            inputs["dec"] = undefined /*out*/;
            inputs["hex"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RandomId.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomId resources.
 */
export interface RandomIdState {
    /**
     * @deprecated Use b64_url for old behavior, or b64_std for standard base64 encoding
     */
    readonly b64?: pulumi.Input<string>;
    /**
     * The generated id presented in base64 without additional transformations.
     */
    readonly b64Std?: pulumi.Input<string>;
    /**
     * The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
     */
    readonly b64Url?: pulumi.Input<string>;
    /**
     * The number of random bytes to produce. The
     * minimum value is 1, which produces eight bits of randomness.
     */
    readonly byteLength?: pulumi.Input<number>;
    /**
     * The generated id presented in non-padded decimal digits.
     */
    readonly dec?: pulumi.Input<string>;
    /**
     * The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
     */
    readonly hex?: pulumi.Input<string>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    readonly keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * Arbitrary string to prefix the output value with. This
     * string is supplied as-is, meaning it is not guaranteed to be URL-safe or
     * base64 encoded.
     */
    readonly prefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomId resource.
 */
export interface RandomIdArgs {
    /**
     * The number of random bytes to produce. The
     * minimum value is 1, which produces eight bits of randomness.
     */
    readonly byteLength: pulumi.Input<number>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    readonly keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * Arbitrary string to prefix the output value with. This
     * string is supplied as-is, meaning it is not guaranteed to be URL-safe or
     * base64 encoded.
     */
    readonly prefix?: pulumi.Input<string>;
}
