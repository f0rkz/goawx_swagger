# \UsersApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeList**](UsersApi.md#UsersMeList) | **Get** /api/v2/me/ |  Retrieve Information about the current User
[**UsersUsersAccessListList**](UsersApi.md#UsersUsersAccessListList) | **Get** /api/v2/users/{id}/access_list/ |  List Users
[**UsersUsersActivityStreamList**](UsersApi.md#UsersUsersActivityStreamList) | **Get** /api/v2/users/{id}/activity_stream/ |  List Activity Streams for a User
[**UsersUsersAdminOfOrganizationsList**](UsersApi.md#UsersUsersAdminOfOrganizationsList) | **Get** /api/v2/users/{id}/admin_of_organizations/ |  List Organizations Administered by this User
[**UsersUsersApplicationsCreate**](UsersApi.md#UsersUsersApplicationsCreate) | **Post** /api/v2/users/{id}/applications/ |  Create an Application
[**UsersUsersApplicationsList**](UsersApi.md#UsersUsersApplicationsList) | **Get** /api/v2/users/{id}/applications/ |  List Applications
[**UsersUsersAuthorizedTokensCreate**](UsersApi.md#UsersUsersAuthorizedTokensCreate) | **Post** /api/v2/users/{id}/authorized_tokens/ |  Create an Access Token for a User
[**UsersUsersAuthorizedTokensList**](UsersApi.md#UsersUsersAuthorizedTokensList) | **Get** /api/v2/users/{id}/authorized_tokens/ |  List Access Tokens for a User
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /api/v2/users/ |  Create a User
[**UsersUsersCredentialsCreate**](UsersApi.md#UsersUsersCredentialsCreate) | **Post** /api/v2/users/{id}/credentials/ |  Create a Credential for a User
[**UsersUsersCredentialsList**](UsersApi.md#UsersUsersCredentialsList) | **Get** /api/v2/users/{id}/credentials/ |  List Credentials for a User
[**UsersUsersDelete**](UsersApi.md#UsersUsersDelete) | **Delete** /api/v2/users/{id}/ |  Delete a User
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /api/v2/users/ |  List Users
[**UsersUsersOrganizationsList**](UsersApi.md#UsersUsersOrganizationsList) | **Get** /api/v2/users/{id}/organizations/ |  List Organizations for a User
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /api/v2/users/{id}/ |  Update a User
[**UsersUsersPersonalTokensCreate**](UsersApi.md#UsersUsersPersonalTokensCreate) | **Post** /api/v2/users/{id}/personal_tokens/ |  Create an Access Token for a User
[**UsersUsersPersonalTokensList**](UsersApi.md#UsersUsersPersonalTokensList) | **Get** /api/v2/users/{id}/personal_tokens/ |  List Access Tokens for a User
[**UsersUsersProjectsList**](UsersApi.md#UsersUsersProjectsList) | **Get** /api/v2/users/{id}/projects/ |  List Projects for a User
[**UsersUsersRead**](UsersApi.md#UsersUsersRead) | **Get** /api/v2/users/{id}/ |  Retrieve a User
[**UsersUsersRolesCreate**](UsersApi.md#UsersUsersRolesCreate) | **Post** /api/v2/users/{id}/roles/ |  Associate Roles with this User
[**UsersUsersRolesList**](UsersApi.md#UsersUsersRolesList) | **Get** /api/v2/users/{id}/roles/ |  List Roles for a User
[**UsersUsersTeamsList**](UsersApi.md#UsersUsersTeamsList) | **Get** /api/v2/users/{id}/teams/ |  List Teams for a User
[**UsersUsersTokensCreate**](UsersApi.md#UsersUsersTokensCreate) | **Post** /api/v2/users/{id}/tokens/ |  Create an Access Token for a User
[**UsersUsersTokensList**](UsersApi.md#UsersUsersTokensList) | **Get** /api/v2/users/{id}/tokens/ |  List Access Tokens for a User
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /api/v2/users/{id}/ |  Update a User


# **UsersMeList**
> UsersMeList(ctx, optional)
 Retrieve Information about the current User

 Make a GET request to retrieve user information about the current user.  One result should be returned containing the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    Use the primary URL for the user (/api/v2/users/N/) to modify the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersMeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersMeListOpts struct

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

# **UsersUsersAccessListList**
> UsersUsersAccessListList(ctx, id, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersAccessListListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersAccessListListOpts struct

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

# **UsersUsersActivityStreamList**
> UsersUsersActivityStreamList(ctx, id, optional)
 List Activity Streams for a User

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersActivityStreamListOpts struct

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

# **UsersUsersAdminOfOrganizationsList**
> UsersUsersAdminOfOrganizationsList(ctx, id, optional)
 List Organizations Administered by this User

 Make a GET request to this resource to retrieve a list of organizations of which the selected user is an admin.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of organizations found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more organization records.    ## Results  Each organization data structure includes the following fields:  * `id`: Database ID for this organization. (integer) * `type`: Data type for this organization. (choice) * `url`: URL for this organization. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this organization was created. (datetime) * `modified`: Timestamp when this organization was last modified. (datetime) * `name`: Name of this organization. (string) * `description`: Optional description of this organization. (string) * `max_hosts`: Maximum number of hosts allowed to be managed by this organization. (integer) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run by this organization. (id)    ## Sorting  To specify that organizations are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersAdminOfOrganizationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersAdminOfOrganizationsListOpts struct

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

# **UsersUsersApplicationsCreate**
> UsersUsersApplicationsCreate(ctx, id, optional)
 Create an Application

 Make a POST request to this resource with the following application fields to create a new application:          * `name`: Name of this application. (string, required) * `description`: Optional description of this application. (string, default=`\"\"`)   * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice, required)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string, default=`\"\"`) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice, required)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean, default=`False`) * `organization`: Organization containing this application. (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersApplicationsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersApplicationsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data69**](Data69.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersApplicationsList**
> UsersUsersApplicationsList(ctx, id, optional)
 List Applications

 Make a GET request to this resource to retrieve the list of applications.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of applications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more application records.    ## Results  Each application data structure includes the following fields:  * `id`: Database ID for this application. (integer) * `type`: Data type for this application. (choice) * `url`: URL for this application. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this application was created. (datetime) * `modified`: Timestamp when this application was last modified. (datetime) * `name`: Name of this application. (string) * `description`: Optional description of this application. (string) * `client_id`:  (string) * `client_secret`: Used for more stringent verification of access to an application when creating a token. (string) * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean) * `organization`: Organization containing this application. (id)    ## Sorting  To specify that applications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersApplicationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersApplicationsListOpts struct

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

# **UsersUsersAuthorizedTokensCreate**
> UsersUsersAuthorizedTokensCreate(ctx, id, optional)
 Create an Access Token for a User

 Make a POST request to this resource with the following access token fields to create a new access token associated with this user.          * `description`: Optional description of this access token. (string, default=`\"\"`)    * `application`:  (id, required)  * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersAuthorizedTokensCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersAuthorizedTokensCreateOpts struct

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

# **UsersUsersAuthorizedTokensList**
> UsersUsersAuthorizedTokensList(ctx, id, optional)
 List Access Tokens for a User

 Make a GET request to this resource to retrieve a list of access tokens associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of access tokens found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more access token records.    ## Results  Each access token data structure includes the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)    ## Sorting  To specify that access tokens are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersAuthorizedTokensListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersAuthorizedTokensListOpts struct

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

# **UsersUsersCreate**
> UsersUsersCreate(ctx, optional)
 Create a User

 Make a POST request to this resource with the following user fields to create a new user:          * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string, required) * `first_name`:  (string, default=`\"\"`) * `last_name`:  (string, default=`\"\"`) * `email`:  (string, default=`\"\"`) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean, default=`False`) * `is_system_auditor`:  (boolean, default=`False`) * `password`: Write-only field used to change the password. (string, default=`\"\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersUsersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersCreateOpts struct

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

# **UsersUsersCredentialsCreate**
> UsersUsersCredentialsCreate(ctx, id, optional)
 Create a Credential for a User

 Make a POST request to this resource with the following credential fields to create a new credential associated with this user.          * `name`: Name of this credential. (string, required) * `description`: Optional description of this credential. (string, default=`\"\"`) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id, required)  * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json, default=`{}`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersCredentialsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersCredentialsCreateOpts struct

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

# **UsersUsersCredentialsList**
> UsersUsersCredentialsList(ctx, id, optional)
 List Credentials for a User

 Make a GET request to this resource to retrieve a list of credentials associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of credentials found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more credential records.    ## Results  Each credential data structure includes the following fields:  * `id`: Database ID for this credential. (integer) * `type`: Data type for this credential. (choice) * `url`: URL for this credential. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this credential was created. (datetime) * `modified`: Timestamp when this credential was last modified. (datetime) * `name`: Name of this credential. (string) * `description`: Optional description of this credential. (string) * `credential_type`: Specify the type of credential you want to create. Refer to the documentation for details on each type. (id) * `managed`:  (boolean) * `inputs`: Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax. (json) * `kind`:  (field) * `cloud`:  (field) * `kubernetes`:  (field)     ## Sorting  To specify that credentials are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersCredentialsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersCredentialsListOpts struct

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

# **UsersUsersDelete**
> UsersUsersDelete(ctx, id, optional)
 Delete a User

 Make a DELETE request to this resource to delete this user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersDeleteOpts struct

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

# **UsersUsersList**
> UsersUsersList(ctx, optional)
 List Users

 Make a GET request to this resource to retrieve the list of users.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of users found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more user records.    ## Results  Each user data structure includes the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)    ## Sorting  To specify that users are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=username  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-username  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=username,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersListOpts struct

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

# **UsersUsersOrganizationsList**
> UsersUsersOrganizationsList(ctx, id, optional)
 List Organizations for a User

 Make a GET request to this resource to retrieve a list of organizations associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of organizations found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more organization records.    ## Results  Each organization data structure includes the following fields:  * `id`: Database ID for this organization. (integer) * `type`: Data type for this organization. (choice) * `url`: URL for this organization. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this organization was created. (datetime) * `modified`: Timestamp when this organization was last modified. (datetime) * `name`: Name of this organization. (string) * `description`: Optional description of this organization. (string) * `max_hosts`: Maximum number of hosts allowed to be managed by this organization. (integer) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run by this organization. (id)    ## Sorting  To specify that organizations are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersOrganizationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersOrganizationsListOpts struct

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

# **UsersUsersPartialUpdate**
> UsersUsersPartialUpdate(ctx, id, optional)
 Update a User

 Make a PUT or PATCH request to this resource to update this user.  The following fields may be modified:          * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string, required) * `first_name`:  (string, default=`\"\"`) * `last_name`:  (string, default=`\"\"`) * `email`:  (string, default=`\"\"`) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean, default=`False`) * `is_system_auditor`:  (boolean, default=`False`) * `password`: Write-only field used to change the password. (string, default=`\"\"`)            For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersPartialUpdateOpts struct

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

# **UsersUsersPersonalTokensCreate**
> UsersUsersPersonalTokensCreate(ctx, id, optional)
 Create an Access Token for a User

 Make a POST request to this resource with the following access token fields to create a new access token associated with this user.          * `description`: Optional description of this access token. (string, default=`\"\"`)      * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersPersonalTokensCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersPersonalTokensCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Data70**](Data70.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPersonalTokensList**
> UsersUsersPersonalTokensList(ctx, id, optional)
 List Access Tokens for a User

 Make a GET request to this resource to retrieve a list of access tokens associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of access tokens found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more access token records.    ## Results  Each access token data structure includes the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)    ## Sorting  To specify that access tokens are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersPersonalTokensListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersPersonalTokensListOpts struct

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

# **UsersUsersProjectsList**
> UsersUsersProjectsList(ctx, id, optional)
 List Projects for a User

 Make a GET request to this resource to retrieve a list of projects associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of projects found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more project records.    ## Results  Each project data structure includes the following fields:  * `id`: Database ID for this project. (integer) * `type`: Data type for this project. (choice) * `url`: URL for this project. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this project was created. (datetime) * `modified`: Timestamp when this project was last modified. (datetime) * `name`: Name of this project. (string) * `description`: Optional description of this project. (string) * `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string) * `scm_type`: Specifies the source control system used to store the project. (choice)     - `\"\"`: Manual     - `git`: Git     - `svn`: Subversion     - `insights`: Red Hat Insights     - `archive`: Remote Archive * `scm_url`: The location where the project is stored. (string) * `scm_branch`: Specific branch, tag or commit to checkout. (string) * `scm_refspec`: For git projects, an additional refspec to fetch. (string) * `scm_clean`: Discard any local changes before syncing the project. (boolean) * `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean) * `scm_delete_on_update`: Delete the project before syncing. (boolean) * `credential`:  (id) * `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer) * `scm_revision`: The last revision fetched by a project update (string) * `last_job_run`:  (datetime) * `last_job_failed`:  (boolean) * `next_job_run`:  (datetime) * `status`:  (choice)     - `new`: New     - `pending`: Pending     - `waiting`: Waiting     - `running`: Running     - `successful`: Successful     - `failed`: Failed     - `error`: Error     - `canceled`: Canceled     - `never updated`: Never Updated     - `ok`: OK     - `missing`: Missing * `organization`: The organization used to determine access to this template. (id) * `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean) * `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer) * `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean) * `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string) * `default_environment`: The default execution environment for jobs run using this project. (id) * `last_update_failed`:  (boolean) * `last_updated`:  (datetime)    ## Sorting  To specify that projects are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersProjectsListOpts struct

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

# **UsersUsersRead**
> UsersUsersRead(ctx, id, optional)
 Retrieve a User

 Make GET request to this resource to retrieve a single user record containing the following fields:  * `id`: Database ID for this user. (integer) * `type`: Data type for this user. (choice) * `url`: URL for this user. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this user was created. (datetime) * `modified`: Timestamp when this user was last modified. (datetime) * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string) * `first_name`:  (string) * `last_name`:  (string) * `email`:  (string) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean) * `is_system_auditor`:  (boolean)  * `ldap_dn`:  (string) * `last_login`:  (datetime) * `external_account`: Set if the account is managed by an external service (field)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersReadOpts struct

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

# **UsersUsersRolesCreate**
> UsersUsersRolesCreate(ctx, id, optional)
 Associate Roles with this User

 Make a POST request to this resource to add or remove a role from this user. The following fields may be modified:     * `id`: The Role ID to add to the user. (int, required)    * `disassociate`: Provide if you want to remove the role. (any value, optional)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersRolesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersRolesCreateOpts struct

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

# **UsersUsersRolesList**
> UsersUsersRolesList(ctx, id, optional)
 List Roles for a User

 Make a GET request to this resource to retrieve a list of roles associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of roles found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more role records.    ## Results  Each role data structure includes the following fields:  * `id`: Database ID for this role. (integer) * `type`: Data type for this role. (choice) * `url`: URL for this role. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `name`: Name of this role. (field) * `description`: Optional description of this role. (field)    ## Sorting  To specify that roles are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersRolesListOpts struct

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

# **UsersUsersTeamsList**
> UsersUsersTeamsList(ctx, id, optional)
 List Teams for a User

 Make a GET request to this resource to retrieve a list of teams associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of teams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more team records.    ## Results  Each team data structure includes the following fields:  * `id`: Database ID for this team. (integer) * `type`: Data type for this team. (choice) * `url`: URL for this team. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this team was created. (datetime) * `modified`: Timestamp when this team was last modified. (datetime) * `name`: Name of this team. (string) * `description`: Optional description of this team. (string) * `organization`:  (id)    ## Sorting  To specify that teams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersTeamsListOpts struct

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

# **UsersUsersTokensCreate**
> UsersUsersTokensCreate(ctx, id, optional)
 Create an Access Token for a User

 Make a POST request to this resource with the following access token fields to create a new access token associated with this user.          * `description`: Optional description of this access token. (string, default=`\"\"`)    * `application`:  (id, default=``)  * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersTokensCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersTokensCreateOpts struct

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

# **UsersUsersTokensList**
> UsersUsersTokensList(ctx, id, optional)
 List Access Tokens for a User

 Make a GET request to this resource to retrieve a list of access tokens associated with the selected user.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of access tokens found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more access token records.    ## Results  Each access token data structure includes the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)    ## Sorting  To specify that access tokens are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersTokensListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersTokensListOpts struct

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

# **UsersUsersUpdate**
> UsersUsersUpdate(ctx, id, optional)
 Update a User

 Make a PUT or PATCH request to this resource to update this user.  The following fields may be modified:          * `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string, required) * `first_name`:  (string, default=`\"\"`) * `last_name`:  (string, default=`\"\"`) * `email`:  (string, default=`\"\"`) * `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean, default=`False`) * `is_system_auditor`:  (boolean, default=`False`) * `password`: Write-only field used to change the password. (string, default=`\"\"`)          For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***UsersApiUsersUsersUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data68**](Data68.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

