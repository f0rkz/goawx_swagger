# \InventoriesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoriesInventoriesAccessListList**](InventoriesApi.md#InventoriesInventoriesAccessListList) | **Get** /api/v2/inventories/{id}/access_list/ |  List Users
[**InventoriesInventoriesActivityStreamList**](InventoriesApi.md#InventoriesInventoriesActivityStreamList) | **Get** /api/v2/inventories/{id}/activity_stream/ |  List Activity Streams for an Inventory
[**InventoriesInventoriesAdHocCommandsCreate**](InventoriesApi.md#InventoriesInventoriesAdHocCommandsCreate) | **Post** /api/v2/inventories/{id}/ad_hoc_commands/ |  Create an Ad Hoc Command for an Inventory
[**InventoriesInventoriesAdHocCommandsList**](InventoriesApi.md#InventoriesInventoriesAdHocCommandsList) | **Get** /api/v2/inventories/{id}/ad_hoc_commands/ |  List Ad Hoc Commands for an Inventory
[**InventoriesInventoriesCopyCreate**](InventoriesApi.md#InventoriesInventoriesCopyCreate) | **Post** /api/v2/inventories/{id}/copy/ | 
[**InventoriesInventoriesCopyList**](InventoriesApi.md#InventoriesInventoriesCopyList) | **Get** /api/v2/inventories/{id}/copy/ | 
[**InventoriesInventoriesCreate**](InventoriesApi.md#InventoriesInventoriesCreate) | **Post** /api/v2/inventories/ |  Create an Inventory
[**InventoriesInventoriesDelete**](InventoriesApi.md#InventoriesInventoriesDelete) | **Delete** /api/v2/inventories/{id}/ |  Delete an Inventory
[**InventoriesInventoriesGroupsCreate**](InventoriesApi.md#InventoriesInventoriesGroupsCreate) | **Post** /api/v2/inventories/{id}/groups/ |  Create a Group for an Inventory
[**InventoriesInventoriesGroupsList**](InventoriesApi.md#InventoriesInventoriesGroupsList) | **Get** /api/v2/inventories/{id}/groups/ |  List Groups for an Inventory
[**InventoriesInventoriesHostsCreate**](InventoriesApi.md#InventoriesInventoriesHostsCreate) | **Post** /api/v2/inventories/{id}/hosts/ |  Create a Host for an Inventory
[**InventoriesInventoriesHostsList**](InventoriesApi.md#InventoriesInventoriesHostsList) | **Get** /api/v2/inventories/{id}/hosts/ |  List Hosts for an Inventory
[**InventoriesInventoriesInstanceGroupsCreate**](InventoriesApi.md#InventoriesInventoriesInstanceGroupsCreate) | **Post** /api/v2/inventories/{id}/instance_groups/ |  Create an Instance Group for an Inventory
[**InventoriesInventoriesInstanceGroupsList**](InventoriesApi.md#InventoriesInventoriesInstanceGroupsList) | **Get** /api/v2/inventories/{id}/instance_groups/ |  List Instance Groups for an Inventory
[**InventoriesInventoriesInventorySourcesCreate**](InventoriesApi.md#InventoriesInventoriesInventorySourcesCreate) | **Post** /api/v2/inventories/{id}/inventory_sources/ |  Create an Inventory Source for an Inventory
[**InventoriesInventoriesInventorySourcesList**](InventoriesApi.md#InventoriesInventoriesInventorySourcesList) | **Get** /api/v2/inventories/{id}/inventory_sources/ |  List Inventory Sources for an Inventory
[**InventoriesInventoriesJobTemplatesList**](InventoriesApi.md#InventoriesInventoriesJobTemplatesList) | **Get** /api/v2/inventories/{id}/job_templates/ |  List Job Templates for an Inventory
[**InventoriesInventoriesList**](InventoriesApi.md#InventoriesInventoriesList) | **Get** /api/v2/inventories/ |  List Inventories
[**InventoriesInventoriesObjectRolesList**](InventoriesApi.md#InventoriesInventoriesObjectRolesList) | **Get** /api/v2/inventories/{id}/object_roles/ |  List Roles for an Inventory
[**InventoriesInventoriesPartialUpdate**](InventoriesApi.md#InventoriesInventoriesPartialUpdate) | **Patch** /api/v2/inventories/{id}/ |  Update an Inventory
[**InventoriesInventoriesRead**](InventoriesApi.md#InventoriesInventoriesRead) | **Get** /api/v2/inventories/{id}/ |  Retrieve an Inventory
[**InventoriesInventoriesRootGroupsCreate**](InventoriesApi.md#InventoriesInventoriesRootGroupsCreate) | **Post** /api/v2/inventories/{id}/root_groups/ | 
[**InventoriesInventoriesRootGroupsList**](InventoriesApi.md#InventoriesInventoriesRootGroupsList) | **Get** /api/v2/inventories/{id}/root_groups/ |  List Root Groups for an Inventory
[**InventoriesInventoriesScriptRead**](InventoriesApi.md#InventoriesInventoriesScriptRead) | **Get** /api/v2/inventories/{id}/script/ | Generate inventory group and host data as needed for an inventory script.
[**InventoriesInventoriesTreeRead**](InventoriesApi.md#InventoriesInventoriesTreeRead) | **Get** /api/v2/inventories/{id}/tree/ |  Group Tree for an Inventory
[**InventoriesInventoriesUpdate**](InventoriesApi.md#InventoriesInventoriesUpdate) | **Put** /api/v2/inventories/{id}/ |  Update an Inventory
[**InventoriesInventoriesUpdateInventorySourcesCreate**](InventoriesApi.md#InventoriesInventoriesUpdateInventorySourcesCreate) | **Post** /api/v2/inventories/{id}/update_inventory_sources/ |  Update Inventory Sources
[**InventoriesInventoriesUpdateInventorySourcesRead**](InventoriesApi.md#InventoriesInventoriesUpdateInventorySourcesRead) | **Get** /api/v2/inventories/{id}/update_inventory_sources/ |  Update Inventory Sources
[**InventoriesInventoriesVariableDataPartialUpdate**](InventoriesApi.md#InventoriesInventoriesVariableDataPartialUpdate) | **Patch** /api/v2/inventories/{id}/variable_data/ |  Update Inventory Variable Data
[**InventoriesInventoriesVariableDataRead**](InventoriesApi.md#InventoriesInventoriesVariableDataRead) | **Get** /api/v2/inventories/{id}/variable_data/ |  Retrieve Inventory Variable Data
[**InventoriesInventoriesVariableDataUpdate**](InventoriesApi.md#InventoriesInventoriesVariableDataUpdate) | **Put** /api/v2/inventories/{id}/variable_data/ |  Update Inventory Variable Data


# **InventoriesInventoriesAccessListList**
> InventoriesInventoriesAccessListList(ctx, id, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesAccessListListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesAccessListListOpts struct

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

# **InventoriesInventoriesActivityStreamList**
> InventoriesInventoriesActivityStreamList(ctx, id, optional)
 List Activity Streams for an Inventory

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesActivityStreamListOpts struct

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

# **InventoriesInventoriesAdHocCommandsCreate**
> InventoriesInventoriesAdHocCommandsCreate(ctx, id, optional)
 Create an Ad Hoc Command for an Inventory

 Make a POST request to this resource with the following ad hoc command fields to create a new ad hoc command associated with this inventory.             * `execution_environment`: The container image to be used for execution. (id, default=``)           * `job_type`:  (choice)     - `run`: Run (default)     - `check`: Check  * `limit`:  (string, default=`\"\"`) * `credential`:  (id, default=``) * `module_name`:  (choice)     - `command` (default)     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string, default=`\"\"`) * `forks`:  (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (Normal) (default)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string, default=`\"\"`) * `become_enabled`:  (boolean, default=`False`) * `diff_mode`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesAdHocCommandsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesAdHocCommandsCreateOpts struct

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

# **InventoriesInventoriesAdHocCommandsList**
> InventoriesInventoriesAdHocCommandsList(ctx, id, optional)
 List Ad Hoc Commands for an Inventory

 Make a GET request to this resource to retrieve a list of ad hoc commands associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of ad hoc commands found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more ad hoc command records.    ## Results  Each ad hoc command data structure includes the following fields:  * `id`: Database ID for this ad hoc command. (integer) * `type`: Data type for this ad hoc command. (choice) * `url`: URL for this ad hoc command. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command was created. (datetime) * `modified`: Timestamp when this ad hoc command was last modified. (datetime) * `name`: Name of this ad hoc command. (string) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `controller_node`: The instance that managed the execution environment. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `limit`:  (string) * `credential`:  (id) * `module_name`:  (choice)     - `command`     - `shell`     - `yum`     - `apt`     - `apt_key`     - `apt_repository`     - `apt_rpm`     - `service`     - `group`     - `user`     - `mount`     - `ping`     - `selinux`     - `setup`     - `win_ping`     - `win_service`     - `win_updates`     - `win_group`     - `win_user` * `module_args`:  (string) * `forks`:  (integer) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (string) * `become_enabled`:  (boolean) * `diff_mode`:  (boolean)    ## Sorting  To specify that ad hoc commands are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesAdHocCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesAdHocCommandsListOpts struct

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

# **InventoriesInventoriesCopyCreate**
> InventoriesInventoriesCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data22**](Data22.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoriesInventoriesCopyList**
> InventoriesInventoriesCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesCopyListOpts struct

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

# **InventoriesInventoriesCreate**
> InventoriesInventoriesCreate(ctx, optional)
 Create an Inventory

 Make a POST request to this resource with the following inventory fields to create a new inventory:          * `name`: Name of this inventory. (string, required) * `description`: Optional description of this inventory. (string, default=`\"\"`) * `organization`: Organization containing this inventory. (id, required) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory. (default)     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string, default=`\"\"`) * `variables`: Inventory variables in JSON or YAML format. (json, default=``)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoriesApiInventoriesInventoriesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesCreateOpts struct

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

# **InventoriesInventoriesDelete**
> InventoriesInventoriesDelete(ctx, id, optional)
 Delete an Inventory

 Make a DELETE request to this resource to delete this inventory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesDeleteOpts struct

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

# **InventoriesInventoriesGroupsCreate**
> InventoriesInventoriesGroupsCreate(ctx, id, optional)
 Create a Group for an Inventory

 Make a POST request to this resource with the following group fields to create a new group associated with this inventory.          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`)  * `variables`: Group variables in JSON or YAML format. (json, default=``)         # Remove Inventory Groups:  Make a POST request to this resource with `id` and `disassociate` fields to delete the associated group.      {         \"id\": 123,         \"disassociate\": true     }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesGroupsCreateOpts struct

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

# **InventoriesInventoriesGroupsList**
> InventoriesInventoriesGroupsList(ctx, id, optional)
 List Groups for an Inventory

 Make a GET request to this resource to retrieve a list of groups associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesGroupsListOpts struct

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

# **InventoriesInventoriesHostsCreate**
> InventoriesInventoriesHostsCreate(ctx, id, optional)
 Create a Host for an Inventory

 Make a POST request to this resource with the following host fields to create a new host associated with this inventory.          * `name`: Name of this host. (string, required) * `description`: Optional description of this host. (string, default=`\"\"`)  * `enabled`: Is this host online and available for running jobs? (boolean, default=`True`) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string, default=`\"\"`) * `variables`: Host variables in JSON or YAML format. (json, default=``)              # Remove Inventory Hosts:  Make a POST request to this resource with `id` and `disassociate` fields to delete the associated host.      {         \"id\": 123,         \"disassociate\": true     }

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesHostsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesHostsCreateOpts struct

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

# **InventoriesInventoriesHostsList**
> InventoriesInventoriesHostsList(ctx, id, optional)
 List Hosts for an Inventory

 Make a GET request to this resource to retrieve a list of hosts associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of hosts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more host records.    ## Results  Each host data structure includes the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)    ## Sorting  To specify that hosts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesHostsListOpts struct

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

# **InventoriesInventoriesInstanceGroupsCreate**
> InventoriesInventoriesInstanceGroupsCreate(ctx, id, optional)
 Create an Instance Group for an Inventory

 Make a POST request to this resource with the following instance group fields to create a new instance group associated with this inventory.       * `name`: Name of this instance group. (string, required)          * `is_container_group`: Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. (boolean, default=``) * `credential`:  (id, default=``) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer, default=`0`) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer, default=`0`) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json, default=``) * `pod_spec_override`:  (string, default=`\"\"`)          # Add Instance Groups for an Inventory:  Make a POST request to this resource with only an `id` field to associate an existing instance group with this inventory.  # Remove Instance Groups from this Inventory:  Make a POST request to this resource with `id` and `disassociate` fields to remove the instance group from this inventory  without deleting the instance group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesInstanceGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesInstanceGroupsCreateOpts struct

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

# **InventoriesInventoriesInstanceGroupsList**
> InventoriesInventoriesInstanceGroupsList(ctx, id, optional)
 List Instance Groups for an Inventory

 Make a GET request to this resource to retrieve a list of instance groups associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of instance groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more instance group records.    ## Results  Each instance group data structure includes the following fields:  * `id`: Database ID for this instance group. (integer) * `type`: Data type for this instance group. (choice) * `url`: URL for this instance group. (string) * `related`: Data structure with URLs of related resources. (object) * `name`: Name of this instance group. (string) * `created`: Timestamp when this instance group was created. (datetime) * `modified`: Timestamp when this instance group was last modified. (datetime) * `capacity`:  (field) * `committed_capacity`:  (field) * `consumed_capacity`:  (field) * `percent_capacity_remaining`:  (field) * `jobs_running`: Count of jobs in the running or waiting state that are targeted for this instance group (integer) * `jobs_total`: Count of all jobs that target this instance group (integer) * `instances`:  (field) * `is_container_group`: Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. (boolean) * `credential`:  (id) * `policy_instance_percentage`: Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. (integer) * `policy_instance_minimum`: Static minimum number of Instances that will be automatically assign to this group when new instances come online. (integer) * `policy_instance_list`: List of exact-match Instances that will be assigned to this group (json) * `pod_spec_override`:  (string) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)    ## Sorting  To specify that instance groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesInstanceGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesInstanceGroupsListOpts struct

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

# **InventoriesInventoriesInventorySourcesCreate**
> InventoriesInventoriesInventorySourcesCreate(ctx, id, optional)
 Create an Inventory Source for an Inventory

 Make a POST request to this resource with the following inventory source fields to create a new inventory source associated with this inventory.          * `name`: Name of this inventory source. (string, required) * `description`: Optional description of this inventory source. (string, default=`\"\"`) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string, default=`\"\"`) * `source_vars`: Inventory source variables in YAML or JSON format. (string, default=`\"\"`) * `credential`: Cloud credential to use for inventory updates. (integer, default=`None`) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string, default=`\"\"`) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string, default=`\"\"`) * `host_filter`: Regex where only matching hosts will be imported. (string, default=`\"\"`) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean, default=`False`) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean, default=`False`)  * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO) (default)     - `2`: 2 (DEBUG)     * `execution_environment`: The container image to be used for execution. (id, default=``)  * `update_on_launch`:  (boolean, default=`False`) * `update_cache_timeout`:  (integer, default=`0`) * `source_project`: Project containing inventory file used as source. (id, default=``) * `update_on_project_update`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesInventorySourcesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesInventorySourcesCreateOpts struct

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

# **InventoriesInventoriesInventorySourcesList**
> InventoriesInventoriesInventorySourcesList(ctx, id, optional)
 List Inventory Sources for an Inventory

 Make a GET request to this resource to retrieve a list of inventory sources associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory source records.    ## Results  Each inventory source data structure includes the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that inventory sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesInventorySourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesInventorySourcesListOpts struct

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

# **InventoriesInventoriesJobTemplatesList**
> InventoriesInventoriesJobTemplatesList(ctx, id, optional)
 List Job Templates for an Inventory

 Make a GET request to this resource to retrieve a list of job templates associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of job templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more job template records.    ## Results  Each job template data structure includes the following fields:  * `id`: Database ID for this job template. (integer) * `type`: Data type for this job template. (choice) * `url`: URL for this job template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job template was created. (datetime) * `modified`: Timestamp when this job template was last modified. (datetime) * `name`: Name of this job template. (string) * `description`: Optional description of this job template. (string) * `job_type`:  (choice)     - `run`: Run     - `check`: Check * `inventory`:  (id) * `project`:  (id) * `playbook`:  (string) * `scm_branch`: Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. (string) * `forks`:  (integer) * `limit`:  (string) * `verbosity`:  (choice)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `extra_vars`:  (json) * `job_tags`:  (string) * `force_handlers`:  (boolean) * `skip_tags`:  (string) * `start_at_task`:  (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `use_fact_cache`: If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. (boolean) * `organization`: The organization used to determine access to this template. (id) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated * `execution_environment`: The container image to be used for execution. (id) * `host_config_key`:  (string) * `ask_scm_branch_on_launch`:  (boolean) * `ask_diff_mode_on_launch`:  (boolean) * `ask_variables_on_launch`:  (boolean) * `ask_limit_on_launch`:  (boolean) * `ask_tags_on_launch`:  (boolean) * `ask_skip_tags_on_launch`:  (boolean) * `ask_job_type_on_launch`:  (boolean) * `ask_verbosity_on_launch`:  (boolean) * `ask_inventory_on_launch`:  (boolean) * `ask_credential_on_launch`:  (boolean) * `survey_enabled`:  (boolean) * `become_enabled`:  (boolean) * `diff_mode`: If enabled, textual changes made to any templated files on the host are shown in the standard output (boolean) * `allow_simultaneous`:  (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `job_slice_count`: The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. (integer) * `webhook_service`: Service that webhook requests will be accepted from (choice)     - `\"\"`: ---------     - `github`: GitHub     - `gitlab`: GitLab * `webhook_credential`: Personal Access Token for posting back the status to the service API (id)    ## Sorting  To specify that job templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesJobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesJobTemplatesListOpts struct

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

# **InventoriesInventoriesList**
> InventoriesInventoriesList(ctx, optional)
 List Inventories

 Make a GET request to this resource to retrieve the list of inventories.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventories found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory records.    ## Results  Each inventory data structure includes the following fields:  * `id`: Database ID for this inventory. (integer) * `type`: Data type for this inventory. (choice) * `url`: URL for this inventory. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory was created. (datetime) * `modified`: Timestamp when this inventory was last modified. (datetime) * `name`: Name of this inventory. (string) * `description`: Optional description of this inventory. (string) * `organization`: Organization containing this inventory. (id) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory.     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string) * `variables`: Inventory variables in JSON or YAML format. (json) * `has_active_failures`: This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed. (boolean) * `total_hosts`: This field is deprecated and will be removed in a future release. Total number of hosts in this inventory. (integer) * `hosts_with_active_failures`: This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures. (integer) * `total_groups`: This field is deprecated and will be removed in a future release. Total number of groups in this inventory. (integer) * `has_inventory_sources`: This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources. (boolean) * `total_inventory_sources`: Total number of external inventory sources configured within this inventory. (integer) * `inventory_sources_with_failures`: Number of external inventory sources in this inventory with failures. (integer) * `pending_deletion`: Flag indicating the inventory is being deleted. (boolean)    ## Sorting  To specify that inventories are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventoriesApiInventoriesInventoriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesListOpts struct

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

# **InventoriesInventoriesObjectRolesList**
> InventoriesInventoriesObjectRolesList(ctx, id, optional)
 List Roles for an Inventory

 Make a GET request to this resource to retrieve a list of roles associated with the selected inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesObjectRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesObjectRolesListOpts struct

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

# **InventoriesInventoriesPartialUpdate**
> InventoriesInventoriesPartialUpdate(ctx, id, optional)
 Update an Inventory

 Make a PUT or PATCH request to this resource to update this inventory.  The following fields may be modified:          * `name`: Name of this inventory. (string, required) * `description`: Optional description of this inventory. (string, default=`\"\"`) * `organization`: Organization containing this inventory. (id, required) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory. (default)     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string, default=`\"\"`) * `variables`: Inventory variables in JSON or YAML format. (json, default=``)                 For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesPartialUpdateOpts struct

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

# **InventoriesInventoriesRead**
> InventoriesInventoriesRead(ctx, id, optional)
 Retrieve an Inventory

 Make GET request to this resource to retrieve a single inventory record containing the following fields:  * `id`: Database ID for this inventory. (integer) * `type`: Data type for this inventory. (choice) * `url`: URL for this inventory. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory was created. (datetime) * `modified`: Timestamp when this inventory was last modified. (datetime) * `name`: Name of this inventory. (string) * `description`: Optional description of this inventory. (string) * `organization`: Organization containing this inventory. (id) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory.     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string) * `variables`: Inventory variables in JSON or YAML format. (json) * `has_active_failures`: This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed. (boolean) * `total_hosts`: This field is deprecated and will be removed in a future release. Total number of hosts in this inventory. (integer) * `hosts_with_active_failures`: This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures. (integer) * `total_groups`: This field is deprecated and will be removed in a future release. Total number of groups in this inventory. (integer) * `has_inventory_sources`: This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources. (boolean) * `total_inventory_sources`: Total number of external inventory sources configured within this inventory. (integer) * `inventory_sources_with_failures`: Number of external inventory sources in this inventory with failures. (integer) * `pending_deletion`: Flag indicating the inventory is being deleted. (boolean)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesReadOpts struct

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

# **InventoriesInventoriesRootGroupsCreate**
> InventoriesInventoriesRootGroupsCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesRootGroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesRootGroupsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data23**](Data23.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoriesInventoriesRootGroupsList**
> InventoriesInventoriesRootGroupsList(ctx, id, optional)
 List Root Groups for an Inventory

 Make a GET request to this resource to retrieve a list of root (top-level) groups associated with this inventory.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesRootGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesRootGroupsListOpts struct

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

# **InventoriesInventoriesScriptRead**
> InventoriesInventoriesScriptRead(ctx, id)
Generate inventory group and host data as needed for an inventory script.

 Refer to [Dynamic Inventory](http://docs.ansible.com/intro_dynamic_inventory.html) for more information on inventory scripts.  ## List Response  Make a GET request to this resource without query parameters to retrieve a JSON object containing groups, including the hosts, children and variables for each group.  The response data is equivalent to that returned by passing the `--list` argument to an inventory script.  Specify a query string of `?hostvars=1` to retrieve the JSON object above including all host variables.  The `['_meta']['hostvars']` object in the response contains an entry for each host with its variables.  This response format can be used with Ansible 1.3 and later to avoid making a separate API request for each host.  Refer to [Tuning the External Inventory Script](http://docs.ansible.com/developing_inventory.html#tuning-the-external-inventory-script) for more information on this feature.  By default, the inventory script will only return hosts that are enabled in the inventory.  This feature allows disabled hosts to be skipped when running jobs without removing them from the inventory.  Specify a query string of `?all=1` to return all hosts, including disabled ones.  Specify a query string of `?towervars=1` to add variables to the hostvars of each host that specifies its enabled state and database ID.  Specify a query string of `?subset=slice2of5` to produce an inventory that has a restricted number of hosts according to the rules of job slicing.  To apply multiple query strings, join them with the `&` character, like `?hostvars=1&all=1`.  ## Host Response  Make a GET request to this resource with a query string similar to `?host=HOSTNAME` to retrieve a JSON object containing host variables for the specified host.  The response data is equivalent to that returned by passing the `--host HOSTNAME` argument to an inventory script.

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoriesInventoriesTreeRead**
> InventoriesInventoriesTreeRead(ctx, id)
 Group Tree for an Inventory

 Make a GET request to this resource to retrieve a hierarchical view of groups associated with the selected inventory.  The resulting data structure contains a list of root groups, with each group also containing a list of its children.  ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json) * `children`:  (field)

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

# **InventoriesInventoriesUpdate**
> InventoriesInventoriesUpdate(ctx, id, optional)
 Update an Inventory

 Make a PUT or PATCH request to this resource to update this inventory.  The following fields may be modified:          * `name`: Name of this inventory. (string, required) * `description`: Optional description of this inventory. (string, default=`\"\"`) * `organization`: Organization containing this inventory. (id, required) * `kind`: Kind of inventory being represented. (choice)     - `\"\"`: Hosts have a direct link to this inventory. (default)     - `smart`: Hosts for inventory generated using the host_filter property. * `host_filter`: Filter that will be applied to the hosts of this inventory. (string, default=`\"\"`) * `variables`: Inventory variables in JSON or YAML format. (json, default=``)               For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesUpdateOpts struct

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

# **InventoriesInventoriesUpdateInventorySourcesCreate**
> InventoriesInventoriesUpdateInventorySourcesCreate(ctx, id)
 Update Inventory Sources

 Make a GET request to this resource to determine if any of the inventory sources for this inventory can be updated. The response will include the following fields for each inventory source:  * `inventory_source`: ID of the inventory_source   (integer, read-only) * `can_update`: Flag indicating if this inventory source can be updated   (boolean, read-only)  Make a POST request to this resource to update the inventory sources. The response status code will be a 202. The response will contain the follow fields for each of the individual inventory sources:  * `status`: `started` or message why the update could not be started.   (string, read-only) * `inventory_update`: ID of the inventory update job that was started.   (integer, read-only) * `project_update`: ID of the project update job that was started if this inventory source is an SCM source.   (interger, read-only, optional)

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

# **InventoriesInventoriesUpdateInventorySourcesRead**
> InventoriesInventoriesUpdateInventorySourcesRead(ctx, id, optional)
 Update Inventory Sources

 Make a GET request to this resource to determine if any of the inventory sources for this inventory can be updated. The response will include the following fields for each inventory source:  * `inventory_source`: ID of the inventory_source   (integer, read-only) * `can_update`: Flag indicating if this inventory source can be updated   (boolean, read-only)  Make a POST request to this resource to update the inventory sources. The response status code will be a 202. The response will contain the follow fields for each of the individual inventory sources:  * `status`: `started` or message why the update could not be started.   (string, read-only) * `inventory_update`: ID of the inventory update job that was started.   (integer, read-only) * `project_update`: ID of the project update job that was started if this inventory source is an SCM source.   (interger, read-only, optional)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesUpdateInventorySourcesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesUpdateInventorySourcesReadOpts struct

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

# **InventoriesInventoriesVariableDataPartialUpdate**
> InventoriesInventoriesVariableDataPartialUpdate(ctx, id, optional)
 Update Inventory Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a inventory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesVariableDataPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesVariableDataPartialUpdateOpts struct

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

# **InventoriesInventoriesVariableDataRead**
> InventoriesInventoriesVariableDataRead(ctx, id, optional)
 Retrieve Inventory Variable Data

 Make a GET request to this resource to retrieve all variables defined for a inventory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesVariableDataReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesVariableDataReadOpts struct

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

# **InventoriesInventoriesVariableDataUpdate**
> InventoriesInventoriesVariableDataUpdate(ctx, id, optional)
 Update Inventory Variable Data

 Make a PUT or PATCH request to this resource to update variables defined for a inventory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventoriesApiInventoriesInventoriesVariableDataUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoriesApiInventoriesInventoriesVariableDataUpdateOpts struct

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

