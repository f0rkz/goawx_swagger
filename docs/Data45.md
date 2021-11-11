# Data45

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSimultaneous** | **bool** |  | [optional] [default to null]
**AskCredentialOnLaunch** | **bool** |  | [optional] [default to null]
**AskDiffModeOnLaunch** | **bool** |  | [optional] [default to null]
**AskInventoryOnLaunch** | **bool** |  | [optional] [default to null]
**AskJobTypeOnLaunch** | **bool** |  | [optional] [default to null]
**AskLimitOnLaunch** | **bool** |  | [optional] [default to null]
**AskScmBranchOnLaunch** | **bool** |  | [optional] [default to null]
**AskSkipTagsOnLaunch** | **bool** |  | [optional] [default to null]
**AskTagsOnLaunch** | **bool** |  | [optional] [default to null]
**AskVariablesOnLaunch** | **bool** |  | [optional] [default to null]
**AskVerbosityOnLaunch** | **bool** |  | [optional] [default to null]
**BecomeEnabled** | **bool** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DiffMode** | **bool** | If enabled, textual changes made to any templated files on the host are shown in the standard output | [optional] [default to null]
**ExecutionEnvironment** | **int32** | The container image to be used for execution. | [optional] [default to null]
**ExtraVars** | **string** |  | [optional] [default to null]
**ForceHandlers** | **bool** |  | [optional] [default to null]
**Forks** | **int32** |  | [optional] [default to null]
**HostConfigKey** | **string** |  | [optional] [default to null]
**Inventory** | **int32** |  | [optional] [default to null]
**JobSliceCount** | **int32** | The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. | [optional] [default to null]
**JobTags** | **string** |  | [optional] [default to null]
**JobType** | **string** |  | [optional] [default to null]
**Limit** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Playbook** | **string** |  | [optional] [default to null]
**Project** | **string** |  | [optional] [default to null]
**ScmBranch** | **string** | Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. | [optional] [default to null]
**SkipTags** | **string** |  | [optional] [default to null]
**StartAtTask** | **string** |  | [optional] [default to null]
**SurveyEnabled** | **bool** |  | [optional] [default to null]
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] [default to null]
**UseFactCache** | **bool** | If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. | [optional] [default to null]
**Verbosity** | **string** |  | [optional] [default to null]
**WebhookCredential** | **int32** | Personal Access Token for posting back the status to the service API | [optional] [default to null]
**WebhookService** | **string** | Service that webhook requests will be accepted from | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


