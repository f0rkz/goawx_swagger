# \InventorySourcesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventorySourcesInventorySourcesActivityStreamList**](InventorySourcesApi.md#InventorySourcesInventorySourcesActivityStreamList) | **Get** /api/v2/inventory_sources/{id}/activity_stream/ |  List Activity Streams for an Inventory Source
[**InventorySourcesInventorySourcesCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesCreate) | **Post** /api/v2/inventory_sources/ |  Create an Inventory Source
[**InventorySourcesInventorySourcesCredentialsCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesCredentialsCreate) | **Post** /api/v2/inventory_sources/{id}/credentials/ |  Create a Credential for an Inventory Source
[**InventorySourcesInventorySourcesCredentialsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesCredentialsList) | **Get** /api/v2/inventory_sources/{id}/credentials/ |  List Credentials for an Inventory Source
[**InventorySourcesInventorySourcesDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesDelete) | **Delete** /api/v2/inventory_sources/{id}/ |  Delete an Inventory Source
[**InventorySourcesInventorySourcesGroupsDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesGroupsDelete) | **Delete** /api/v2/inventory_sources/{id}/groups/ |  Create a Group for an Inventory Source
[**InventorySourcesInventorySourcesGroupsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesGroupsList) | **Get** /api/v2/inventory_sources/{id}/groups/ |  List Groups for an Inventory Source
[**InventorySourcesInventorySourcesHostsDelete**](InventorySourcesApi.md#InventorySourcesInventorySourcesHostsDelete) | **Delete** /api/v2/inventory_sources/{id}/hosts/ |  Create a Host for an Inventory Source
[**InventorySourcesInventorySourcesHostsList**](InventorySourcesApi.md#InventorySourcesInventorySourcesHostsList) | **Get** /api/v2/inventory_sources/{id}/hosts/ |  List Hosts for an Inventory Source
[**InventorySourcesInventorySourcesInventoryUpdatesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesInventoryUpdatesList) | **Get** /api/v2/inventory_sources/{id}/inventory_updates/ |  List Inventory Updates for an Inventory Source
[**InventorySourcesInventorySourcesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesList) | **Get** /api/v2/inventory_sources/ |  List Inventory Sources
[**InventorySourcesInventorySourcesNotificationTemplatesErrorCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesErrorCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_error/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesErrorList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesErrorList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_error/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesStartedCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesStartedCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_started/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesStartedList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesStartedList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_started/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate) | **Post** /api/v2/inventory_sources/{id}/notification_templates_success/ |  Create a Notification Template for an Inventory Source
[**InventorySourcesInventorySourcesNotificationTemplatesSuccessList**](InventorySourcesApi.md#InventorySourcesInventorySourcesNotificationTemplatesSuccessList) | **Get** /api/v2/inventory_sources/{id}/notification_templates_success/ |  List Notification Templates for an Inventory Source
[**InventorySourcesInventorySourcesPartialUpdate**](InventorySourcesApi.md#InventorySourcesInventorySourcesPartialUpdate) | **Patch** /api/v2/inventory_sources/{id}/ |  Update an Inventory Source
[**InventorySourcesInventorySourcesRead**](InventorySourcesApi.md#InventorySourcesInventorySourcesRead) | **Get** /api/v2/inventory_sources/{id}/ |  Retrieve an Inventory Source
[**InventorySourcesInventorySourcesSchedulesCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesSchedulesCreate) | **Post** /api/v2/inventory_sources/{id}/schedules/ |  Create a Schedule for an Inventory Source
[**InventorySourcesInventorySourcesSchedulesList**](InventorySourcesApi.md#InventorySourcesInventorySourcesSchedulesList) | **Get** /api/v2/inventory_sources/{id}/schedules/ |  List Schedules for an Inventory Source
[**InventorySourcesInventorySourcesUpdate0**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdate0) | **Put** /api/v2/inventory_sources/{id}/ |  Update an Inventory Source
[**InventorySourcesInventorySourcesUpdateCreate**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdateCreate) | **Post** /api/v2/inventory_sources/{id}/update/ |  Update Inventory Source
[**InventorySourcesInventorySourcesUpdateRead**](InventorySourcesApi.md#InventorySourcesInventorySourcesUpdateRead) | **Get** /api/v2/inventory_sources/{id}/update/ |  Update Inventory Source


# **InventorySourcesInventorySourcesActivityStreamList**
> InventorySourcesInventorySourcesActivityStreamList(ctx, id, optional)
 List Activity Streams for an Inventory Source

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesActivityStreamListOpts struct

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

# **InventorySourcesInventorySourcesCreate**
> InventorySourcesInventorySourcesCreate(ctx, optional)
 Create an Inventory Source

 Make a POST request to this resource with the following inventory source fields to create a new inventory source:          * `name`: Name of this inventory source. (string, required) * `description`: Optional description of this inventory source. (string, default=`\"\"`) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string, default=`\"\"`) * `source_vars`: Inventory source variables in YAML or JSON format. (string, default=`\"\"`) * `credential`: Cloud credential to use for inventory updates. (integer, default=`None`) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string, default=`\"\"`) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string, default=`\"\"`) * `host_filter`: Regex where only matching hosts will be imported. (string, default=`\"\"`) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean, default=`False`) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean, default=`False`)  * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO) (default)     - `2`: 2 (DEBUG)     * `execution_environment`: The container image to be used for execution. (id, default=``) * `inventory`:  (id, required) * `update_on_launch`:  (boolean, default=`False`) * `update_cache_timeout`:  (integer, default=`0`) * `source_project`: Project containing inventory file used as source. (id, default=``) * `update_on_project_update`:  (boolean, default=`False`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesCreateOpts struct

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

# **InventorySourcesInventorySourcesCredentialsCreate**
> InventorySourcesInventorySourcesCredentialsCreate(ctx, id, optional)
 Create a Credential for an Inventory Source

 Make a POST request to this resource with the following credential fields to create a new credential associated with this inventory source.          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`:  (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)            # Add Credentials for an Inventory Source:  Make a POST request to this resource with only an `id` field to associate an existing credential with this inventory source.  # Remove Credentials from this Inventory Source:  Make a POST request to this resource with `id` and `disassociate` fields to remove the credential from this inventory source  without deleting the credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesCredentialsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data25**](Data25.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventorySourcesInventorySourcesCredentialsList**
> InventorySourcesInventorySourcesCredentialsList(ctx, id, optional)
 List Credentials for an Inventory Source

 Make a GET request to this resource to retrieve a list of credentials associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`:  (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)    ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesCredentialsListOpts struct

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

# **InventorySourcesInventorySourcesDelete**
> InventorySourcesInventorySourcesDelete(ctx, id, optional)
 Delete an Inventory Source

 Make a DELETE request to this resource to delete this inventory source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesDeleteOpts struct

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

# **InventorySourcesInventorySourcesGroupsDelete**
> InventorySourcesInventorySourcesGroupsDelete(ctx, id, optional)
 Create a Group for an Inventory Source

 Make a POST request to this resource with the following group fields to create a new group associated with this inventory source.          * `name`: Name of this group. (string, required) * `description`: Optional description of this group. (string, default=`\"\"`) * `inventory`:  (id, required) * `variables`: Group variables in JSON or YAML format. (json, default=``)          # Delete all groups of this Inventory Source:  Make a DELETE request to this resource to delete all groups show in the list. The Inventory Source will not be deleted by this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesGroupsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesGroupsDeleteOpts struct

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

# **InventorySourcesInventorySourcesGroupsList**
> InventorySourcesInventorySourcesGroupsList(ctx, id, optional)
 List Groups for an Inventory Source

 Make a GET request to this resource to retrieve a list of groups associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of groups found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more group records.    ## Results  Each group data structure includes the following fields:  * `id`: Database ID for this group. (integer) * `type`: Data type for this group. (choice) * `url`: URL for this group. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this group was created. (datetime) * `modified`: Timestamp when this group was last modified. (datetime) * `name`: Name of this group. (string) * `description`: Optional description of this group. (string) * `inventory`:  (id) * `variables`: Group variables in JSON or YAML format. (json)    ## Sorting  To specify that groups are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesGroupsListOpts struct

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

# **InventorySourcesInventorySourcesHostsDelete**
> InventorySourcesInventorySourcesHostsDelete(ctx, id, optional)
 Create a Host for an Inventory Source

 Make a POST request to this resource with the following host fields to create a new host associated with this inventory source.          * `name`: Name of this host. (string, required) * `description`: Optional description of this host. (string, default=`\"\"`) * `inventory`:  (id, required) * `enabled`: Is this host online and available for running jobs? (boolean, default=`True`) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string, default=`\"\"`) * `variables`: Host variables in JSON or YAML format. (json, default=``)               # Delete all hosts of this Inventory Source:  Make a DELETE request to this resource to delete all hosts show in the list. The Inventory Source will not be deleted by this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesHostsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesHostsDeleteOpts struct

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

# **InventorySourcesInventorySourcesHostsList**
> InventorySourcesInventorySourcesHostsList(ctx, id, optional)
 List Hosts for an Inventory Source

 Make a GET request to this resource to retrieve a list of hosts associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of hosts found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more host records.    ## Results  Each host data structure includes the following fields:  * `id`: Database ID for this host. (integer) * `type`: Data type for this host. (choice) * `url`: URL for this host. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this host was created. (datetime) * `modified`: Timestamp when this host was last modified. (datetime) * `name`: Name of this host. (string) * `description`: Optional description of this host. (string) * `inventory`:  (id) * `enabled`: Is this host online and available for running jobs? (boolean) * `instance_id`: The value used by the remote inventory source to uniquely identify the host (string) * `variables`: Host variables in JSON or YAML format. (json) * `has_active_failures`:  (field) * `has_inventory_sources`:  (field) * `last_job`:  (id) * `last_job_host_summary`:  (id) * `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)    ## Sorting  To specify that hosts are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesHostsListOpts struct

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

# **InventorySourcesInventorySourcesInventoryUpdatesList**
> InventorySourcesInventorySourcesInventoryUpdatesList(ctx, id, optional)
 List Inventory Updates for an Inventory Source

 Make a GET request to this resource to retrieve a list of inventory updates associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory updates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory update records.    ## Results  Each inventory update data structure includes the following fields:  * `id`: Database ID for this inventory update. (integer) * `type`: Data type for this inventory update. (choice) * `url`: URL for this inventory update. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory update was created. (datetime) * `modified`: Timestamp when this inventory update was last modified. (datetime) * `name`: Name of this inventory update. (string) * `description`: Optional description of this inventory update. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `execution_node`: The node the job executed on. (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`:  (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `inventory`:  (id) * `inventory_source`:  (id) * `license_error`:  (boolean) * `org_host_limit_error`:  (boolean) * `source_project_update`: Inventory files from this Project Update were used for the inventory update. (id) * `instance_group`: The Instance group the job was run under (id)    ## Sorting  To specify that inventory updates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesInventoryUpdatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesInventoryUpdatesListOpts struct

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

# **InventorySourcesInventorySourcesList**
> InventorySourcesInventorySourcesList(ctx, optional)
 List Inventory Sources

 Make a GET request to this resource to retrieve the list of inventory sources.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of inventory sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more inventory source records.    ## Results  Each inventory source data structure includes the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that inventory sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesListOpts struct

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

# **InventorySourcesInventorySourcesNotificationTemplatesErrorCreate**
> InventorySourcesInventorySourcesNotificationTemplatesErrorCreate(ctx, id, optional)
 Create a Notification Template for an Inventory Source

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this inventory source.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for an Inventory Source:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this inventory source.  # Remove Notification Templates from this Inventory Source:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this inventory source  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesErrorCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesErrorCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data26**](Data26.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventorySourcesInventorySourcesNotificationTemplatesErrorList**
> InventorySourcesInventorySourcesNotificationTemplatesErrorList(ctx, id, optional)
 List Notification Templates for an Inventory Source

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesErrorListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesErrorListOpts struct

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

# **InventorySourcesInventorySourcesNotificationTemplatesStartedCreate**
> InventorySourcesInventorySourcesNotificationTemplatesStartedCreate(ctx, id, optional)
 Create a Notification Template for an Inventory Source

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this inventory source.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for an Inventory Source:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this inventory source.  # Remove Notification Templates from this Inventory Source:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this inventory source  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesStartedCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesStartedCreateOpts struct

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

# **InventorySourcesInventorySourcesNotificationTemplatesStartedList**
> InventorySourcesInventorySourcesNotificationTemplatesStartedList(ctx, id, optional)
 List Notification Templates for an Inventory Source

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesStartedListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesStartedListOpts struct

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

# **InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate**
> InventorySourcesInventorySourcesNotificationTemplatesSuccessCreate(ctx, id, optional)
 Create a Notification Template for an Inventory Source

 Make a POST request to this resource with the following notification template fields to create a new notification template associated with this inventory source.          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         # Add Notification Templates for an Inventory Source:  Make a POST request to this resource with only an `id` field to associate an existing notification template with this inventory source.  # Remove Notification Templates from this Inventory Source:  Make a POST request to this resource with `id` and `disassociate` fields to remove the notification template from this inventory source  without deleting the notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesSuccessCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesSuccessCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data27**](Data27.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventorySourcesInventorySourcesNotificationTemplatesSuccessList**
> InventorySourcesInventorySourcesNotificationTemplatesSuccessList(ctx, id, optional)
 List Notification Templates for an Inventory Source

 Make a GET request to this resource to retrieve a list of notification templates associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesSuccessListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesNotificationTemplatesSuccessListOpts struct

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

# **InventorySourcesInventorySourcesPartialUpdate**
> InventorySourcesInventorySourcesPartialUpdate(ctx, id, optional)
 Update an Inventory Source

 Make a PUT or PATCH request to this resource to update this inventory source.  The following fields may be modified:          * `name`: Name of this inventory source. (string, required) * `description`: Optional description of this inventory source. (string, default=`\"\"`) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string, default=`\"\"`) * `source_vars`: Inventory source variables in YAML or JSON format. (string, default=`\"\"`) * `credential`: Cloud credential to use for inventory updates. (integer, default=`None`) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string, default=`\"\"`) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string, default=`\"\"`) * `host_filter`: Regex where only matching hosts will be imported. (string, default=`\"\"`) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean, default=`False`) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean, default=`False`)  * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO) (default)     - `2`: 2 (DEBUG)     * `execution_environment`: The container image to be used for execution. (id, default=``) * `inventory`:  (id, required) * `update_on_launch`:  (boolean, default=`False`) * `update_cache_timeout`:  (integer, default=`0`) * `source_project`: Project containing inventory file used as source. (id, default=``) * `update_on_project_update`:  (boolean, default=`False`)           For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesPartialUpdateOpts struct

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

