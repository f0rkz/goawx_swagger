# Data24

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | **int32** | Cloud credential to use for inventory updates. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**EnabledValue** | **string** | Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var&#x3D;\&quot;status.power_state\&quot;and enabled_value&#x3D;\&quot;powered_on\&quot; with host variables:{   \&quot;status\&quot;: {     \&quot;power_state\&quot;: \&quot;powered_on\&quot;,     \&quot;created\&quot;: \&quot;2018-02-01T08:00:00.000000Z:00\&quot;,     \&quot;healthy\&quot;: true    },    \&quot;name\&quot;: \&quot;foobar\&quot;,    \&quot;ip_address\&quot;: \&quot;192.168.2.1\&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled | [optional] [default to null]
**EnabledVar** | **string** | Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \&quot;foo.bar\&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\&quot;foo\&quot;, {}).get(\&quot;bar\&quot;, default) | [optional] [default to null]
**ExecutionEnvironment** | **int32** | The container image to be used for execution. | [optional] [default to null]
**HostFilter** | **string** | Regex where only matching hosts will be imported. | [optional] [default to null]
**Inventory** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**Overwrite** | **bool** | Overwrite local groups and hosts from remote inventory source. | [optional] [default to null]
**OverwriteVars** | **bool** | Overwrite local variables from remote inventory source. | [optional] [default to null]
**Source** | **string** |  | [optional] [default to null]
**SourcePath** | **string** |  | [optional] [default to null]
**SourceProject** | **string** | Project containing inventory file used as source. | [optional] [default to null]
**SourceVars** | **string** | Inventory source variables in YAML or JSON format. | [optional] [default to null]
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] [default to null]
**UpdateCacheTimeout** | **int32** |  | [optional] [default to null]
**UpdateOnLaunch** | **bool** |  | [optional] [default to null]
**UpdateOnProjectUpdate** | **bool** |  | [optional] [default to null]
**Verbosity** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


