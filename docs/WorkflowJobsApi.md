# \WorkflowJobsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowJobsWorkflowJobsActivityStreamList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsActivityStreamList) | **Get** /api/v2/workflow_jobs/{id}/activity_stream/ |  List Activity Streams for a Workflow Job
[**WorkflowJobsWorkflowJobsCancelCreate**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsCancelCreate) | **Post** /api/v2/workflow_jobs/{id}/cancel/ |  Cancel Workflow Job
[**WorkflowJobsWorkflowJobsCancelRead**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsCancelRead) | **Get** /api/v2/workflow_jobs/{id}/cancel/ |  Cancel Workflow Job
[**WorkflowJobsWorkflowJobsDelete**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsDelete) | **Delete** /api/v2/workflow_jobs/{id}/ |  Delete a Workflow Job
[**WorkflowJobsWorkflowJobsLabelsList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsLabelsList) | **Get** /api/v2/workflow_jobs/{id}/labels/ |  List Labels for a Workflow Job
[**WorkflowJobsWorkflowJobsList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsList) | **Get** /api/v2/workflow_jobs/ |  List Workflow Jobs
[**WorkflowJobsWorkflowJobsNotificationsList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsNotificationsList) | **Get** /api/v2/workflow_jobs/{id}/notifications/ |  List Notifications for a Workflow Job
[**WorkflowJobsWorkflowJobsRead**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsRead) | **Get** /api/v2/workflow_jobs/{id}/ |  Retrieve a Workflow Job
[**WorkflowJobsWorkflowJobsRelaunchCreate**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsRelaunchCreate) | **Post** /api/v2/workflow_jobs/{id}/relaunch/ | Relaunch a workflow job
[**WorkflowJobsWorkflowJobsRelaunchList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsRelaunchList) | **Get** /api/v2/workflow_jobs/{id}/relaunch/ | Relaunch a workflow job
[**WorkflowJobsWorkflowJobsWorkflowNodesList**](WorkflowJobsApi.md#WorkflowJobsWorkflowJobsWorkflowNodesList) | **Get** /api/v2/workflow_jobs/{id}/workflow_nodes/ |  List Workflow Job Nodes for a Workflow Job


# **WorkflowJobsWorkflowJobsActivityStreamList**
> WorkflowJobsWorkflowJobsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Workflow Job

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected workflow job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsActivityStreamListOpts struct

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

# **WorkflowJobsWorkflowJobsCancelCreate**
> WorkflowJobsWorkflowJobsCancelCreate(ctx, id)
 Cancel Workflow Job

 Make a GET request to this resource to determine if the workflow job can be canceled. The response will include the following field:  * `can_cancel`: Indicates whether this workflow job is in a state that can   be canceled (boolean, read-only)  Make a POST request to this endpoint to submit a request to cancel a pending or running workflow job.  The response status code will be 202 if the request to cancel was successfully submitted, or 405 if the workflow job cannot be canceled.

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

# **WorkflowJobsWorkflowJobsCancelRead**
> WorkflowJobsWorkflowJobsCancelRead(ctx, id, optional)
 Cancel Workflow Job

 Make a GET request to this resource to determine if the workflow job can be canceled. The response will include the following field:  * `can_cancel`: Indicates whether this workflow job is in a state that can   be canceled (boolean, read-only)  Make a POST request to this endpoint to submit a request to cancel a pending or running workflow job.  The response status code will be 202 if the request to cancel was successfully submitted, or 405 if the workflow job cannot be canceled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsCancelReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsCancelReadOpts struct

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

# **WorkflowJobsWorkflowJobsDelete**
> WorkflowJobsWorkflowJobsDelete(ctx, id, optional)
 Delete a Workflow Job

 Make a DELETE request to this resource to delete this workflow job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsDeleteOpts struct

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

# **WorkflowJobsWorkflowJobsLabelsList**
> WorkflowJobsWorkflowJobsLabelsList(ctx, id, optional)
 List Labels for a Workflow Job

 Make a GET request to this resource to retrieve a list of labels associated with the selected workflow job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of labels found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more label records.    ## Results  Each label data structure includes the following fields:  * `id`: Database ID for this label. (integer) * `type`: Data type for this label. (choice) * `url`: URL for this label. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this label was created. (datetime) * `modified`: Timestamp when this label was last modified. (datetime) * `name`: Name of this label. (string) * `organization`: Organization this label belongs to. (id)    ## Sorting  To specify that labels are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsLabelsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsLabelsListOpts struct

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

# **WorkflowJobsWorkflowJobsList**
> WorkflowJobsWorkflowJobsList(ctx, optional)
 List Workflow Jobs

 Make a GET request to this resource to retrieve the list of workflow jobs.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow jobs found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job records.    ## Results  Each workflow job data structure includes the following fields:  * `id`: Database ID for this workflow job. (integer) * `type`: Data type for this workflow job. (choice) * `url`: URL for this workflow job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job was created. (datetime) * `modified`: Timestamp when this workflow job was last modified. (datetime) * `name`: Name of this workflow job. (string) * `description`: Optional description of this workflow job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `workflow_job_template`:  (id) * `extra_vars`:  (json) * `allow_simultaneous`:  (boolean) * `job_template`: If automatically created for a sliced job run, the job template the workflow job was created from. (id) * `is_sliced_job`:  (boolean) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `limit`:  (string) * `scm_branch`:  (string) * `webhook_service`: Service that webhook requests will be accepted from (choice)     - `\"\"`: ---------     - `github`: GitHub     - `gitlab`: GitLab * `webhook_credential`: Personal Access Token for posting back the status to the service API (id) * `webhook_guid`: Unique identifier of the event that triggered this webhook (string)    ## Sorting  To specify that workflow jobs are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsListOpts struct

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

# **WorkflowJobsWorkflowJobsNotificationsList**
> WorkflowJobsWorkflowJobsNotificationsList(ctx, id, optional)
 List Notifications for a Workflow Job

 Make a GET request to this resource to retrieve a list of notifications associated with the selected workflow job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notifications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * `id`: Database ID for this notification. (integer) * `type`: Data type for this notification. (choice) * `url`: URL for this notification. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification was created. (datetime) * `modified`: Timestamp when this notification was last modified. (datetime) * `notification_template`:  (id) * `error`:  (string) * `status`:  (choice)     - `pending`: Pending     - `successful`: Successful     - `failed`: Failed * `notifications_sent`:  (integer) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `recipients`:  (string) * `subject`:  (string) * `body`: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsNotificationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsNotificationsListOpts struct

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

# **WorkflowJobsWorkflowJobsRead**
> WorkflowJobsWorkflowJobsRead(ctx, id, optional)
 Retrieve a Workflow Job

 Make GET request to this resource to retrieve a single workflow job record containing the following fields:  * `id`: Database ID for this workflow job. (integer) * `type`: Data type for this workflow job. (choice) * `url`: URL for this workflow job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job was created. (datetime) * `modified`: Timestamp when this workflow job was last modified. (datetime) * `name`: Name of this workflow job. (string) * `description`: Optional description of this workflow job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_args`:  (string) * `job_cwd`:  (string) * `job_env`:  (json) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `result_traceback`:  (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `workflow_job_template`:  (id) * `extra_vars`:  (json) * `allow_simultaneous`:  (boolean) * `job_template`: If automatically created for a sliced job run, the job template the workflow job was created from. (id) * `is_sliced_job`:  (boolean) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `limit`:  (string) * `scm_branch`:  (string) * `webhook_service`: Service that webhook requests will be accepted from (choice)     - `\"\"`: ---------     - `github`: GitHub     - `gitlab`: GitLab * `webhook_credential`: Personal Access Token for posting back the status to the service API (id) * `webhook_guid`: Unique identifier of the event that triggered this webhook (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsReadOpts struct

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

# **WorkflowJobsWorkflowJobsRelaunchCreate**
> WorkflowJobsWorkflowJobsRelaunchCreate(ctx, id)
Relaunch a workflow job

 Make a POST request to this endpoint to launch a workflow job identical to the parent workflow job. This will spawn jobs, project updates, or inventory updates based on the unified job templates referenced in the workflow nodes in the workflow job. No POST data is accepted for this action.  If successful, the response status code will be 201 and serialized data of the new workflow job will be returned.

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

# **WorkflowJobsWorkflowJobsRelaunchList**
> WorkflowJobsWorkflowJobsRelaunchList(ctx, id, optional)
Relaunch a workflow job

 Make a POST request to this endpoint to launch a workflow job identical to the parent workflow job. This will spawn jobs, project updates, or inventory updates based on the unified job templates referenced in the workflow nodes in the workflow job. No POST data is accepted for this action.  If successful, the response status code will be 201 and serialized data of the new workflow job will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsRelaunchListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsRelaunchListOpts struct

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

# **WorkflowJobsWorkflowJobsWorkflowNodesList**
> WorkflowJobsWorkflowJobsWorkflowNodesList(ctx, id, optional)
 List Workflow Job Nodes for a Workflow Job

 Make a GET request to this resource to retrieve a list of workflow job nodes associated with the selected workflow job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job node records.    ## Results  Each workflow job node data structure includes the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)    ## Sorting  To specify that workflow job nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobsApiWorkflowJobsWorkflowJobsWorkflowNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobsApiWorkflowJobsWorkflowJobsWorkflowNodesListOpts struct

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

