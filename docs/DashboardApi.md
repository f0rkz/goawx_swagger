# \DashboardApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardDashboardGraphsJobsList**](DashboardApi.md#DashboardDashboardGraphsJobsList) | **Get** /api/v2/dashboard/graphs/jobs/ |  View Statistics for Job Runs
[**DashboardDashboardList**](DashboardApi.md#DashboardDashboardList) | **Get** /api/v2/dashboard/ | Show Dashboard Details


# **DashboardDashboardGraphsJobsList**
> DashboardDashboardGraphsJobsList(ctx, )
 View Statistics for Job Runs

 Make a GET request to this resource to retrieve aggregate statistics about job runs suitable for graphing.  ## Parmeters and Filtering  The `period` of the data can be adjusted with:      ?period=month  Where `month` can be replaced with `week`, `two_weeks`, or `day`.  `month` is the default.  The type of job can be filtered with:      ?job_type=all  Where `all` can be replaced with `inv_sync`, `playbook_run` or `scm_update`.  `all` is the default.  ## Results  Data will be returned in the following format:      \"jobs\": {             \"successful\": [                 [                     1402808400.0,                      9                 ], ... ],             \"failed\": [              [                     1402808400.0,                      3                 ], ... ]     }  Each element contains an epoch timestamp represented in seconds and a numerical value indicating the number of events during that time period

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardDashboardList**
> DashboardDashboardList(ctx, )
Show Dashboard Details



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

