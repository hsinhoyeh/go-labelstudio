# \StorageAzureApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesAzureCreate**](StorageAzureApi.md#ApiStoragesAzureCreate) | **Post** /api/storages/azure/ | Create new storage
[**ApiStoragesAzureDelete**](StorageAzureApi.md#ApiStoragesAzureDelete) | **Delete** /api/storages/azure/{id} | Delete import storage
[**ApiStoragesAzureList**](StorageAzureApi.md#ApiStoragesAzureList) | **Get** /api/storages/azure/ | Get all import storage
[**ApiStoragesAzurePartialUpdate**](StorageAzureApi.md#ApiStoragesAzurePartialUpdate) | **Patch** /api/storages/azure/{id} | Update import storage
[**ApiStoragesAzureRead**](StorageAzureApi.md#ApiStoragesAzureRead) | **Get** /api/storages/azure/{id} | Get import storage
[**ApiStoragesAzureSyncCreate**](StorageAzureApi.md#ApiStoragesAzureSyncCreate) | **Post** /api/storages/azure/{id}/sync | Sync import storage
[**ApiStoragesAzureValidateCreate**](StorageAzureApi.md#ApiStoragesAzureValidateCreate) | **Post** /api/storages/azure/validate | Validate import storage
[**ApiStoragesExportAzureCreate**](StorageAzureApi.md#ApiStoragesExportAzureCreate) | **Post** /api/storages/export/azure | Create export storage
[**ApiStoragesExportAzureDelete**](StorageAzureApi.md#ApiStoragesExportAzureDelete) | **Delete** /api/storages/export/azure/{id} | Delete export storage
[**ApiStoragesExportAzureList**](StorageAzureApi.md#ApiStoragesExportAzureList) | **Get** /api/storages/export/azure | Get all export storage
[**ApiStoragesExportAzurePartialUpdate**](StorageAzureApi.md#ApiStoragesExportAzurePartialUpdate) | **Patch** /api/storages/export/azure/{id} | Update export storage
[**ApiStoragesExportAzureRead**](StorageAzureApi.md#ApiStoragesExportAzureRead) | **Get** /api/storages/export/azure/{id} | Get export storage
[**ApiStoragesExportAzureSyncCreate**](StorageAzureApi.md#ApiStoragesExportAzureSyncCreate) | **Post** /api/storages/export/azure/{id}/sync | Sync export storage
[**ApiStoragesExportAzureValidateCreate**](StorageAzureApi.md#ApiStoragesExportAzureValidateCreate) | **Post** /api/storages/export/azure/validate | Validate export storage


# **ApiStoragesAzureCreate**
> AzureBlobImportStorage ApiStoragesAzureCreate(ctx, data)
Create new storage

Get new Azure import storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AzureBlobImportStorage**](AzureBlobImportStorage.md)|  | 

### Return type

[**AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzureDelete**
> ApiStoragesAzureDelete(ctx, id)
Delete import storage

Delete a specific Azure import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob import storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzureList**
> []AzureBlobImportStorage ApiStoragesAzureList(ctx, optional)
Get all import storage

Get list of all Azure import storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageAzureApiApiStoragesAzureListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageAzureApiApiStoragesAzureListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzurePartialUpdate**
> AzureBlobImportStorage ApiStoragesAzurePartialUpdate(ctx, id, data)
Update import storage

Update a specific Azure import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob import storage. | 
  **data** | [**AzureBlobImportStorage**](AzureBlobImportStorage.md)|  | 

### Return type

[**AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzureRead**
> AzureBlobImportStorage ApiStoragesAzureRead(ctx, id)
Get import storage

Get a specific Azure import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob import storage. | 

### Return type

[**AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzureSyncCreate**
> AzureBlobImportStorage ApiStoragesAzureSyncCreate(ctx, id, data)
Sync import storage

Sync tasks from an Azure import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AzureBlobImportStorage**](AzureBlobImportStorage.md)|  | 

### Return type

[**AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesAzureValidateCreate**
> AzureBlobImportStorage ApiStoragesAzureValidateCreate(ctx, data)
Validate import storage

Validate a specific Azure import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AzureBlobImportStorage**](AzureBlobImportStorage.md)|  | 

### Return type

[**AzureBlobImportStorage**](AzureBlobImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureCreate**
> AzureBlobExportStorage ApiStoragesExportAzureCreate(ctx, data)
Create export storage

Create a new Azure export storage connection to store annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AzureBlobExportStorage**](AzureBlobExportStorage.md)|  | 

### Return type

[**AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureDelete**
> ApiStoragesExportAzureDelete(ctx, id)
Delete export storage

Delete a specific Azure export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob export storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureList**
> []AzureBlobExportStorage ApiStoragesExportAzureList(ctx, optional)
Get all export storage

Get a list of all Azure export storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageAzureApiApiStoragesExportAzureListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageAzureApiApiStoragesExportAzureListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzurePartialUpdate**
> AzureBlobExportStorage ApiStoragesExportAzurePartialUpdate(ctx, id, data)
Update export storage

Update a specific Azure export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob export storage. | 
  **data** | [**AzureBlobExportStorage**](AzureBlobExportStorage.md)|  | 

### Return type

[**AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureRead**
> AzureBlobExportStorage ApiStoragesExportAzureRead(ctx, id)
Get export storage

Get a specific Azure export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this azure blob export storage. | 

### Return type

[**AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureSyncCreate**
> AzureBlobExportStorage ApiStoragesExportAzureSyncCreate(ctx, id, data)
Sync export storage

Sync tasks from an Azure export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**AzureBlobExportStorage**](AzureBlobExportStorage.md)|  | 

### Return type

[**AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportAzureValidateCreate**
> AzureBlobExportStorage ApiStoragesExportAzureValidateCreate(ctx, data)
Validate export storage

Validate a specific Azure export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AzureBlobExportStorage**](AzureBlobExportStorage.md)|  | 

### Return type

[**AzureBlobExportStorage**](AzureBlobExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

