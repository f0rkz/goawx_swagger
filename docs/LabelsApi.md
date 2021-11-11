# \LabelsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelsLabelsCreate**](LabelsApi.md#LabelsLabelsCreate) | **Post** /api/v2/labels/ |  Create a Label
[**LabelsLabelsList**](LabelsApi.md#LabelsLabelsList) | **Get** /api/v2/labels/ |  List Labels
[**LabelsLabelsPartialUpdate**](LabelsApi.md#LabelsLabelsPartialUpdate) | **Patch** /api/v2/labels/{id}/ |  Update a Label
[**LabelsLabelsRead**](LabelsApi.md#LabelsLabelsRead) | **Get** /api/v2/labels/{id}/ |  Retrieve a Label
[**LabelsLabelsUpdate**](LabelsApi.md#LabelsLabelsUpdate) | **Put** /api/v2/labels/{id}/ |  Update a Label


# **LabelsLabelsCreate**
> LabelsLabelsCreate(ctx, optional)
 Create a Label

 Make a POST request to this resource with the following label fields to create a new label:          * `name`: Name of this label. (string, required) * `organization`: Organization this label belongs to. (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LabelsApiLabelsLabelsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiLabelsLabelsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data34**](Data34.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsLabelsList**
> LabelsLabelsList(ctx, optional)
 List Labels

 Make a GET request to this resource to retrieve the list of labels.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of labels found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more label records.    ## Results  Each label data structure includes the following fields:  * `id`: Database ID for this label. (integer) * `type`: Data type for this label. (choice) * `url`: URL for this label. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this label was created. (datetime) * `modified`: Timestamp when this label was last modified. (datetime) * `name`: Name of this label. (string) * `organization`: Organization this label belongs to. (id)    ## Sorting  To specify that labels are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LabelsApiLabelsLabelsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiLabelsLabelsListOpts struct

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

# **LabelsLabelsPartialUpdate**
> LabelsLabelsPartialUpdate(ctx, id, optional)
 Update a Label

 Make a PUT or PATCH request to this resource to update this label.  The following fields may be modified:          * `name`: Name of this label. (string, required) * `organization`: Organization this label belongs to. (id, required)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***LabelsApiLabelsLabelsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiLabelsLabelsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data36**](Data36.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsLabelsRead**
> LabelsLabelsRead(ctx, id, optional)
 Retrieve a Label

 Make GET request to this resource to retrieve a single label record containing the following fields:  * `id`: Database ID for this label. (integer) * `type`: Data type for this label. (choice) * `url`: URL for this label. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this label was created. (datetime) * `modified`: Timestamp when this label was last modified. (datetime) * `name`: Name of this label. (string) * `organization`: Organization this label belongs to. (id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***LabelsApiLabelsLabelsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiLabelsLabelsReadOpts struct

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

# **LabelsLabelsUpdate**
> LabelsLabelsUpdate(ctx, id, optional)
 Update a Label

 Make a PUT or PATCH request to this resource to update this label.  The following fields may be modified:          * `name`: Name of this label. (string, required) * `organization`: Organization this label belongs to. (id, required)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***LabelsApiLabelsLabelsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiLabelsLabelsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data35**](Data35.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

