# \StorageS3Api

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesExportS3Create**](StorageS3Api.md#ApiStoragesExportS3Create) | **Post** /api/storages/export/s3 | Create export storage
[**ApiStoragesExportS3Delete**](StorageS3Api.md#ApiStoragesExportS3Delete) | **Delete** /api/storages/export/s3/{id} | Delete export storage
[**ApiStoragesExportS3List**](StorageS3Api.md#ApiStoragesExportS3List) | **Get** /api/storages/export/s3 | Get all export storage
[**ApiStoragesExportS3PartialUpdate**](StorageS3Api.md#ApiStoragesExportS3PartialUpdate) | **Patch** /api/storages/export/s3/{id} | Update export storage
[**ApiStoragesExportS3Read**](StorageS3Api.md#ApiStoragesExportS3Read) | **Get** /api/storages/export/s3/{id} | Get export storage
[**ApiStoragesExportS3SyncCreate**](StorageS3Api.md#ApiStoragesExportS3SyncCreate) | **Post** /api/storages/export/s3/{id}/sync | Sync export storage
[**ApiStoragesExportS3ValidateCreate**](StorageS3Api.md#ApiStoragesExportS3ValidateCreate) | **Post** /api/storages/export/s3/validate | Validate export storage
[**ApiStoragesS3Create**](StorageS3Api.md#ApiStoragesS3Create) | **Post** /api/storages/s3/ | Create new storage
[**ApiStoragesS3Delete**](StorageS3Api.md#ApiStoragesS3Delete) | **Delete** /api/storages/s3/{id} | Delete import storage
[**ApiStoragesS3List**](StorageS3Api.md#ApiStoragesS3List) | **Get** /api/storages/s3/ | Get import storage
[**ApiStoragesS3PartialUpdate**](StorageS3Api.md#ApiStoragesS3PartialUpdate) | **Patch** /api/storages/s3/{id} | Update import storage
[**ApiStoragesS3Read**](StorageS3Api.md#ApiStoragesS3Read) | **Get** /api/storages/s3/{id} | Get import storage
[**ApiStoragesS3SyncCreate**](StorageS3Api.md#ApiStoragesS3SyncCreate) | **Post** /api/storages/s3/{id}/sync | Sync import storage
[**ApiStoragesS3ValidateCreate**](StorageS3Api.md#ApiStoragesS3ValidateCreate) | **Post** /api/storages/s3/validate | Validate import storage


# **ApiStoragesExportS3Create**
> S3ExportStorage ApiStoragesExportS3Create(ctx, data)
Create export storage

Create a new S3 export storage connection to store annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**S3ExportStorage**](S3ExportStorage.md)|  | 

### Return type

[**S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3Delete**
> ApiStoragesExportS3Delete(ctx, id)
Delete export storage

Delete a specific S3 export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 export storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3List**
> []S3ExportStorage ApiStoragesExportS3List(ctx, optional)
Get all export storage

Get a list of all S3 export storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageS3ApiApiStoragesExportS3ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageS3ApiApiStoragesExportS3ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3PartialUpdate**
> S3ExportStorage ApiStoragesExportS3PartialUpdate(ctx, id, data)
Update export storage

Update a specific S3 export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 export storage. | 
  **data** | [**S3ExportStorage**](S3ExportStorage.md)|  | 

### Return type

[**S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3Read**
> S3ExportStorage ApiStoragesExportS3Read(ctx, id)
Get export storage

Get a specific S3 export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 export storage. | 

### Return type

[**S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3SyncCreate**
> S3ExportStorage ApiStoragesExportS3SyncCreate(ctx, id, data)
Sync export storage

Sync tasks from an S3 export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**S3ExportStorage**](S3ExportStorage.md)|  | 

### Return type

[**S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportS3ValidateCreate**
> S3ExportStorage ApiStoragesExportS3ValidateCreate(ctx, data)
Validate export storage

Validate a specific S3 export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**S3ExportStorage**](S3ExportStorage.md)|  | 

### Return type

[**S3ExportStorage**](S3ExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3Create**
> S3ImportStorage ApiStoragesS3Create(ctx, data)
Create new storage

Get new S3 import storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**S3ImportStorage**](S3ImportStorage.md)|  | 

### Return type

[**S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3Delete**
> ApiStoragesS3Delete(ctx, id)
Delete import storage

Delete a specific S3 import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 import storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3List**
> []S3ImportStorage ApiStoragesS3List(ctx, optional)
Get import storage

Get a list of all S3 import storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageS3ApiApiStoragesS3ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageS3ApiApiStoragesS3ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3PartialUpdate**
> S3ImportStorage ApiStoragesS3PartialUpdate(ctx, id, data)
Update import storage

Update a specific S3 import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 import storage. | 
  **data** | [**S3ImportStorage**](S3ImportStorage.md)|  | 

### Return type

[**S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3Read**
> S3ImportStorage ApiStoragesS3Read(ctx, id)
Get import storage

Get a specific S3 import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this s3 import storage. | 

### Return type

[**S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3SyncCreate**
> S3ImportStorage ApiStoragesS3SyncCreate(ctx, id, data)
Sync import storage

Sync tasks from an S3 import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**S3ImportStorage**](S3ImportStorage.md)|  | 

### Return type

[**S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesS3ValidateCreate**
> S3ImportStorage ApiStoragesS3ValidateCreate(ctx, data)
Validate import storage

Validate a specific S3 import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**S3ImportStorage**](S3ImportStorage.md)|  | 

### Return type

[**S3ImportStorage**](S3ImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

