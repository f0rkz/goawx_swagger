# \TeamsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTeamsAccessListList**](TeamsApi.md#TeamsTeamsAccessListList) | **Get** /api/v2/teams/{id}/access_list/ |  List Users
[**TeamsTeamsActivityStreamList**](TeamsApi.md#TeamsTeamsActivityStreamList) | **Get** /api/v2/teams/{id}/activity_stream/ |  List Activity Streams for a Team
[**TeamsTeamsCreate**](TeamsApi.md#TeamsTeamsCreate) | **Post** /api/v2/teams/ |  Create a Team
[**TeamsTeamsCredentialsCreate**](TeamsApi.md#TeamsTeamsCredentialsCreate) | **Post** /api/v2/teams/{id}/credentials/ |  Create a Credential for a Team
[**TeamsTeamsCredentialsList**](TeamsApi.md#TeamsTeamsCredentialsList) | **Get** /api/v2/teams/{id}/credentials/ |  List Credentials for a Team
[**TeamsTeamsDelete**](TeamsApi.md#TeamsTeamsDelete) | **Delete** /api/v2/teams/{id}/ |  Delete a Team
[**TeamsTeamsList**](TeamsApi.md#TeamsTeamsList) | **Get** /api/v2/teams/ |  List Teams
[**TeamsTeamsObjectRolesList**](TeamsApi.md#TeamsTeamsObjectRolesList) | **Get** /api/v2/teams/{id}/object_roles/ |  List Roles for a Team
[**TeamsTeamsPartialUpdate**](TeamsApi.md#TeamsTeamsPartialUpdate) | **Patch** /api/v2/teams/{id}/ |  Update a Team
[**TeamsTeamsProjectsList**](TeamsApi.md#TeamsTeamsProjectsList) | **Get** /api/v2/teams/{id}/projects/ |  List Projects for a Team
[**TeamsTeamsRead**](TeamsApi.md#TeamsTeamsRead) | **Get** /api/v2/teams/{id}/ |  Retrieve a Team
[**TeamsTeamsRolesCreate**](TeamsApi.md#TeamsTeamsRolesCreate) | **Post** /api/v2/teams/{id}/roles/ |  Associate Roles with this Team
[**TeamsTeamsRolesList**](TeamsApi.md#TeamsTeamsRolesList) | **Get** /api/v2/teams/{id}/roles/ |  List Roles for a Team
[**TeamsTeamsUpdate**](TeamsApi.md#TeamsTeamsUpdate) | **Put** /api/v2/teams/{id}/ |  Update a Team
[**TeamsTeamsUsersCreate**](TeamsApi.md#TeamsTeamsUsersCreate) | **Post** /api/v2/teams/{id}/users/ |  Create a User for a Team
[**TeamsTeamsUsersList**](TeamsApi.md#TeamsTeamsUsersList) | **Get** /api/v2/teams/{id}/users/ |  List Users for a Team


# **TeamsTeamsAccessListList**
> TeamsTeamsAccessListList(ctx, id, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsAccessListListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsAccessListListOpts struct

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

# **TeamsTeamsActivityStreamList**
> TeamsTeamsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Team

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsActivityStreamListOpts struct

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

# **TeamsTeamsCreate**
> TeamsTeamsCreate(ctx, optional)
 Create a Team

 Make a POST request to this resource with the following team fields to create a new team:          * `name`: Name of this team. (string, required) * `description`: Optional description of this team. (string, default=`\"\"`) * `organization`:  (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiTeamsTeamsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**optional.Interface of Data63**](Data63.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsTeamsCredentialsCreate**
> TeamsTeamsCredentialsCreate(ctx, id, optional)
 Create a Credential for a Team

 Make a POST request to this resource with the following credential fields to create a new credential associated with this team.          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsCredentialsCreateOpts struct

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

# **TeamsTeamsCredentialsList**
> TeamsTeamsCredentialsList(ctx, id, optional)
 List Credentials for a Team

 Make a GET request to this resource to retrieve a list of credentials associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)     ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsCredentialsListOpts struct

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

# **TeamsTeamsDelete**
> TeamsTeamsDelete(ctx, id, optional)
 Delete a Team

 Make a DELETE request to this resource to delete this team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsDeleteOpts struct

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

# **TeamsTeamsList**
> TeamsTeamsList(ctx, optional)
 List Teams

 Make a GET request to this resource to retrieve the list of teams.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of teams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more team records.    ## Results  Each team data structure includes the following fields:  * `id`: Database ID for this team. (integer) * `type`: Data type for this team. (choice) * `url`: URL for this team. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this team was created. (datetime) * `modified`: Timestamp when this team was last modified. (datetime) * `name`: Name of this team. (string) * `description`: Optional description of this team. (string) * `organization`:  (id)    ## Sorting  To specify that teams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiTeamsTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsListOpts struct

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

# **TeamsTeamsObjectRolesList**
> TeamsTeamsObjectRolesList(ctx, id, optional)
 List Roles for a Team

 Make a GET request to this resource to retrieve a list of roles associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsObjectRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsObjectRolesListOpts struct

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

# **TeamsTeamsPartialUpdate**
> TeamsTeamsPartialUpdate(ctx, id, optional)
 Update a Team

 Make a PUT or PATCH request to this resource to update this team.  The following fields may be modified:          * `name`: Name of this team. (string, required) * `description`: Optional description of this team. (string, default=`\"\"`) * `organization`:  (id, required)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data65**](Data65.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsTeamsProjectsList**
> TeamsTeamsProjectsList(ctx, id, optional)
 List Projects for a Team

 Make a GET request to this resource to retrieve a list of projects associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of projects found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project records.    ## Results  Each project data structure includes the following fields:  * `id`: Database ID for this project. (integer) * `type`: Data type for this project. (choice) * `url`: URL for this project. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project was created. (datetime) * `modified`: Timestamp when this project was last modified. (datetime) * `name`: Name of this project. (string) * `description`: Optional description of this project. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The last revision fetched by a project update (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing * `organization`: The organization used to determine access to this template. (id) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run using this project. (id) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that projects are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsProjectsListOpts struct

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

# **TeamsTeamsRead**
> TeamsTeamsRead(ctx, id, optional)
 Retrieve a Team

 Make GET request to this resource to retrieve a single team record containing the following fields:  * `id`: Database ID for this team. (integer) * `type`: Data type for this team. (choice) * `url`: URL for this team. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this team was created. (datetime) * `modified`: Timestamp when this team was last modified. (datetime) * `name`: Name of this team. (string) * `description`: Optional description of this team. (string) * `organization`:  (id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsReadOpts struct

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

# **TeamsTeamsRolesCreate**
> TeamsTeamsRolesCreate(ctx, id, optional)
 Associate Roles with this Team

 Make a POST request to this resource to add or remove a role from this team. The following fields may be modified:     * `id`: The Role ID to add to the team. (int, required)    * `disassociate`: Provide if you want to remove the role. (any value, optional)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsRolesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsRolesCreateOpts struct

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

# **TeamsTeamsRolesList**
> TeamsTeamsRolesList(ctx, id, optional)
 List Roles for a Team

 Make a GET request to this resource to retrieve a list of roles associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsRolesListOpts struct

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

# **TeamsTeamsUpdate**
> TeamsTeamsUpdate(ctx, id, optional)
 Update a Team

 Make a PUT or PATCH request to this resource to update this team.  The following fields may be modified:          * `name`: Name of this team. (string, required) * `description`: Optional description of this team. (string, default=`\"\"`) * `organization`:  (id, required)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data64**](Data64.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsTeamsUsersCreate**
> TeamsTeamsUsersCreate(ctx, id, optional)
 Create a User for a Team

 Make a POST request to this resource with the following user fields to create a new user associated with this team.          * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string, required) * `first_name`:  (string, default=`\"\"`) * `last_name`:  (string, default=`\"\"`) * `email`:  (string, default=`\"\"`) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean, default=`False`) * `is_system_auditor`:  (boolean, default=`False`) * `password`: Write-only field used to change the password. (string, default=`\"\"`)            # Add Users for a Team:  Make a POST request to this resource with only an `id` field to associate an existing user with this team.  # Remove Users from this Team:  Make a POST request to this resource with `id` and `disassociate` fields to remove the user from this team  without deleting the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsUsersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsUsersCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data66**](Data66.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsTeamsUsersList**
> TeamsTeamsUsersList(ctx, id, optional)
 List Users for a Team

 Make a GET request to this resource to retrieve a list of users associated with the selected team.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***TeamsApiTeamsTeamsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsTeamsUsersListOpts struct

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

