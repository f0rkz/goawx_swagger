# \ExecutionEnvironmentsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecutionEnvironmentsExecutionEnvironmentsActivityStreamList**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsActivityStreamList) | **Get** /api/v2/execution_environments/{id}/activity_stream/ |  List Activity Streams for an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsCopyCreate**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsCopyCreate) | **Post** /api/v2/execution_environments/{id}/copy/ | 
[**ExecutionEnvironmentsExecutionEnvironmentsCopyList**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsCopyList) | **Get** /api/v2/execution_environments/{id}/copy/ | 
[**ExecutionEnvironmentsExecutionEnvironmentsCreate**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsCreate) | **Post** /api/v2/execution_environments/ |  Create an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsDelete**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsDelete) | **Delete** /api/v2/execution_environments/{id}/ |  Delete an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsList**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsList) | **Get** /api/v2/execution_environments/ |  List Execution Environments
[**ExecutionEnvironmentsExecutionEnvironmentsPartialUpdate**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsPartialUpdate) | **Patch** /api/v2/execution_environments/{id}/ |  Update an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsRead**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsRead) | **Get** /api/v2/execution_environments/{id}/ |  Retrieve an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesList**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesList) | **Get** /api/v2/execution_environments/{id}/unified_job_templates/ |  List Unified Job Templates for an Execution Environment
[**ExecutionEnvironmentsExecutionEnvironmentsUpdate**](ExecutionEnvironmentsApi.md#ExecutionEnvironmentsExecutionEnvironmentsUpdate) | **Put** /api/v2/execution_environments/{id}/ |  Update an Execution Environment


# **ExecutionEnvironmentsExecutionEnvironmentsActivityStreamList**
> ExecutionEnvironmentsExecutionEnvironmentsActivityStreamList(ctx, id, optional)
 List Activity Streams for an Execution Environment

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected execution environment.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsCopyCreate**
> ExecutionEnvironmentsExecutionEnvironmentsCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data8**](Data8.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutionEnvironmentsExecutionEnvironmentsCopyList**
> ExecutionEnvironmentsExecutionEnvironmentsCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCopyListOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsCreate**
> ExecutionEnvironmentsExecutionEnvironmentsCreate(ctx, optional)
 Create an Execution Environment

 Make a POST request to this resource with the following execution environment fields to create a new execution environment:          * `name`: Name of this execution environment. (string, required) * `description`: Optional description of this execution environment. (string, default=`\"\"`) * `organization`: The organization used to determine access to this execution environment. (id, default=``) * `image`: The full image location, including the container registry, image name, and version tag. (string, required)  * `credential`:  (id, default=``) * `pull`: Pull image before running? (choice)     - `\"\"`: --------- (default)     - `always`: Always pull container before running.     - `missing`: Only pull the image if not present before running.     - `never`: Never pull container before running.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data5**](Data5.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutionEnvironmentsExecutionEnvironmentsDelete**
> ExecutionEnvironmentsExecutionEnvironmentsDelete(ctx, id, optional)
 Delete an Execution Environment

 Make a DELETE request to this resource to delete this execution environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsDeleteOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsList**
> ExecutionEnvironmentsExecutionEnvironmentsList(ctx, optional)
 List Execution Environments

 Make a GET request to this resource to retrieve the list of execution environments.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of execution environments found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more execution environment records.    ## Results  Each execution environment data structure includes the following fields:  * `id`: Database ID for this execution environment. (integer) * `type`: Data type for this execution environment. (choice) * `url`: URL for this execution environment. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this execution environment was created. (datetime) * `modified`: Timestamp when this execution environment was last modified. (datetime) * `name`: Name of this execution environment. (string) * `description`: Optional description of this execution environment. (string) * `organization`: The organization used to determine access to this execution environment. (id) * `image`: The full image location, including the container registry, image name, and version tag. (string) * `managed`:  (boolean) * `credential`:  (id) * `pull`: Pull image before running? (choice)     - `\"\"`: ---------     - `always`: Always pull container before running.     - `missing`: Only pull the image if not present before running.     - `never`: Never pull container before running.    ## Sorting  To specify that execution environments are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsListOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsPartialUpdate**
> ExecutionEnvironmentsExecutionEnvironmentsPartialUpdate(ctx, id, optional)
 Update an Execution Environment

 Make a PUT or PATCH request to this resource to update this execution environment.  The following fields may be modified:          * `name`: Name of this execution environment. (string, required) * `description`: Optional description of this execution environment. (string, default=`\"\"`) * `organization`: The organization used to determine access to this execution environment. (id, default=``) * `image`: The full image location, including the container registry, image name, and version tag. (string, required)  * `credential`:  (id, default=``) * `pull`: Pull image before running? (choice)     - `\"\"`: --------- (default)     - `always`: Always pull container before running.     - `missing`: Only pull the image if not present before running.     - `never`: Never pull container before running.         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data7**](Data7.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutionEnvironmentsExecutionEnvironmentsRead**
> ExecutionEnvironmentsExecutionEnvironmentsRead(ctx, id, optional)
 Retrieve an Execution Environment

 Make GET request to this resource to retrieve a single execution environment record containing the following fields:  * `id`: Database ID for this execution environment. (integer) * `type`: Data type for this execution environment. (choice) * `url`: URL for this execution environment. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this execution environment was created. (datetime) * `modified`: Timestamp when this execution environment was last modified. (datetime) * `name`: Name of this execution environment. (string) * `description`: Optional description of this execution environment. (string) * `organization`: The organization used to determine access to this execution environment. (id) * `image`: The full image location, including the container registry, image name, and version tag. (string) * `managed`:  (boolean) * `credential`:  (id) * `pull`: Pull image before running? (choice)     - `\"\"`: ---------     - `always`: Always pull container before running.     - `missing`: Only pull the image if not present before running.     - `never`: Never pull container before running.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsReadOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesList**
> ExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesList(ctx, id, optional)
 List Unified Job Templates for an Execution Environment

 Make a GET request to this resource to retrieve a list of unified job templates associated with the selected execution environment.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of unified job templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more unified job template records.    ## Results  Each unified job template data structure includes the following fields:  * `id`: Database ID for this unified job template. (integer) * `type`: Data type for this unified job template. (choice) * `url`: URL for this unified job template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this unified job template was created. (datetime) * `modified`: Timestamp when this unified job template was last modified. (datetime) * `name`: Name of this unified job template. (string) * `description`: Optional description of this unified job template. (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing     - `none`: No External Source     - `updating`: Updating * `execution_environment`: The container image to be used for execution. (id)    ## Sorting  To specify that unified job templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsUnifiedJobTemplatesListOpts struct

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

# **ExecutionEnvironmentsExecutionEnvironmentsUpdate**
> ExecutionEnvironmentsExecutionEnvironmentsUpdate(ctx, id, optional)
 Update an Execution Environment

 Make a PUT or PATCH request to this resource to update this execution environment.  The following fields may be modified:          * `name`: Name of this execution environment. (string, required) * `description`: Optional description of this execution environment. (string, default=`\"\"`) * `organization`: The organization used to determine access to this execution environment. (id, default=``) * `image`: The full image location, including the container registry, image name, and version tag. (string, required)  * `credential`:  (id, default=``) * `pull`: Pull image before running? (choice)     - `\"\"`: --------- (default)     - `always`: Always pull container before running.     - `missing`: Only pull the image if not present before running.     - `never`: Never pull container before running.       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutionEnvironmentsApiExecutionEnvironmentsExecutionEnvironmentsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data6**](Data6.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

