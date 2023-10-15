# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetIAMAPIKeyResult',
    'AwaitableGetIAMAPIKeyResult',
    'get_iamapi_key',
    'get_iamapi_key_output',
]

@pulumi.output_type
class GetIAMAPIKeyResult:
    """
    A collection of values returned by getIAMAPIKey.
    """
    def __init__(__self__, id=None, key=None, name=None, role_id=None, timeouts=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if role_id and not isinstance(role_id, str):
            raise TypeError("Expected argument 'role_id' to be a str")
        pulumi.set(__self__, "role_id", role_id)
        if timeouts and not isinstance(timeouts, dict):
            raise TypeError("Expected argument 'timeouts' to be a dict")
        pulumi.set(__self__, "timeouts", timeouts)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The IAM API Key to match.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        IAM API Key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> str:
        """
        IAM API Key role ID.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def timeouts(self) -> Optional['outputs.GetIAMAPIKeyTimeoutsResult']:
        return pulumi.get(self, "timeouts")


class AwaitableGetIAMAPIKeyResult(GetIAMAPIKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIAMAPIKeyResult(
            id=self.id,
            key=self.key,
            name=self.name,
            role_id=self.role_id,
            timeouts=self.timeouts)


def get_iamapi_key(key: Optional[str] = None,
                   timeouts: Optional[pulumi.InputType['GetIAMAPIKeyTimeoutsArgs']] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIAMAPIKeyResult:
    """
    Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) API Key.

    Corresponding resource: exoscale_iam_role.


    :param str key: The IAM API Key to match.
    """
    __args__ = dict()
    __args__['key'] = key
    __args__['timeouts'] = timeouts
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getIAMAPIKey:getIAMAPIKey', __args__, opts=opts, typ=GetIAMAPIKeyResult).value

    return AwaitableGetIAMAPIKeyResult(
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        name=pulumi.get(__ret__, 'name'),
        role_id=pulumi.get(__ret__, 'role_id'),
        timeouts=pulumi.get(__ret__, 'timeouts'))


@_utilities.lift_output_func(get_iamapi_key)
def get_iamapi_key_output(key: Optional[pulumi.Input[str]] = None,
                          timeouts: Optional[pulumi.Input[Optional[pulumi.InputType['GetIAMAPIKeyTimeoutsArgs']]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIAMAPIKeyResult]:
    """
    Fetch Exoscale [IAM](https://community.exoscale.com/documentation/iam/) API Key.

    Corresponding resource: exoscale_iam_role.


    :param str key: The IAM API Key to match.
    """
    ...