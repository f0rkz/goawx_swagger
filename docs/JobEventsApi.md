# \JobEventsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobEventsJobEventsChildrenList**](JobEventsApi.md#JobEventsJobEventsChildrenList) | **Get** /api/v2/job_events/{id}/children/ |  List Job Events for a Job Event
[**JobEventsJobEventsRead**](JobEventsApi.md#JobEventsJobEventsRead) | **Get** /api/v2/job_events/{id}/ |  Retrieve a Job Event


# **JobEventsJobEventsChildrenList**
> JobEventsJobEventsChildrenList(ctx, id, optional)
 List Job Events for a Job Event

 Make a GET request to this resource to retrieve a list of job events associated with the selected job event.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job event records.    ## Results  Each job event data structure includes the following fields:  * `id`: Database ID for this job event. (integer) * `type`: Data type for this job event. (choice) * `url`: URL for this job event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job event was created. (datetime) * `modified`: Timestamp when this job event was last modified. (datetime) * `job`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_start`: Host Started     - `runner_on_ok`: Host OK     - `runner_on_error`: Host Failure     - `runner_on_skipped`: Host Skipped     - `runner_on_unreachable`: Host Unreachable     - `runner_on_no_hosts`: No Hosts Remaining     - `runner_on_async_poll`: Host Polling     - `runner_on_async_ok`: Host Async OK     - `runner_on_async_failed`: Host Async Failure     - `runner_item_on_ok`: Item OK     - `runner_item_on_failed`: Item Failed     - `runner_item_on_skipped`: Item Skipped     - `runner_retry`: Host Retry     - `runner_on_file_diff`: File Difference     - `playbook_on_start`: Playbook Started     - `playbook_on_notify`: Running Handlers     - `playbook_on_include`: Including File     - `playbook_on_no_hosts_matched`: No Hosts Matched     - `playbook_on_no_hosts_remaining`: No Hosts Remaining     - `playbook_on_task_start`: Task Started     - `playbook_on_vars_prompt`: Variables Prompted     - `playbook_on_setup`: Gathering Facts     - `playbook_on_import_for_host`: internal: on Import for Host     - `playbook_on_not_import_for_host`: internal: on Not Import for Host     - `playbook_on_play_start`: Play Started     - `playbook_on_stats`: Playbook Complete     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `event_level`:  (integer) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `parent_uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `playbook`:  (string) * `play`:  (string) * `task`:  (string) * `role`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that job events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***JobEventsApiJobEventsJobEventsChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobEventsApiJobEventsJobEventsChildrenListOpts struct

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

# **JobEventsJobEventsRead**
> JobEventsJobEventsRead(ctx, id, optional)
 Retrieve a Job Event

 Make GET request to this resource to retrieve a single job event record containing the following fields:  * `id`: Database ID for this job event. (integer) * `type`: Data type for this job event. (choice) * `url`: URL for this job event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job event was created. (datetime) * `modified`: Timestamp when this job event was last modified. (datetime) * `job`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_start`: Host Started     - `runner_on_ok`: Host OK     - `runner_on_error`: Host Failure     - `runner_on_skipped`: Host Skipped     - `runner_on_unreachable`: Host Unreachable     - `runner_on_no_hosts`: No Hosts Remaining     - `runner_on_async_poll`: Host Polling     - `runner_on_async_ok`: Host Async OK     - `runner_on_async_failed`: Host Async Failure     - `runner_item_on_ok`: Item OK     - `runner_item_on_failed`: Item Failed     - `runner_item_on_skipped`: Item Skipped     - `runner_retry`: Host Retry     - `runner_on_file_diff`: File Difference     - `playbook_on_start`: Playbook Started     - `playbook_on_notify`: Running Handlers     - `playbook_on_include`: Including File     - `playbook_on_no_hosts_matched`: No Hosts Matched     - `playbook_on_no_hosts_remaining`: No Hosts Remaining     - `playbook_on_task_start`: Task Started     - `playbook_on_vars_prompt`: Variables Prompted     - `playbook_on_setup`: Gathering Facts     - `playbook_on_import_for_host`: internal: on Import for Host     - `playbook_on_not_import_for_host`: internal: on Not Import for Host     - `playbook_on_play_start`: Play Started     - `playbook_on_stats`: Playbook Complete     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `event_level`:  (integer) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `parent_uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `playbook`:  (string) * `play`:  (string) * `task`:  (string) * `role`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***JobEventsApiJobEventsJobEventsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobEventsApiJobEventsJobEventsReadOpts struct

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

