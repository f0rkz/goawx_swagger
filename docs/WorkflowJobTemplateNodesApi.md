# \WorkflowJobTemplateNodesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/always_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/always_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/ |  Create a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/create_approval_template/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead) | **Get** /api/v2/workflow_job_template_nodes/{id}/create_approval_template/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/credentials/ |  Create a Credential for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList) | **Get** /api/v2/workflow_job_template_nodes/{id}/credentials/ |  List Credentials for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete) | **Delete** /api/v2/workflow_job_template_nodes/{id}/ |  Delete a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/failure_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/failure_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesList) | **Get** /api/v2/workflow_job_template_nodes/ |  List Workflow Job Template Nodes
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate) | **Patch** /api/v2/workflow_job_template_nodes/{id}/ |  Update a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead) | **Get** /api/v2/workflow_job_template_nodes/{id}/ |  Retrieve a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate) | **Post** /api/v2/workflow_job_template_nodes/{id}/success_nodes/ |  Create a Workflow Job Template Node for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList) | **Get** /api/v2/workflow_job_template_nodes/{id}/success_nodes/ |  List Workflow Job Template Nodes for a Workflow Job Template Node
[**WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate**](WorkflowJobTemplateNodesApi.md#WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate) | **Put** /api/v2/workflow_job_template_nodes/{id}/ |  Update a Workflow Job Template Node


# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreate(ctx, id, optional)
 Create a Workflow Job Template Node for a Workflow Job Template Node

 Make a POST request to this resource with the following workflow job template node fields to create a new workflow job template node associated with this workflow job template node.          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)         # Add Workflow Job Template Nodes for a Workflow Job Template Node:  Make a POST request to this resource with only an `id` field to associate an existing workflow job template node with this workflow job template node.  # Remove Workflow Job Template Nodes from this Workflow Job Template Node:  Make a POST request to this resource with `id` and `disassociate` fields to remove the workflow job template node from this workflow job template node  without deleting the workflow job template node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList(ctx, id, optional)
 List Workflow Job Template Nodes for a Workflow Job Template Node

 Make a GET request to this resource to retrieve a list of workflow job template nodes associated with the selected workflow job template node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job template nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job template node records.    ## Results  Each workflow job template node data structure includes the following fields:  * `id`: Database ID for this workflow job template node. (integer) * `type`: Data type for this workflow job template node. (choice) * `url`: URL for this workflow job template node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job template node was created. (datetime) * `modified`: Timestamp when this workflow job template node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)    ## Sorting  To specify that workflow job template nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreate(ctx, optional)
 Create a Workflow Job Template Node

 Make a POST request to this resource with the following workflow job template node fields to create a new workflow job template node:          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data74**](Data74.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreate(ctx, id, optional)
 Retrieve a Workflow Job Template Node

 Make GET request to this resource to retrieve a single workflow job template node record containing the following fields:  * `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer) * `name`: Name of this workflow approval template. (string) * `description`: Optional description of this workflow approval template. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateRead(ctx, id, optional)
 Retrieve a Workflow Job Template Node

 Make GET request to this resource to retrieve a single workflow job template node record containing the following fields:  * `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer) * `name`: Name of this workflow approval template. (string) * `description`: Optional description of this workflow approval template. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreate(ctx, id, optional)
 Create a Credential for a Workflow Job Template Node

 Make a POST request to this resource with the following credential fields to create a new credential associated with this workflow job template node.          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`:  (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)            # Add Credentials for a Workflow Job Template Node:  Make a POST request to this resource with only an `id` field to associate an existing credential with this workflow job template node.  # Remove Credentials from this Workflow Job Template Node:  Make a POST request to this resource with `id` and `disassociate` fields to remove the credential from this workflow job template node  without deleting the credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsCreateOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsList(ctx, id, optional)
 List Credentials for a Workflow Job Template Node

 Make a GET request to this resource to retrieve a list of credentials associated with the selected workflow job template node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`:  (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)    ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesCredentialsListOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesDelete(ctx, id, optional)
 Delete a Workflow Job Template Node

 Make a DELETE request to this resource to delete this workflow job template node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesDeleteOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreate(ctx, id, optional)
 Create a Workflow Job Template Node for a Workflow Job Template Node

 Make a POST request to this resource with the following workflow job template node fields to create a new workflow job template node associated with this workflow job template node.          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)         # Add Workflow Job Template Nodes for a Workflow Job Template Node:  Make a POST request to this resource with only an `id` field to associate an existing workflow job template node with this workflow job template node.  # Remove Workflow Job Template Nodes from this Workflow Job Template Node:  Make a POST request to this resource with `id` and `disassociate` fields to remove the workflow job template node from this workflow job template node  without deleting the workflow job template node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesCreateOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesList(ctx, id, optional)
 List Workflow Job Template Nodes for a Workflow Job Template Node

 Make a GET request to this resource to retrieve a list of workflow job template nodes associated with the selected workflow job template node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job template nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job template node records.    ## Results  Each workflow job template node data structure includes the following fields:  * `id`: Database ID for this workflow job template node. (integer) * `type`: Data type for this workflow job template node. (choice) * `url`: URL for this workflow job template node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job template node was created. (datetime) * `modified`: Timestamp when this workflow job template node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)    ## Sorting  To specify that workflow job template nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesFailureNodesListOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesList**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesList(ctx, optional)
 List Workflow Job Template Nodes

 Make a GET request to this resource to retrieve the list of workflow job template nodes.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job template nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job template node records.    ## Results  Each workflow job template node data structure includes the following fields:  * `id`: Database ID for this workflow job template node. (integer) * `type`: Data type for this workflow job template node. (choice) * `url`: URL for this workflow job template node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job template node was created. (datetime) * `modified`: Timestamp when this workflow job template node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)    ## Sorting  To specify that workflow job template nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesListOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdate(ctx, id, optional)
 Update a Workflow Job Template Node

 Make a PUT or PATCH request to this resource to update this workflow job template node.  The following fields may be modified:          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data76**](Data76.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesRead(ctx, id, optional)
 Retrieve a Workflow Job Template Node

 Make GET request to this resource to retrieve a single workflow job template node record containing the following fields:  * `id`: Database ID for this workflow job template node. (integer) * `type`: Data type for this workflow job template node. (choice) * `url`: URL for this workflow job template node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job template node was created. (datetime) * `modified`: Timestamp when this workflow job template node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesReadOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreate(ctx, id, optional)
 Create a Workflow Job Template Node for a Workflow Job Template Node

 Make a POST request to this resource with the following workflow job template node fields to create a new workflow job template node associated with this workflow job template node.          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)         # Add Workflow Job Template Nodes for a Workflow Job Template Node:  Make a POST request to this resource with only an `id` field to associate an existing workflow job template node with this workflow job template node.  # Remove Workflow Job Template Nodes from this Workflow Job Template Node:  Make a POST request to this resource with `id` and `disassociate` fields to remove the workflow job template node from this workflow job template node  without deleting the workflow job template node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesCreateOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesList(ctx, id, optional)
 List Workflow Job Template Nodes for a Workflow Job Template Node

 Make a GET request to this resource to retrieve a list of workflow job template nodes associated with the selected workflow job template node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job template nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job template node records.    ## Results  Each workflow job template node data structure includes the following fields:  * `id`: Database ID for this workflow job template node. (integer) * `type`: Data type for this workflow job template node. (choice) * `url`: URL for this workflow job template node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job template node was created. (datetime) * `modified`: Timestamp when this workflow job template node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)    ## Sorting  To specify that workflow job template nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesSuccessNodesListOpts struct

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

# **WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate**
> WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate(ctx, id, optional)
 Update a Workflow Job Template Node

 Make a PUT or PATCH request to this resource to update this workflow job template node.  The following fields may be modified:          * `extra_data`:  (json, default=`{}`) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``) * `scm_branch`:  (string, default=`\"\"`) * `job_type`:  (choice)     - `None`: --------- (default)     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string, default=`\"\"`) * `skip_tags`:  (string, default=`\"\"`) * `limit`:  (string, default=`\"\"`) * `diff_mode`:  (boolean, default=`None`) * `verbosity`:  (choice)     - `None`: --------- (default)     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `workflow_job_template`:  (id, required) * `unified_job_template`:  (id, default=``)    * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`) * `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`\"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\"`)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobTemplateNodesApiWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data75**](Data75.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

