# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .random_id import *
from .random_integer import *
from .random_password import *
from .random_pet import *
from .random_shuffle import *
from .random_string import *
from .random_uuid import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "random",
  "mod": "index/randomString",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomString:RandomString": "RandomString"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomUuid",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomUuid:RandomUuid": "RandomUuid"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomId",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomId:RandomId": "RandomId"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomInteger",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomInteger:RandomInteger": "RandomInteger"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomPassword",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomPassword:RandomPassword": "RandomPassword"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomPet",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomPet:RandomPet": "RandomPet"
  }
 },
 {
  "pkg": "random",
  "mod": "index/randomShuffle",
  "fqn": "pulumi_random",
  "classes": {
   "random:index/randomShuffle:RandomShuffle": "RandomShuffle"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "random",
  "token": "pulumi:providers:random",
  "fqn": "pulumi_random",
  "class": "Provider"
 }
]
"""
)
