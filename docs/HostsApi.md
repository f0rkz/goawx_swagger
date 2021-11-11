# \HostsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostsHostsActivityStreamList**](HostsApi.md#HostsHostsActivityStreamList) | **Get** /api/v2/hosts/{id}/activity_stream/ |  List Activity Streams for a Host
[**HostsHostsAdHocCommandEventsList**](HostsApi.md#HostsHostsAdHocCommandEventsList) | **Get** /api/v2/hosts/{id}/ad_hoc_command_events/ |  List Ad Hoc Command Events for a Host
[**HostsHostsAdHocCommandsCreate**](HostsApi.md#HostsHostsAdHocCommandsCreate) | **Post** /api/v2/hosts/{id}/ad_hoc_commands/ |  Create an Ad Hoc Command for a Host
[**HostsHostsAdHocCommandsList**](HostsApi.md#HostsHostsAdHocCommandsList) | **Get** /api/v2/hosts/{id}/ad_hoc_commands/ |  List Ad Hoc Commands for a Host
[**HostsHostsAllGroupsList**](HostsApi.md#HostsHostsAllGroupsList) | **Get** /api/v2/hosts/{id}/all_groups/ |  List All Groups for a Host
[**HostsHostsAnsibleFactsRead**](HostsApi.md#HostsHostsAnsibleFactsRead) | **Get** /api/v2/hosts/{id}/ansible_facts/ |  Retrieve a Host
[**HostsHostsCreate**](HostsApi.md#HostsHostsCreate) | **Post** /api/v2/hosts/ | &#x60;host_filter&#x60; is available on this endpoint. The filter supports: relational queries, &#x60;and&#x60; &#x60;or&#x60; boolean logic, as well as expression grouping via &#x60;()&#x60;.
[**HostsHostsDelete**](HostsApi.md#HostsHostsDelete) | **Delete** /api/v2/hosts/{id}/ |  Delete a Host
[**HostsHostsGroupsCreate**](HostsApi.md#HostsHostsGroupsCreate) | **Post** /api/v2/hosts/{id}/groups/ |  Create a Group for a Host
[**HostsHostsGroupsList**](HostsApi.md#HostsHostsGroupsList) | **Get** /api/v2/hosts/{id}/groups/ |  List Groups for a Host
[**HostsHostsInventorySourcesList**](HostsApi.md#HostsHostsInventorySourcesList) | **Get** /api/v2/hosts/{id}/inventory_sources/ |  List Inventory Sources for a Host
[**HostsHostsJobEventsList**](HostsApi.md#HostsHostsJobEventsList) | **Get** /api/v2/hosts/{id}/job_events/ |  List Job Events for a Host
[**HostsHostsJobHostSummariesList**](HostsApi.md#HostsHostsJobHostSummariesList) | **Get** /api/v2/hosts/{id}/job_host_summaries/ |  List Job Host Summaries for a Host
[**HostsHostsList**](HostsApi.md#HostsHostsList) | **Get** /api/v2/hosts/ |  List Hosts
[**HostsHostsPartialUpdate**](HostsApi.md#HostsHostsPartialUpdate) | **Patch** /api/v2/hosts/{id}/ |  Update a Host
[**HostsHostsRead**](HostsApi.md#HostsHostsRead) | **Get** /api/v2/hosts/{id}/ |  Retrieve a Host
[**HostsHostsSmartInventoriesList**](HostsApi.md#HostsHostsSmartInventoriesList) | **Get** /api/v2/hosts/{id}/smart_inventories/ |  List Inventories for a Host
[**HostsHostsUpdate**](HostsApi.md#HostsHostsUpdate) | **Put** /api/v2/hosts/{id}/ |  Update a Host
[**HostsHostsVariableDataPartialUpdate**](HostsApi.md#HostsHostsVariableDataPartialUpdate) | **Patch** /api/v2/hosts/{id}/variable_data/ |  Update Host Variable Data
[**HostsHostsVariableDataRead**](HostsApi.md#HostsHostsVariableDataRead) | **Get** /api/v2/hosts/{id}/variable_data/ |  Retrieve Host Variable Data
[**HostsHostsVariableDataUpdate**](HostsApi.md#HostsHostsVariableDataUpdate) | **Put** /api/v2/hosts/{id}/variable_data/ |  Update Host Variable Data


# **HostsHostsActivityStreamList**
> HostsHostsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Host

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsActivityStreamListOpts struct

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

# **HostsHostsAdHocCommandEventsList**
> HostsHostsAdHocCommandEventsList(ctx, id, optional)
 List Ad Hoc Command Events for a Host

 Make a GET request to this resource to retrieve a list of ad hoc command events associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc command events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command event records.    ## Results  Each ad hoc command event data structure includes the following fields:  * `id`: Database ID for this ad hoc command event. (integer) * `type`: Data type for this ad hoc command event. (choice) * `url`: URL for this ad hoc command event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command event was created. (datetime) * `modified`: Timestamp when this ad hoc command event was last modified. (datetime) * `ad_hoc_command`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_ok`: Host OK     - `runner_on_unreachable`: Host Unreachable     - `runner_on_skipped`: Host Skipped     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that ad hoc command events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsAdHocCommandEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsAdHocCommandEventsListOpts struct

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

# **HostsHostsAdHocCommandsCreate**
> HostsHostsAdHocCommandsCreate(ctx, id, optional)
 Create an Ad Hoc Command for a Host

 Make a POST request to this resource with the following ad hoc command fields to create a new ad hoc command associated with this host.             * `execution_environment`: The container image to be used for execution. (id, default=``)           * `job_type`:  (choice)     - `run`: Run (default)     - `check`: Check * `inventory`:  (id, default=``) * `limit`:  (string, default=`\"\"`) * `credential`:  (id, default=``) * `module_name`:  (choice)     - `command` (default)     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string, default=`\"\"`) * `forks`:  (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (Normal) (default)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string, default=`\"\"`) * `become_enabled`:  (boolean, default=`False`) * `diff_mode`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsAdHocCommandsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsAdHocCommandsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data16**](Data16.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsHostsAdHocCommandsList**
> HostsHostsAdHocCommandsList(ctx, id, optional)
 List Ad Hoc Commands for a Host

 Make a GET request to this resource to retrieve a list of ad hoc commands associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc commands found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command records.    ## Results  Each ad hoc command data structure includes the following fields:  * `id`: Database ID for this ad hoc command. (integer) * `type`: Data type for this ad hoc command. (choice) * `url`: URL for this ad hoc command. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command was created. (datetime) * `modified`: Timestamp when this ad hoc command was last modified. (datetime) * `name`: Name of this ad hoc command. (string) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the execution environment. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `limit`:  (string) * `credential`:  (id) * `module_name`:  (choice)     - `command`     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string) * `forks`:  (integer) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string) * `become_enabled`:  (boolean) * `diff_mode`:  (boolean)    ## Sorting  To specify that ad hoc commands are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsAdHocCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsAdHocCommandsListOpts struct

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

# **HostsHostsAllGroupsList**
> HostsHostsAllGroupsList(ctx, id, optional)
 List All Groups for a Host

 Make a GET request to this resource to retrieve a list of all groups of which the selected host is directly or indirectly a member.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsAllGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsAllGroupsListOpts struct

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

# **HostsHostsAnsibleFactsRead**
> HostsHostsAnsibleFactsRead(ctx, id, optional)
 Retrieve a Host

 Make GET request to this resource to retrieve a single host record containing the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsAnsibleFactsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsAnsibleFactsReadOpts struct

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

# **HostsHostsCreate**
> HostsHostsCreate(ctx, optional)
`host_filter` is available on this endpoint. The filter supports: relational queries, `and` `or` boolean logic, as well as expression grouping via `()`.

     ?host_filter=name=my_host     ?host_filter=name=\"my host\" or name=my_host     ?host_filter=groups__name=\"my group\"     ?host_filter=name=my_host and groups__name=\"my group\"     ?host_filter=name=my_host and groups__name=\"my group\"     ?host_filter=(name=my_host and groups__name=\"my group\") or (name=my_host2 and groups__name=my_group2)  `host_filter` can also be used to query JSON data in the related `ansible_facts`. `__` may be used to traverse JSON dictionaries. `[]` may be used to traverse JSON arrays.      ?host_filter=ansible_facts__ansible_processor_vcpus=8     ?host_filter=ansible_facts__ansible_processor_vcpus=8 and name=\"my_host\" and ansible_facts__ansible_lo__ipv6[]__scope=host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsApiHostsHostsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data14**](Data14.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsHostsDelete**
> HostsHostsDelete(ctx, id, optional)
 Delete a Host

 Make a DELETE request to this resource to delete this host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsDeleteOpts struct

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

# **HostsHostsGroupsCreate**
> HostsHostsGroupsCreate(ctx, id, optional)
 Create a Group for a Host

 Make a POST request to this resource with the following group fields to create a new group associated with this host.          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)         # Add Groups for a Host:  Make a POST request to this resource with only an `id` field to associate an existing group with this host.  # Remove Groups from this Host:  Make a POST request to this resource with `id` and `disassociate` fields to remove the group from this host  without deleting the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsGroupsCreateOpts struct

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

# **HostsHostsGroupsList**
> HostsHostsGroupsList(ctx, id, optional)
 List Groups for a Host

 Make a GET request to this resource to retrieve a list of groups associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsGroupsListOpts struct

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

# **HostsHostsInventorySourcesList**
> HostsHostsInventorySourcesList(ctx, id, optional)
 List Inventory Sources for a Host

 Make a GET request to this resource to retrieve a list of inventory sources associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory source records.    ## Results  Each inventory source data structure includes the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that inventory sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsInventorySourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsInventorySourcesListOpts struct

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

# **HostsHostsJobEventsList**
> HostsHostsJobEventsList(ctx, id, optional)
 List Job Events for a Host

 Make a GET request to this resource to retrieve a list of job events associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job event records.    ## Results  Each job event data structure includes the following fields:  * `id`: Database ID for this job event. (integer) * `type`: Data type for this job event. (choice) * `url`: URL for this job event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job event was created. (datetime) * `modified`: Timestamp when this job event was last modified. (datetime) * `job`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_start`: Host Started     - `runner_on_ok`: Host OK     - `runner_on_error`: Host Failure     - `runner_on_skipped`: Host Skipped     - `runner_on_unreachable`: Host Unreachable     - `runner_on_no_hosts`: No Hosts Remaining     - `runner_on_async_poll`: Host Polling     - `runner_on_async_ok`: Host Async OK     - `runner_on_async_failed`: Host Async Failure     - `runner_item_on_ok`: Item OK     - `runner_item_on_failed`: Item Failed     - `runner_item_on_skipped`: Item Skipped     - `runner_retry`: Host Retry     - `runner_on_file_diff`: File Difference     - `playbook_on_start`: Playbook Started     - `playbook_on_notify`: Running Handlers     - `playbook_on_include`: Including File     - `playbook_on_no_hosts_matched`: No Hosts Matched     - `playbook_on_no_hosts_remaining`: No Hosts Remaining     - `playbook_on_task_start`: Task Started     - `playbook_on_vars_prompt`: Variables Prompted     - `playbook_on_setup`: Gathering Facts     - `playbook_on_import_for_host`: internal: on Import for Host     - `playbook_on_not_import_for_host`: internal: on Not Import for Host     - `playbook_on_play_start`: Play Started     - `playbook_on_stats`: Playbook Complete     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `event_level`:  (integer) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `parent_uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `playbook`:  (string) * `play`:  (string) * `task`:  (string) * `role`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that job events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsJobEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsJobEventsListOpts struct

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

# **HostsHostsJobHostSummariesList**
> HostsHostsJobHostSummariesList(ctx, id, optional)
 List Job Host Summaries for a Host

 Make a GET request to this resource to retrieve a list of job host summaries associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job host summaries found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job host summary records.    ## Results  Each job host summary data structure includes the following fields:  * `id`: Database ID for this job host summary. (integer) * `type`: Data type for this job host summary. (choice) * `url`: URL for this job host summary. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job host summary was created. (datetime) * `modified`: Timestamp when this job host summary was last modified. (datetime) * `job`:  (id) * `host`:  (id) * `host_name`:  (string) * `changed`:  (integer) * `dark`:  (integer) * `failures`:  (integer) * `ok`:  (integer) * `processed`:  (integer) * `skipped`:  (integer) * `failed`:  (boolean) * `ignored`:  (integer) * `rescued`:  (integer)    ## Sorting  To specify that job host summaries are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsJobHostSummariesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsJobHostSummariesListOpts struct

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

# **HostsHostsList**
> HostsHostsList(ctx, optional)
 List Hosts

 Make a GET request to this resource to retrieve the list of hosts.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of hosts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more host records.    ## Results  Each host data structure includes the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)    ## Sorting  To specify that hosts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsApiHostsHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsListOpts struct

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

# **HostsHostsPartialUpdate**
> HostsHostsPartialUpdate(ctx, id, optional)
 Update a Host

 Make a PUT or PATCH request to this resource to update this host.  The following fields may be modified:          * `name`: Name of this host. (string, required) * `description`: Optional description of this host. (string, default=`\"\"`) * `inventory`:  (id, required) * `enabled`: Is this host online and available for running jobs? (boolean, default=`True`) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string, default=`\"\"`) * `variables`: Host variables in JSON or YAML format. (json, default=``)              For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data15**](Data15.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsHostsRead**
> HostsHostsRead(ctx, id, optional)
 Retrieve a Host

 Make GET request to this resource to retrieve a single host record containing the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsReadOpts struct

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

# **HostsHostsSmartInventoriesList**
> HostsHostsSmartInventoriesList(ctx, id, optional)
 List Inventories for a Host

 Make a GET request to this resource to retrieve a list of inventories associated with the selected host.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventories found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory records.    ## Results  Each inventory data structure includes the following fields:  * `id`: Database ID for this inventory. (integer) * `type`: Data type for this inventory. (choice) * `url`: URL for this inventory. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory was created. (datetime) * `modified`: Timestamp when this inventory was last modified. (datetime) * `name`: Name of this inventory. (string) * `description`: Optional description of this inventory. (string) * `organization`: Organization containing this inventory. (id) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory.     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string) * `variables`: Inventory variables in JSON or YAML format. (json) * `has_active_failures`: This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed. (boolean) * `total_hosts`: This field is deprecated and will be removed in a future release. Total number of hosts in this inventory. (integer) * `hosts_with_active_failures`: This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures. (integer) * `total_groups`: This field is deprecated and will be removed in a future release. Total number of groups in this inventory. (integer) * `has_inventory_sources`: This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources. (boolean) * `total_inventory_sources`: Total number of external inventory sources configured within this inventory. (integer) * `inventory_sources_with_failures`: Number of external inventory sources in this inventory with failures. (integer) * `pending_deletion`: Flag indicating the inventory is being deleted. (boolean)    ## Sorting  To specify that inventories are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsSmartInventoriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsSmartInventoriesListOpts struct

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

# **HostsHostsUpdate**
> HostsHostsUpdate(ctx, id, optional)
 Update a Host

 Make a PUT or PATCH request to this resource to update this host.  The following fields may be modified:          * `name`: Name of this host. (string, required) * `description`: Optional description of this host. (string, default=`\"\"`) * `inventory`:  (id, required) * `enabled`: Is this host online and available for running jobs? (boolean, default=`True`) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string, default=`\"\"`) * `variables`: Host variables in JSON or YAML format. (json, default=``)            For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsUpdateOpts struct

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

# **HostsHostsVariableDataPartialUpdate**
> HostsHostsVariableDataPartialUpdate(ctx, id, optional)
 Update Host Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsVariableDataPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsVariableDataPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data18**](Data18.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsHostsVariableDataRead**
> HostsHostsVariableDataRead(ctx, id, optional)
 Retrieve Host Variable Data

 Make a GET request to this resource to retrieve all variables defined for a host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsVariableDataReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsVariableDataReadOpts struct

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

# **HostsHostsVariableDataUpdate**
> HostsHostsVariableDataUpdate(ctx, id, optional)
 Update Host Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***HostsApiHostsHostsVariableDataUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsHostsVariableDataUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data17**](Data17.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