# **InventorySourcesInventorySourcesRead**
> InventorySourcesInventorySourcesRead(ctx, id, optional)
 Retrieve an Inventory Source

 Make GET request to this resource to retrieve a single inventory source record containing the following fields:  * `id`: Database ID for this inventory source. (integer) * `type`: Data type for this inventory source. (choice) * `url`: URL for this inventory source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this inventory source was created. (datetime) * `modified`: Timestamp when this inventory source was last modified. (datetime) * `name`: Name of this inventory source. (string) * `description`: Optional description of this inventory source. (string) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string) * `source_vars`: Inventory source variables in YAML or JSON format. (string) * `credential`: Cloud credential to use for inventory updates. (integer) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string) * `host_filter`: Regex where only matching hosts will be imported. (string) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO)     - `2`: 2 (DEBUG) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `none`: No External Source * `execution_environment`: The container image to be used for execution. (id) * `inventory`:  (id) * `update_on_launch`:  (boolean) * `update_cache_timeout`:  (integer) * `source_project`: Project containing inventory file used as source. (id) * `update_on_project_update`:  (boolean) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesReadOpts struct

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

# **InventorySourcesInventorySourcesSchedulesCreate**
> InventorySourcesInventorySourcesSchedulesCreate(ctx, id, optional)
 Create a Schedule for an Inventory Source

 Make a POST request to this resource with the following schedule fields to create a new schedule associated with this inventory source.   * `rrule`: A value representing the schedules iCal recurrence rule. (string, required)        * `name`: Name of this schedule. (string, required) * `description`: Optional description of this schedule. (string, default=`\"\"`) * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug)  * `enabled`: Enables processing of this schedule. (boolean, default=`True`)            POST requests to this resource must include a proper `rrule` value following a particular format and conforming to subset of allowed rules.  The following lists the expected format and details of our rrules:  * DTSTART is required and must follow the following format: DTSTART:YYYYMMDDTHHMMSSZ * DTSTART is expected to be in UTC * INTERVAL is required * SECONDLY is not supported * RRULE must precede the rule statements * BYDAY is supported but not BYDAY with a numerical prefix * BYYEARDAY and BYWEEKNO are not supported * Only one rrule statement per schedule is supported * COUNT must be < 1000  Here are some example rrules:      \"DTSTART:20500331T055000Z RRULE:FREQ=MINUTELY;INTERVAL=10;COUNT=5\"     \"DTSTART:20240331T075000Z RRULE:FREQ=DAILY;INTERVAL=1;COUNT=1\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MINUTELY;INTERVAL=1;UNTIL=20230401T075000Z\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=5;BYDAY=MO\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYMONTHDAY=6\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=4;BYDAY=SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=-1;BYDAY=MO,TU,WE,TH,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=MONTHLY;INTERVAL=1;BYSETPOS=-1;BYDAY=MO,TU,WE,TH,FR,SA,SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=YEARLY;INTERVAL=1;BYMONTH=4;BYMONTHDAY=1\"     \"DTSTART:20140331T075000Z RRULE:FREQ=YEARLY;INTERVAL=1;BYSETPOS=-1;BYMONTH=8;BYDAY=SU\"     \"DTSTART:20140331T075000Z RRULE:FREQ=WEEKLY;INTERVAL=1;UNTIL=20230401T075000Z;BYDAY=MO,WE,FR\"     \"DTSTART:20140331T075000Z RRULE:FREQ=HOURLY;INTERVAL=1;UNTIL=20230610T075000Z\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesSchedulesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesSchedulesCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data28**](Data28.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventorySourcesInventorySourcesSchedulesList**
