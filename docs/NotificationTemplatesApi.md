# \NotificationTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationTemplatesNotificationTemplatesCopyCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCopyCreate) | **Post** /api/v2/notification_templates/{id}/copy/ | 
[**NotificationTemplatesNotificationTemplatesCopyList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCopyList) | **Get** /api/v2/notification_templates/{id}/copy/ | 
[**NotificationTemplatesNotificationTemplatesCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesCreate) | **Post** /api/v2/notification_templates/ |  Create a Notification Template
[**NotificationTemplatesNotificationTemplatesDelete**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesDelete) | **Delete** /api/v2/notification_templates/{id}/ |  Delete a Notification Template
[**NotificationTemplatesNotificationTemplatesList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesList) | **Get** /api/v2/notification_templates/ |  List Notification Templates
[**NotificationTemplatesNotificationTemplatesNotificationsList**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesNotificationsList) | **Get** /api/v2/notification_templates/{id}/notifications/ |  List Notifications for a Notification Template
[**NotificationTemplatesNotificationTemplatesPartialUpdate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesPartialUpdate) | **Patch** /api/v2/notification_templates/{id}/ |  Update a Notification Template
[**NotificationTemplatesNotificationTemplatesRead**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesRead) | **Get** /api/v2/notification_templates/{id}/ |  Retrieve a Notification Template
[**NotificationTemplatesNotificationTemplatesTestCreate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesTestCreate) | **Post** /api/v2/notification_templates/{id}/test/ | Test a Notification Template
[**NotificationTemplatesNotificationTemplatesUpdate**](NotificationTemplatesApi.md#NotificationTemplatesNotificationTemplatesUpdate) | **Put** /api/v2/notification_templates/{id}/ |  Update a Notification Template


# **NotificationTemplatesNotificationTemplatesCopyCreate**
> NotificationTemplatesNotificationTemplatesCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data40**](Data40.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationTemplatesNotificationTemplatesCopyList**
> NotificationTemplatesNotificationTemplatesCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCopyListOpts struct

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

# **NotificationTemplatesNotificationTemplatesCreate**
> NotificationTemplatesNotificationTemplatesCreate(ctx, optional)
 Create a Notification Template

 Make a POST request to this resource with the following notification template fields to create a new notification template:          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data37**](Data37.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationTemplatesNotificationTemplatesDelete**
> NotificationTemplatesNotificationTemplatesDelete(ctx, id, optional)
 Delete a Notification Template

 Make a DELETE request to this resource to delete this notification template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesDeleteOpts struct

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

# **NotificationTemplatesNotificationTemplatesList**
> NotificationTemplatesNotificationTemplatesList(ctx, optional)
 List Notification Templates

 Make a GET request to this resource to retrieve the list of notification templates.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notification templates found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification template records.    ## Results  Each notification template data structure includes the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)    ## Sorting  To specify that notification templates are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesListOpts struct

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

# **NotificationTemplatesNotificationTemplatesNotificationsList**
> NotificationTemplatesNotificationTemplatesNotificationsList(ctx, id, optional)
 List Notifications for a Notification Template

 Make a GET request to this resource to retrieve a list of notifications associated with the selected notification template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of notifications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more notification records.    ## Results  Each notification data structure includes the following fields:  * `id`: Database ID for this notification. (integer) * `type`: Data type for this notification. (choice) * `url`: URL for this notification. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification was created. (datetime) * `modified`: Timestamp when this notification was last modified. (datetime) * `notification_template`:  (id) * `error`:  (string) * `status`:  (choice)     - `pending`: Pending     - `successful`: Successful     - `failed`: Failed * `notifications_sent`:  (integer) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `recipients`:  (string) * `subject`:  (string) * `body`: Notification body (json)    ## Sorting  To specify that notifications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesNotificationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesNotificationsListOpts struct

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

# **NotificationTemplatesNotificationTemplatesPartialUpdate**
> NotificationTemplatesNotificationTemplatesPartialUpdate(ctx, id, optional)
 Update a Notification Template

 Make a PUT or PATCH request to this resource to update this notification template.  The following fields may be modified:          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data39**](Data39.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationTemplatesNotificationTemplatesRead**
> NotificationTemplatesNotificationTemplatesRead(ctx, id, optional)
 Retrieve a Notification Template

 Make GET request to this resource to retrieve a single notification template record containing the following fields:  * `id`: Database ID for this notification template. (integer) * `type`: Data type for this notification template. (choice) * `url`: URL for this notification template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this notification template was created. (datetime) * `modified`: Timestamp when this notification template was last modified. (datetime) * `name`: Name of this notification template. (string) * `description`: Optional description of this notification template. (string) * `organization`:  (id) * `notification_type`:  (choice)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json) * `messages`: Optional custom messages for notification template. (json)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesReadOpts struct

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

# **NotificationTemplatesNotificationTemplatesTestCreate**
> NotificationTemplatesNotificationTemplatesTestCreate(ctx, id)
Test a Notification Template



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

# **NotificationTemplatesNotificationTemplatesUpdate**
> NotificationTemplatesNotificationTemplatesUpdate(ctx, id, optional)
 Update a Notification Template

 Make a PUT or PATCH request to this resource to update this notification template.  The following fields may be modified:          * `name`: Name of this notification template. (string, required) * `description`: Optional description of this notification template. (string, default=`\"\"`) * `organization`:  (id, required) * `notification_type`:  (choice, required)     - `email`: Email     - `grafana`: Grafana     - `irc`: IRC     - `mattermost`: Mattermost     - `pagerduty`: Pagerduty     - `rocketchat`: Rocket.Chat     - `slack`: Slack     - `twilio`: Twilio     - `webhook`: Webhook * `notification_configuration`:  (json, default=`{}`) * `messages`: Optional custom messages for notification template. (json, default=`{&#39;started&#39;: None, &#39;success&#39;: None, &#39;error&#39;: None, &#39;workflow_approval&#39;: None}`)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***NotificationTemplatesApiNotificationTemplatesNotificationTemplatesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationTemplatesApiNotificationTemplatesNotificationTemplatesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data38**](Data38.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

