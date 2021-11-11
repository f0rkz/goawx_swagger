# \AdHocCommandEventsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocCommandEventsAdHocCommandEventsRead**](AdHocCommandEventsApi.md#AdHocCommandEventsAdHocCommandEventsRead) | **Get** /api/v2/ad_hoc_command_events/{id}/ |  Retrieve an Ad Hoc Command Event


# **AdHocCommandEventsAdHocCommandEventsRead**
> AdHocCommandEventsAdHocCommandEventsRead(ctx, id, optional)
 Retrieve an Ad Hoc Command Event

 Make GET request to this resource to retrieve a single ad hoc command event record containing the following fields:  * `id`: Database ID for this ad hoc command event. (integer) * `type`: Data type for this ad hoc command event. (choice) * `url`: URL for this ad hoc command event. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this ad hoc command event was created. (datetime) * `modified`: Timestamp when this ad hoc command event was last modified. (datetime) * `ad_hoc_command`:  (id) * `event`:  (choice)     - `runner_on_failed`: Host Failed     - `runner_on_ok`: Host OK     - `runner_on_unreachable`: Host Unreachable     - `runner_on_skipped`: Host Skipped     - `debug`: Debug     - `verbose`: Verbose     - `deprecated`: Deprecated     - `warning`: Warning     - `system_warning`: System Warning     - `error`: Error * `counter`:  (integer) * `event_display`:  (string) * `event_data`:  (json) * `failed`:  (boolean) * `changed`:  (boolean) * `uuid`:  (string) * `host`:  (id) * `host_name`:  (string) * `stdout`:  (string) * `start_line`:  (integer) * `end_line`:  (integer) * `verbosity`:  (integer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdHocCommandEventsApiAdHocCommandEventsAdHocCommandEventsReadOpts struct

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

