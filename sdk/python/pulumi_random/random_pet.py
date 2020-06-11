# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class RandomPet(pulumi.CustomResource):
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new id to be generated. See
    the main provider documentation for more information.
    """
    length: pulumi.Output[float]
    """
    The length (in words) of the pet name.
    """
    prefix: pulumi.Output[str]
    """
    A string to prefix the name with.
    """
    separator: pulumi.Output[str]
    """
    The character to separate words in the pet name.
    """
    def __init__(__self__, resource_name, opts=None, keepers=None, length=None, prefix=None, separator=None, __props__=None, __name__=None, __opts__=None):
        """
        The resource `RandomPet` generates random pet names that are intended to be
        used as unique identifiers for other resources.

        This resource can be used in conjunction with resources that have
        the `create_before_destroy` lifecycle flag set, to avoid conflicts with
        unique names during the brief period where both the old and new resources
        exist concurrently.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_random as random

        server_random_pet = random.RandomPet("serverRandomPet", keepers={
            "ami_id": var["ami_id"],
        })
        server_instance = aws.ec2.Instance("serverInstance",
            ami=server_random_pet.keepers["amiId"],
            tags={
                "Name": server_random_pet.id.apply(lambda id: f"web-server-{id}"),
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[float] length: The length (in words) of the pet name.
        :param pulumi.Input[str] prefix: A string to prefix the name with.
        :param pulumi.Input[str] separator: The character to separate words in the pet name.
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
            __props__['length'] = length
            __props__['prefix'] = prefix
            __props__['separator'] = separator
        super(RandomPet, __self__).__init__(
            'random:index/randomPet:RandomPet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, keepers=None, length=None, prefix=None, separator=None):
        """
        Get an existing RandomPet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new id to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[float] length: The length (in words) of the pet name.
        :param pulumi.Input[str] prefix: A string to prefix the name with.
        :param pulumi.Input[str] separator: The character to separate words in the pet name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["keepers"] = keepers
        __props__["length"] = length
        __props__["prefix"] = prefix
        __props__["separator"] = separator
        return RandomPet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

