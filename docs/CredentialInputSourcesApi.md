# \CredentialInputSourcesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialInputSourcesCredentialInputSourcesCreate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesCreate) | **Post** /api/v2/credential_input_sources/ |  Create a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesDelete**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesDelete) | **Delete** /api/v2/credential_input_sources/{id}/ |  Delete a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesList**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesList) | **Get** /api/v2/credential_input_sources/ |  List Credential Input Sources
[**CredentialInputSourcesCredentialInputSourcesPartialUpdate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesPartialUpdate) | **Patch** /api/v2/credential_input_sources/{id}/ |  Update a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesRead**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesRead) | **Get** /api/v2/credential_input_sources/{id}/ |  Retrieve a Credential Input Source
[**CredentialInputSourcesCredentialInputSourcesUpdate**](CredentialInputSourcesApi.md#CredentialInputSourcesCredentialInputSourcesUpdate) | **Put** /api/v2/credential_input_sources/{id}/ |  Update a Credential Input Source


# **CredentialInputSourcesCredentialInputSourcesCreate**
> CredentialInputSourcesCredentialInputSourcesCreate(ctx, optional)
 Create a Credential Input Source

 Make a POST request to this resource with the following credential input source fields to create a new credential input source:          * `description`: Optional description of this credential input source. (string, default=`\"\"`) * `input_field_name`:  (string, required) * `metadata`:  (json, default=`{}`) * `target_credential`:  (id, required) * `source_credential`:  (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesCreateOpts struct

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

# **CredentialInputSourcesCredentialInputSourcesDelete**
> CredentialInputSourcesCredentialInputSourcesDelete(ctx, id, optional)
 Delete a Credential Input Source

 Make a DELETE request to this resource to delete this credential input source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesDeleteOpts struct

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

# **CredentialInputSourcesCredentialInputSourcesList**
> CredentialInputSourcesCredentialInputSourcesList(ctx, optional)
 List Credential Input Sources

 Make a GET request to this resource to retrieve the list of credential input sources.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credential input sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential input source records.    ## Results  Each credential input source data structure includes the following fields:  * `id`: Database ID for this credential input source. (integer) * `type`: Data type for this credential input source. (choice) * `url`: URL for this credential input source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential input source was created. (datetime) * `modified`: Timestamp when this credential input source was last modified. (datetime) * `description`: Optional description of this credential input source. (string) * `input_field_name`:  (string) * `metadata`:  (json) * `target_credential`:  (id) * `source_credential`:  (id)    ## Sorting  To specify that credential input sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesListOpts struct

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

# **CredentialInputSourcesCredentialInputSourcesPartialUpdate**
> CredentialInputSourcesCredentialInputSourcesPartialUpdate(ctx, id, optional)
 Update a Credential Input Source

 Make a PUT or PATCH request to this resource to update this credential input source.  The following fields may be modified:          * `description`: Optional description of this credential input source. (string, default=`\"\"`) * `input_field_name`:  (string, required) * `metadata`:  (json, default=`{}`) * `target_credential`:  (id, required) * `source_credential`:  (id, required)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesPartialUpdateOpts struct

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

# **CredentialInputSourcesCredentialInputSourcesRead**
> CredentialInputSourcesCredentialInputSourcesRead(ctx, id, optional)
 Retrieve a Credential Input Source

 Make GET request to this resource to retrieve a single credential input source record containing the following fields:  * `id`: Database ID for this credential input source. (integer) * `type`: Data type for this credential input source. (choice) * `url`: URL for this credential input source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential input source was created. (datetime) * `modified`: Timestamp when this credential input source was last modified. (datetime) * `description`: Optional description of this credential input source. (string) * `input_field_name`:  (string) * `metadata`:  (json) * `target_credential`:  (id) * `source_credential`:  (id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesReadOpts struct

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

# **CredentialInputSourcesCredentialInputSourcesUpdate**
> CredentialInputSourcesCredentialInputSourcesUpdate(ctx, id, optional)
 Update a Credential Input Source

 Make a PUT or PATCH request to this resource to update this credential input source.  The following fields may be modified:          * `description`: Optional description of this credential input source. (string, default=`\"\"`) * `input_field_name`:  (string, required) * `metadata`:  (json, default=`{}`) * `target_credential`:  (id, required) * `source_credential`:  (id, required)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialInputSourcesApiCredentialInputSourcesCredentialInputSourcesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data1**](Data1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

