# \UnifiedJobTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnifiedJobTemplatesUnifiedJobTemplatesList**](UnifiedJobTemplatesApi.md#UnifiedJobTemplatesUnifiedJobTemplatesList) | **Get** /api/v2/unified_job_templates/ |  List Unified Job Templates


# **UnifiedJobTemplatesUnifiedJobTemplatesList**
> UnifiedJobTemplatesUnifiedJobTemplatesList(ctx, optional)
 List Unified Job Templates

 Make a GET request to this resource to retrieve the list of unified job templates.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of unified job templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more unified job template records.    ## Results  Each unified job template data structure includes the following fields:  * `id`: Database ID for this unified job template. (integer) * `type`: Data type for this unified job template. (choice) * `url`: URL for this unified job template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this unified job template was created. (datetime) * `modified`: Timestamp when this unified job template was last modified. (datetime) * `name`: Name of this unified job template. (string) * `description`: Optional description of this unified job template. (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing     - `none`: No External Source     - `updating`: Updating * `execution_environment`: The container image to be used for execution. (id)    ## Sorting  To specify that unified job templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnifiedJobTemplatesApiUnifiedJobTemplatesUnifiedJobTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnifiedJobTemplatesApiUnifiedJobTemplatesUnifiedJobTemplatesListOpts struct

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

