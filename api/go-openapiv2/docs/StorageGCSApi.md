# \StorageGCSApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesExportGcsCreate**](StorageGCSApi.md#ApiStoragesExportGcsCreate) | **Post** /api/storages/export/gcs | Create export storage
[**ApiStoragesExportGcsDelete**](StorageGCSApi.md#ApiStoragesExportGcsDelete) | **Delete** /api/storages/export/gcs/{id} | Delete export storage
[**ApiStoragesExportGcsList**](StorageGCSApi.md#ApiStoragesExportGcsList) | **Get** /api/storages/export/gcs | Get all export storage
[**ApiStoragesExportGcsPartialUpdate**](StorageGCSApi.md#ApiStoragesExportGcsPartialUpdate) | **Patch** /api/storages/export/gcs/{id} | Update export storage
[**ApiStoragesExportGcsRead**](StorageGCSApi.md#ApiStoragesExportGcsRead) | **Get** /api/storages/export/gcs/{id} | Get export storage
[**ApiStoragesExportGcsSyncCreate**](StorageGCSApi.md#ApiStoragesExportGcsSyncCreate) | **Post** /api/storages/export/gcs/{id}/sync | Sync export storage
[**ApiStoragesExportGcsValidateCreate**](StorageGCSApi.md#ApiStoragesExportGcsValidateCreate) | **Post** /api/storages/export/gcs/validate | Validate export storage
[**ApiStoragesGcsCreate**](StorageGCSApi.md#ApiStoragesGcsCreate) | **Post** /api/storages/gcs/ | Create import storage
[**ApiStoragesGcsDelete**](StorageGCSApi.md#ApiStoragesGcsDelete) | **Delete** /api/storages/gcs/{id} | Delete import storage
[**ApiStoragesGcsList**](StorageGCSApi.md#ApiStoragesGcsList) | **Get** /api/storages/gcs/ | Get all import storage
[**ApiStoragesGcsPartialUpdate**](StorageGCSApi.md#ApiStoragesGcsPartialUpdate) | **Patch** /api/storages/gcs/{id} | Update import storage
[**ApiStoragesGcsRead**](StorageGCSApi.md#ApiStoragesGcsRead) | **Get** /api/storages/gcs/{id} | Get import storage
[**ApiStoragesGcsSyncCreate**](StorageGCSApi.md#ApiStoragesGcsSyncCreate) | **Post** /api/storages/gcs/{id}/sync | Sync import storage
[**ApiStoragesGcsValidateCreate**](StorageGCSApi.md#ApiStoragesGcsValidateCreate) | **Post** /api/storages/gcs/validate | Validate import storage


# **ApiStoragesExportGcsCreate**
> GcsExportStorage ApiStoragesExportGcsCreate(ctx, data)
Create export storage

Create a new GCS export storage connection to store annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GcsExportStorage**](GcsExportStorage.md)|  | 

### Return type

[**GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsDelete**
> ApiStoragesExportGcsDelete(ctx, id)
Delete export storage

Delete a specific GCS export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs export storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsList**
> []GcsExportStorage ApiStoragesExportGcsList(ctx, optional)
Get all export storage

Get a list of all GCS export storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageGCSApiApiStoragesExportGcsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageGCSApiApiStoragesExportGcsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsPartialUpdate**
> GcsExportStorage ApiStoragesExportGcsPartialUpdate(ctx, id, data)
Update export storage

Update a specific GCS export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs export storage. | 
  **data** | [**GcsExportStorage**](GcsExportStorage.md)|  | 

### Return type

[**GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsRead**
> GcsExportStorage ApiStoragesExportGcsRead(ctx, id)
Get export storage

Get a specific GCS export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs export storage. | 

### Return type

[**GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsSyncCreate**
> GcsExportStorage ApiStoragesExportGcsSyncCreate(ctx, id, data)
Sync export storage

Sync tasks from an GCS export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**GcsExportStorage**](GcsExportStorage.md)|  | 

### Return type

[**GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportGcsValidateCreate**
> GcsExportStorage ApiStoragesExportGcsValidateCreate(ctx, data)
Validate export storage

Validate a specific GCS export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GcsExportStorage**](GcsExportStorage.md)|  | 

### Return type

[**GcsExportStorage**](GCSExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsCreate**
> GcsImportStorage ApiStoragesGcsCreate(ctx, data)
Create import storage

Create a new GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GcsImportStorage**](GcsImportStorage.md)|  | 

### Return type

[**GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsDelete**
> ApiStoragesGcsDelete(ctx, id)
Delete import storage

Delete a specific GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs import storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsList**
> []GcsImportStorage ApiStoragesGcsList(ctx, optional)
Get all import storage

Get a list of all GCS import storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageGCSApiApiStoragesGcsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageGCSApiApiStoragesGcsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsPartialUpdate**
> GcsImportStorage ApiStoragesGcsPartialUpdate(ctx, id, data)
Update import storage

Update a specific GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs import storage. | 
  **data** | [**GcsImportStorage**](GcsImportStorage.md)|  | 

### Return type

[**GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsRead**
> GcsImportStorage ApiStoragesGcsRead(ctx, id)
Get import storage

Get a specific GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this gcs import storage. | 

### Return type

[**GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsSyncCreate**
> GcsImportStorage ApiStoragesGcsSyncCreate(ctx, id, data)
Sync import storage

Sync tasks from an GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**GcsImportStorage**](GcsImportStorage.md)|  | 

### Return type

[**GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesGcsValidateCreate**
> GcsImportStorage ApiStoragesGcsValidateCreate(ctx, data)
Validate import storage

Validate a specific GCS import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GcsImportStorage**](GcsImportStorage.md)|  | 

### Return type

[**GcsImportStorage**](GCSImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

