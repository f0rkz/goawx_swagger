# \JobHostSummariesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobHostSummariesJobHostSummariesRead**](JobHostSummariesApi.md#JobHostSummariesJobHostSummariesRead) | **Get** /api/v2/job_host_summaries/{id}/ |  Retrieve a Job Host Summary


# **JobHostSummariesJobHostSummariesRead**
> JobHostSummariesJobHostSummariesRead(ctx, id, optional)
 Retrieve a Job Host Summary

 Make GET request to this resource to retrieve a single job host summary record containing the following fields:  * `id`: Database ID for this job host summary. (integer) * `type`: Data type for this job host summary. (choice) * `url`: URL for this job host summary. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this job host summary was created. (datetime) * `modified`: Timestamp when this job host summary was last modified. (datetime) * `job`:  (id) * `host`:  (id) * `host_name`:  (string) * `changed`:  (integer) * `dark`:  (integer) * `failures`:  (integer) * `ok`:  (integer) * `processed`:  (integer) * `skipped`:  (integer) * `failed`:  (boolean) * `ignored`:  (integer) * `rescued`:  (integer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***JobHostSummariesApiJobHostSummariesJobHostSummariesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobHostSummariesApiJobHostSummariesJobHostSummariesReadOpts struct

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

