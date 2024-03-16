# \DataManagerApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDmActionsCreate**](DataManagerApi.md#ApiDmActionsCreate) | **Post** /api/dm/actions/ | Post actions
[**ApiDmActionsList**](DataManagerApi.md#ApiDmActionsList) | **Get** /api/dm/actions/ | Get actions
[**ApiDmColumnsList**](DataManagerApi.md#ApiDmColumnsList) | **Get** /api/dm/columns/ | Get data manager columns
[**ApiDmProjectList**](DataManagerApi.md#ApiDmProjectList) | **Get** /api/dm/project/ | Get project state
[**ApiDmViewsCreate**](DataManagerApi.md#ApiDmViewsCreate) | **Post** /api/dm/views/ | Create view
[**ApiDmViewsDelete**](DataManagerApi.md#ApiDmViewsDelete) | **Delete** /api/dm/views/{id}/ | Delete view
[**ApiDmViewsList**](DataManagerApi.md#ApiDmViewsList) | **Get** /api/dm/views/ | List views
[**ApiDmViewsPartialUpdate**](DataManagerApi.md#ApiDmViewsPartialUpdate) | **Patch** /api/dm/views/{id}/ | Update view
[**ApiDmViewsRead**](DataManagerApi.md#ApiDmViewsRead) | **Get** /api/dm/views/{id}/ | Get view details
[**ApiDmViewsReset**](DataManagerApi.md#ApiDmViewsReset) | **Delete** /api/dm/views/reset/ | Reset project views
[**ApiDmViewsUpdate**](DataManagerApi.md#ApiDmViewsUpdate) | **Put** /api/dm/views/{id}/ | Put view


# **ApiDmActionsCreate**
> ApiDmActionsCreate(ctx, )
Post actions

Perform an action with the selected items from a specific view.

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

# **ApiDmActionsList**
> ApiDmActionsList(ctx, )
Get actions

Retrieve all the registered actions with descriptions that data manager can use.

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

# **ApiDmColumnsList**
> ApiDmColumnsList(ctx, )
Get data manager columns

Retrieve the data manager columns available for the tasks in a specific project.

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

# **ApiDmProjectList**
> ApiDmProjectList(ctx, )
Get project state

Retrieve the project state for the data manager.

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

# **ApiDmViewsCreate**
> View ApiDmViewsCreate(ctx, data)
Create view

Create a view for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**View**](View.md)|  | 

### Return type

[**View**](View.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsDelete**
> ApiDmViewsDelete(ctx, id)
Delete view

Delete a specific view by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsList**
> []View ApiDmViewsList(ctx, optional)
List views

List all views for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataManagerApiApiDmViewsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataManagerApiApiDmViewsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]View**](View.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsPartialUpdate**
> View ApiDmViewsPartialUpdate(ctx, data, id)
Update view

Update view data with additional filters and other information for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**View**](View.md)|  | 
  **id** | **string**| View ID | 

### Return type

[**View**](View.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsRead**
> View ApiDmViewsRead(ctx, id)
Get view details

Get the details about a specific view in the data manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 

### Return type

[**View**](View.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsReset**
> ApiDmViewsReset(ctx, data)
Reset project views

Reset all views for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ViewReset**](ViewReset.md)|  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDmViewsUpdate**
> View ApiDmViewsUpdate(ctx, data, id)
Put view

Overwrite view data with updated filters and other information for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**View**](View.md)|  | 
  **id** | **string**| View ID | 

### Return type

[**View**](View.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

