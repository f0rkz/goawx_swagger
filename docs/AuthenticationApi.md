# \AuthenticationApi

All URIs are relative to *https://null*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationApplicationsActivityStreamList**](AuthenticationApi.md#AuthenticationApplicationsActivityStreamList) | **Get** /api/v2/applications/{id}/activity_stream/ |  List Activity Streams for an Application
[**AuthenticationApplicationsCreate0**](AuthenticationApi.md#AuthenticationApplicationsCreate0) | **Post** /api/v2/applications/ |  Create an Application
[**AuthenticationApplicationsDelete0**](AuthenticationApi.md#AuthenticationApplicationsDelete0) | **Delete** /api/v2/applications/{id}/ |  Delete an Application
[**AuthenticationApplicationsList0**](AuthenticationApi.md#AuthenticationApplicationsList0) | **Get** /api/v2/applications/ |  List Applications
[**AuthenticationApplicationsPartialUpdate0**](AuthenticationApi.md#AuthenticationApplicationsPartialUpdate0) | **Patch** /api/v2/applications/{id}/ |  Update an Application
[**AuthenticationApplicationsRead0**](AuthenticationApi.md#AuthenticationApplicationsRead0) | **Get** /api/v2/applications/{id}/ |  Retrieve an Application
[**AuthenticationApplicationsTokensCreate0**](AuthenticationApi.md#AuthenticationApplicationsTokensCreate0) | **Post** /api/v2/applications/{id}/tokens/ |  Create an Access Token for an Application
[**AuthenticationApplicationsTokensList0**](AuthenticationApi.md#AuthenticationApplicationsTokensList0) | **Get** /api/v2/applications/{id}/tokens/ |  List Access Tokens for an Application
[**AuthenticationApplicationsUpdate0**](AuthenticationApi.md#AuthenticationApplicationsUpdate0) | **Put** /api/v2/applications/{id}/ |  Update an Application
[**AuthenticationOList**](AuthenticationApi.md#AuthenticationOList) | **Get** /api/o/ |  Token Handling using OAuth2
[**AuthenticationTokensActivityStreamList**](AuthenticationApi.md#AuthenticationTokensActivityStreamList) | **Get** /api/v2/tokens/{id}/activity_stream/ |  List Activity Streams for an Access Token
[**AuthenticationTokensCreate0**](AuthenticationApi.md#AuthenticationTokensCreate0) | **Post** /api/v2/tokens/ |  Create an Access Token
[**AuthenticationTokensDelete**](AuthenticationApi.md#AuthenticationTokensDelete) | **Delete** /api/v2/tokens/{id}/ |  Delete an Access Token
[**AuthenticationTokensList0**](AuthenticationApi.md#AuthenticationTokensList0) | **Get** /api/v2/tokens/ |  List Access Tokens
[**AuthenticationTokensPartialUpdate**](AuthenticationApi.md#AuthenticationTokensPartialUpdate) | **Patch** /api/v2/tokens/{id}/ |  Update an Access Token
[**AuthenticationTokensRead**](AuthenticationApi.md#AuthenticationTokensRead) | **Get** /api/v2/tokens/{id}/ |  Retrieve an Access Token
[**AuthenticationTokensUpdate**](AuthenticationApi.md#AuthenticationTokensUpdate) | **Put** /api/v2/tokens/{id}/ |  Update an Access Token


# **AuthenticationApplicationsActivityStreamList**
> AuthenticationApplicationsActivityStreamList(ctx, id, optional)
 List Activity Streams for an Application

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected application.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsActivityStreamListOpts struct

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

# **AuthenticationApplicationsCreate0**
> AuthenticationApplicationsCreate0(ctx, optional)
 Create an Application

 Make a POST request to this resource with the following application fields to create a new application:          * `name`: Name of this application. (string, required) * `description`: Optional description of this application. (string, default=`\"\"`)   * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice, required)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string, default=`\"\"`) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice, required)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean, default=`False`) * `organization`: Organization containing this application. (id, required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationApiAuthenticationApplicationsCreate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsCreate0Opts struct

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

# **AuthenticationApplicationsDelete0**
> AuthenticationApplicationsDelete0(ctx, id, optional)
 Delete an Application

 Make a DELETE request to this resource to delete this application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsDelete0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsDelete0Opts struct

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

# **AuthenticationApplicationsList0**
> AuthenticationApplicationsList0(ctx, optional)
 List Applications

 Make a GET request to this resource to retrieve the list of applications.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of applications found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more application records.    ## Results  Each application data structure includes the following fields:  * `id`: Database ID for this application. (integer) * `type`: Data type for this application. (choice) * `url`: URL for this application. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this application was created. (datetime) * `modified`: Timestamp when this application was last modified. (datetime) * `name`: Name of this application. (string) * `description`: Optional description of this application. (string) * `client_id`:  (string) * `client_secret`: Used for more stringent verification of access to an application when creating a token. (string) * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean) * `organization`: Organization containing this application. (id)    ## Sorting  To specify that applications are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationApiAuthenticationApplicationsList0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsList0Opts struct

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

# **AuthenticationApplicationsPartialUpdate0**
> AuthenticationApplicationsPartialUpdate0(ctx, id, optional)
 Update an Application

 Make a PUT or PATCH request to this resource to update this application.  The following fields may be modified:          * `name`: Name of this application. (string, required) * `description`: Optional description of this application. (string, default=`\"\"`)   * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice, required)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string, default=`\"\"`) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice, required)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean, default=`False`) * `organization`: Organization containing this application. (id, required)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsPartialUpdate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsPartialUpdate0Opts struct

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

# **AuthenticationApplicationsRead0**
> AuthenticationApplicationsRead0(ctx, id, optional)
 Retrieve an Application

 Make GET request to this resource to retrieve a single application record containing the following fields:  * `id`: Database ID for this application. (integer) * `type`: Data type for this application. (choice) * `url`: URL for this application. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this application was created. (datetime) * `modified`: Timestamp when this application was last modified. (datetime) * `name`: Name of this application. (string) * `description`: Optional description of this application. (string) * `client_id`:  (string) * `client_secret`: Used for more stringent verification of access to an application when creating a token. (string) * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean) * `organization`: Organization containing this application. (id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsRead0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsRead0Opts struct

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

# **AuthenticationApplicationsTokensCreate0**
> AuthenticationApplicationsTokensCreate0(ctx, id, optional)
 Create an Access Token for an Application

 Make a POST request to this resource with the following access token fields to create a new access token associated with this application.          * `description`: Optional description of this access token. (string, default=`\"\"`)      * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsTokensCreate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsTokensCreate0Opts struct

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

# **AuthenticationApplicationsTokensList0**
> AuthenticationApplicationsTokensList0(ctx, id, optional)
 List Access Tokens for an Application

 Make a GET request to this resource to retrieve a list of access tokens associated with the selected application.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of access tokens found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more access token records.    ## Results  Each access token data structure includes the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)    ## Sorting  To specify that access tokens are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsTokensList0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsTokensList0Opts struct

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

# **AuthenticationApplicationsUpdate0**
> AuthenticationApplicationsUpdate0(ctx, id, optional)
 Update an Application

 Make a PUT or PATCH request to this resource to update this application.  The following fields may be modified:          * `name`: Name of this application. (string, required) * `description`: Optional description of this application. (string, default=`\"\"`)   * `client_type`: Set to Public or Confidential depending on how secure the client device is. (choice, required)     - `confidential`: Confidential     - `public`: Public * `redirect_uris`: Allowed URIs list, space separated (string, default=`\"\"`) * `authorization_grant_type`: The Grant type the user must use for acquire tokens for this application. (choice, required)     - `authorization-code`: Authorization code     - `password`: Resource owner password-based * `skip_authorization`: Set True to skip authorization step for completely trusted applications. (boolean, default=`False`) * `organization`: Organization containing this application. (id, required)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationApplicationsUpdate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationApplicationsUpdate0Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data**](Data.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationOList**
> AuthenticationOList(ctx, )
 Token Handling using OAuth2

 This page lists OAuth 2 utility endpoints used for authorization, token refresh and revoke. Note endpoints other than `/api/o/authorize/` are not meant to be used in browsers and do not support HTTP GET. The endpoints here strictly follow [RFC specs for OAuth2](https://tools.ietf.org/html/rfc6749), so please use that for detailed reference. Note AWX net location default to `http://localhost:8013` in examples:   ## Create Token for an Application using Authorization code grant type Given an application \"AuthCodeApp\" of grant type `authorization-code`,  from the client app, the user makes a GET to the Authorize endpoint with   * `response_type` * `client_id` * `redirect_uris` * `scope`    AWX will respond with the authorization `code` and `state` to the redirect_uri specified in the application. The client application will then make a POST to the `api/o/token/` endpoint on AWX with  * `code` * `client_id` * `client_secret` * `grant_type` * `redirect_uri`  AWX will respond with the `access_token`, `token_type`, `refresh_token`, and `expires_in`. For more information on testing this flow, refer to [django-oauth-toolkit](http://django-oauth-toolkit.readthedocs.io/en/latest/tutorial/tutorial_01.html#test-your-authorization-server).   ## Create Token for an Application using Password grant type  Log in is not required for `password` grant type, so a simple `curl` can be used to acquire a personal access token via `/api/o/token/` with   * `grant_type`: Required to be \"password\" * `username` * `password` * `client_id`: Associated application must have grant_type \"password\" * `client_secret`  For example:  ```bash curl -X POST \\   -H \"Content-Type: application/x-www-form-urlencoded\" \\   -d \"grant_type=password&username=<username>&password=<password>&scope=read\" \\   -u \"gwSPoasWSdNkMDtBN3Hu2WYQpPWCO9SwUEsKK22l:fI6ZpfocHYBGfm1tP92r0yIgCyfRdDQt0Tos9L8a4fNsJjQQMwp9569e IaUBsaVDgt2eiwOGe0bg5m5vCSstClZmtdy359RVx2rQK5YlIWyPlrolpt2LEpVeKXWaiybo\" \\   http://localhost:8013/api/o/token/ -i ``` In the above post request, parameters `username` and `password` are username and password of the related AWX user of the underlying application, and the authentication information is of format `<client_id>:<client_secret>`, where `client_id` and `client_secret` are the corresponding fields of underlying application.  Upon success, access token, refresh token and other information are given in the response body in JSON

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

# **AuthenticationTokensActivityStreamList**
> AuthenticationTokensActivityStreamList(ctx, id, optional)
 List Activity Streams for an Access Token

 Make a GET request to this resource to retrieve a list of activity streams associated with the selected access token.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of activity streams found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more activity stream records.    ## Results  Each activity stream data structure includes the following fields:  * `id`: Database ID for this activity stream. (integer) * `type`: Data type for this activity stream. (choice) * `url`: URL for this activity stream. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `timestamp`:  (datetime) * `operation`: The action taken with respect to the given object(s). (choice)     - `create`: Entity Created     - `update`: Entity Updated     - `delete`: Entity Deleted     - `associate`: Entity Associated with another Entity     - `disassociate`: Entity was Disassociated with another Entity * `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json) * `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string) * `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string) * `object_association`: When present, shows the field name of the role or relationship that changed. (field) * `action_node`: The cluster node the activity took place on. (string) * `object_type`: When present, shows the model on which the role or relationship was defined. (field)    ## Sorting  To specify that activity streams are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationTokensActivityStreamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensActivityStreamListOpts struct

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

# **AuthenticationTokensCreate0**
> AuthenticationTokensCreate0(ctx, optional)
 Create an Access Token

 Make a POST request to this resource with the following access token fields to create a new access token:          * `description`: Optional description of this access token. (string, default=`\"\"`)    * `application`:  (id, default=``)  * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationApiAuthenticationTokensCreate0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensCreate0Opts struct

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

# **AuthenticationTokensDelete**
> AuthenticationTokensDelete(ctx, id, optional)
 Delete an Access Token

 Make a DELETE request to this resource to delete this access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationTokensDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensDeleteOpts struct

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

# **AuthenticationTokensList0**
> AuthenticationTokensList0(ctx, optional)
 List Access Tokens

 Make a GET request to this resource to retrieve the list of access tokens.  The resulting data structure contains:      {         \"count\": 99,         \"next\": null,         \"previous\": null,         \"results\": [             ...         ]     }  The `count` field indicates the total number of access tokens found for the given query.  The `next` and `previous` fields provides links to additional results if there are more than will fit on a single page.  The `results` list contains zero or more access token records.    ## Results  Each access token data structure includes the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)    ## Sorting  To specify that access tokens are returned in a particular order, use the `order_by` query string parameter on the GET request.      ?order_by=name  Prefix the field name with a dash `-` to sort in reverse:      ?order_by=-name  Multiple sorting fields may be specified by separating the field names with a comma `,`:      ?order_by=name,some_other_field  ## Pagination  Use the `page_size` query string parameter to change the number of results returned for each request.  Use the `page` query string parameter to retrieve a particular page of results.      ?page_size=100&page=2  The `previous` and `next` links returned with the results will set these query string parameters automatically.  ## Searching  Use the `search` query string parameter to perform a case-insensitive search within all designated text fields of a model.      ?search=findme  (_Added in Ansible Tower 3.1.0_) Search across related fields:      ?related__search=findme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationApiAuthenticationTokensList0Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensList0Opts struct

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

# **AuthenticationTokensPartialUpdate**
> AuthenticationTokensPartialUpdate(ctx, id, optional)
 Update an Access Token

 Make a PUT or PATCH request to this resource to update this access token.  The following fields may be modified:          * `description`: Optional description of this access token. (string, default=`\"\"`)      * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)         For a PATCH request, include only the fields that are being modified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationTokensPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensPartialUpdateOpts struct

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

# **AuthenticationTokensRead**
> AuthenticationTokensRead(ctx, id, optional)
 Retrieve an Access Token

 Make GET request to this resource to retrieve a single access token record containing the following fields:  * `id`: Database ID for this access token. (integer) * `type`: Data type for this access token. (choice) * `url`: URL for this access token. (string) * `related`: Data structure with URLs of related resources. (object) * `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object) * `created`: Timestamp when this access token was created. (datetime) * `modified`: Timestamp when this access token was last modified. (datetime) * `description`: Optional description of this access token. (string) * `user`: The user representing the token owner (id) * `token`:  (string) * `refresh_token`:  (field) * `application`:  (id) * `expires`:  (datetime) * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationTokensReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensReadOpts struct

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

# **AuthenticationTokensUpdate**
> AuthenticationTokensUpdate(ctx, id, optional)
 Update an Access Token

 Make a PUT or PATCH request to this resource to update this access token.  The following fields may be modified:          * `description`: Optional description of this access token. (string, default=`\"\"`)      * `scope`: Allowed scopes, further restricts user&#39;s permissions. Must be a simple space-separated string with allowed scopes [&#39;read&#39;, &#39;write&#39;]. (string, default=`\"write\"`)       For a PUT request, include **all** fields in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AuthenticationApiAuthenticationTokensUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationTokensUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| A search term. | 
 **data** | [**optional.Interface of Data67**](Data67.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

