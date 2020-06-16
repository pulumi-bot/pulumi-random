// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Random
{
    /// <summary>
    /// The resource `random..RandomShuffle` generates a random permutation of a list
    /// of strings given as an argument.
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Random = Pulumi.Random;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var az = new Random.RandomShuffle("az", new Random.RandomShuffleArgs
    ///         {
    ///             Inputs = 
    ///             {
    ///                 "us-west-1a",
    ///                 "us-west-1c",
    ///                 "us-west-1d",
    ///                 "us-west-1e",
    ///             },
    ///             ResultCount = 2,
    ///         });
    ///         var example = new Aws.Elb.LoadBalancer("example", new Aws.Elb.LoadBalancerArgs
    ///         {
    ///             AvailabilityZones = az.Results,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RandomShuffle : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of strings to shuffle.
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<string>> Inputs { get; private set; } = null!;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The number of results to return. Defaults to
        /// the number of items in the `input` list. If fewer items are requested,
        /// some elements will be excluded from the result. If more items are requested,
        /// items will be repeated in the result but not more frequently than the number
        /// of items in the input list.
        /// </summary>
        [Output("resultCount")]
        public Output<int?> ResultCount { get; private set; } = null!;

        /// <summary>
        /// Random permutation of the list of strings given in `input`.
        /// </summary>
        [Output("results")]
        public Output<ImmutableArray<string>> Results { get; private set; } = null!;

        [Output("seed")]
        public Output<string?> Seed { get; private set; } = null!;


        /// <summary>
        /// Create a RandomShuffle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomShuffle(string name, RandomShuffleArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomShuffle:RandomShuffle", name, args ?? new RandomShuffleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomShuffle(string name, Input<string> id, RandomShuffleState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomShuffle:RandomShuffle", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RandomShuffle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomShuffle Get(string name, Input<string> id, RandomShuffleState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomShuffle(name, id, state, options);
        }
    }

    public sealed class RandomShuffleArgs : Pulumi.ResourceArgs
    {
        [Input("inputs", required: true)]
        private InputList<string>? _inputs;

        /// <summary>
        /// The list of strings to shuffle.
        /// </summary>
        public InputList<string> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<string>());
            set => _inputs = value;
        }

        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        /// <summary>
        /// The number of results to return. Defaults to
        /// the number of items in the `input` list. If fewer items are requested,
        /// some elements will be excluded from the result. If more items are requested,
        /// items will be repeated in the result but not more frequently than the number
        /// of items in the input list.
        /// </summary>
        [Input("resultCount")]
        public Input<int>? ResultCount { get; set; }

        [Input("seed")]
        public Input<string>? Seed { get; set; }

        public RandomShuffleArgs()
        {
        }
    }

    public sealed class RandomShuffleState : Pulumi.ResourceArgs
    {
        [Input("inputs")]
        private InputList<string>? _inputs;

        /// <summary>
        /// The list of strings to shuffle.
        /// </summary>
        public InputList<string> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<string>());
            set => _inputs = value;
        }

        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        /// <summary>
        /// The number of results to return. Defaults to
        /// the number of items in the `input` list. If fewer items are requested,
        /// some elements will be excluded from the result. If more items are requested,
        /// items will be repeated in the result but not more frequently than the number
        /// of items in the input list.
        /// </summary>
        [Input("resultCount")]
        public Input<int>? ResultCount { get; set; }

        [Input("results")]
        private InputList<string>? _results;

        /// <summary>
        /// Random permutation of the list of strings given in `input`.
        /// </summary>
        public InputList<string> Results
        {
            get => _results ?? (_results = new InputList<string>());
            set => _results = value;
        }

        [Input("seed")]
        public Input<string>? Seed { get; set; }

        public RandomShuffleState()
        {
        }
    }
}
