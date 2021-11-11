# \CredentialsApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialsCredentialsAccessListList**](CredentialsApi.md#CredentialsCredentialsAccessListList) | **Get** /api/v2/credentials/{id}/access_list/ |  List Users
[**CredentialsCredentialsActivityStreamList**](CredentialsApi.md#CredentialsCredentialsActivityStreamList) | **Get** /api/v2/credentials/{id}/activity_stream/ |  List Activity Streams for a Credential
[**CredentialsCredentialsCopyCreate**](CredentialsApi.md#CredentialsCredentialsCopyCreate) | **Post** /api/v2/credentials/{id}/copy/ | 
[**CredentialsCredentialsCopyList**](CredentialsApi.md#CredentialsCredentialsCopyList) | **Get** /api/v2/credentials/{id}/copy/ | 
[**CredentialsCredentialsCreate**](CredentialsApi.md#CredentialsCredentialsCreate) | **Post** /api/v2/credentials/ |  Create a Credential
[**CredentialsCredentialsDelete**](CredentialsApi.md#CredentialsCredentialsDelete) | **Delete** /api/v2/credentials/{id}/ |  Delete a Credential
[**CredentialsCredentialsInputSourcesCreate**](CredentialsApi.md#CredentialsCredentialsInputSourcesCreate) | **Post** /api/v2/credentials/{id}/input_sources/ |  Create a Credential Input Source for a Credential
[**CredentialsCredentialsInputSourcesList**](CredentialsApi.md#CredentialsCredentialsInputSourcesList) | **Get** /api/v2/credentials/{id}/input_sources/ |  List Credential Input Sources for a Credential
[**CredentialsCredentialsList**](CredentialsApi.md#CredentialsCredentialsList) | **Get** /api/v2/credentials/ |  List Credentials
[**CredentialsCredentialsObjectRolesList**](CredentialsApi.md#CredentialsCredentialsObjectRolesList) | **Get** /api/v2/credentials/{id}/object_roles/ |  List Roles for a Credential
[**CredentialsCredentialsOwnerTeamsList**](CredentialsApi.md#CredentialsCredentialsOwnerTeamsList) | **Get** /api/v2/credentials/{id}/owner_teams/ |  List Teams for a Credential
[**CredentialsCredentialsOwnerUsersList**](CredentialsApi.md#CredentialsCredentialsOwnerUsersList) | **Get** /api/v2/credentials/{id}/owner_users/ |  List Users for a Credential
[**CredentialsCredentialsPartialUpdate**](CredentialsApi.md#CredentialsCredentialsPartialUpdate) | **Patch** /api/v2/credentials/{id}/ |  Update a Credential
[**CredentialsCredentialsRead**](CredentialsApi.md#CredentialsCredentialsRead) | **Get** /api/v2/credentials/{id}/ |  Retrieve a Credential
[**CredentialsCredentialsTestCreate**](CredentialsApi.md#CredentialsCredentialsTestCreate) | **Post** /api/v2/credentials/{id}/test/ |  Retrieve a Credential
[**CredentialsCredentialsTestRead**](CredentialsApi.md#CredentialsCredentialsTestRead) | **Get** /api/v2/credentials/{id}/test/ |  Retrieve a Credential
[**CredentialsCredentialsUpdate**](CredentialsApi.md#CredentialsCredentialsUpdate) | **Put** /api/v2/credentials/{id}/ |  Update a Credential


# **CredentialsCredentialsAccessListList**
> CredentialsCredentialsAccessListList(ctx, id, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsAccessListListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsAccessListListOpts struct

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

# **CredentialsCredentialsActivityStreamList**
> CredentialsCredentialsActivityStreamList(ctx, id, optional)
 List Activity Streams for a Credential

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected credential.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsActivityStreamListOpts struct

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

# **CredentialsCredentialsCopyCreate**
> CredentialsCredentialsCopyCreate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsCopyCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsCopyCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data4**](Data4.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CredentialsCredentialsCopyList**
> CredentialsCredentialsCopyList(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsCopyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsCopyListOpts struct

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

# **CredentialsCredentialsCreate**
> CredentialsCredentialsCreate(ctx, optional)
 Create a Credential

 Make a POST request to this resource with the following credential fields to create a new credential:          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`: Inherit permissions from organization roles. If provided on creation, do not give either user or team. (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)    * `user`: Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation. (id, default=`None`) * `team`: Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation. (id, default=`None`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialsApiCredentialsCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsCreateOpts struct

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

# **CredentialsCredentialsDelete**
> CredentialsCredentialsDelete(ctx, id, optional)
 Delete a Credential

 Make a DELETE request to this resource to delete this credential.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsDeleteOpts struct

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

# **CredentialsCredentialsInputSourcesCreate**
> CredentialsCredentialsInputSourcesCreate(ctx, id, optional)
 Create a Credential Input Source for a Credential

 Make a POST request to this resource with the following credential input source fields to create a new credential input source associated with this credential.          * `description`: Optional description of this credential input source. (string, default=`\"\"`) * `input_field_name`:  (string, required) * `metadata`:  (json, default=`{}`)  * `source_credential`:  (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsInputSourcesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsInputSourcesCreateOpts struct

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

# **CredentialsCredentialsInputSourcesList**
> CredentialsCredentialsInputSourcesList(ctx, id, optional)
 List Credential Input Sources for a Credential

 Make a GET request to this resource to retrieve a list of credential input sources associated with the selected credential.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credential input sources found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential input source records.    ## Results  Each credential input source data structure includes the following fields:  * `id`: Database ID for this credential input source. (integer) * `type`: Data type for this credential input source. (choice) * `url`: URL for this credential input source. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential input source was created. (datetime) * `modified`: Timestamp when this credential input source was last modified. (datetime) * `description`: Optional description of this credential input source. (string) * `input_field_name`:  (string) * `metadata`:  (json) * `target_credential`:  (id) * `source_credential`:  (id)    ## Sorting  To specify that credential input sources are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsInputSourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsInputSourcesListOpts struct

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

# **CredentialsCredentialsList**
> CredentialsCredentialsList(ctx, optional)
 List Credentials

 Make a GET request to this resource to retrieve the list of credentials.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`: Inherit permissions from organization roles. If provided on creation, do not give either user or team. (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)      ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialsApiCredentialsCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsListOpts struct

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

# **CredentialsCredentialsObjectRolesList**
> CredentialsCredentialsObjectRolesList(ctx, id, optional)
 List Roles for a Credential

 Make a GET request to this resource to retrieve a list of roles associated with the selected credential.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsObjectRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsObjectRolesListOpts struct

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

# **CredentialsCredentialsOwnerTeamsList**
> CredentialsCredentialsOwnerTeamsList(ctx, id, optional)
 List Teams for a Credential

 Make a GET request to this resource to retrieve a list of teams associated with the selected credential.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of teams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more team records.    ## Results  Each team data structure includes the following fields:  * `id`: Database ID for this team. (integer) * `type`: Data type for this team. (choice) * `url`: URL for this team. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this team was created. (datetime) * `modified`: Timestamp when this team was last modified. (datetime) * `name`: Name of this team. (string) * `description`: Optional description of this team. (string) * `organization`:  (id)    ## Sorting  To specify that teams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsOwnerTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsOwnerTeamsListOpts struct

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

# **CredentialsCredentialsOwnerUsersList**
> CredentialsCredentialsOwnerUsersList(ctx, id, optional)
 List Users for a Credential

 Make a GET request to this resource to retrieve a list of users associated with the selected credential.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsOwnerUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsOwnerUsersListOpts struct

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

# **CredentialsCredentialsPartialUpdate**
> CredentialsCredentialsPartialUpdate(ctx, id, optional)
 Update a Credential

 Make a PUT or PATCH request to this resource to update this credential.  The following fields may be modified:          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`:  (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)            For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CredentialsCredentialsRead**
> CredentialsCredentialsRead(ctx, id, optional)
 Retrieve a Credential

 Make GET request to this resource to retrieve a single credential record containing the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `organization`:  (id) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsReadOpts struct

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

# **CredentialsCredentialsTestCreate**
> CredentialsCredentialsTestCreate(ctx, id, optional)
 Retrieve a Credential

 Make GET request to this resource to retrieve a single credential record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsTestCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsTestCreateOpts struct

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

# **CredentialsCredentialsTestRead**
> CredentialsCredentialsTestRead(ctx, id, optional)
 Retrieve a Credential

 Make GET request to this resource to retrieve a single credential record containing the following fields:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsTestReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsTestReadOpts struct

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

# **CredentialsCredentialsUpdate**
> CredentialsCredentialsUpdate(ctx, id, optional)
 Update a Credential

 Make a PUT or PATCH request to this resource to update this credential.  The following fields may be modified:          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `organization`:  (id, default=`None`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)          For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CredentialsApiCredentialsCredentialsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CredentialsApiCredentialsCredentialsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | **optional.Interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

