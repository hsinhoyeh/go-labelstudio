# \PredictionsApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPredictionsCreate**](PredictionsApi.md#ApiPredictionsCreate) | **Post** /api/predictions/ | Create prediction
[**ApiPredictionsDelete**](PredictionsApi.md#ApiPredictionsDelete) | **Delete** /api/predictions/{id}/ | Delete prediction
[**ApiPredictionsList**](PredictionsApi.md#ApiPredictionsList) | **Get** /api/predictions/ | List predictions
[**ApiPredictionsPartialUpdate**](PredictionsApi.md#ApiPredictionsPartialUpdate) | **Patch** /api/predictions/{id}/ | Update prediction
[**ApiPredictionsRead**](PredictionsApi.md#ApiPredictionsRead) | **Get** /api/predictions/{id}/ | Get prediction details
[**ApiPredictionsUpdate**](PredictionsApi.md#ApiPredictionsUpdate) | **Put** /api/predictions/{id}/ | Put prediction


# **ApiPredictionsCreate**
> Prediction ApiPredictionsCreate(ctx, data)
Create prediction

Create a prediction for a specific task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Prediction**](Prediction.md)|  | 

### Return type

[**Prediction**](Prediction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPredictionsDelete**
> ApiPredictionsDelete(ctx, id)
Delete prediction

Delete a prediction by prediction ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Prediction ID | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPredictionsList**
> []Prediction ApiPredictionsList(ctx, )
List predictions

List all predictions and their IDs.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Prediction**](Prediction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPredictionsPartialUpdate**
> Prediction ApiPredictionsPartialUpdate(ctx, data, id)
Update prediction

Update prediction data by prediction ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Prediction**](Prediction.md)|  | 
  **id** | **int32**| Prediction ID | 

### Return type

[**Prediction**](Prediction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPredictionsRead**
> Prediction ApiPredictionsRead(ctx, id)
Get prediction details

Get details about a specific prediction by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Prediction ID | 

### Return type

[**Prediction**](Prediction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPredictionsUpdate**
> Prediction ApiPredictionsUpdate(ctx, data, id)
Put prediction

Overwrite prediction data by prediction ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Prediction**](Prediction.md)|  | 
  **id** | **int32**| Prediction ID | 

### Return type

[**Prediction**](Prediction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

