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
    /// &gt; **Note:** Requires random provider version &gt;= 2.2.0
    /// 
    /// Identical to random.RandomString with the exception that the
    /// result is treated as sensitive and, thus, _not_ displayed in console output.
    /// 
    /// &gt; **Note:** All attributes including the generated password will be stored in
    /// the raw state as plain-text. [Read more about sensitive data in
    /// state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// This resource *does* use a cryptographic random number generator.
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
    ///         var password = new Random.RandomPassword("password", new Random.RandomPasswordArgs
    ///         {
    ///             Length = 16,
    ///             Special = true,
    ///             OverrideSpecial = "_%@",
    ///         });
    ///         var example = new Aws.Rds.Instance("example", new Aws.Rds.InstanceArgs
    ///         {
    ///             InstanceClass = "db.t3.micro",
    ///             AllocatedStorage = 64,
    ///             Engine = "mysql",
    ///             Username = "someone",
    ///             Password = random_string.Password.Result,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RandomPassword : Pulumi.CustomResource
    {
        [Output("keepers")]
        public Output<ImmutableDictionary<string, object>?> Keepers { get; private set; } = null!;

        [Output("length")]
        public Output<int> Length { get; private set; } = null!;

        [Output("lower")]
        public Output<bool?> Lower { get; private set; } = null!;

        [Output("minLower")]
        public Output<int?> MinLower { get; private set; } = null!;

        [Output("minNumeric")]
        public Output<int?> MinNumeric { get; private set; } = null!;

        [Output("minSpecial")]
        public Output<int?> MinSpecial { get; private set; } = null!;

        [Output("minUpper")]
        public Output<int?> MinUpper { get; private set; } = null!;

        [Output("number")]
        public Output<bool?> Number { get; private set; } = null!;

        [Output("overrideSpecial")]
        public Output<string?> OverrideSpecial { get; private set; } = null!;

        [Output("result")]
        public Output<string> Result { get; private set; } = null!;

        [Output("special")]
        public Output<bool?> Special { get; private set; } = null!;

        [Output("upper")]
        public Output<bool?> Upper { get; private set; } = null!;


        /// <summary>
        /// Create a RandomPassword resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RandomPassword(string name, RandomPasswordArgs args, CustomResourceOptions? options = null)
            : base("random:index/randomPassword:RandomPassword", name, args ?? new RandomPasswordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RandomPassword(string name, Input<string> id, RandomPasswordState? state = null, CustomResourceOptions? options = null)
            : base("random:index/randomPassword:RandomPassword", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RandomPassword resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RandomPassword Get(string name, Input<string> id, RandomPasswordState? state = null, CustomResourceOptions? options = null)
        {
            return new RandomPassword(name, id, state, options);
        }
    }

    public sealed class RandomPasswordArgs : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        [Input("length", required: true)]
        public Input<int> Length { get; set; } = null!;

        [Input("lower")]
        public Input<bool>? Lower { get; set; }

        [Input("minLower")]
        public Input<int>? MinLower { get; set; }

        [Input("minNumeric")]
        public Input<int>? MinNumeric { get; set; }

        [Input("minSpecial")]
        public Input<int>? MinSpecial { get; set; }

        [Input("minUpper")]
        public Input<int>? MinUpper { get; set; }

        [Input("number")]
        public Input<bool>? Number { get; set; }

        [Input("overrideSpecial")]
        public Input<string>? OverrideSpecial { get; set; }

        [Input("special")]
        public Input<bool>? Special { get; set; }

        [Input("upper")]
        public Input<bool>? Upper { get; set; }

        public RandomPasswordArgs()
        {
        }
    }

    public sealed class RandomPasswordState : Pulumi.ResourceArgs
    {
        [Input("keepers")]
        private InputMap<object>? _keepers;
        public InputMap<object> Keepers
        {
            get => _keepers ?? (_keepers = new InputMap<object>());
            set => _keepers = value;
        }

        [Input("length")]
        public Input<int>? Length { get; set; }

        [Input("lower")]
        public Input<bool>? Lower { get; set; }

        [Input("minLower")]
        public Input<int>? MinLower { get; set; }

        [Input("minNumeric")]
        public Input<int>? MinNumeric { get; set; }

        [Input("minSpecial")]
        public Input<int>? MinSpecial { get; set; }

        [Input("minUpper")]
        public Input<int>? MinUpper { get; set; }

        [Input("number")]
        public Input<bool>? Number { get; set; }

        [Input("overrideSpecial")]
        public Input<string>? OverrideSpecial { get; set; }

        [Input("result")]
        public Input<string>? Result { get; set; }

        [Input("special")]
        public Input<bool>? Special { get; set; }

        [Input("upper")]
        public Input<bool>? Upper { get; set; }

        public RandomPasswordState()
        {
        }
    }
}
