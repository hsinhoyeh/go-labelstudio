# \MachineLearningApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMlCreate**](MachineLearningApi.md#ApiMlCreate) | **Post** /api/ml/ | Add ML Backend
[**ApiMlDelete**](MachineLearningApi.md#ApiMlDelete) | **Delete** /api/ml/{id} | Remove ML Backend
[**ApiMlInteractiveAnnotatingCreate**](MachineLearningApi.md#ApiMlInteractiveAnnotatingCreate) | **Post** /api/ml/{id}/interactive-annotating | Request Interactive Annotation
[**ApiMlList**](MachineLearningApi.md#ApiMlList) | **Get** /api/ml/ | List ML backends
[**ApiMlPartialUpdate**](MachineLearningApi.md#ApiMlPartialUpdate) | **Patch** /api/ml/{id} | Update ML Backend
[**ApiMlRead**](MachineLearningApi.md#ApiMlRead) | **Get** /api/ml/{id} | Get ML Backend
[**ApiMlTrainCreate**](MachineLearningApi.md#ApiMlTrainCreate) | **Post** /api/ml/{id}/train | Train
[**ApiMlVersionsRead**](MachineLearningApi.md#ApiMlVersionsRead) | **Get** /api/ml/{id}/versions | Get model versions


# **ApiMlCreate**
> Data ApiMlCreate(ctx, data)
Add ML Backend

 Add an ML backend to a project using the Label Studio UI or by sending a POST request using the following cURL  command: ```bash curl -X POST -H 'Content-type: application/json' https://localhost:8080/api/ml -H 'Authorization: Token abc123'\\ --data '{\"url\": \"http://localhost:9090\", \"project\": {project_id}}'  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Data**](Data.md)|  | 

### Return type

[**Data**](data.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlDelete**
> ApiMlDelete(ctx, id)
Remove ML Backend

 Remove an existing ML backend connection by ID. For example, use the following cURL command: ```bash curl -X DELETE https://localhost:8080/api/ml/{ml_backend_ID} -H 'Authorization: Token abc123' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this ml backend. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlInteractiveAnnotatingCreate**
> ApiMlInteractiveAnnotatingCreate(ctx, data, id)
Request Interactive Annotation

 Send a request to the machine learning backend set up to be used for interactive preannotations to retrieve a predicted region based on annotator input.  See [set up machine learning](https://labelstud.io/guide/ml.html#Get-interactive-preannotations) for more. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MlInteractiveAnnotatingRequest**](MlInteractiveAnnotatingRequest.md)|  | 
  **id** | **int32**| A unique integer value identifying this ML backend. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlList**
> []MlBackend ApiMlList(ctx, optional)
List ML backends

 List all configured ML backends for a specific project by ID. Use the following cURL command: ```bash curl https://localhost:8080/api/ml?project={project_id} -H 'Authorization: Token abc123' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MachineLearningApiApiMlListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MachineLearningApiApiMlListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]MlBackend**](MLBackend.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlPartialUpdate**
> MlBackend ApiMlPartialUpdate(ctx, id, data)
Update ML Backend

 Update ML backend parameters using the Label Studio UI or by sending a PATCH request using the following cURL command: ```bash curl -X PATCH -H 'Content-type: application/json' https://localhost:8080/api/ml/{ml_backend_ID} -H 'Authorization: Token abc123'\\ --data '{\"url\": \"http://localhost:9091\"}'  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this ml backend. | 
  **data** | [**MlBackend**](MlBackend.md)|  | 

### Return type

[**MlBackend**](MLBackend.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlRead**
> MlBackend ApiMlRead(ctx, id)
Get ML Backend

 Get details about a specific ML backend connection by ID. For example, make a GET request using the following cURL command: ```bash curl https://localhost:8080/api/ml/{ml_backend_ID} -H 'Authorization: Token abc123' 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this ml backend. | 

### Return type

[**MlBackend**](MLBackend.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlTrainCreate**
> ApiMlTrainCreate(ctx, data, id)
Train

 After you add an ML backend, call this API with the ML backend ID to start training with  already-labeled tasks.   Get the ML backend ID by [listing the ML backends for a project](https://labelstud.io/api/#operation/api_ml_list). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Data1**](Data1.md)|  | 
  **id** | **int32**| A unique integer value identifying this ML backend. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiMlVersionsRead**
> ApiMlVersionsRead(ctx, id)
Get model versions

Get available versions of the model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

