# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class RandomUuid(pulumi.CustomResource):
    keepers: pulumi.Output[dict]
    """
    Arbitrary map of values that, when changed, will
    trigger a new uuid to be generated. See
    the main provider documentation for more information.
    """
    result: pulumi.Output[str]
    """
    The generated uuid presented in string format.
    """
    def __init__(__self__, resource_name, opts=None, keepers=None, __props__=None, __name__=None, __opts__=None):
        """
        The resource `.RandomUuid` generates random uuid string that is intended to be
        used as unique identifiers for other resources.

        This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
        for use with services needed a unique string identifier.

        ## Example Usage

        The following example shows how to generate a unique name for an Azure Resource Group.

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        test_random_uuid = random.RandomUuid("testRandomUuid")
        test_resource_group = azure.core.ResourceGroup("testResourceGroup", location="Central US")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new uuid to be generated. See
               the main provider documentation for more information.
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
            __props__['result'] = None
        super(RandomUuid, __self__).__init__(
            'random:index/randomUuid:RandomUuid',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, keepers=None, result=None):
        """
        Get an existing RandomUuid resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] keepers: Arbitrary map of values that, when changed, will
               trigger a new uuid to be generated. See
               the main provider documentation for more information.
        :param pulumi.Input[str] result: The generated uuid presented in string format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["keepers"] = keepers
        __props__["result"] = result
        return RandomUuid(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
