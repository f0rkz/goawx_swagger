# \CredentialTypesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialTypesCredentialTypesActivityStreamList**](CredentialTypesApi.md#CredentialTypesCredentialTypesActivityStreamList) | **Get** /api/v2/credential_types/{id}/activity_stream/ |  List Activity Streams for a Credential Type
[**CredentialTypesCredentialTypesCreate**](CredentialTypesApi.md#CredentialTypesCredentialTypesCreate) | **Post** /api/v2/credential_types/ |  Create a Credential Type
[**CredentialTypesCredentialTypesCredentialsCreate**](CredentialTypesApi.md#CredentialTypesCredentialTypesCredentialsCreate) | **Post** /api/v2/credential_types/{id}/credentials/ |  Create a Credential for a Credential Type
[**CredentialTypesCredentialTypesCredentialsList**](CredentialTypesApi.md#CredentialTypesCredentialTypesCredentialsList) | **Get** /api/v2/credential_types/{id}/credentials/ |  List Credentials for a Credential Type
[**CredentialTypesCredentialTypesDelete**](CredentialTypesApi.md#CredentialTypesCredentialTypesDelete) | **Delete** /api/v2/credential_types/{id}/ |  Delete a Credential Type
[**CredentialTypesCredentialTypesList**](CredentialTypesApi.md#CredentialTypesCredentialTypesList) | **Get** /api/v2/credential_types/ |  List Credential Types
[**CredentialTypesCredentialTypesPartialUpdate**](CredentialTypesApi.md#CredentialTypesCredentialTypesPartialUpdate) | **Patch** /api/v2/credential_types/{id}/ |  Update a Credential Type
[**CredentialTypesCredentialTypesRead**](CredentialTypesApi.md#CredentialTypesCredentialTypesRead) | **Get** /api/v2/credential_types/{id}/ |  Retrieve a Credential Type
[**CredentialTypesCredentialTypesTestCreate**](CredentialTypesApi.md#CredentialTypesCredentialTypesTestCreate) | **Post** /api/v2/credential_types/{id}/test/ |  Retrieve a Credential Type
[**CredentialTypesCredentialTypesTestRead**](CredentialTypesApi.md#CredentialTypesCredentialTypesTestRead) | **Get** /api/v2/credential_types/{id}/test/ |  Retrieve a Credential Type
[**CredentialTypesCredentialTypesUpdate**](CredentialTypesApi.md#CredentialTypesCredentialTypesUpdate) | **Put** /api/v2/credential_types/{id}/ |  Update a Credential Type


# **CredentialTypesCredentialTypesActivityStreamList**
> CredentialTypesCredentialTypesActivityStreamList(ctx, id, optional)
 List Activity Streams for a Credential Type

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected credential type.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesActivityStreamListOpts struct

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

# **CredentialTypesCredentialTypesCreate**
> CredentialTypesCredentialTypesCreate(ctx, optional)
 Create a Credential Type

 Make a POST request to this resource with the following credential type fields to create a new credential type:          * `name`: Name of this credential type. (string, required) * `description`: Optional description of this credential type. (string, default=`\"\"`) * `kind`:  (choice, required)     - `net`: Network     - `cloud`: Cloud   * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`) * `injectors`: Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesCreateOpts struct

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

# **CredentialTypesCredentialTypesCredentialsCreate**
> CredentialTypesCredentialTypesCredentialsCreate(ctx, id, optional)
 Create a Credential for a Credential Type

 Make a POST request to this resource with the following credential fields to create a new credential associated with this credential type.          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`:  (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesCredentialsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data3**](Data3.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CredentialTypesCredentialTypesCredentialsList**
> CredentialTypesCredentialTypesCredentialsList(ctx, id, optional)
 List Credentials for a Credential Type

 Make a GET request to this resource to retrieve a list of credentials associated with the selected credential type.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`:  (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)    ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesCredentialsListOpts struct

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

# **CredentialTypesCredentialTypesDelete**
> CredentialTypesCredentialTypesDelete(ctx, id, optional)
 Delete a Credential Type

 Make a DELETE request to this resource to delete this credential type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesDeleteOpts struct

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

# **CredentialTypesCredentialTypesList**
> CredentialTypesCredentialTypesList(ctx, optional)
 List Credential Types

 Make a GET request to this resource to retrieve the list of credential types.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credential types found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential type records.    ## Results  Each credential type data structure includes the following fields:  * `id`: Database ID for this credential type. (integer) * `type`: Data type for this credential type. (choice) * `url`: URL for this credential type. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential type was created. (datetime) * `modified`: Timestamp when this credential type was last modified. (datetime) * `name`: Name of this credential type. (string) * `description`: Optional description of this credential type. (string) * `kind`:  (choice)     - `ssh`: Machine     - `vault`: Vault     - `net`: Network     - `scm`: Source Control     - `cloud`: Cloud     - `registry`: Container Registry     - `token`: Personal Access Token     - `insights`: Insights     - `external`: External     - `kubernetes`: Kubernetes     - `galaxy`: Galaxy/Automation Hub * `namespace`:  (string) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `injectors`: Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax. (json)    ## Sorting  To specify that credential types are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesListOpts struct

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

# **CredentialTypesCredentialTypesPartialUpdate**
> CredentialTypesCredentialTypesPartialUpdate(ctx, id, optional)
 Update a Credential Type

 Make a PUT or PATCH request to this resource to update this credential type.  The following fields may be modified:          * `name`: Name of this credential type. (string, required) * `description`: Optional description of this credential type. (string, default=`\"\"`) * `kind`:  (choice, required)     - `net`: Network     - `cloud`: Cloud   * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`) * `injectors`: Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesPartialUpdateOpts struct

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

# **CredentialTypesCredentialTypesRead**
> CredentialTypesCredentialTypesRead(ctx, id, optional)
 Retrieve a Credential Type

 Make GET request to this resource to retrieve a single credential type record containing the following fields:  * `id`: Database ID for this credential type. (integer) * `type`: Data type for this credential type. (choice) * `url`: URL for this credential type. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential type was created. (datetime) * `modified`: Timestamp when this credential type was last modified. (datetime) * `name`: Name of this credential type. (string) * `description`: Optional description of this credential type. (string) * `kind`:  (choice)     - `ssh`: Machine     - `vault`: Vault     - `net`: Network     - `scm`: Source Control     - `cloud`: Cloud     - `registry`: Container Registry     - `token`: Personal Access Token     - `insights`: Insights     - `external`: External     - `kubernetes`: Kubernetes     - `galaxy`: Galaxy/Automation Hub * `namespace`:  (string) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `injectors`: Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax. (json)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesReadOpts struct

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

# **CredentialTypesCredentialTypesTestCreate**
> CredentialTypesCredentialTypesTestCreate(ctx, id, optional)
 Retrieve a Credential Type

 Make GET request to this resource to retrieve a single credential type record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesTestCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesTestCreateOpts struct

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

# **CredentialTypesCredentialTypesTestRead**
> CredentialTypesCredentialTypesTestRead(ctx, id, optional)
 Retrieve a Credential Type

 Make GET request to this resource to retrieve a single credential type record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesTestReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesTestReadOpts struct

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

# **CredentialTypesCredentialTypesUpdate**
> CredentialTypesCredentialTypesUpdate(ctx, id, optional)
 Update a Credential Type

 Make a PUT or PATCH request to this resource to update this credential type.  The following fields may be modified:          * `name`: Name of this credential type. (string, required) * `description`: Optional description of this credential type. (string, default=`\"\"`) * `kind`:  (choice, required)     - `net`: Network     - `cloud`: Cloud   * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`) * `injectors`: Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialTypesApiCredentialTypesCredentialTypesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialTypesApiCredentialTypesCredentialTypesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data2**](Data2.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

