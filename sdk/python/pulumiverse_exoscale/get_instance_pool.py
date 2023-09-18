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

__all__ = [
    'GetInstancePoolResult',
    'AwaitableGetInstancePoolResult',
    'get_instance_pool',
    'get_instance_pool_output',
]

@pulumi.output_type
class GetInstancePoolResult:
    """
    A collection of values returned by getInstancePool.
    """
    def __init__(__self__, affinity_group_ids=None, deploy_target_id=None, description=None, disk_size=None, elastic_ip_ids=None, id=None, instance_prefix=None, instance_type=None, instances=None, ipv6=None, key_pair=None, labels=None, name=None, network_ids=None, security_group_ids=None, size=None, state=None, template_id=None, user_data=None, zone=None):
        if affinity_group_ids and not isinstance(affinity_group_ids, list):
            raise TypeError("Expected argument 'affinity_group_ids' to be a list")
        pulumi.set(__self__, "affinity_group_ids", affinity_group_ids)
        if deploy_target_id and not isinstance(deploy_target_id, str):
            raise TypeError("Expected argument 'deploy_target_id' to be a str")
        pulumi.set(__self__, "deploy_target_id", deploy_target_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size and not isinstance(disk_size, int):
            raise TypeError("Expected argument 'disk_size' to be a int")
        pulumi.set(__self__, "disk_size", disk_size)
        if elastic_ip_ids and not isinstance(elastic_ip_ids, list):
            raise TypeError("Expected argument 'elastic_ip_ids' to be a list")
        pulumi.set(__self__, "elastic_ip_ids", elastic_ip_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_prefix and not isinstance(instance_prefix, str):
            raise TypeError("Expected argument 'instance_prefix' to be a str")
        pulumi.set(__self__, "instance_prefix", instance_prefix)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if ipv6 and not isinstance(ipv6, bool):
            raise TypeError("Expected argument 'ipv6' to be a bool")
        pulumi.set(__self__, "ipv6", ipv6)
        if key_pair and not isinstance(key_pair, str):
            raise TypeError("Expected argument 'key_pair' to be a str")
        pulumi.set(__self__, "key_pair", key_pair)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_ids and not isinstance(network_ids, list):
            raise TypeError("Expected argument 'network_ids' to be a list")
        pulumi.set(__self__, "network_ids", network_ids)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if template_id and not isinstance(template_id, str):
            raise TypeError("Expected argument 'template_id' to be a str")
        pulumi.set(__self__, "template_id", template_id)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="affinityGroupIds")
    def affinity_group_ids(self) -> Sequence[str]:
        """
        The list of attached exoscale*anti*affinity_group (IDs).
        """
        return pulumi.get(self, "affinity_group_ids")

    @property
    @pulumi.getter(name="deployTargetId")
    def deploy_target_id(self) -> str:
        """
        The deploy target ID.
        """
        return pulumi.get(self, "deploy_target_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The instance pool description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> int:
        """
        The managed instances disk size.
        """
        return pulumi.get(self, "disk_size")

    @property
    @pulumi.getter(name="elasticIpIds")
    def elastic_ip_ids(self) -> Sequence[str]:
        """
        The list of attached exoscale*elastic*ip (IDs).
        """
        return pulumi.get(self, "elastic_ip_ids")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The instance pool ID to match (conflicts with `name`).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instancePrefix")
    def instance_prefix(self) -> str:
        """
        The string used to prefix the managed instances name.
        """
        return pulumi.get(self, "instance_prefix")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The managed instances type.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetInstancePoolInstanceResult']:
        """
        The list of managed instances. Structure is documented below.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def ipv6(self) -> bool:
        """
        Whether IPv6 is enabled on managed instances.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> str:
        """
        The exoscale*ssh*key (name) authorized on the managed instances.
        """
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        A map of key/value labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The pool name to match (conflicts with `id`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkIds")
    def network_ids(self) -> Sequence[str]:
        """
        The list of attached exoscale*private*network (IDs).
        """
        return pulumi.get(self, "network_ids")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        """
        The list of attached exoscale*security*group (IDs).
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The number managed instances.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The pool state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> str:
        """
        The managed instances exoscale*compute*template ID.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> str:
        """
        [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        """
        return pulumi.get(self, "zone")


class AwaitableGetInstancePoolResult(GetInstancePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancePoolResult(
            affinity_group_ids=self.affinity_group_ids,
            deploy_target_id=self.deploy_target_id,
            description=self.description,
            disk_size=self.disk_size,
            elastic_ip_ids=self.elastic_ip_ids,
            id=self.id,
            instance_prefix=self.instance_prefix,
            instance_type=self.instance_type,
            instances=self.instances,
            ipv6=self.ipv6,
            key_pair=self.key_pair,
            labels=self.labels,
            name=self.name,
            network_ids=self.network_ids,
            security_group_ids=self.security_group_ids,
            size=self.size,
            state=self.state,
            template_id=self.template_id,
            user_data=self.user_data,
            zone=self.zone)


def get_instance_pool(id: Optional[str] = None,
                      labels: Optional[Mapping[str, str]] = None,
                      name: Optional[str] = None,
                      zone: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancePoolResult:
    """
    Use this data source to access information about an existing resource.

    :param str id: The instance pool ID to match (conflicts with `name`).
    :param Mapping[str, str] labels: A map of key/value labels.
    :param str name: The pool name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['labels'] = labels
    __args__['name'] = name
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('exoscale:index/getInstancePool:getInstancePool', __args__, opts=opts, typ=GetInstancePoolResult).value

    return AwaitableGetInstancePoolResult(
        affinity_group_ids=pulumi.get(__ret__, 'affinity_group_ids'),
        deploy_target_id=pulumi.get(__ret__, 'deploy_target_id'),
        description=pulumi.get(__ret__, 'description'),
        disk_size=pulumi.get(__ret__, 'disk_size'),
        elastic_ip_ids=pulumi.get(__ret__, 'elastic_ip_ids'),
        id=pulumi.get(__ret__, 'id'),
        instance_prefix=pulumi.get(__ret__, 'instance_prefix'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        instances=pulumi.get(__ret__, 'instances'),
        ipv6=pulumi.get(__ret__, 'ipv6'),
        key_pair=pulumi.get(__ret__, 'key_pair'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        network_ids=pulumi.get(__ret__, 'network_ids'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        size=pulumi.get(__ret__, 'size'),
        state=pulumi.get(__ret__, 'state'),
        template_id=pulumi.get(__ret__, 'template_id'),
        user_data=pulumi.get(__ret__, 'user_data'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_instance_pool)
def get_instance_pool_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                             labels: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             zone: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancePoolResult]:
    """
    Use this data source to access information about an existing resource.

    :param str id: The instance pool ID to match (conflicts with `name`).
    :param Mapping[str, str] labels: A map of key/value labels.
    :param str name: The pool name to match (conflicts with `id`).
    :param str zone: The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
    """
    ...
