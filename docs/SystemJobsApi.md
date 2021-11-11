# \SystemJobsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemJobsSystemJobsCancelCreate**](SystemJobsApi.md#SystemJobsSystemJobsCancelCreate) | **Post** /api/v2/system_jobs/{id}/cancel/ |  Retrieve a System Job
[**SystemJobsSystemJobsCancelRead**](SystemJobsApi.md#SystemJobsSystemJobsCancelRead) | **Get** /api/v2/system_jobs/{id}/cancel/ |  Retrieve a System Job
[**SystemJobsSystemJobsDelete**](SystemJobsApi.md#SystemJobsSystemJobsDelete) | **Delete** /api/v2/system_jobs/{id}/ |  Delete a System Job
[**SystemJobsSystemJobsEventsList**](SystemJobsApi.md#SystemJobsSystemJobsEventsList) | **Get** /api/v2/system_jobs/{id}/events/ |  List System Job Events for a System Job
[**SystemJobsSystemJobsList**](SystemJobsApi.md#SystemJobsSystemJobsList) | **Get** /api/v2/system_jobs/ |  List System Jobs
[**SystemJobsSystemJobsNotificationsList**](SystemJobsApi.md#SystemJobsSystemJobsNotificationsList) | **Get** /api/v2/system_jobs/{id}/notifications/ |  List Notifications for a System Job
[**SystemJobsSystemJobsRead**](SystemJobsApi.md#SystemJobsSystemJobsRead) | **Get** /api/v2/system_jobs/{id}/ |  Retrieve a System Job


# **SystemJobsSystemJobsCancelCreate**
> SystemJobsSystemJobsCancelCreate(ctx, id)
 Retrieve a System Job

 Make GET request to this resource to retrieve a single system job record containing the following fields:  * `can_cancel`:  (boolean)

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

# **SystemJobsSystemJobsCancelRead**
> SystemJobsSystemJobsCancelRead(ctx, id, optional)
 Retrieve a System Job

 Make GET request to this resource to retrieve a single system job record containing the following fields:  * `can_cancel`:  (boolean)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobsApiSystemJobsSystemJobsCancelReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsCancelReadOpts struct

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

# **SystemJobsSystemJobsDelete**
> SystemJobsSystemJobsDelete(ctx, id, optional)
 Delete a System Job

 Make a DELETE request to this resource to delete this system job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobsApiSystemJobsSystemJobsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsDeleteOpts struct

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

# **SystemJobsSystemJobsEventsList**
> SystemJobsSystemJobsEventsList(ctx, id, optional)
 List System Job Events for a System Job

 Make a GET request to this resource to retrieve a list of system job events associated with the selected system job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of system job events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more system job event records.    ## Results  Each system job event data structure includes the following fields:  * `id`: Database ID for this system job event. (integer) * `type`: Data type for this system job event. (choice) * `url`: URL for this system job event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job event was created. (datetime) * `modified`: Timestamp when this system job event was last modified. (datetime) * `event`:  (field) * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (field) * `changed`:  (field) * `uuid`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer) * `system_job`:  (id)    ## Sorting  To specify that system job events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobsApiSystemJobsSystemJobsEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsEventsListOpts struct

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

# **SystemJobsSystemJobsList**
> SystemJobsSystemJobsList(ctx, optional)
 List System Jobs

 Make a GET request to this resource to retrieve the list of system jobs.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of system jobs found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more system job records.    ## Results  Each system job data structure includes the following fields:  * `id`: Database ID for this system job. (integer) * `type`: Data type for this system job. (choice) * `url`: URL for this system job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job was created. (datetime) * `modified`: Timestamp when this system job was last modified. (datetime) * `name`: Name of this system job. (string) * `description`: Optional description of this system job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `system_job_template`:  (id) * `job_type`:  (choice)     - `\"\"`: ---------     - `cleanup_jobs`: Remove jobs older than a certain number of days     - `cleanup_activitystream`: Remove activity stream entries older than a certain number of days     - `cleanup_sessions`: Removes expired browser sessions from the database     - `cleanup_tokens`: Removes expired OAuth 2 access tokens and refresh tokens * `extra_vars`:  (string) * `result_stdout`:  (field)    ## Sorting  To specify that system jobs are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemJobsApiSystemJobsSystemJobsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsListOpts struct

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

# **SystemJobsSystemJobsNotificationsList**
> SystemJobsSystemJobsNotificationsList(ctx, id, optional)
 List Notifications for a System Job

 Make a GET request to this resource to retrieve a list of notifications associated with the selected system job.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notifications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * `id`: Database ID for this notification. (integer) * `type`: Data type for this notification. (choice) * `url`: URL for this notification. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification was created. (datetime) * `modified`: Timestamp when this notification was last modified. (datetime) * `notification_template`:  (id) * `error`:  (string) * `status`:  (choice)     - `pending`: Pending     - `successful`: Successful     - `failed`: Failed * `notifications_sent`:  (integer) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `recipients`:  (string) * `subject`:  (string) * `body`: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobsApiSystemJobsSystemJobsNotificationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsNotificationsListOpts struct

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

# **SystemJobsSystemJobsRead**
> SystemJobsSystemJobsRead(ctx, id, optional)
 Retrieve a System Job

 Make GET request to this resource to retrieve a single system job record containing the following fields:  * `id`: Database ID for this system job. (integer) * `type`: Data type for this system job. (choice) * `url`: URL for this system job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job was created. (datetime) * `modified`: Timestamp when this system job was last modified. (datetime) * `name`: Name of this system job. (string) * `description`: Optional description of this system job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_args`:  (string) * `job_cwd`:  (string) * `job_env`:  (json) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `result_traceback`:  (string) * `event_processing_finished`: Indicates whether all of the events generated by this unified job have been saved to the database. (boolean) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `system_job_template`:  (id) * `job_type`:  (choice)     - `\"\"`: ---------     - `cleanup_jobs`: Remove jobs older than a certain number of days     - `cleanup_activitystream`: Remove activity stream entries older than a certain number of days     - `cleanup_sessions`: Removes expired browser sessions from the database     - `cleanup_tokens`: Removes expired OAuth 2 access tokens and refresh tokens * `extra_vars`:  (string) * `result_stdout`:  (field)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobsApiSystemJobsSystemJobsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobsApiSystemJobsSystemJobsReadOpts struct

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

