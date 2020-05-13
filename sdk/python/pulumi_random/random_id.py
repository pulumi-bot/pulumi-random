# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class RandomId(pulumi.CustomResource):
    b64: pulumi.Output[str]
    b64_std: pulumi.Output[str]
    """
    The generated id presented in base64 without additional transformations.
    """
    b64_url: pulumi.Output[str]
    """
    The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
    """
    byte_length: pulumi.Output[float]
    """
    The number of random bytes to produce. The
    minimum value is 1, which produces eight bits of randomness.
    """
    dec: pulumi.Output[str]
    """
    The generated id presented in non-padded decimal digits.
    """
    hex: pulumi.Output[str]
    """
    The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
    """
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new id to be generated. See
    the main provider documentation for more information.
    """
    prefix: pulumi.Output[str]
    """
    Arbitrary string to prefix the output value with. This
    string is supplied as-is, meaning it is not guaranteed to be URL-safe or
    base64 encoded.
    """
    def __init__(__self__, resource_name, opts=None, byte_length=None, keepers=None, prefix=None, __props__=None, __name__=None, __opts__=None):
        """
        The resource `.RandomId` generates random numbers that are intended to be
        used as unique identifiers for other resources.

        This resource *does* use a cryptographic random number generator in order
        to minimize the chance of collisions, making the results of this resource
        when a 16-byte identifier is requested of equivalent uniqueness to a
        type-4 UUID.

        This resource can be used in conjunction with resources that have
        the `create_before_destroy` lifecycle flag set to avoid conflicts with
        unique names during the brief period where both the old and new resources
        exist concurrently.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        server_random_id = random.RandomId("serverRandomId",
            byte_length=8,
            keepers={
                "ami_id": var["ami_id"],
            })
        server_instance = aws.ec2.Instance("serverInstance",
            ami=server_random_id.keepers["amiId"],
            tags={
                "Name": server_random_id.hex.apply(lambda hex: f"web-server {hex}"),
            })
        ```

        > This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/id.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] byte_length: The number of random bytes to produce. The
               minimum value is 1, which produces eight bits of randomness.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[str] prefix: Arbitrary string to prefix the output value with. This
               string is supplied as-is, meaning it is not guaranteed to be URL-safe or
               base64 encoded.
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

            if byte_length is None:
                raise TypeError("Missing required property 'byte_length'")
            __props__['byte_length'] = byte_length
            __props__['keepers'] = keepers
            __props__['prefix'] = prefix
            __props__['b64'] = None
            __props__['b64_std'] = None
            __props__['b64_url'] = None
            __props__['dec'] = None
            __props__['hex'] = None
        super(RandomId, __self__).__init__(
            'random:index/randomId:RandomId',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, b64=None, b64_std=None, b64_url=None, byte_length=None, dec=None, hex=None, keepers=None, prefix=None):
        """
        Get an existing RandomId resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] b64_std: The generated id presented in base64 without additional transformations.
        :param pulumi.Input[str] b64_url: The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters `_` and `-`.
        :param pulumi.Input[float] byte_length: The number of random bytes to produce. The
               minimum value is 1, which produces eight bits of randomness.
        :param pulumi.Input[str] dec: The generated id presented in non-padded decimal digits.
        :param pulumi.Input[str] hex: The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[str] prefix: Arbitrary string to prefix the output value with. This
               string is supplied as-is, meaning it is not guaranteed to be URL-safe or
               base64 encoded.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["b64"] = b64
        __props__["b64_std"] = b64_std
        __props__["b64_url"] = b64_url
        __props__["byte_length"] = byte_length
        __props__["dec"] = dec
        __props__["hex"] = hex
        __props__["keepers"] = keepers
        __props__["prefix"] = prefix
        return RandomId(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

