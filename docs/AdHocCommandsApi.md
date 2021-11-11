# \AdHocCommandsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocCommandsAdHocCommandsActivityStreamList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsActivityStreamList) | **Get** /api/v2/ad_hoc_commands/{id}/activity_stream/ |  List Activity Streams for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCancelCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCancelCreate) | **Post** /api/v2/ad_hoc_commands/{id}/cancel/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCancelRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCancelRead) | **Get** /api/v2/ad_hoc_commands/{id}/cancel/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsCreate) | **Post** /api/v2/ad_hoc_commands/ |  Create an Ad Hoc Command
[**AdHocCommandsAdHocCommandsDelete**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsDelete) | **Delete** /api/v2/ad_hoc_commands/{id}/ |  Delete an Ad Hoc Command
[**AdHocCommandsAdHocCommandsEventsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsEventsList) | **Get** /api/v2/ad_hoc_commands/{id}/events/ |  List Ad Hoc Command Events for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsList) | **Get** /api/v2/ad_hoc_commands/ |  List Ad Hoc Commands
[**AdHocCommandsAdHocCommandsNotificationsList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsNotificationsList) | **Get** /api/v2/ad_hoc_commands/{id}/notifications/ |  List Notifications for an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRead) | **Get** /api/v2/ad_hoc_commands/{id}/ |  Retrieve an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRelaunchCreate**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRelaunchCreate) | **Post** /api/v2/ad_hoc_commands/{id}/relaunch/ | Relaunch an Ad Hoc Command
[**AdHocCommandsAdHocCommandsRelaunchList**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsRelaunchList) | **Get** /api/v2/ad_hoc_commands/{id}/relaunch/ | Relaunch an Ad Hoc Command
[**AdHocCommandsAdHocCommandsStdoutRead**](AdHocCommandsApi.md#AdHocCommandsAdHocCommandsStdoutRead) | **Get** /api/v2/ad_hoc_commands/{id}/stdout/ |  Retrieve Ad Hoc Command Stdout


# **AdHocCommandsAdHocCommandsActivityStreamList**
> AdHocCommandsAdHocCommandsActivityStreamList(ctx, id, optional)
 List Activity Streams for an Ad Hoc Command

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected ad hoc command.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsActivityStreamListOpts struct

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

# **AdHocCommandsAdHocCommandsCancelCreate**
> AdHocCommandsAdHocCommandsCancelCreate(ctx, id)
 Retrieve an Ad Hoc Command

 Make GET request to this resource to retrieve a single ad hoc command record containing the following fields:  * `can_cancel`:  (boolean)

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

# **AdHocCommandsAdHocCommandsCancelRead**
> AdHocCommandsAdHocCommandsCancelRead(ctx, id, optional)
 Retrieve an Ad Hoc Command

 Make GET request to this resource to retrieve a single ad hoc command record containing the following fields:  * `can_cancel`:  (boolean)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsCancelReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsCancelReadOpts struct

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

# **AdHocCommandsAdHocCommandsCreate**
> AdHocCommandsAdHocCommandsCreate(ctx, optional)
 Create an Ad Hoc Command

 Make a POST request to this resource with the following ad hoc command fields to create a new ad hoc command:             * `execution_environment`: The container image to be used for execution. (id, default=``)           * `job_type`:  (choice)     - `run`: Run (default)     - `check`: Check * `inventory`:  (id, default=``) * `limit`:  (string, default=`\"\"`) * `credential`:  (id, default=``) * `module_name`:  (choice)     - `command` (default)     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string, default=`\"\"`) * `forks`:  (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (Normal) (default)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string, default=`\"\"`) * `become_enabled`:  (boolean, default=`False`) * `diff_mode`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsCreateOpts struct

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

# **AdHocCommandsAdHocCommandsDelete**
> AdHocCommandsAdHocCommandsDelete(ctx, id, optional)
 Delete an Ad Hoc Command

 Make a DELETE request to this resource to delete this ad hoc command.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsDeleteOpts struct

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

# **AdHocCommandsAdHocCommandsEventsList**
> AdHocCommandsAdHocCommandsEventsList(ctx, id, optional)
 List Ad Hoc Command Events for an Ad Hoc Command

 Make a GET request to this resource to retrieve a list of ad hoc command events associated with the selected ad hoc command.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc command events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command event records.    ## Results  Each ad hoc command event data structure includes the following fields:  * `id`: Database ID for this ad hoc command event. (integer) * `type`: Data type for this ad hoc command event. (choice) * `url`: URL for this ad hoc command event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command event was created. (datetime) * `modified`: Timestamp when this ad hoc command event was last modified. (datetime) * `ad_hoc_command`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_ok`: Host OK     - `runner_on_unreachable`: Host Unreachable     - `runner_on_skipped`: Host Skipped     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that ad hoc command events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsEventsListOpts struct

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

