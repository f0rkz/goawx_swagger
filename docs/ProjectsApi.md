# \ProjectsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectsAccessListList**](ProjectsApi.md#ProjectsProjectsAccessListList) | **Get** /api/v2/projects/{id}/access_list/ |  List Users
[**ProjectsProjectsActivityStreamList**](ProjectsApi.md#ProjectsProjectsActivityStreamList) | **Get** /api/v2/projects/{id}/activity_stream/ |  List Activity Streams for a Project
[**ProjectsProjectsCopyCreate**](ProjectsApi.md#ProjectsProjectsCopyCreate) | **Post** /api/v2/projects/{id}/copy/ | 
[**ProjectsProjectsCopyList**](ProjectsApi.md#ProjectsProjectsCopyList) | **Get** /api/v2/projects/{id}/copy/ | 
[**ProjectsProjectsCreate**](ProjectsApi.md#ProjectsProjectsCreate) | **Post** /api/v2/projects/ |  Create a Project
[**ProjectsProjectsDelete**](ProjectsApi.md#ProjectsProjectsDelete) | **Delete** /api/v2/projects/{id}/ |  Delete a Project
[**ProjectsProjectsInventoriesRead**](ProjectsApi.md#ProjectsProjectsInventoriesRead) | **Get** /api/v2/projects/{id}/inventories/ |  Retrieve a Project
[**ProjectsProjectsList**](ProjectsApi.md#ProjectsProjectsList) | **Get** /api/v2/projects/ |  List Projects
[**ProjectsProjectsNotificationTemplatesErrorCreate**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesErrorCreate) | **Post** /api/v2/projects/{id}/notification_templates_error/ |  Create a Notification Template for a Project
[**ProjectsProjectsNotificationTemplatesErrorList**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesErrorList) | **Get** /api/v2/projects/{id}/notification_templates_error/ |  List Notification Templates for a Project
[**ProjectsProjectsNotificationTemplatesStartedCreate**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesStartedCreate) | **Post** /api/v2/projects/{id}/notification_templates_started/ |  Create a Notification Template for a Project
[**ProjectsProjectsNotificationTemplatesStartedList**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesStartedList) | **Get** /api/v2/projects/{id}/notification_templates_started/ |  List Notification Templates for a Project
[**ProjectsProjectsNotificationTemplatesSuccessCreate**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesSuccessCreate) | **Post** /api/v2/projects/{id}/notification_templates_success/ |  Create a Notification Template for a Project
[**ProjectsProjectsNotificationTemplatesSuccessList**](ProjectsApi.md#ProjectsProjectsNotificationTemplatesSuccessList) | **Get** /api/v2/projects/{id}/notification_templates_success/ |  List Notification Templates for a Project
[**ProjectsProjectsObjectRolesList**](ProjectsApi.md#ProjectsProjectsObjectRolesList) | **Get** /api/v2/projects/{id}/object_roles/ |  List Roles for a Project
[**ProjectsProjectsPartialUpdate**](ProjectsApi.md#ProjectsProjectsPartialUpdate) | **Patch** /api/v2/projects/{id}/ |  Update a Project
[**ProjectsProjectsPlaybooksRead**](ProjectsApi.md#ProjectsProjectsPlaybooksRead) | **Get** /api/v2/projects/{id}/playbooks/ |  Retrieve Project Playbooks
[**ProjectsProjectsProjectUpdatesList**](ProjectsApi.md#ProjectsProjectsProjectUpdatesList) | **Get** /api/v2/projects/{id}/project_updates/ |  List Project Updates for a Project
[**ProjectsProjectsRead**](ProjectsApi.md#ProjectsProjectsRead) | **Get** /api/v2/projects/{id}/ |  Retrieve a Project
[**ProjectsProjectsSchedulesCreate**](ProjectsApi.md#ProjectsProjectsSchedulesCreate) | **Post** /api/v2/projects/{id}/schedules/ |  Create a Schedule for a Project
[**ProjectsProjectsSchedulesList**](ProjectsApi.md#ProjectsProjectsSchedulesList) | **Get** /api/v2/projects/{id}/schedules/ |  List Schedules for a Project
[**ProjectsProjectsScmInventorySourcesList**](ProjectsApi.md#ProjectsProjectsScmInventorySourcesList) | **Get** /api/v2/projects/{id}/scm_inventory_sources/ |  List Inventory Sources for a Project
[**ProjectsProjectsTeamsList**](ProjectsApi.md#ProjectsProjectsTeamsList) | **Get** /api/v2/projects/{id}/teams/ |  List Teams
[**ProjectsProjectsUpdate0**](ProjectsApi.md#ProjectsProjectsUpdate0) | **Put** /api/v2/projects/{id}/ |  Update a Project
[**ProjectsProjectsUpdateCreate**](ProjectsApi.md#ProjectsProjectsUpdateCreate) | **Post** /api/v2/projects/{id}/update/ |  Update Project
[**ProjectsProjectsUpdateRead**](ProjectsApi.md#ProjectsProjectsUpdateRead) | **Get** /api/v2/projects/{id}/update/ |  Update Project


# **ProjectsProjectsAccessListList**
> ProjectsProjectsAccessListList(ctx, id, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsAccessListListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsAccessListListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsActivityStreamList**
> ProjectsProjectsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Project

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsActivityStreamListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsCopyCreate**
> ProjectsProjectsCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data53**](Data53.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsCopyList**
> ProjectsProjectsCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsCopyListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsCreate**
> ProjectsProjectsCreate(ctx, optional)
 Create a Project

 Make a POST request to this resource with the following project fields to create a new project:          * `name`: Name of this project. (string, required) * `description`: Optional description of this project. (string, default=`\"\"`) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string, default=`\"\"`) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual (default)     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string, default=`\"\"`) * `scm_branch`: Specific branch, tag or commit to checkout. (string, default=`\"\"`) * `scm_refspec`: For git projects, an additional refspec to fetch. (string, default=`\"\"`) * `scm_clean`: Discard any local changes before syncing the project. (boolean, default=`False`) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean, default=`False`) * `scm_delete_on_update`: Delete the project before syncing. (boolean, default=`False`) * `credential`:  (id, default=``) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)      * `organization`: The organization used to determine access to this template. (id, default=``) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean, default=`False`) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer, default=`0`) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean, default=`False`)  * `default_environment`: The default execution environment for jobs run using this project. (id, default=``)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiProjectsProjectsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsDelete**
> ProjectsProjectsDelete(ctx, id, optional)
 Delete a Project

 Make a DELETE request to this resource to delete this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsInventoriesRead**
> ProjectsProjectsInventoriesRead(ctx, id, optional)
 Retrieve a Project

 Make GET request to this resource to retrieve a single project record containing the following fields:  * `inventory_files`: Array of inventory files and directories available within this project, not comprehensive. (json)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsInventoriesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsInventoriesReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsList**
> ProjectsProjectsList(ctx, optional)
 List Projects

 Make a GET request to this resource to retrieve the list of projects.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of projects found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project records.    ## Results  Each project data structure includes the following fields:  * `id`: Database ID for this project. (integer) * `type`: Data type for this project. (choice) * `url`: URL for this project. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project was created. (datetime) * `modified`: Timestamp when this project was last modified. (datetime) * `name`: Name of this project. (string) * `description`: Optional description of this project. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The last revision fetched by a project update (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing * `organization`: The organization used to determine access to this template. (id) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run using this project. (id) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that projects are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiProjectsProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesErrorCreate**
> ProjectsProjectsNotificationTemplatesErrorCreate(ctx, id, optional)
 Create a Notification Template for a Project

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this project.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a Project:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this project.  # Remove Notification Templates from this Project:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this project  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesErrorCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesErrorCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data54**](Data54.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesErrorList**
> ProjectsProjectsNotificationTemplatesErrorList(ctx, id, optional)
 List Notification Templates for a Project

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesErrorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesErrorListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesStartedCreate**
> ProjectsProjectsNotificationTemplatesStartedCreate(ctx, id, optional)
 Create a Notification Template for a Project

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this project.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a Project:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this project.  # Remove Notification Templates from this Project:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this project  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesStartedCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesStartedCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesStartedList**
> ProjectsProjectsNotificationTemplatesStartedList(ctx, id, optional)
 List Notification Templates for a Project

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesStartedListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesStartedListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesSuccessCreate**
> ProjectsProjectsNotificationTemplatesSuccessCreate(ctx, id, optional)
 Create a Notification Template for a Project

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this project.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a Project:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this project.  # Remove Notification Templates from this Project:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this project  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesSuccessCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesSuccessCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data55**](Data55.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsNotificationTemplatesSuccessList**
> ProjectsProjectsNotificationTemplatesSuccessList(ctx, id, optional)
 List Notification Templates for a Project

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsNotificationTemplatesSuccessListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsNotificationTemplatesSuccessListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsObjectRolesList**
> ProjectsProjectsObjectRolesList(ctx, id, optional)
 List Roles for a Project

 Make a GET request to this resource to retrieve a list of roles associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsObjectRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsObjectRolesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsPartialUpdate**
> ProjectsProjectsPartialUpdate(ctx, id, optional)
 Update a Project

 Make a PUT or PATCH request to this resource to update this project.  The following fields may be modified:          * `name`: Name of this project. (string, required) * `description`: Optional description of this project. (string, default=`\"\"`) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string, default=`\"\"`) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual (default)     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string, default=`\"\"`) * `scm_branch`: Specific branch, tag or commit to checkout. (string, default=`\"\"`) * `scm_refspec`: For git projects, an additional refspec to fetch. (string, default=`\"\"`) * `scm_clean`: Discard any local changes before syncing the project. (boolean, default=`False`) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean, default=`False`) * `scm_delete_on_update`: Delete the project before syncing. (boolean, default=`False`) * `credential`:  (id, default=``) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)      * `organization`: The organization used to determine access to this template. (id, default=``) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean, default=`False`) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer, default=`0`) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean, default=`False`)  * `default_environment`: The default execution environment for jobs run using this project. (id, default=``)           For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsPlaybooksRead**
> ProjectsProjectsPlaybooksRead(ctx, id, optional)
 Retrieve Project Playbooks

 Make GET request to this resource to retrieve a list of playbooks available for a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsPlaybooksReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsPlaybooksReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsProjectUpdatesList**
> ProjectsProjectsProjectUpdatesList(ctx, id, optional)
 List Project Updates for a Project

 Make a GET request to this resource to retrieve a list of project updates associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of project updates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project update records.    ## Results  Each project update data structure includes the following fields:  * `id`: Database ID for this project update. (integer) * `type`: Data type for this project update. (choice) * `url`: URL for this project update. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project update was created. (datetime) * `modified`: Timestamp when this project update was last modified. (datetime) * `name`: Name of this project update. (string) * `description`: Optional description of this project update. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The SCM Revision discovered by this update for the given project and branch. (string) * `project`:  (id) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `job_tags`: Parts of the project update playbook that will be run. (string)    ## Sorting  To specify that project updates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsProjectUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsProjectUpdatesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsRead**
> ProjectsProjectsRead(ctx, id, optional)
 Retrieve a Project

 Make GET request to this resource to retrieve a single project record containing the following fields:  * `id`: Database ID for this project. (integer) * `type`: Data type for this project. (choice) * `url`: URL for this project. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project was created. (datetime) * `modified`: Timestamp when this project was last modified. (datetime) * `name`: Name of this project. (string) * `description`: Optional description of this project. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The last revision fetched by a project update (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing * `organization`: The organization used to determine access to this template. (id) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run using this project. (id) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsSchedulesCreate**
> ProjectsProjectsSchedulesCreate(ctx, id, optional)
 Create a Schedule for a Project

 Make a POST request to this resource with the following schedule fields to create a new schedule associated with this project.   * `rrule`: A value representing the schedules iCal recurrence rule. (string, required)        * `name`: Name of this schedule. (string, required) * `description`: Optional description of this schedule. (string, default=`\"\"`) * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug)  * `enabled`: Enables processing of this schedule. (boolean, default=`True`)            POST requests to this resource must include a proper `rrule` value following a particular format and conforming to subset of allowed rules.  The following lists the expected format and details of our rrules:  * DTSTART is required and must follow the following format: DTSTART:YYYYMMDDTHHMMSSZ * DTSTART is expected to be in UTC * INTERVAL is required * SECONDLY is not supported * RRULE must precede the rule statements * BYDAY is supported but not BYDAY with a numerical prefix * BYYEARDAY and BYWEEKNO are not supported * Only one rrule statement per schedule is supported * COUNT must be < 1000  Here are some example rrules:      \"DTSTART:20500331T055000Z RRULE:FREQ=MINUTELY;INTERVAL=10;COUNT=5\"     \"DTSTART:20240331T075000Z RRULE:FREQ=DAILY;INTERVAL=1;COUNT=1\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MINUTELY;INTERVAL=1;UNTIL=20230401T075000Z\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=5;BYDAY=MO\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYMONTHDAY=6\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=4;BYDAY=SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=-1;BYDAY=MO,TU,WE,TH,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=-1;BYDAY=MO,TU,WE,TH,FR,SA,SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=YEARLY;INTERVAL=1;BYMONTH=4;BYMONTHDAY=1\"     \"DTSTART:20140331T075000Z RRULE:FREQ=YEARLY;INTERVAL=1;BYSETPOS=-1;BYMONTH=8;BYDAY=SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=1;UNTIL=20230401T075000Z;BYDAY=MO,WE,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=HOURLY;INTERVAL=1;UNTIL=20230610T075000Z\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsSchedulesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsSchedulesCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data56**](Data56.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsSchedulesList**
> ProjectsProjectsSchedulesList(ctx, id, optional)
 List Schedules for a Project

 Make a GET request to this resource to retrieve a list of schedules associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of schedules found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more schedule records.    ## Results  Each schedule data structure includes the following fields:  * `rrule`: A value representing the schedules iCal recurrence rule. (string) * `id`: Database ID for this schedule. (integer) * `type`: Data type for this schedule. (choice) * `url`: URL for this schedule. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this schedule was created. (datetime) * `modified`: Timestamp when this schedule was last modified. (datetime) * `name`: Name of this schedule. (string) * `description`: Optional description of this schedule. (string) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `unified_job_template`:  (id) * `enabled`: Enables processing of this schedule. (boolean) * `dtstart`: The first occurrence of the schedule occurs on or after this time. (datetime) * `dtend`: The last occurrence of the schedule occurs before this time, aftewards the schedule expires. (datetime) * `next_run`: The next time that the scheduled action will run. (datetime) * `timezone`:  (field) * `until`:  (field)    ## Sorting  To specify that schedules are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsSchedulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsSchedulesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsScmInventorySourcesList**
> ProjectsProjectsScmInventorySourcesList(ctx, id, optional)
 List Inventory Sources for a Project

 Make a GET request to this resource to retrieve a list of inventory sources associated with the selected project.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory source records.    ## Results  Each inventory source data structure includes the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that inventory sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsScmInventorySourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsScmInventorySourcesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsTeamsList**
> ProjectsProjectsTeamsList(ctx, id, optional)
 List Teams

 Make a GET request to this resource to retrieve the list of teams.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of teams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more team records.    ## Results  Each team data structure includes the following fields:  * `id`: Database ID for this team. (integer) * `type`: Data type for this team. (choice) * `url`: URL for this team. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this team was created. (datetime) * `modified`: Timestamp when this team was last modified. (datetime) * `name`: Name of this team. (string) * `description`: Optional description of this team. (string) * `organization`:  (id)    ## Sorting  To specify that teams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsTeamsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsUpdate0**
> ProjectsProjectsUpdate0(ctx, id, optional)
 Update a Project

 Make a PUT or PATCH request to this resource to update this project.  The following fields may be modified:          * `name`: Name of this project. (string, required) * `description`: Optional description of this project. (string, default=`\"\"`) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string, default=`\"\"`) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual (default)     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string, default=`\"\"`) * `scm_branch`: Specific branch, tag or commit to checkout. (string, default=`\"\"`) * `scm_refspec`: For git projects, an additional refspec to fetch. (string, default=`\"\"`) * `scm_clean`: Discard any local changes before syncing the project. (boolean, default=`False`) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean, default=`False`) * `scm_delete_on_update`: Delete the project before syncing. (boolean, default=`False`) * `credential`:  (id, default=``) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)      * `organization`: The organization used to determine access to this template. (id, default=``) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean, default=`False`) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer, default=`0`) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean, default=`False`)  * `default_environment`: The default execution environment for jobs run using this project. (id, default=``)         For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsUpdate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsUpdate0Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data52**](Data52.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsUpdateCreate**
> ProjectsProjectsUpdateCreate(ctx, id)
 Update Project

 Make a GET request to this resource to determine if the project can be updated from its SCM source.  The response will include the following field:  * `can_update`: Flag indicating if this project can be updated (boolean,   read-only)  Make a POST request to this resource to update the project.  If the project cannot be updated, a 405 status code will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectsUpdateRead**
> ProjectsProjectsUpdateRead(ctx, id, optional)
 Update Project

 Make a GET request to this resource to determine if the project can be updated from its SCM source.  The response will include the following field:  * `can_update`: Flag indicating if this project can be updated (boolean,   read-only)  Make a POST request to this resource to update the project.  If the project cannot be updated, a 405 status code will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectsApiProjectsProjectsUpdateReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsProjectsUpdateReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

