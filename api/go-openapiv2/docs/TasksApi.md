# \TasksApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTasksCreate**](TasksApi.md#ApiTasksCreate) | **Post** /api/tasks/ | Create task
[**ApiTasksDelete**](TasksApi.md#ApiTasksDelete) | **Delete** /api/tasks/{id}/ | Delete task
[**ApiTasksList**](TasksApi.md#ApiTasksList) | **Get** /api/tasks/ | Get tasks list
[**ApiTasksPartialUpdate**](TasksApi.md#ApiTasksPartialUpdate) | **Patch** /api/tasks/{id}/ | Update task
[**ApiTasksRead**](TasksApi.md#ApiTasksRead) | **Get** /api/tasks/{id}/ | Get task


# **ApiTasksCreate**
> BaseTask ApiTasksCreate(ctx, data)
Create task

Create a new labeling task in Label Studio.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BaseTask**](BaseTask.md)|  | 

### Return type

[**BaseTask**](BaseTask.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksDelete**
> ApiTasksDelete(ctx, id)
Delete task

Delete a task in Label Studio. This action cannot be undone!

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Task ID | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksList**
> []BaseTask ApiTasksList(ctx, optional)
Get tasks list

 Retrieve a list of tasks with pagination for a specific view or project, by using filters and ordering. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TasksApiApiTasksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiApiTasksListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.Int32**| View ID | 
 **project** | **optional.Int32**| Project ID | 
 **resolveUri** | **optional.Bool**| Resolve task data URIs using Cloud Storage | 

### Return type

[**[]BaseTask**](BaseTask.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksPartialUpdate**
> TaskSimple ApiTasksPartialUpdate(ctx, data, id)
Update task

Update the attributes of an existing labeling task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**TaskSimple**](TaskSimple.md)|  | 
  **id** | **string**| Task ID | 

### Return type

[**TaskSimple**](TaskSimple.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksRead**
> DataManagerTaskSerializer ApiTasksRead(ctx, id)
Get task

 Get task data, metadata, annotations and other attributes for a specific labeling task by task ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Task ID | 

### Return type

[**DataManagerTaskSerializer**](data_manager_task_serializer.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

