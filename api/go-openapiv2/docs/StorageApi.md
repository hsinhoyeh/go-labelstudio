# \StorageApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesExportList**](StorageApi.md#ApiStoragesExportList) | **Get** /api/storages/export | List all export storages from the project
[**ApiStoragesExportTypesList**](StorageApi.md#ApiStoragesExportTypesList) | **Get** /api/storages/export/types | List all export storages types
[**ApiStoragesList**](StorageApi.md#ApiStoragesList) | **Get** /api/storages/ | List all import storages from the project
[**ApiStoragesTypesList**](StorageApi.md#ApiStoragesTypesList) | **Get** /api/storages/types | List all import storages types


# **ApiStoragesExportList**
> ApiStoragesExportList(ctx, project)
List all export storages from the project

Retrieve a list of the export storages of all types with their IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **int32**| A unique integer value identifying your project. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportTypesList**
> ApiStoragesExportTypesList(ctx, )
List all export storages types

Retrieve a list of the export storages types.

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

# **ApiStoragesList**
> ApiStoragesList(ctx, project)
List all import storages from the project

Retrieve a list of the import storages of all types with their IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **int32**| A unique integer value identifying your project. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesTypesList**
> ApiStoragesTypesList(ctx, )
List all import storages types

Retrieve a list of the import storages types.

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

