# \UsersApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCurrentUserResetTokenCreate**](UsersApi.md#ApiCurrentUserResetTokenCreate) | **Post** /api/current-user/reset-token/ | Reset user token
[**ApiCurrentUserTokenList**](UsersApi.md#ApiCurrentUserTokenList) | **Get** /api/current-user/token | Get user token
[**ApiCurrentUserWhoamiRead**](UsersApi.md#ApiCurrentUserWhoamiRead) | **Get** /api/current-user/whoami | Retrieve my user
[**ApiUsersCreate**](UsersApi.md#ApiUsersCreate) | **Post** /api/users/ | Create new user
[**ApiUsersDelete**](UsersApi.md#ApiUsersDelete) | **Delete** /api/users/{id}/ | Delete user
[**ApiUsersList**](UsersApi.md#ApiUsersList) | **Get** /api/users/ | List users
[**ApiUsersPartialUpdate**](UsersApi.md#ApiUsersPartialUpdate) | **Patch** /api/users/{id}/ | Update user details
[**ApiUsersRead**](UsersApi.md#ApiUsersRead) | **Get** /api/users/{id}/ | Get user info


# **ApiCurrentUserResetTokenCreate**
> InlineResponse201 ApiCurrentUserResetTokenCreate(ctx, )
Reset user token

Reset the user token for the current user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCurrentUserTokenList**
> ApiCurrentUserTokenList(ctx, )
Get user token

Get a user token to authenticate to the API as the current user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCurrentUserWhoamiRead**
> BaseUser ApiCurrentUserWhoamiRead(ctx, )
Retrieve my user

Retrieve details of the account that you are using to access the API.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BaseUser**](BaseUser.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersCreate**
> BaseUser ApiUsersCreate(ctx, data)
Create new user

Create a user in Label Studio.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BaseUser**](BaseUser.md)|  | 

### Return type

[**BaseUser**](BaseUser.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersDelete**
> ApiUsersDelete(ctx, id)
Delete user

Delete a specific Label Studio user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User ID | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersList**
> []BaseUser ApiUsersList(ctx, )
List users

List the users that exist on the Label Studio server.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]BaseUser**](BaseUser.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersPartialUpdate**
> BaseUser ApiUsersPartialUpdate(ctx, data, id)
Update user details

 Update details for a specific user, such as their name or contact information, in Label Studio. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BaseUser**](BaseUser.md)|  | 
  **id** | **int32**| User ID | 

### Return type

[**BaseUser**](BaseUser.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiUsersRead**
> BaseUser ApiUsersRead(ctx, id)
Get user info

Get info about a specific Label Studio user, based on the user ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User ID | 

### Return type

[**BaseUser**](BaseUser.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

