# \WorkflowApprovalsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowApprovalsWorkflowApprovalsApproveCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsApproveCreate) | **Post** /api/v2/workflow_approvals/{id}/approve/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsApproveRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsApproveRead) | **Get** /api/v2/workflow_approvals/{id}/approve/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsCreate) | **Post** /api/v2/workflow_approvals/ |  Create a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDelete**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDelete) | **Delete** /api/v2/workflow_approvals/{id}/ |  Delete a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDenyCreate**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDenyCreate) | **Post** /api/v2/workflow_approvals/{id}/deny/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsDenyRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsDenyRead) | **Get** /api/v2/workflow_approvals/{id}/deny/ |  Retrieve a Workflow Approval
[**WorkflowApprovalsWorkflowApprovalsList**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsList) | **Get** /api/v2/workflow_approvals/ |  List Workflow Approvals
[**WorkflowApprovalsWorkflowApprovalsRead**](WorkflowApprovalsApi.md#WorkflowApprovalsWorkflowApprovalsRead) | **Get** /api/v2/workflow_approvals/{id}/ |  Retrieve a Workflow Approval


# **WorkflowApprovalsWorkflowApprovalsApproveCreate**
> WorkflowApprovalsWorkflowApprovalsApproveCreate(ctx, id)
 Retrieve a Workflow Approval

 Make GET request to this resource to retrieve a single workflow approval record containing the following fields:

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

# **WorkflowApprovalsWorkflowApprovalsApproveRead**
> WorkflowApprovalsWorkflowApprovalsApproveRead(ctx, id, optional)
 Retrieve a Workflow Approval

 Make GET request to this resource to retrieve a single workflow approval record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsApproveReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsApproveReadOpts struct

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

# **WorkflowApprovalsWorkflowApprovalsCreate**
> WorkflowApprovalsWorkflowApprovalsCreate(ctx, optional)
 Create a Workflow Approval

 Make a POST request to this resource with the following workflow approval fields to create a new workflow approval:          * `name`: Name of this workflow approval. (string, required) * `description`: Optional description of this workflow approval. (string, default=`\"\"`)    * `execution_environment`: The container image to be used for execution. (id, default=``)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data73**](Data73.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkflowApprovalsWorkflowApprovalsDelete**
> WorkflowApprovalsWorkflowApprovalsDelete(ctx, id, optional)
 Delete a Workflow Approval

 Make a DELETE request to this resource to delete this workflow approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsDeleteOpts struct

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

# **WorkflowApprovalsWorkflowApprovalsDenyCreate**
> WorkflowApprovalsWorkflowApprovalsDenyCreate(ctx, id)
 Retrieve a Workflow Approval

 Make GET request to this resource to retrieve a single workflow approval record containing the following fields:

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

# **WorkflowApprovalsWorkflowApprovalsDenyRead**
> WorkflowApprovalsWorkflowApprovalsDenyRead(ctx, id, optional)
 Retrieve a Workflow Approval

 Make GET request to this resource to retrieve a single workflow approval record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsDenyReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsDenyReadOpts struct

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

# **WorkflowApprovalsWorkflowApprovalsList**
> WorkflowApprovalsWorkflowApprovalsList(ctx, optional)
 List Workflow Approvals

 Make a GET request to this resource to retrieve the list of workflow approvals.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow approvals found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow approval records.    ## Results  Each workflow approval data structure includes the following fields:  * `id`: Database ID for this workflow approval. (integer) * `type`: Data type for this workflow approval. (choice) * `url`: URL for this workflow approval. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow approval was created. (datetime) * `modified`: Timestamp when this workflow approval was last modified. (datetime) * `name`: Name of this workflow approval. (string) * `description`: Optional description of this workflow approval. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `can_approve_or_deny`:  (field) * `approval_expiration`:  (field) * `timed_out`:  (boolean)    ## Sorting  To specify that workflow approvals are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsListOpts struct

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

# **WorkflowApprovalsWorkflowApprovalsRead**
> WorkflowApprovalsWorkflowApprovalsRead(ctx, id, optional)
 Retrieve a Workflow Approval

 Make GET request to this resource to retrieve a single workflow approval record containing the following fields:  * `id`: Database ID for this workflow approval. (integer) * `type`: Data type for this workflow approval. (choice) * `url`: URL for this workflow approval. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow approval was created. (datetime) * `modified`: Timestamp when this workflow approval was last modified. (datetime) * `name`: Name of this workflow approval. (string) * `description`: Optional description of this workflow approval. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_args`:  (string) * `job_cwd`:  (string) * `job_env`:  (json) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `result_traceback`:  (string) * `event_processing_finished`: Indicates whether all of the events generated by this unified job have been saved to the database. (boolean) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `can_approve_or_deny`:  (field) * `approval_expiration`:  (field) * `timed_out`:  (boolean)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalsApiWorkflowApprovalsWorkflowApprovalsReadOpts struct

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

