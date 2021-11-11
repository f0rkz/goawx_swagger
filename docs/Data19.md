# Data19

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | **int32** |  | [optional] [default to null]
**IsContainerGroup** | **bool** | Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. | [optional] [default to null]
**Name** | **string** |  | [default to null]
**PodSpecOverride** | **string** |  | [optional] [default to null]
**PolicyInstanceList** | **[]string** | List of exact-match Instances that will be assigned to this group | [optional] [default to null]
**PolicyInstanceMinimum** | **int32** | Static minimum number of Instances that will be automatically assign to this group when new instances come online. | [optional] [default to null]
**PolicyInstancePercentage** | **int32** | Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


