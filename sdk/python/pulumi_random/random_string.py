# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class RandomString(pulumi.CustomResource):
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new id to be generated. See
    the main provider documentation for more information.
    """
    length: pulumi.Output[float]
    """
    The length of the string desired
    """
    lower: pulumi.Output[bool]
    """
    (default true) Include lowercase alphabet characters
    in random string.
    """
    min_lower: pulumi.Output[float]
    """
    (default 0) Minimum number of lowercase alphabet
    characters in random string.
    """
    min_numeric: pulumi.Output[float]
    """
    (default 0) Minimum number of numeric characters
    in random string.
    """
    min_special: pulumi.Output[float]
    """
    (default 0) Minimum number of special characters
    in random string.
    """
    min_upper: pulumi.Output[float]
    """
    (default 0) Minimum number of uppercase alphabet
    characters in random string.
    """
    number: pulumi.Output[bool]
    """
    (default true) Include numeric characters in random
    string.
    """
    override_special: pulumi.Output[str]
    """
    Supply your own list of special characters to
    use for string generation.  This overrides the default character list in the special
    argument.  The special argument must still be set to true for any overwritten
    characters to be used in generation.
    """
    result: pulumi.Output[str]
    """
    Random string generated.
    """
    special: pulumi.Output[bool]
    """
    (default true) Include special characters in random
    string. These are `!@#$%&*()-_=+[]{}<>:?`
    """
    upper: pulumi.Output[bool]
    """
    (default true) Include uppercase alphabet characters
    in random string.
    """
    def __init__(__self__, resource_name, opts=None, keepers=None, length=None, lower=None, min_lower=None, min_numeric=None, min_special=None, min_upper=None, number=None, override_special=None, special=None, upper=None, __props__=None, __name__=None, __opts__=None):
        """
        The resource `RandomString` generates a random permutation of alphanumeric
        characters and optionally special characters.

        This resource *does* use a cryptographic random number generator.

        Historically this resource's intended usage has been ambiguous as the original example
        used it in a password. For backwards compatibility it will
        continue to exist. For unique ids please use random_id, for sensitive
        random values please use random_password.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_random as random

        random = random.RandomString("random",
            length=16,
            override_special="/@£$",
            special=True)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[float] length: The length of the string desired
        :param pulumi.Input[bool] lower: (default true) Include lowercase alphabet characters
               in random string.
        :param pulumi.Input[float] min_lower: (default 0) Minimum number of lowercase alphabet
               characters in random string.
        :param pulumi.Input[float] min_numeric: (default 0) Minimum number of numeric characters
               in random string.
        :param pulumi.Input[float] min_special: (default 0) Minimum number of special characters
               in random string.
        :param pulumi.Input[float] min_upper: (default 0) Minimum number of uppercase alphabet
               characters in random string.
        :param pulumi.Input[bool] number: (default true) Include numeric characters in random
               string.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to
               use for string generation.  This overrides the default character list in the special
               argument.  The special argument must still be set to true for any overwritten
               characters to be used in generation.
        :param pulumi.Input[bool] special: (default true) Include special characters in random
               string. These are `!@#$%&*()-_=+[]{}<>:?`
        :param pulumi.Input[bool] upper: (default true) Include uppercase alphabet characters
               in random string.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['keepers'] = keepers
            if length is None:
                raise TypeError("Missing required property 'length'")
            __props__['length'] = length
            __props__['lower'] = lower
            __props__['min_lower'] = min_lower
            __props__['min_numeric'] = min_numeric
            __props__['min_special'] = min_special
            __props__['min_upper'] = min_upper
            __props__['number'] = number
            __props__['override_special'] = override_special
            __props__['special'] = special
            __props__['upper'] = upper
            __props__['result'] = None
        super(RandomString, __self__).__init__(
            'random:index/randomString:RandomString',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, keepers=None, length=None, lower=None, min_lower=None, min_numeric=None, min_special=None, min_upper=None, number=None, override_special=None, result=None, special=None, upper=None):
        """
        Get an existing RandomString resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[float] length: The length of the string desired
        :param pulumi.Input[bool] lower: (default true) Include lowercase alphabet characters
               in random string.
        :param pulumi.Input[float] min_lower: (default 0) Minimum number of lowercase alphabet
               characters in random string.
        :param pulumi.Input[float] min_numeric: (default 0) Minimum number of numeric characters
               in random string.
        :param pulumi.Input[float] min_special: (default 0) Minimum number of special characters
               in random string.
        :param pulumi.Input[float] min_upper: (default 0) Minimum number of uppercase alphabet
               characters in random string.
        :param pulumi.Input[bool] number: (default true) Include numeric characters in random
               string.
        :param pulumi.Input[str] override_special: Supply your own list of special characters to
               use for string generation.  This overrides the default character list in the special
               argument.  The special argument must still be set to true for any overwritten
               characters to be used in generation.
        :param pulumi.Input[str] result: Random string generated.
        :param pulumi.Input[bool] special: (default true) Include special characters in random
               string. These are `!@#$%&*()-_=+[]{}<>:?`
        :param pulumi.Input[bool] upper: (default true) Include uppercase alphabet characters
               in random string.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["keepers"] = keepers
        __props__["length"] = length
        __props__["lower"] = lower
        __props__["min_lower"] = min_lower
        __props__["min_numeric"] = min_numeric
        __props__["min_special"] = min_special
        __props__["min_upper"] = min_upper
        __props__["number"] = number
        __props__["override_special"] = override_special
        __props__["result"] = result
        __props__["special"] = special
        __props__["upper"] = upper
        return RandomString(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

