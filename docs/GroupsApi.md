# \GroupsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGroupsActivityStreamList**](GroupsApi.md#GroupsGroupsActivityStreamList) | **Get** /api/v2/groups/{id}/activity_stream/ |  List Activity Streams for a Group
[**GroupsGroupsAdHocCommandsCreate**](GroupsApi.md#GroupsGroupsAdHocCommandsCreate) | **Post** /api/v2/groups/{id}/ad_hoc_commands/ |  Create an Ad Hoc Command for a Group
[**GroupsGroupsAdHocCommandsList**](GroupsApi.md#GroupsGroupsAdHocCommandsList) | **Get** /api/v2/groups/{id}/ad_hoc_commands/ |  List Ad Hoc Commands for a Group
[**GroupsGroupsAllHostsList**](GroupsApi.md#GroupsGroupsAllHostsList) | **Get** /api/v2/groups/{id}/all_hosts/ |  List All Hosts for a Group
[**GroupsGroupsChildrenCreate**](GroupsApi.md#GroupsGroupsChildrenCreate) | **Post** /api/v2/groups/{id}/children/ |  Create a Group for a Group
[**GroupsGroupsChildrenList**](GroupsApi.md#GroupsGroupsChildrenList) | **Get** /api/v2/groups/{id}/children/ |  List Groups for a Group
[**GroupsGroupsCreate**](GroupsApi.md#GroupsGroupsCreate) | **Post** /api/v2/groups/ |  Create a Group
[**GroupsGroupsDelete**](GroupsApi.md#GroupsGroupsDelete) | **Delete** /api/v2/groups/{id}/ |  Delete a Group
[**GroupsGroupsHostsCreate**](GroupsApi.md#GroupsGroupsHostsCreate) | **Post** /api/v2/groups/{id}/hosts/ |  Create a Host for a Group
[**GroupsGroupsHostsList**](GroupsApi.md#GroupsGroupsHostsList) | **Get** /api/v2/groups/{id}/hosts/ |  List Hosts for a Group
[**GroupsGroupsInventorySourcesList**](GroupsApi.md#GroupsGroupsInventorySourcesList) | **Get** /api/v2/groups/{id}/inventory_sources/ |  List Inventory Sources for a Group
[**GroupsGroupsJobEventsList**](GroupsApi.md#GroupsGroupsJobEventsList) | **Get** /api/v2/groups/{id}/job_events/ |  List Job Events for a Group
[**GroupsGroupsJobHostSummariesList**](GroupsApi.md#GroupsGroupsJobHostSummariesList) | **Get** /api/v2/groups/{id}/job_host_summaries/ |  List Job Host Summaries for a Group
[**GroupsGroupsList**](GroupsApi.md#GroupsGroupsList) | **Get** /api/v2/groups/ |  List Groups
[**GroupsGroupsPartialUpdate**](GroupsApi.md#GroupsGroupsPartialUpdate) | **Patch** /api/v2/groups/{id}/ |  Update a Group
[**GroupsGroupsPotentialChildrenList**](GroupsApi.md#GroupsGroupsPotentialChildrenList) | **Get** /api/v2/groups/{id}/potential_children/ |  List Potential Child Groups for a Group
[**GroupsGroupsRead**](GroupsApi.md#GroupsGroupsRead) | **Get** /api/v2/groups/{id}/ |  Retrieve a Group
[**GroupsGroupsUpdate**](GroupsApi.md#GroupsGroupsUpdate) | **Put** /api/v2/groups/{id}/ |  Update a Group
[**GroupsGroupsVariableDataPartialUpdate**](GroupsApi.md#GroupsGroupsVariableDataPartialUpdate) | **Patch** /api/v2/groups/{id}/variable_data/ |  Update Group Variable Data
[**GroupsGroupsVariableDataRead**](GroupsApi.md#GroupsGroupsVariableDataRead) | **Get** /api/v2/groups/{id}/variable_data/ |  Retrieve Group Variable Data
[**GroupsGroupsVariableDataUpdate**](GroupsApi.md#GroupsGroupsVariableDataUpdate) | **Put** /api/v2/groups/{id}/variable_data/ |  Update Group Variable Data


# **GroupsGroupsActivityStreamList**
> GroupsGroupsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Group

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsActivityStreamListOpts struct

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

# **GroupsGroupsAdHocCommandsCreate**
> GroupsGroupsAdHocCommandsCreate(ctx, id, optional)
 Create an Ad Hoc Command for a Group

 Make a POST request to this resource with the following ad hoc command fields to create a new ad hoc command associated with this group.             * `execution_environment`: The container image to be used for execution. (id, default=``)           * `job_type`:  (choice)     - `run`: Run (default)     - `check`: Check * `inventory`:  (id, default=``) * `limit`:  (string, default=`\"\"`) * `credential`:  (id, default=``) * `module_name`:  (choice)     - `command` (default)     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string, default=`\"\"`) * `forks`:  (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (Normal) (default)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string, default=`\"\"`) * `become_enabled`:  (boolean, default=`False`) * `diff_mode`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsAdHocCommandsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsAdHocCommandsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data11**](Data11.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupsAdHocCommandsList**
> GroupsGroupsAdHocCommandsList(ctx, id, optional)
 List Ad Hoc Commands for a Group

 Make a GET request to this resource to retrieve a list of ad hoc commands associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc commands found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command records.    ## Results  Each ad hoc command data structure includes the following fields:  * `id`: Database ID for this ad hoc command. (integer) * `type`: Data type for this ad hoc command. (choice) * `url`: URL for this ad hoc command. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command was created. (datetime) * `modified`: Timestamp when this ad hoc command was last modified. (datetime) * `name`: Name of this ad hoc command. (string) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the execution environment. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `limit`:  (string) * `credential`:  (id) * `module_name`:  (choice)     - `command`     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string) * `forks`:  (integer) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string) * `become_enabled`:  (boolean) * `diff_mode`:  (boolean)    ## Sorting  To specify that ad hoc commands are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsAdHocCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsAdHocCommandsListOpts struct

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

# **GroupsGroupsAllHostsList**
> GroupsGroupsAllHostsList(ctx, id, optional)
 List All Hosts for a Group

 Make a GET request to this resource to retrieve a list of all hosts directly or indirectly belonging to this group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of hosts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more host records.    ## Results  Each host data structure includes the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)    ## Sorting  To specify that hosts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsAllHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsAllHostsListOpts struct

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

# **GroupsGroupsChildrenCreate**
> GroupsGroupsChildrenCreate(ctx, id, optional)
 Create a Group for a Group

 Make a POST request to this resource with the following group fields to create a new group associated with this group.          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)         # Add Groups for a Group:  Make a POST request to this resource with only an `id` field to associate an existing group with this group.  # Remove Groups from this Group:  Make a POST request to this resource with `id` and `disassociate` fields to remove the group from this group  without deleting the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsChildrenCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsChildrenCreateOpts struct

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

# **GroupsGroupsChildrenList**
> GroupsGroupsChildrenList(ctx, id, optional)
 List Groups for a Group

 Make a GET request to this resource to retrieve a list of groups associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsChildrenListOpts struct

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

# **GroupsGroupsCreate**
> GroupsGroupsCreate(ctx, optional)
 Create a Group

 Make a POST request to this resource with the following group fields to create a new group:          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data9**](Data9.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupsDelete**
> GroupsGroupsDelete(ctx, id, optional)
 Delete a Group

 Make a DELETE request to this resource to delete this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsDeleteOpts struct

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

# **GroupsGroupsHostsCreate**
> GroupsGroupsHostsCreate(ctx, id, optional)
 Create a Host for a Group

 Make a POST request to this resource with the following host fields to create a new host associated with this group.          * `name`: Name of this host. (string, required) * `description`: Optional description of this host. (string, default=`\"\"`) * `inventory`:  (id, required) * `enabled`: Is this host online and available for running jobs? (boolean, default=`True`) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string, default=`\"\"`) * `variables`: Host variables in JSON or YAML format. (json, default=``)              # Add Hosts for a Group:  Make a POST request to this resource with only an `id` field to associate an existing host with this group.  # Remove Hosts from this Group:  Make a POST request to this resource with `id` and `disassociate` fields to remove the host from this group  without deleting the host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsHostsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsHostsCreateOpts struct

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

# **GroupsGroupsHostsList**
> GroupsGroupsHostsList(ctx, id, optional)
 List Hosts for a Group

 Make a GET request to this resource to retrieve a list of hosts associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of hosts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more host records.    ## Results  Each host data structure includes the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)    ## Sorting  To specify that hosts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsHostsListOpts struct

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

# **GroupsGroupsInventorySourcesList**
> GroupsGroupsInventorySourcesList(ctx, id, optional)
 List Inventory Sources for a Group

 Make a GET request to this resource to retrieve a list of inventory sources associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory source records.    ## Results  Each inventory source data structure includes the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that inventory sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsInventorySourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsInventorySourcesListOpts struct

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

# **GroupsGroupsJobEventsList**
> GroupsGroupsJobEventsList(ctx, id, optional)
 List Job Events for a Group

 Make a GET request to this resource to retrieve a list of job events associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job events found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job event records.    ## Results  Each job event data structure includes the following fields:  * `id`: Database ID for this job event. (integer) * `type`: Data type for this job event. (choice) * `url`: URL for this job event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job event was created. (datetime) * `modified`: Timestamp when this job event was last modified. (datetime) * `job`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_start`: Host Started     - `runner_on_ok`: Host OK     - `runner_on_error`: Host Failure     - `runner_on_skipped`: Host Skipped     - `runner_on_unreachable`: Host Unreachable     - `runner_on_no_hosts`: No Hosts Remaining     - `runner_on_async_poll`: Host Polling     - `runner_on_async_ok`: Host Async OK     - `runner_on_async_failed`: Host Async Failure     - `runner_item_on_ok`: Item OK     - `runner_item_on_failed`: Item Failed     - `runner_item_on_skipped`: Item Skipped     - `runner_retry`: Host Retry     - `runner_on_file_diff`: File Difference     - `playbook_on_start`: Playbook Started     - `playbook_on_notify`: Running Handlers     - `playbook_on_include`: Including File     - `playbook_on_no_hosts_matched`: No Hosts Matched     - `playbook_on_no_hosts_remaining`: No Hosts Remaining     - `playbook_on_task_start`: Task Started     - `playbook_on_vars_prompt`: Variables Prompted     - `playbook_on_setup`: Gathering Facts     - `playbook_on_import_for_host`: internal: on Import for Host     - `playbook_on_not_import_for_host`: internal: on Not Import for Host     - `playbook_on_play_start`: Play Started     - `playbook_on_stats`: Playbook Complete     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `event_level`:  (integer) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `parent_uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `playbook`:  (string) * `play`:  (string) * `task`:  (string) * `role`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)    ## Sorting  To specify that job events are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsJobEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsJobEventsListOpts struct

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

# **GroupsGroupsJobHostSummariesList**
> GroupsGroupsJobHostSummariesList(ctx, id, optional)
 List Job Host Summaries for a Group

 Make a GET request to this resource to retrieve a list of job host summaries associated with the selected group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job host summaries found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job host summary records.    ## Results  Each job host summary data structure includes the following fields:  * `id`: Database ID for this job host summary. (integer) * `type`: Data type for this job host summary. (choice) * `url`: URL for this job host summary. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job host summary was created. (datetime) * `modified`: Timestamp when this job host summary was last modified. (datetime) * `job`:  (id) * `host`:  (id) * `host_name`:  (string) * `changed`:  (integer) * `dark`:  (integer) * `failures`:  (integer) * `ok`:  (integer) * `processed`:  (integer) * `skipped`:  (integer) * `failed`:  (boolean) * `ignored`:  (integer) * `rescued`:  (integer)    ## Sorting  To specify that job host summaries are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsJobHostSummariesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsJobHostSummariesListOpts struct

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

# **GroupsGroupsList**
> GroupsGroupsList(ctx, optional)
 List Groups

 Make a GET request to this resource to retrieve the list of groups.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsListOpts struct

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

# **GroupsGroupsPartialUpdate**
> GroupsGroupsPartialUpdate(ctx, id, optional)
 Update a Group

 Make a PUT or PATCH request to this resource to update this group.  The following fields may be modified:          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data10**](Data10.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupsPotentialChildrenList**
> GroupsGroupsPotentialChildrenList(ctx, id, optional)
 List Potential Child Groups for a Group

 Make a GET request to this resource to retrieve a list of groups available to be added as children of the current group.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsPotentialChildrenListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsPotentialChildrenListOpts struct

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

# **GroupsGroupsRead**
> GroupsGroupsRead(ctx, id, optional)
 Retrieve a Group

 Make GET request to this resource to retrieve a single group record containing the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsReadOpts struct

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

# **GroupsGroupsUpdate**
> GroupsGroupsUpdate(ctx, id, optional)
 Update a Group

 Make a PUT or PATCH request to this resource to update this group.  The following fields may be modified:          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsUpdateOpts struct

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

# **GroupsGroupsVariableDataPartialUpdate**
> GroupsGroupsVariableDataPartialUpdate(ctx, id, optional)
 Update Group Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsVariableDataPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsVariableDataPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data13**](Data13.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupsVariableDataRead**
> GroupsGroupsVariableDataRead(ctx, id, optional)
 Retrieve Group Variable Data

 Make a GET request to this resource to retrieve all variables defined for a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsVariableDataReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsVariableDataReadOpts struct

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

# **GroupsGroupsVariableDataUpdate**
> GroupsGroupsVariableDataUpdate(ctx, id, optional)
 Update Group Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***GroupsApiGroupsGroupsVariableDataUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsGroupsVariableDataUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data12**](Data12.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

