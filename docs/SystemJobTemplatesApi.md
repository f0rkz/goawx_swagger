# \SystemJobTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemJobTemplatesSystemJobTemplatesJobsList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesJobsList) | **Get** /api/v2/system_job_templates/{id}/jobs/ |  List System Jobs for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesLaunchCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesLaunchCreate) | **Post** /api/v2/system_job_templates/{id}/launch/ | Launch a Job Template
[**SystemJobTemplatesSystemJobTemplatesLaunchList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesLaunchList) | **Get** /api/v2/system_job_templates/{id}/launch/ | Launch a Job Template
[**SystemJobTemplatesSystemJobTemplatesList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesList) | **Get** /api/v2/system_job_templates/ |  List System Job Templates
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_error/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_error/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_started/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_started/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate) | **Post** /api/v2/system_job_templates/{id}/notification_templates_success/ |  Create a Notification Template for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList) | **Get** /api/v2/system_job_templates/{id}/notification_templates_success/ |  List Notification Templates for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesRead**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesRead) | **Get** /api/v2/system_job_templates/{id}/ |  Retrieve a System Job Template
[**SystemJobTemplatesSystemJobTemplatesSchedulesCreate**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesSchedulesCreate) | **Post** /api/v2/system_job_templates/{id}/schedules/ |  Create a Schedule for a System Job Template
[**SystemJobTemplatesSystemJobTemplatesSchedulesList**](SystemJobTemplatesApi.md#SystemJobTemplatesSystemJobTemplatesSchedulesList) | **Get** /api/v2/system_job_templates/{id}/schedules/ |  List Schedules for a System Job Template


# **SystemJobTemplatesSystemJobTemplatesJobsList**
> SystemJobTemplatesSystemJobTemplatesJobsList(ctx, id, optional)
 List System Jobs for a System Job Template

 Make a GET request to this resource to retrieve a list of system jobs associated with the selected system job template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of system jobs found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more system job records.    ## Results  Each system job data structure includes the following fields:  * `id`: Database ID for this system job. (integer) * `type`: Data type for this system job. (choice) * `url`: URL for this system job. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job was created. (datetime) * `modified`: Timestamp when this system job was last modified. (datetime) * `name`: Name of this system job. (string) * `description`: Optional description of this system job. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `system_job_template`:  (id) * `job_type`:  (choice)     - `\"\"`: ---------     - `cleanup_jobs`: Remove jobs older than a certain number of days     - `cleanup_activitystream`: Remove activity stream entries older than a certain number of days     - `cleanup_sessions`: Removes expired browser sessions from the database     - `cleanup_tokens`: Removes expired OAuth 2 access tokens and refresh tokens * `extra_vars`:  (string) * `result_stdout`:  (field)    ## Sorting  To specify that system jobs are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesJobsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesJobsListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesLaunchCreate**
> SystemJobTemplatesSystemJobTemplatesLaunchCreate(ctx, id)
Launch a Job Template

 Make a POST request to this resource to launch the system job template.  Variables specified inside of the parameter `extra_vars` are passed to the system job task as command line parameters. These tasks can be run manually on the host system via the `awx-manage` command.  For example on `cleanup_jobs` and `cleanup_activitystream`:  `{\"extra_vars\": {\"days\": 30}}`  Which will act on data older than 30 days.  For `cleanup_activitystream` and `cleanup_jobs` commands, providing `\"dry_run\": true` inside of `extra_vars` will show items that will be removed without deleting them.  Each individual system job task has its own default values, which are applicable either when running it from the command line or launching its system job template with empty `extra_vars`.   - Defaults for `cleanup_activitystream`: days=90  - Defaults for `cleanup_jobs`: days=90  If successful, the response status code will be 202.  If the job cannot be launched, a 405 status code will be returned.

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

# **SystemJobTemplatesSystemJobTemplatesLaunchList**
> SystemJobTemplatesSystemJobTemplatesLaunchList(ctx, id, optional)
Launch a Job Template

 Make a POST request to this resource to launch the system job template.  Variables specified inside of the parameter `extra_vars` are passed to the system job task as command line parameters. These tasks can be run manually on the host system via the `awx-manage` command.  For example on `cleanup_jobs` and `cleanup_activitystream`:  `{\"extra_vars\": {\"days\": 30}}`  Which will act on data older than 30 days.  For `cleanup_activitystream` and `cleanup_jobs` commands, providing `\"dry_run\": true` inside of `extra_vars` will show items that will be removed without deleting them.  Each individual system job task has its own default values, which are applicable either when running it from the command line or launching its system job template with empty `extra_vars`.   - Defaults for `cleanup_activitystream`: days=90  - Defaults for `cleanup_jobs`: days=90  If successful, the response status code will be 202.  If the job cannot be launched, a 405 status code will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesLaunchListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesLaunchListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesList**
> SystemJobTemplatesSystemJobTemplatesList(ctx, optional)
 List System Job Templates

 Make a GET request to this resource to retrieve the list of system job templates.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of system job templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more system job template records.    ## Results  Each system job template data structure includes the following fields:  * `id`: Database ID for this system job template. (integer) * `type`: Data type for this system job template. (choice) * `url`: URL for this system job template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job template was created. (datetime) * `modified`: Timestamp when this system job template was last modified. (datetime) * `name`: Name of this system job template. (string) * `description`: Optional description of this system job template. (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing     - `none`: No External Source     - `updating`: Updating * `execution_environment`: The container image to be used for execution. (id) * `job_type`:  (choice)     - `\"\"`: ---------     - `cleanup_jobs`: Remove jobs older than a certain number of days     - `cleanup_activitystream`: Remove activity stream entries older than a certain number of days     - `cleanup_sessions`: Removes expired browser sessions from the database     - `cleanup_tokens`: Removes expired OAuth 2 access tokens and refresh tokens    ## Sorting  To specify that system job templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate(ctx, id, optional)
 Create a Notification Template for a System Job Template

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this system job template.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a System Job Template:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this system job template.  # Remove Notification Templates from this System Job Template:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this system job template  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data60**](Data60.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList(ctx, id, optional)
 List Notification Templates for a System Job Template

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected system job template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate(ctx, id, optional)
 Create a Notification Template for a System Job Template

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this system job template.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a System Job Template:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this system job template.  # Remove Notification Templates from this System Job Template:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this system job template  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedList(ctx, id, optional)
 List Notification Templates for a System Job Template

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected system job template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreate(ctx, id, optional)
 Create a Notification Template for a System Job Template

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this system job template.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for a System Job Template:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this system job template.  # Remove Notification Templates from this System Job Template:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this system job template  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data61**](Data61.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList**
> SystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessList(ctx, id, optional)
 List Notification Templates for a System Job Template

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected system job template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesNotificationTemplatesSuccessListOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesRead**
> SystemJobTemplatesSystemJobTemplatesRead(ctx, id, optional)
 Retrieve a System Job Template

 Make GET request to this resource to retrieve a single system job template record containing the following fields:  * `id`: Database ID for this system job template. (integer) * `type`: Data type for this system job template. (choice) * `url`: URL for this system job template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this system job template was created. (datetime) * `modified`: Timestamp when this system job template was last modified. (datetime) * `name`: Name of this system job template. (string) * `description`: Optional description of this system job template. (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing     - `none`: No External Source     - `updating`: Updating * `execution_environment`: The container image to be used for execution. (id) * `job_type`:  (choice)     - `\"\"`: ---------     - `cleanup_jobs`: Remove jobs older than a certain number of days     - `cleanup_activitystream`: Remove activity stream entries older than a certain number of days     - `cleanup_sessions`: Removes expired browser sessions from the database     - `cleanup_tokens`: Removes expired OAuth 2 access tokens and refresh tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesReadOpts struct

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

# **SystemJobTemplatesSystemJobTemplatesSchedulesCreate**
> SystemJobTemplatesSystemJobTemplatesSchedulesCreate(ctx, id, optional)
 Create a Schedule for a System Job Template

 Make a POST request to this resource with the following schedule fields to create a new schedule associated with this system job template.   * `rrule`: A value representing the schedules iCal recurrence rule. (string, required)        * `name`: Name of this schedule. (string, required) * `description`: Optional description of this schedule. (string, default=`\"\"`) * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug)  * `enabled`: Enables processing of this schedule. (boolean, default=`True`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesSchedulesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesSchedulesCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data62**](Data62.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemJobTemplatesSystemJobTemplatesSchedulesList**
> SystemJobTemplatesSystemJobTemplatesSchedulesList(ctx, id, optional)
 List Schedules for a System Job Template

 Make a GET request to this resource to retrieve a list of schedules associated with the selected system job template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of schedules found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more schedule records.    ## Results  Each schedule data structure includes the following fields:  * `rrule`: A value representing the schedules iCal recurrence rule. (string) * `id`: Database ID for this schedule. (integer) * `type`: Data type for this schedule. (choice) * `url`: URL for this schedule. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this schedule was created. (datetime) * `modified`: Timestamp when this schedule was last modified. (datetime) * `name`: Name of this schedule. (string) * `description`: Optional description of this schedule. (string) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `unified_job_template`:  (id) * `enabled`: Enables processing of this schedule. (boolean) * `dtstart`: The first occurrence of the schedule occurs on or after this time. (datetime) * `dtend`: The last occurrence of the schedule occurs before this time, aftewards the schedule expires. (datetime) * `next_run`: The next time that the scheduled action will run. (datetime) * `timezone`:  (field) * `until`:  (field)    ## Sorting  To specify that schedules are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesSchedulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemJobTemplatesApiSystemJobTemplatesSystemJobTemplatesSchedulesListOpts struct

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

