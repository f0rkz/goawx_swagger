# Data52

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOverride** | **bool** | Allow changing the SCM branch or revision in a job template that uses this project. | [optional] [default to null]
**Credential** | **int32** |  | [optional] [default to null]
**DefaultEnvironment** | **int32** | The default execution environment for jobs run using this project. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**LocalPath** | **string** | Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Organization** | **int32** | The organization used to determine access to this template. | [optional] [default to null]
**ScmBranch** | **string** | Specific branch, tag or commit to checkout. | [optional] [default to null]
**ScmClean** | **bool** | Discard any local changes before syncing the project. | [optional] [default to null]
**ScmDeleteOnUpdate** | **bool** | Delete the project before syncing. | [optional] [default to null]
**ScmRefspec** | **string** | For git projects, an additional refspec to fetch. | [optional] [default to null]
**ScmTrackSubmodules** | **bool** | Track submodules latest commits on defined branch. | [optional] [default to null]
**ScmType** | **string** | Specifies the source control system used to store the project. | [optional] [default to null]
**ScmUpdateCacheTimeout** | **int32** | The number of seconds after the last project update ran that a new project update will be launched as a job dependency. | [optional] [default to null]
**ScmUpdateOnLaunch** | **bool** | Update the project when a job is launched that uses the project. | [optional] [default to null]
**ScmUrl** | **string** | The location where the project is stored. | [optional] [default to null]
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