> InventorySourcesInventorySourcesSchedulesList(ctx, id, optional)
 List Schedules for an Inventory Source

 Make a GET request to this resource to retrieve a list of schedules associated with the selected inventory source.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of schedules found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more schedule records.    ## Results  Each schedule data structure includes the following fields:  * `rrule`: A value representing the schedules iCal recurrence rule. (string) * `id`: Database ID for this schedule. (integer) * `type`: Data type for this schedule. (choice) * `url`: URL for this schedule. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this schedule was created. (datetime) * `modified`: Timestamp when this schedule was last modified. (datetime) * `name`: Name of this schedule. (string) * `description`: Optional description of this schedule. (string) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `unified_job_template`:  (id) * `enabled`: Enables processing of this schedule. (boolean) * `dtstart`: The first occurrence of the schedule occurs on or after this time. (datetime) * `dtend`: The last occurrence of the schedule occurs before this time, aftewards the schedule expires. (datetime) * `next_run`: The next time that the scheduled action will run. (datetime) * `timezone`:  (field) * `until`:  (field)    ## Sorting  To specify that schedules are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesSchedulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesSchedulesListOpts struct

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

# **InventorySourcesInventorySourcesUpdate0**
> InventorySourcesInventorySourcesUpdate0(ctx, id, optional)
 Update an Inventory Source

 Make a PUT or PATCH request to this resource to update this inventory source.  The following fields may be modified:          * `name`: Name of this inventory source. (string, required) * `description`: Optional description of this inventory source. (string, default=`\"\"`) * `source`:  (choice)     - `file`: File, Directory or Script     - `scm`: Sourced from a Project     - `ec2`: Amazon EC2     - `gce`: Google Compute Engine     - `azure_rm`: Microsoft Azure Resource Manager     - `vmware`: VMware vCenter     - `satellite6`: Red Hat Satellite 6     - `openstack`: OpenStack     - `rhv`: Red Hat Virtualization     - `controller`: Red Hat Ansible Automation Platform     - `insights`: Red Hat Insights * `source_path`:  (string, default=`\"\"`) * `source_vars`: Inventory source variables in YAML or JSON format. (string, default=`\"\"`) * `credential`: Cloud credential to use for inventory updates. (integer, default=`None`) * `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string, default=`\"\"`) * `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2018-02-01T08:00:00.000000Z:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string, default=`\"\"`) * `host_filter`: Regex where only matching hosts will be imported. (string, default=`\"\"`) * `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean, default=`False`) * `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean, default=`False`)  * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`) * `verbosity`:  (choice)     - `0`: 0 (WARNING)     - `1`: 1 (INFO) (default)     - `2`: 2 (DEBUG)     * `execution_environment`: The container image to be used for execution. (id, default=``) * `inventory`:  (id, required) * `update_on_launch`:  (boolean, default=`False`) * `update_cache_timeout`:  (integer, default=`0`) * `source_project`: Project containing inventory file used as source. (id, default=``) * `update_on_project_update`:  (boolean, default=`False`)         For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesUpdate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesUpdate0Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data24**](Data24.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventorySourcesInventorySourcesUpdateCreate**
> InventorySourcesInventorySourcesUpdateCreate(ctx, id)
 Update Inventory Source

 Make a GET request to this resource to determine if the group can be updated from its inventory source.  The response will include the following field:  * `can_update`: Flag indicating if this inventory source can be updated   (boolean, read-only)  Make a POST request to this resource to update the inventory source.  If successful, the response status code will be 202.  If the inventory source is not defined or cannot be updated, a 405 status code will be returned.

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

# **InventorySourcesInventorySourcesUpdateRead**
> InventorySourcesInventorySourcesUpdateRead(ctx, id, optional)
 Update Inventory Source

 Make a GET request to this resource to determine if the group can be updated from its inventory source.  The response will include the following field:  * `can_update`: Flag indicating if this inventory source can be updated   (boolean, read-only)  Make a POST request to this resource to update the inventory source.  If successful, the response status code will be 202.  If the inventory source is not defined or cannot be updated, a 405 status code will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***InventorySourcesApiInventorySourcesInventorySourcesUpdateReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventorySourcesApiInventorySourcesInventorySourcesUpdateReadOpts struct

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

