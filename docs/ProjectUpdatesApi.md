# \ProjectUpdatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectUpdatesProjectUpdatesCancelCreate**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesCancelCreate) | **Post** /api/v2/project_updates/{id}/cancel/ |  Cancel Project Update
[**ProjectUpdatesProjectUpdatesCancelRead**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesCancelRead) | **Get** /api/v2/project_updates/{id}/cancel/ |  Cancel Project Update
[**ProjectUpdatesProjectUpdatesDelete**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesDelete) | **Delete** /api/v2/project_updates/{id}/ |  Delete a Project Update
[**ProjectUpdatesProjectUpdatesEventsList**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesEventsList) | **Get** /api/v2/project_updates/{id}/events/ |  List Project Update Events for a Project Update
[**ProjectUpdatesProjectUpdatesList**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesList) | **Get** /api/v2/project_updates/ |  List Project Updates
[**ProjectUpdatesProjectUpdatesNotificationsList**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesNotificationsList) | **Get** /api/v2/project_updates/{id}/notifications/ |  List Notifications for a Project Update
[**ProjectUpdatesProjectUpdatesRead**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesRead) | **Get** /api/v2/project_updates/{id}/ |  Retrieve a Project Update
[**ProjectUpdatesProjectUpdatesScmInventoryUpdatesList**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesScmInventoryUpdatesList) | **Get** /api/v2/project_updates/{id}/scm_inventory_updates/ |  List Inventory Updates for a Project Update
[**ProjectUpdatesProjectUpdatesStdoutRead**](ProjectUpdatesApi.md#ProjectUpdatesProjectUpdatesStdoutRead) | **Get** /api/v2/project_updates/{id}/stdout/ |  Retrieve Project Update Stdout


# **ProjectUpdatesProjectUpdatesCancelCreate**
> ProjectUpdatesProjectUpdatesCancelCreate(ctx, id)
 Cancel Project Update

 Make a GET request to this resource to determine if the project update can be canceled.  The response will include the following field:  * `can_cancel`: Indicates whether this update can be canceled (boolean,   read-only)  Make a POST request to this resource to cancel a pending or running project update.  The response status code will be 202 if successful, or 405 if the update cannot be canceled.

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

# **ProjectUpdatesProjectUpdatesCancelRead**
> ProjectUpdatesProjectUpdatesCancelRead(ctx, id, optional)
 Cancel Project Update

 Make a GET request to this resource to determine if the project update can be canceled.  The response will include the following field:  * `can_cancel`: Indicates whether this update can be canceled (boolean,   read-only)  Make a POST request to this resource to cancel a pending or running project update.  The response status code will be 202 if successful, or 405 if the update cannot be canceled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesCancelReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesCancelReadOpts struct

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

# **ProjectUpdatesProjectUpdatesDelete**
> ProjectUpdatesProjectUpdatesDelete(ctx, id, optional)
 Delete a Project Update

 Make a DELETE request to this resource to delete this project update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesDeleteOpts struct

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

# **ProjectUpdatesProjectUpdatesEventsList**
> ProjectUpdatesProjectUpdatesEventsList(ctx, id, optional)
 List Project Update Events for a Project Update

 Make a GET request to this resource to retrieve a list of project update events associated with the selected project update.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of project update events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project update event records.    ## Results  Each project update event data structure includes the following fields:  * `id`: Database ID for this project update event. (integer) * `type`: Data type for this project update event. (choice) * `url`: URL for this project update event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project update event was created. (datetime) * `modified`: Timestamp when this project update event was last modified. (datetime) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_start`: Host Started     - `runner_on_ok`: Host OK     - `runner_on_error`: Host Failure     - `runner_on_skipped`: Host Skipped     - `runner_on_unreachable`: Host Unreachable     - `runner_on_no_hosts`: No Hosts Remaining     - `runner_on_async_poll`: Host Polling     - `runner_on_async_ok`: Host Async OK     - `runner_on_async_failed`: Host Async Failure     - `runner_item_on_ok`: Item OK     - `runner_item_on_failed`: Item Failed     - `runner_item_on_skipped`: Item Skipped     - `runner_retry`: Host Retry     - `runner_on_file_diff`: File Difference     - `playbook_on_start`: Playbook Started     - `playbook_on_notify`: Running Handlers     - `playbook_on_include`: Including File     - `playbook_on_no_hosts_matched`: No Hosts Matched     - `playbook_on_no_hosts_remaining`: No Hosts Remaining     - `playbook_on_task_start`: Task Started     - `playbook_on_vars_prompt`: Variables Prompted     - `playbook_on_setup`: Gathering Facts     - `playbook_on_import_for_host`: internal: on Import for Host     - `playbook_on_not_import_for_host`: internal: on Not Import for Host     - `playbook_on_play_start`: Play Started     - `playbook_on_stats`: Playbook Complete     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `event_level`:  (integer) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host_name`:  (field) * `playbook`:  (string) * `play`:  (string) * `task`:  (string) * `role`:  (string) * `stdout`:  (field) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer) * `project_update`:  (id)    ## Sorting  To specify that project update events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesEventsListOpts struct

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

# **ProjectUpdatesProjectUpdatesList**
> ProjectUpdatesProjectUpdatesList(ctx, optional)
 List Project Updates

 Make a GET request to this resource to retrieve the list of project updates.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of project updates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project update records.    ## Results  Each project update data structure includes the following fields:  * `id`: Database ID for this project update. (integer) * `type`: Data type for this project update. (choice) * `url`: URL for this project update. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project update was created. (datetime) * `modified`: Timestamp when this project update was last modified. (datetime) * `name`: Name of this project update. (string) * `description`: Optional description of this project update. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The SCM Revision discovered by this update for the given project and branch. (string) * `project`:  (id) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `job_tags`: Parts of the project update playbook that will be run. (string)    ## Sorting  To specify that project updates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesListOpts struct

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

# **ProjectUpdatesProjectUpdatesNotificationsList**
> ProjectUpdatesProjectUpdatesNotificationsList(ctx, id, optional)
 List Notifications for a Project Update

 Make a GET request to this resource to retrieve a list of notifications associated with the selected project update.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notifications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * `id`: Database ID for this notification. (integer) * `type`: Data type for this notification. (choice) * `url`: URL for this notification. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification was created. (datetime) * `modified`: Timestamp when this notification was last modified. (datetime) * `notification_template`:  (id) * `error`:  (string) * `status`:  (choice)     - `pending`: Pending     - `successful`: Successful     - `failed`: Failed * `notifications_sent`:  (integer) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `recipients`:  (string) * `subject`:  (string) * `body`: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesNotificationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesNotificationsListOpts struct

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

# **ProjectUpdatesProjectUpdatesRead**
> ProjectUpdatesProjectUpdatesRead(ctx, id, optional)
 Retrieve a Project Update

 Make GET request to this resource to retrieve a single project update record containing the following fields:  * `id`: Database ID for this project update. (integer) * `type`: Data type for this project update. (choice) * `url`: URL for this project update. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project update was created. (datetime) * `modified`: Timestamp when this project update was last modified. (datetime) * `name`: Name of this project update. (string) * `description`: Optional description of this project update. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The SCM Revision discovered by this update for the given project and branch. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_args`:  (string) * `job_cwd`:  (string) * `job_env`:  (json) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `result_traceback`:  (string) * `event_processing_finished`: Indicates whether all of the events generated by this unified job have been saved to the database. (boolean) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `project`:  (id) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `job_tags`: Parts of the project update playbook that will be run. (string) * `host_status_counts`: A count of hosts uniquely assigned to each status. (field) * `playbook_counts`: A count of all plays and tasks for the job run. (field)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesReadOpts struct

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

# **ProjectUpdatesProjectUpdatesScmInventoryUpdatesList**
> ProjectUpdatesProjectUpdatesScmInventoryUpdatesList(ctx, id, optional)
 List Inventory Updates for a Project Update

 Make a GET request to this resource to retrieve a list of inventory updates associated with the selected project update.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory updates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory update records.    ## Results  Each inventory update data structure includes the following fields:  * `id`: Database ID for this inventory update. (integer) * `type`: Data type for this inventory update. (choice) * `url`: URL for this inventory update. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory update was created. (datetime) * `modified`: Timestamp when this inventory update was last modified. (datetime) * `name`: Name of this inventory update. (string) * `description`: Optional description of this inventory update. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`:  (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `inventory`:  (id) * `inventory_source`:  (id) * `license_error`:  (boolean) * `org_host_limit_error`:  (boolean) * `source_project_update`: Inventory files from this Project Update were used for the inventory update. (id) * `instance_group`: The Instance group the job was run under (id)    ## Sorting  To specify that inventory updates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProjectUpdatesApiProjectUpdatesProjectUpdatesScmInventoryUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectUpdatesApiProjectUpdatesProjectUpdatesScmInventoryUpdatesListOpts struct

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

# **ProjectUpdatesProjectUpdatesStdoutRead**
> ProjectUpdatesProjectUpdatesStdoutRead(ctx, id)
 Retrieve Project Update Stdout

 Make GET request to this resource to retrieve the stdout from running this project update.  ## Format  Use the `format` query string parameter to specify the output format.  * Browsable API: `?format=api` * HTML: `?format=html` * Plain Text: `?format=txt` * Plain Text with ANSI color codes: `?format=ansi` * JSON structure: `?format=json` * Downloaded Plain Text: `?format=txt_download` * Downloaded Plain Text with ANSI color codes: `?format=ansi_download`  (_New in Ansible Tower 2.0.0_) When using the Browsable API, HTML and JSON formats, the `start_line` and `end_line` query string parameters can be used to specify a range of line numbers to retrieve.  Use `dark=1` or `dark=0` as a query string parameter to force or disable a dark background.  Files over 1.0 MB (configurable) will not display in the browser. Use the `txt_download` or `ansi_download` formats to download the file directly to view it.

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
 - **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

