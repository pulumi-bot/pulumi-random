// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random.RandomPet` generates random pet names that are intended to be used as unique identifiers for other resources.
 *
 * This resource can be used in conjunction with resources that have the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * // The following example shows how to generate a unique pet name
 * // for an AWS EC2 instance that changes each time a new AMI id is
 * // selected.
 * const serverRandomPet = new random.RandomPet("serverRandomPet", {keepers: {
 *     ami_id: _var.ami_id,
 * }});
 * const serverInstance = new aws.ec2.Instance("serverInstance", {
 *     tags: {
 *         Name: pulumi.interpolate`web-server-${serverRandomPet.id}`,
 *     },
 *     ami: serverRandomPet.keepers.amiId,
 * });
 * // ... (other aws_instance arguments) ...
 * ```
 */
export class RandomPet extends pulumi.CustomResource {
    /**
     * Get an existing RandomPet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomPetState, opts?: pulumi.CustomResourceOptions): RandomPet {
        return new RandomPet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomPet:RandomPet';

    /**
     * Returns true if the given object is an instance of RandomPet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomPet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomPet.__pulumiType;
    }

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The length (in words) of the pet name.
     */
    public readonly length!: pulumi.Output<number | undefined>;
    /**
     * A string to prefix the name with.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * The character to separate words in the pet name.
     */
    public readonly separator!: pulumi.Output<string | undefined>;

    /**
     * Create a RandomPet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RandomPetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomPetArgs | RandomPetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomPetState | undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["length"] = state ? state.length : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
            inputs["separator"] = state ? state.separator : undefined;
        } else {
            const args = argsOrState as RandomPetArgs | undefined;
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["length"] = args ? args.length : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["separator"] = args ? args.separator : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RandomPet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomPet resources.
 */
export interface RandomPetState {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length (in words) of the pet name.
     */
    length?: pulumi.Input<number>;
    /**
     * A string to prefix the name with.
     */
    prefix?: pulumi.Input<string>;
    /**
     * The character to separate words in the pet name.
     */
    separator?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomPet resource.
 */
export interface RandomPetArgs {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length (in words) of the pet name.
     */
    length?: pulumi.Input<number>;
    /**
     * A string to prefix the name with.
     */
    prefix?: pulumi.Input<string>;
    /**
     * The character to separate words in the pet name.
     */
    separator?: pulumi.Input<string>;
}
