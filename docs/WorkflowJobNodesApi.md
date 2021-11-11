# \WorkflowJobNodesApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowJobNodesWorkflowJobNodesAlwaysNodesList**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesAlwaysNodesList) | **Get** /api/v2/workflow_job_nodes/{id}/always_nodes/ |  List Workflow Job Nodes for a Workflow Job Node
[**WorkflowJobNodesWorkflowJobNodesCredentialsList**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesCredentialsList) | **Get** /api/v2/workflow_job_nodes/{id}/credentials/ |  List Credentials for a Workflow Job Node
[**WorkflowJobNodesWorkflowJobNodesFailureNodesList**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesFailureNodesList) | **Get** /api/v2/workflow_job_nodes/{id}/failure_nodes/ |  List Workflow Job Nodes for a Workflow Job Node
[**WorkflowJobNodesWorkflowJobNodesList**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesList) | **Get** /api/v2/workflow_job_nodes/ |  List Workflow Job Nodes
[**WorkflowJobNodesWorkflowJobNodesRead**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesRead) | **Get** /api/v2/workflow_job_nodes/{id}/ |  Retrieve a Workflow Job Node
[**WorkflowJobNodesWorkflowJobNodesSuccessNodesList**](WorkflowJobNodesApi.md#WorkflowJobNodesWorkflowJobNodesSuccessNodesList) | **Get** /api/v2/workflow_job_nodes/{id}/success_nodes/ |  List Workflow Job Nodes for a Workflow Job Node


# **WorkflowJobNodesWorkflowJobNodesAlwaysNodesList**
> WorkflowJobNodesWorkflowJobNodesAlwaysNodesList(ctx, id, optional)
 List Workflow Job Nodes for a Workflow Job Node

 Make a GET request to this resource to retrieve a list of workflow job nodes associated with the selected workflow job node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job node records.    ## Results  Each workflow job node data structure includes the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)    ## Sorting  To specify that workflow job nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesAlwaysNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesAlwaysNodesListOpts struct

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

# **WorkflowJobNodesWorkflowJobNodesCredentialsList**
> WorkflowJobNodesWorkflowJobNodesCredentialsList(ctx, id, optional)
 List Credentials for a Workflow Job Node

 Make a GET request to this resource to retrieve a list of credentials associated with the selected workflow job node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`:  (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)    ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesCredentialsListOpts struct

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

# **WorkflowJobNodesWorkflowJobNodesFailureNodesList**
> WorkflowJobNodesWorkflowJobNodesFailureNodesList(ctx, id, optional)
 List Workflow Job Nodes for a Workflow Job Node

 Make a GET request to this resource to retrieve a list of workflow job nodes associated with the selected workflow job node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job node records.    ## Results  Each workflow job node data structure includes the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)    ## Sorting  To specify that workflow job nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesFailureNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesFailureNodesListOpts struct

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

# **WorkflowJobNodesWorkflowJobNodesList**
> WorkflowJobNodesWorkflowJobNodesList(ctx, optional)
 List Workflow Job Nodes

 Make a GET request to this resource to retrieve the list of workflow job nodes.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job node records.    ## Results  Each workflow job node data structure includes the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)    ## Sorting  To specify that workflow job nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesListOpts struct

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

# **WorkflowJobNodesWorkflowJobNodesRead**
> WorkflowJobNodesWorkflowJobNodesRead(ctx, id, optional)
 Retrieve a Workflow Job Node

 Make GET request to this resource to retrieve a single workflow job node record containing the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesReadOpts struct

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

# **WorkflowJobNodesWorkflowJobNodesSuccessNodesList**
> WorkflowJobNodesWorkflowJobNodesSuccessNodesList(ctx, id, optional)
 List Workflow Job Nodes for a Workflow Job Node

 Make a GET request to this resource to retrieve a list of workflow job nodes associated with the selected workflow job node.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of workflow job nodes found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more workflow job node records.    ## Results  Each workflow job node data structure includes the following fields:  * `id`: Database ID for this workflow job node. (integer) * `type`: Data type for this workflow job node. (choice) * `url`: URL for this workflow job node. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this workflow job node was created. (datetime) * `modified`: Timestamp when this workflow job node was last modified. (datetime) * `extra_data`:  (json) * `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id) * `scm_branch`:  (string) * `job_type`:  (choice)     - `None`: ---------     - `\"\"`: ---------     - `run`: Run     - `check`: Check * `job_tags`:  (string) * `skip_tags`:  (string) * `limit`:  (string) * `diff_mode`:  (boolean) * `verbosity`:  (choice)     - `None`: ---------     - `0`: 0 (Normal)     - `1`: 1 (Verbose)     - `2`: 2 (More Verbose)     - `3`: 3 (Debug)     - `4`: 4 (Connection Debug)     - `5`: 5 (WinRM Debug) * `job`:  (id) * `workflow_job`:  (id) * `unified_job_template`:  (id) * `success_nodes`:  (field) * `failure_nodes`:  (field) * `always_nodes`:  (field) * `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean) * `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean) * `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)    ## Sorting  To specify that workflow job nodes are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesSuccessNodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowJobNodesApiWorkflowJobNodesWorkflowJobNodesSuccessNodesListOpts struct

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

