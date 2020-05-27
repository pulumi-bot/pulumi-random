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
    /// The resource `random..RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.
    /// 
    /// This resource can be used in conjunction with resources that have
    /// the `create_before_destroy` lifecycle flag set, to avoid conflicts with
    /// unique names during the brief period where both the old and new resources
    /// exist concurrently.
    /// 
    /// ## Example Usage
    /// 
    /// 
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
    ///         var priority = new Random.RandomInteger("priority", new Random.RandomIntegerArgs
    ///         {
    ///             Keepers = 
    ///             {
    ///                 { "listener_arn", @var.Listener_arn },
    ///             },
    ///             Max = 50000,
    ///             Min = 1,
    ///         });
    ///         var main = new Aws.Alb.ListenerRule("main", new Aws.Alb.ListenerRuleArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new Aws.Alb.Inputs.ListenerRuleActionArgs
    ///                 {
    ///                     TargetGroupArn = @var.Target_group_arn,
    ///                     Type = "forward",
    ///                 },
    ///             },
    ///             ListenerArn = @var.Listener_arn,
    ///             Priority = priority.Result,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RandomInteger : Pulumi.CustomResource
    {
        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The maximum inclusive value of the range.
        /// </summary>
        [Output("max")]
        public Output<int> Max { get; private set; } = null!;

        /// <summary>
        /// The minimum inclusive value of the range.
        /// </summary>
        [Output("min")]
        public Output<int> Min { get; private set; } = null!;

        /// <summary>
        /// (int) The random Integer result.
        /// </summary>
        [Output("result")]
        public Output<int> Result { get; private set; } = null!;

        /// <summary>
        /// A custom seed to always produce the same value.
        /// </summary>
        [Output("seed")]
        public Output<string?> Seed { get; private set; } = null!;


        /// <summary>
        /// Create a RandomInteger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomInteger(string name, RandomIntegerArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomInteger:RandomInteger", name, args ?? new RandomIntegerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomInteger(string name, Input<string> id, RandomIntegerState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomInteger:RandomInteger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RandomInteger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomInteger Get(string name, Input<string> id, RandomIntegerState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomInteger(name, id, state, options);
        }
    }

    public sealed class RandomIntegerArgs : Pulumi.ResourceArgs
    {
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
        /// The maximum inclusive value of the range.
        /// </summary>
        [Input("max", required: true)]
        public Input<int> Max { get; set; } = null!;

        /// <summary>
        /// The minimum inclusive value of the range.
        /// </summary>
        [Input("min", required: true)]
        public Input<int> Min { get; set; } = null!;

        /// <summary>
        /// A custom seed to always produce the same value.
        /// </summary>
        [Input("seed")]
        public Input<string>? Seed { get; set; }

        public RandomIntegerArgs()
        {
        }
    }

    public sealed class RandomIntegerState : Pulumi.ResourceArgs
    {
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
        /// The maximum inclusive value of the range.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum inclusive value of the range.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        /// <summary>
        /// (int) The random Integer result.
        /// </summary>
        [Input("result")]
        public Input<int>? Result { get; set; }

        /// <summary>
        /// A custom seed to always produce the same value.
        /// </summary>
        [Input("seed")]
        public Input<string>? Seed { get; set; }

        public RandomIntegerState()
        {
        }
    }
}
