# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetBusesResult',
    'AwaitableGetBusesResult',
    'get_buses',
    'get_buses_output',
]

@pulumi.output_type
class GetBusesResult:
    """
    A collection of values returned by getBuses.
    """
    def __init__(__self__, id=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetBusesTagResult']:
        return pulumi.get(self, "tags")


class AwaitableGetBusesResult(GetBusesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBusesResult(
            id=self.id,
            tags=self.tags)


def get_buses(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBusesResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_splight as splight

    buses = splight.get_buses()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('splight:index/getBuses:getBuses', __args__, opts=opts, typ=GetBusesResult).value

    return AwaitableGetBusesResult(
        id=pulumi.get(__ret__, 'id'),
        tags=pulumi.get(__ret__, 'tags'))
def get_buses_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBusesResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_splight as splight

    buses = splight.get_buses()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('splight:index/getBuses:getBuses', __args__, opts=opts, typ=GetBusesResult)
    return __ret__.apply(lambda __response__: GetBusesResult(
        id=pulumi.get(__response__, 'id'),
        tags=pulumi.get(__response__, 'tags')))
