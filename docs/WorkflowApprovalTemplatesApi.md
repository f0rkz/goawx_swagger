# \WorkflowApprovalTemplatesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList) | **Get** /api/v2/workflow_approval_templates/{id}/approvals/ |  List Workflow Approvals for a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete) | **Delete** /api/v2/workflow_approval_templates/{id}/ |  Delete a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate) | **Patch** /api/v2/workflow_approval_templates/{id}/ |  Update a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead) | **Get** /api/v2/workflow_approval_templates/{id}/ |  Retrieve a Workflow Approval Template
[**WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate**](WorkflowApprovalTemplatesApi.md#WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate) | **Put** /api/v2/workflow_approval_templates/{id}/ |  Update a Workflow Approval Template


# **WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList**
> WorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsList(ctx, id, optional)
 List Workflow Approvals for a Workflow Approval Template

 Make a GET request to this resource to retrieve a list of workflow approvals associated with the selected workflow approval template.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow approvals found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow approval records.    ## Results  Each workflow approval data structure includes the following fields:  * `id`: Database ID for this workflow approval. (integer) * `type`: Data type for this workflow approval. (choice) * `url`: URL for this workflow approval. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow approval was created. (datetime) * `modified`: Timestamp when this workflow approval was last modified. (datetime) * `name`: Name of this workflow approval. (string) * `description`: Optional description of this workflow approval. (string) * `unified_job_template`:  (id) * `launch_type`:  (choice)     - `manual`: Manual     - `relaunch`: Relaunch     - `callback`: Callback     - `scheduled`: Scheduled     - `dependency`: Dependency     - `workflow`: Workflow     - `webhook`: Webhook     - `sync`: Sync     - `scm`: SCM Update * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled * `execution_environment`: The container image to be used for execution. (id) * `failed`:  (boolean) * `started`: The date and time the job was queued for starting. (datetime) * `finished`: The date and time the job finished execution. (datetime) * `canceled_on`: The date and time when the cancel request was sent. (datetime) * `elapsed`: Elapsed time in seconds that the job ran. (decimal) * `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string) * `launched_by`:  (field) * `work_unit_id`: The Receptor work unit ID associated with this job. (string) * `can_approve_or_deny`:  (field) * `approval_expiration`:  (field) * `timed_out`:  (boolean)    ## Sorting  To specify that workflow approvals are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesApprovalsListOpts struct

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

# **WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete**
> WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete(ctx, id, optional)
 Delete a Workflow Approval Template

 Make a DELETE request to this resource to delete this workflow approval template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteOpts struct

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

# **WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate**
> WorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdate(ctx, id, optional)
 Update a Workflow Approval Template

 Make a PUT or PATCH request to this resource to update this workflow approval template.  The following fields may be modified:          * `name`: Name of this workflow approval template. (string, required) * `description`: Optional description of this workflow approval template. (string, default=`\"\"`)     * `execution_environment`: The container image to be used for execution. (id, default=``) * `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer, default=`0`)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data72**](Data72.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead**
> WorkflowApprovalTemplatesWorkflowApprovalTemplatesRead(ctx, id, optional)
 Retrieve a Workflow Approval Template

 Make GET request to this resource to retrieve a single workflow approval template record containing the following fields:  * `id`: Database ID for this workflow approval template. (integer) * `type`: Data type for this workflow approval template. (choice) * `url`: URL for this workflow approval template. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow approval template was created. (datetime) * `modified`: Timestamp when this workflow approval template was last modified. (datetime) * `name`: Name of this workflow approval template. (string) * `description`: Optional description of this workflow approval template. (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing     - `none`: No External Source     - `updating`: Updating * `execution_environment`: The container image to be used for execution. (id) * `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesReadOpts struct

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

# **WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate**
> WorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdate(ctx, id, optional)
 Update a Workflow Approval Template

 Make a PUT or PATCH request to this resource to update this workflow approval template.  The following fields may be modified:          * `name`: Name of this workflow approval template. (string, required) * `description`: Optional description of this workflow approval template. (string, default=`\"\"`)     * `execution_environment`: The container image to be used for execution. (id, default=``) * `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer, default=`0`)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApprovalTemplatesApiWorkflowApprovalTemplatesWorkflowApprovalTemplatesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data71**](Data71.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

