// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random..RandomShuffle` generates a random permutation of a list
 * of strings given as an argument.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/shuffle.html.md.
 */
export class RandomShuffle extends pulumi.CustomResource {
    /**
     * Get an existing RandomShuffle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomShuffleState, opts?: pulumi.CustomResourceOptions): RandomShuffle {
        return new RandomShuffle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomShuffle:RandomShuffle';

    /**
     * Returns true if the given object is an instance of RandomShuffle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomShuffle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomShuffle.__pulumiType;
    }

    /**
     * The list of strings to shuffle.
     */
    public readonly inputs!: pulumi.Output<string[]>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Random permutation of the list of strings given in `input`.
     */
    public /*out*/ readonly results!: pulumi.Output<string[]>;
    /**
     * The number of results to return. Defaults to
     * the number of items in the `input` list. If fewer items are requested,
     * some elements will be excluded from the result. If more items are requested,
     * items will be repeated in the result but not more frequently than the number
     * of items in the input list.
     */
    public readonly resultCount!: pulumi.Output<number | undefined>;
    public readonly seed!: pulumi.Output<string | undefined>;

    /**
     * Create a RandomShuffle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomShuffleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomShuffleArgs | RandomShuffleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RandomShuffleState | undefined;
            inputs["inputs"] = state ? state.inputs : undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["results"] = state ? state.results : undefined;
            inputs["resultCount"] = state ? state.resultCount : undefined;
            inputs["seed"] = state ? state.seed : undefined;
        } else {
            const args = argsOrState as RandomShuffleArgs | undefined;
            if (!args || args.inputs === undefined) {
                throw new Error("Missing required property 'inputs'");
            }
            inputs["inputs"] = args ? args.inputs : undefined;
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["resultCount"] = args ? args.resultCount : undefined;
            inputs["seed"] = args ? args.seed : undefined;
            inputs["results"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RandomShuffle.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomShuffle resources.
 */
export interface RandomShuffleState {
    /**
     * The list of strings to shuffle.
     */
    readonly inputs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    readonly keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * Random permutation of the list of strings given in `input`.
     */
    readonly results?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of results to return. Defaults to
     * the number of items in the `input` list. If fewer items are requested,
     * some elements will be excluded from the result. If more items are requested,
     * items will be repeated in the result but not more frequently than the number
     * of items in the input list.
     */
    readonly resultCount?: pulumi.Input<number>;
    readonly seed?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomShuffle resource.
 */
export interface RandomShuffleArgs {
    /**
     * The list of strings to shuffle.
     */
    readonly inputs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    readonly keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The number of results to return. Defaults to
     * the number of items in the `input` list. If fewer items are requested,
     * some elements will be excluded from the result. If more items are requested,
     * items will be repeated in the result but not more frequently than the number
     * of items in the input list.
     */
    readonly resultCount?: pulumi.Input<number>;
    readonly seed?: pulumi.Input<string>;
}
