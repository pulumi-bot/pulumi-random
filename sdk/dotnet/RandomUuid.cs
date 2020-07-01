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
    /// The resource `random.RandomUuid` generates random uuid string that is intended to be
    /// used as unique identifiers for other resources.
    /// 
    /// This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
    /// for use with services needed a unique string identifier.
    /// 
    /// ## Example Usage
    /// 
    /// The following example shows how to generate a unique name for an Azure Resource Group.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// using Random = Pulumi.Random;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testRandomUuid = new Random.RandomUuid("testRandomUuid", new Random.RandomUuidArgs
    ///         {
    ///         });
    ///         var testResourceGroup = new Azure.Core.ResourceGroup("testResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "Central US",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RandomUuid : Pulumi.CustomResource
    {
        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new uuid to be generated.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// The generated uuid presented in string format.
        /// </summary>
        [Output("result")]
        public Output<string> Result { get; private set; } = null!;


        /// <summary>
        /// Create a RandomUuid resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomUuid(string name, RandomUuidArgs? args = null, CustomResourceOptions? options = null)
            : base("random:index/randomUuid:RandomUuid", name, args ?? new RandomUuidArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomUuid(string name, Input<string> id, RandomUuidState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomUuid:RandomUuid", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RandomUuid resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomUuid Get(string name, Input<string> id, RandomUuidState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomUuid(name, id, state, options);
        }
    }

    public sealed class RandomUuidArgs : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new uuid to be generated.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        public RandomUuidArgs()
        {
        }
    }

    public sealed class RandomUuidState : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new uuid to be generated.
        /// </summary>
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        /// <summary>
        /// The generated uuid presented in string format.
        /// </summary>
        [Input("result")]
        public Input<string>? Result { get; set; }

        public RandomUuidState()
        {
        }
    }
}
