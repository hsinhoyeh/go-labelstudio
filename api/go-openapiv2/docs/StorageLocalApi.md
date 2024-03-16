# \StorageLocalApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesExportLocalfilesCreate**](StorageLocalApi.md#ApiStoragesExportLocalfilesCreate) | **Post** /api/storages/export/localfiles | Create export storage
[**ApiStoragesExportLocalfilesDelete**](StorageLocalApi.md#ApiStoragesExportLocalfilesDelete) | **Delete** /api/storages/export/localfiles/{id} | Delete export storage
[**ApiStoragesExportLocalfilesList**](StorageLocalApi.md#ApiStoragesExportLocalfilesList) | **Get** /api/storages/export/localfiles | Get all export storage
[**ApiStoragesExportLocalfilesPartialUpdate**](StorageLocalApi.md#ApiStoragesExportLocalfilesPartialUpdate) | **Patch** /api/storages/export/localfiles/{id} | Update export storage
[**ApiStoragesExportLocalfilesRead**](StorageLocalApi.md#ApiStoragesExportLocalfilesRead) | **Get** /api/storages/export/localfiles/{id} | Get export storage
[**ApiStoragesExportLocalfilesSyncCreate**](StorageLocalApi.md#ApiStoragesExportLocalfilesSyncCreate) | **Post** /api/storages/export/localfiles/{id}/sync | Sync export storage
[**ApiStoragesExportLocalfilesValidateCreate**](StorageLocalApi.md#ApiStoragesExportLocalfilesValidateCreate) | **Post** /api/storages/export/localfiles/validate | Validate export storage
[**ApiStoragesLocalfilesCreate**](StorageLocalApi.md#ApiStoragesLocalfilesCreate) | **Post** /api/storages/localfiles/ | Create import storage
[**ApiStoragesLocalfilesDelete**](StorageLocalApi.md#ApiStoragesLocalfilesDelete) | **Delete** /api/storages/localfiles/{id} | Delete import storage
[**ApiStoragesLocalfilesList**](StorageLocalApi.md#ApiStoragesLocalfilesList) | **Get** /api/storages/localfiles/ | Get all import storage
[**ApiStoragesLocalfilesPartialUpdate**](StorageLocalApi.md#ApiStoragesLocalfilesPartialUpdate) | **Patch** /api/storages/localfiles/{id} | Update import storage
[**ApiStoragesLocalfilesRead**](StorageLocalApi.md#ApiStoragesLocalfilesRead) | **Get** /api/storages/localfiles/{id} | Get import storage
[**ApiStoragesLocalfilesSyncCreate**](StorageLocalApi.md#ApiStoragesLocalfilesSyncCreate) | **Post** /api/storages/localfiles/{id}/sync | Sync import storage
[**ApiStoragesLocalfilesValidateCreate**](StorageLocalApi.md#ApiStoragesLocalfilesValidateCreate) | **Post** /api/storages/localfiles/validate | Validate import storage


# **ApiStoragesExportLocalfilesCreate**
> LocalFilesExportStorage ApiStoragesExportLocalfilesCreate(ctx, data)
Create export storage

Create a new local file export storage connection to store annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalFilesExportStorage**](LocalFilesExportStorage.md)|  | 

### Return type

[**LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesDelete**
> ApiStoragesExportLocalfilesDelete(ctx, id)
Delete export storage

Delete a specific local file export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files export storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesList**
> []LocalFilesExportStorage ApiStoragesExportLocalfilesList(ctx, optional)
Get all export storage

Get a list of all Local export storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageLocalApiApiStoragesExportLocalfilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageLocalApiApiStoragesExportLocalfilesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesPartialUpdate**
> LocalFilesExportStorage ApiStoragesExportLocalfilesPartialUpdate(ctx, id, data)
Update export storage

Update a specific local file export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files export storage. | 
  **data** | [**LocalFilesExportStorage**](LocalFilesExportStorage.md)|  | 

### Return type

[**LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesRead**
> LocalFilesExportStorage ApiStoragesExportLocalfilesRead(ctx, id)
Get export storage

Get a specific local file export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files export storage. | 

### Return type

[**LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesSyncCreate**
> LocalFilesExportStorage ApiStoragesExportLocalfilesSyncCreate(ctx, id, data)
Sync export storage

Sync tasks from a local file export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**LocalFilesExportStorage**](LocalFilesExportStorage.md)|  | 

### Return type

[**LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportLocalfilesValidateCreate**
> LocalFilesExportStorage ApiStoragesExportLocalfilesValidateCreate(ctx, data)
Validate export storage

Validate a specific local file export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalFilesExportStorage**](LocalFilesExportStorage.md)|  | 

### Return type

[**LocalFilesExportStorage**](LocalFilesExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesCreate**
> LocalFilesImportStorage ApiStoragesLocalfilesCreate(ctx, data)
Create import storage

Create a new local file import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalFilesImportStorage**](LocalFilesImportStorage.md)|  | 

### Return type

[**LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesDelete**
> ApiStoragesLocalfilesDelete(ctx, id)
Delete import storage

Delete a specific local import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files import storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesList**
> []LocalFilesImportStorage ApiStoragesLocalfilesList(ctx, optional)
Get all import storage

Get a list of all local file import storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageLocalApiApiStoragesLocalfilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageLocalApiApiStoragesLocalfilesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesPartialUpdate**
> LocalFilesImportStorage ApiStoragesLocalfilesPartialUpdate(ctx, id, data)
Update import storage

Update a specific local file import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files import storage. | 
  **data** | [**LocalFilesImportStorage**](LocalFilesImportStorage.md)|  | 

### Return type

[**LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesRead**
> LocalFilesImportStorage ApiStoragesLocalfilesRead(ctx, id)
Get import storage

Get a specific local file import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this local files import storage. | 

### Return type

[**LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesSyncCreate**
> LocalFilesImportStorage ApiStoragesLocalfilesSyncCreate(ctx, id, data)
Sync import storage

Sync tasks from a local file import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**LocalFilesImportStorage**](LocalFilesImportStorage.md)|  | 

### Return type

[**LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesLocalfilesValidateCreate**
> LocalFilesImportStorage ApiStoragesLocalfilesValidateCreate(ctx, data)
Validate import storage

Validate a specific local file import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalFilesImportStorage**](LocalFilesImportStorage.md)|  | 

### Return type

[**LocalFilesImportStorage**](LocalFilesImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