# **AdHocCommandsAdHocCommandsList**
> AdHocCommandsAdHocCommandsList(ctx, optional)
 List Ad Hoc Commands

 Make a GET request to this resource to retrieve the list of ad hoc commands.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc commands found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command records.    ## Results  Each ad hoc command data structure includes the following fields:  * `id`: Database ID for this ad hoc command. (integer) * `type`: Data type for this ad hoc command. (choice) * `url`: URL for this ad hoc command. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command was created. (datetime) * `modified`: Timestamp when this ad hoc command was last modified. (datetime) * `name`: Name of this ad hoc command. (string) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the execution environment. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `limit`:  (string) * `credential`:  (id) * `module_name`:  (choice)     - `command`     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string) * `forks`:  (integer) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string) * `become_enabled`:  (boolean) * `diff_mode`:  (boolean)    ## Sorting  To specify that ad hoc commands are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsListOpts struct

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

# **AdHocCommandsAdHocCommandsNotificationsList**
> AdHocCommandsAdHocCommandsNotificationsList(ctx, id, optional)
 List Notifications for an Ad Hoc Command

 Make a GET request to this resource to retrieve a list of notifications associated with the selected ad hoc command.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notifications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * `id`: Database ID for this notification. (integer) * `type`: Data type for this notification. (choice) * `url`: URL for this notification. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification was created. (datetime) * `modified`: Timestamp when this notification was last modified. (datetime) * `notification_template`:  (id) * `error`:  (string) * `status`:  (choice)     - `pending`: Pending     - `successful`: Successful     - `failed`: Failed * `notifications_sent`:  (integer) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `recipients`:  (string) * `subject`:  (string) * `body`: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsNotificationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsNotificationsListOpts struct

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

# **AdHocCommandsAdHocCommandsRead**
> AdHocCommandsAdHocCommandsRead(ctx, id, optional)
 Retrieve an Ad Hoc Command

 Make GET request to this resource to retrieve a single ad hoc command record containing the following fields:  * `id`: Database ID for this ad hoc command. (integer) * `type`: Data type for this ad hoc command. (choice) * `url`: URL for this ad hoc command. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command was created. (datetime) * `modified`: Timestamp when this ad hoc command was last modified. (datetime) * `name`: Name of this ad hoc command. (string) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_args`:  (string) * `job_cwd`:  (string) * `job_env`:  (json) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the execution environment. (string) * `result_traceback`:  (string) * `event_processing_finished`: Indicates whether all of the events generated by this unified job have been saved to the database. (boolean) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `limit`:  (string) * `credential`:  (id) * `module_name`:  (choice)     - `command`     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string) * `forks`:  (integer) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string) * `become_enabled`:  (boolean) * `diff_mode`:  (boolean) * `host_status_counts`: A count of hosts uniquely assigned to each status. (field)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsReadOpts struct

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

# **AdHocCommandsAdHocCommandsRelaunchCreate**
> AdHocCommandsAdHocCommandsRelaunchCreate(ctx, id)
Relaunch an Ad Hoc Command

 Make a POST request to this resource to launch a job. If any passwords or variables are required then they should be passed in via POST data.   In order to determine what values are required in order to launch a job based on this job template you may make a GET request to this endpoint.

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

# **AdHocCommandsAdHocCommandsRelaunchList**
> AdHocCommandsAdHocCommandsRelaunchList(ctx, id, optional)
Relaunch an Ad Hoc Command

 Make a POST request to this resource to launch a job. If any passwords or variables are required then they should be passed in via POST data.   In order to determine what values are required in order to launch a job based on this job template you may make a GET request to this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandsApiAdHocCommandsAdHocCommandsRelaunchListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandsApiAdHocCommandsAdHocCommandsRelaunchListOpts struct

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

# **AdHocCommandsAdHocCommandsStdoutRead**
> AdHocCommandsAdHocCommandsStdoutRead(ctx, id)
 Retrieve Ad Hoc Command Stdout

 Make GET request to this resource to retrieve the stdout from running this ad hoc command.  ## Format  Use the `format` query string parameter to specify the output format.  * Browsable API: `?format=api` * HTML: `?format=html` * Plain Text: `?format=txt` * Plain Text with ANSI color codes: `?format=ansi` * JSON structure: `?format=json` * Downloaded Plain Text: `?format=txt_download` * Downloaded Plain Text with ANSI color codes: `?format=ansi_download`  (_New in Ansible Tower 2.0.0_) When using the Browsable API, HTML and JSON formats, the `start_line` and `end_line` query string parameters can be used to specify a range of line numbers to retrieve.  Use `dark=1` or `dark=0` as a query string parameter to force or disable a dark background.  Files over 1.0 MB (configurable) will not display in the browser. Use the `txt_download` or `ansi_download` formats to download the file directly to view it.

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
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

