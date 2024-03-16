# \AnnotationsApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAnnotationsDelete**](AnnotationsApi.md#ApiAnnotationsDelete) | **Delete** /api/annotations/{id}/ | Delete annotation
[**ApiAnnotationsPartialUpdate**](AnnotationsApi.md#ApiAnnotationsPartialUpdate) | **Patch** /api/annotations/{id}/ | Update annotation
[**ApiAnnotationsRead**](AnnotationsApi.md#ApiAnnotationsRead) | **Get** /api/annotations/{id}/ | Get annotation by its ID
[**ApiTasksAnnotationsCreate**](AnnotationsApi.md#ApiTasksAnnotationsCreate) | **Post** /api/tasks/{id}/annotations/ | Create annotation
[**ApiTasksAnnotationsList**](AnnotationsApi.md#ApiTasksAnnotationsList) | **Get** /api/tasks/{id}/annotations/ | Get all task annotations


# **ApiAnnotationsDelete**
> ApiAnnotationsDelete(ctx, id)
Delete annotation

Delete an annotation. This action can't be undone!

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this annotation. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAnnotationsPartialUpdate**
> Annotation ApiAnnotationsPartialUpdate(ctx, id, data)
Update annotation

Update existing attributes on an annotation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this annotation. | 
  **data** | [**Annotation**](Annotation.md)|  | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAnnotationsRead**
> Annotation ApiAnnotationsRead(ctx, id)
Get annotation by its ID

Retrieve a specific annotation for a task using the annotation result ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this annotation. | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksAnnotationsCreate**
> Annotation ApiTasksAnnotationsCreate(ctx, data, id)
Create annotation

 Add annotations to a task like an annotator does. The content of the result field depends on your  labeling configuration. For example, send the following data as part of your POST  request to send an empty annotation with the ID of the user who completed the task:  ```json { \"result\": {}, \"was_cancelled\": true, \"ground_truth\": true, \"lead_time\": 0, \"task\": 0 \"completed_by\": 123 }  ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Annotation**](Annotation.md)|  | 
  **id** | **int32**| Task ID | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTasksAnnotationsList**
> []Annotation ApiTasksAnnotationsList(ctx, id)
Get all task annotations

List all annotations for a task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Task ID | 

### Return type

[**[]Annotation**](Annotation.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

