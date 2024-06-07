# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert import *
from .asset import *
from .asset_attribute import *
from .asset_metadata import *
from .component import *
from .component_routine import *
from .dashboard import *
from .dashboard_chart import *
from .dashboard_tab import *
from .file import *
from .file_folder import *
from .function import *
from .get_asset_kinds import *
from .provider import *
from .secret import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_splight.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_splight.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "splight",
  "mod": "index/alert",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/alert:Alert": "Alert"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/asset",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/asset:Asset": "Asset"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/assetAttribute",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/assetAttribute:AssetAttribute": "AssetAttribute"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/assetMetadata",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/assetMetadata:AssetMetadata": "AssetMetadata"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/component",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/component:Component": "Component"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/componentRoutine",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/componentRoutine:ComponentRoutine": "ComponentRoutine"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/dashboard",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/dashboard:Dashboard": "Dashboard"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/dashboardChart",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/dashboardChart:DashboardChart": "DashboardChart"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/dashboardTab",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/dashboardTab:DashboardTab": "DashboardTab"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/file",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/file:File": "File"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/fileFolder",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/fileFolder:FileFolder": "FileFolder"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/function",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/function:Function": "Function"
  }
 },
 {
  "pkg": "splight",
  "mod": "index/secret",
  "fqn": "pulumi_splight",
  "classes": {
   "splight:index/secret:Secret": "Secret"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "splight",
  "token": "pulumi:providers:splight",
  "fqn": "pulumi_splight",
  "class": "Provider"
 }
]
"""
)
