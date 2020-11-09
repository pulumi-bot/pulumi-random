# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['RandomPassword']


class RandomPassword(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 lower: Optional[pulumi.Input[bool]] = None,
                 min_lower: Optional[pulumi.Input[int]] = None,
                 min_numeric: Optional[pulumi.Input[int]] = None,
                 min_special: Optional[pulumi.Input[int]] = None,
                 min_upper: Optional[pulumi.Input[int]] = None,
                 number: Optional[pulumi.Input[bool]] = None,
                 override_special: Optional[pulumi.Input[str]] = None,
                 special: Optional[pulumi.Input[bool]] = None,
                 upper: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Identical to the `RandomString` resource with the exception that the
        result is treated as sensitive and, thus, _not_ displayed in console output.

        > **Note:** All attributes including the generated password will be stored in
        the raw state as plain-text.

        This resource *does* use a cryptographic random number generator.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        password = random.RandomPassword("password",
            length=16,
            special=True,
            override_special="_%@")
        example = aws.rds.Instance("example",
            instance_class="db.t3.micro",
            allocated_storage=64,
            engine="mysql",
            username="someone",
            password=password.result)
        ```

        ## Import

        Random Password can be imported by specifying the value of the string

        ```sh
         $ pulumi import random:index/randomPassword:RandomPassword password securepassword
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
            opts.version = _utilities.get_version()
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
        super(RandomPassword, __self__).__init__(
            'random:index/randomPassword:RandomPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            length: Optional[pulumi.Input[int]] = None,
            lower: Optional[pulumi.Input[bool]] = None,
            min_lower: Optional[pulumi.Input[int]] = None,
            min_numeric: Optional[pulumi.Input[int]] = None,
            min_special: Optional[pulumi.Input[int]] = None,
            min_upper: Optional[pulumi.Input[int]] = None,
            number: Optional[pulumi.Input[bool]] = None,
            override_special: Optional[pulumi.Input[str]] = None,
            result: Optional[pulumi.Input[str]] = None,
            special: Optional[pulumi.Input[bool]] = None,
            upper: Optional[pulumi.Input[bool]] = None) -> 'RandomPassword':
        """
        Get an existing RandomPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
        return RandomPassword(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter
    def length(self) -> pulumi.Output[int]:
        return pulumi.get(self, "length")

    @property
    @pulumi.getter
    def lower(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "lower")

    @property
    @pulumi.getter(name="minLower")
    def min_lower(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_lower")

    @property
    @pulumi.getter(name="minNumeric")
    def min_numeric(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_numeric")

    @property
    @pulumi.getter(name="minSpecial")
    def min_special(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_special")

    @property
    @pulumi.getter(name="minUpper")
    def min_upper(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_upper")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "number")

    @property
    @pulumi.getter(name="overrideSpecial")
    def override_special(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "override_special")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def special(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "special")

    @property
    @pulumi.getter
    def upper(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "upper")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

