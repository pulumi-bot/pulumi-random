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
    /// The resource `random.RandomId` generates random numbers that are intended to be
    /// used as unique identifiers for other resources.
    /// 
    /// This resource *does* use a cryptographic random number generator in order
    /// to minimize the chance of collisions, making the results of this resource
    /// when a 16-byte identifier is requested of equivalent uniqueness to a
    /// type-4 UUID.
    /// 
    /// This resource can be used in conjunction with resources that have
    /// the `create_before_destroy` lifecycle flag set to avoid conflicts with
    /// unique names during the brief period where both the old and new resources
    /// exist concurrently.
    /// 
    /// ## Example Usage
    /// 
    /// The following example shows how to generate a unique name for an AWS EC2
    /// instance that changes each time a new AMI id is selected.
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
    ///         var serverRandomId = new Random.RandomId("serverRandomId", new Random.RandomIdArgs
    ///         {
    ///             ByteLength = 8,
    ///             Keepers = 
    ///             {
    ///                 { "ami_id", @var.Ami_id },
    ///             },
    ///         });
    ///         var serverInstance = new Aws.Ec2.Instance("serverInstance", new Aws.Ec2.InstanceArgs
    ///         {
    ///             Ami = serverRandomId.Keepers.Apply(keepers =&gt; keepers.AmiId),
    ///             Tags = 
    ///             {
    ///                 { "Name", serverRandomId.Hex.Apply(hex =&gt; Output.Format($"web-server {hex}")) },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Random Ids can be imported using the `b64_url` with an optional `prefix`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example with no prefix
    /// 
    /// ```sh
    ///  $ pulumi import random:index/randomId:RandomId server p-9hUg
    /// ```
    /// 
    ///  Example with prefix (prefix is separated by a `,`)
    /// 
    /// ```sh
    ///  $ pulumi import random:index/randomId:RandomId server my-prefix-,p-9hUg
    /// ```
    /// </summary>
    [RandomResourceType("random:index/randomId:RandomId")]
    public partial class RandomId : Pulumi.CustomResource
    {
        /// <summary>
        /// The generated id presented in base64 without additional transformations.
        /// </summary>
        [Output("b64Std")]
        public Output<string> B64Std { get; private set; } = null!;

        /// <summary>
        /// The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
        /// </summary>
        [Output("b64Url")]
        public Output<string> B64Url { get; private set; } = null!;

        /// <summary>
        /// The number of random bytes to produce. The
        /// minimum value is 1, which produces eight bits of randomness.
        /// </summary>
        [Output("byteLength")]
        public Output<int> ByteLength { get; private set; } = null!;

        /// <summary>
        /// The generated id presented in non-padded decimal digits.
        /// </summary>
        [Output("dec")]
        public Output<string> Dec { get; private set; } = null!;

        /// <summary>
        /// The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
        /// </summary>
        [Output("hex")]
        public Output<string> Hex { get; private set; } = null!;

        /// <summary>
        /// Arbitrary map of values that, when changed, will
        /// trigger a new id to be generated. See
        /// the main provider documentation for more information.
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        /// <summary>
        /// Arbitrary string to prefix the output value with. This
        /// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
        /// base64 encoded.
        /// </summary>
        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;


        /// <summary>
        /// Create a RandomId resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomId(string name, RandomIdArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomId:RandomId", name, args ?? new RandomIdArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomId(string name, Input<string> id, RandomIdState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomId:RandomId", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RandomId resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomId Get(string name, Input<string> id, RandomIdState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomId(name, id, state, options);
        }
    }

    public sealed class RandomIdArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of random bytes to produce. The
        /// minimum value is 1, which produces eight bits of randomness.
        /// </summary>
        [Input("byteLength", required: true)]
        public Input<int> ByteLength { get; set; } = null!;

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
        /// Arbitrary string to prefix the output value with. This
        /// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
        /// base64 encoded.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public RandomIdArgs()
        {
        }
    }

    public sealed class RandomIdState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The generated id presented in base64 without additional transformations.
        /// </summary>
        [Input("b64Std")]
        public Input<string>? B64Std { get; set; }

        /// <summary>
        /// The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
        /// </summary>
        [Input("b64Url")]
        public Input<string>? B64Url { get; set; }

        /// <summary>
        /// The number of random bytes to produce. The
        /// minimum value is 1, which produces eight bits of randomness.
        /// </summary>
        [Input("byteLength")]
        public Input<int>? ByteLength { get; set; }

        /// <summary>
        /// The generated id presented in non-padded decimal digits.
        /// </summary>
        [Input("dec")]
        public Input<string>? Dec { get; set; }

        /// <summary>
        /// The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
        /// </summary>
        [Input("hex")]
        public Input<string>? Hex { get; set; }

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
        /// Arbitrary string to prefix the output value with. This
        /// string is supplied as-is, meaning it is not guaranteed to be URL-safe or
        /// base64 encoded.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public RandomIdState()
        {
        }
    }
}
