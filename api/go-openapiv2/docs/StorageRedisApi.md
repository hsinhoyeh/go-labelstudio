# \StorageRedisApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStoragesExportRedisCreate**](StorageRedisApi.md#ApiStoragesExportRedisCreate) | **Post** /api/storages/export/redis | Create export storage
[**ApiStoragesExportRedisDelete**](StorageRedisApi.md#ApiStoragesExportRedisDelete) | **Delete** /api/storages/export/redis/{id} | Delete export storage
[**ApiStoragesExportRedisList**](StorageRedisApi.md#ApiStoragesExportRedisList) | **Get** /api/storages/export/redis | Get all export storage
[**ApiStoragesExportRedisPartialUpdate**](StorageRedisApi.md#ApiStoragesExportRedisPartialUpdate) | **Patch** /api/storages/export/redis/{id} | Update export storage
[**ApiStoragesExportRedisRead**](StorageRedisApi.md#ApiStoragesExportRedisRead) | **Get** /api/storages/export/redis/{id} | Get export storage
[**ApiStoragesExportRedisSyncCreate**](StorageRedisApi.md#ApiStoragesExportRedisSyncCreate) | **Post** /api/storages/export/redis/{id}/sync | Sync export storage
[**ApiStoragesExportRedisValidateCreate**](StorageRedisApi.md#ApiStoragesExportRedisValidateCreate) | **Post** /api/storages/export/redis/validate | Validate export storage
[**ApiStoragesRedisCreate**](StorageRedisApi.md#ApiStoragesRedisCreate) | **Post** /api/storages/redis/ | Create import storage
[**ApiStoragesRedisDelete**](StorageRedisApi.md#ApiStoragesRedisDelete) | **Delete** /api/storages/redis/{id} | Delete import storage
[**ApiStoragesRedisList**](StorageRedisApi.md#ApiStoragesRedisList) | **Get** /api/storages/redis/ | Get all import storage
[**ApiStoragesRedisPartialUpdate**](StorageRedisApi.md#ApiStoragesRedisPartialUpdate) | **Patch** /api/storages/redis/{id} | Update import storage
[**ApiStoragesRedisRead**](StorageRedisApi.md#ApiStoragesRedisRead) | **Get** /api/storages/redis/{id} | Get import storage
[**ApiStoragesRedisSyncCreate**](StorageRedisApi.md#ApiStoragesRedisSyncCreate) | **Post** /api/storages/redis/{id}/sync | Sync import storage
[**ApiStoragesRedisValidateCreate**](StorageRedisApi.md#ApiStoragesRedisValidateCreate) | **Post** /api/storages/redis/validate | Validate import storage


# **ApiStoragesExportRedisCreate**
> RedisExportStorage ApiStoragesExportRedisCreate(ctx, data)
Create export storage

Create a new Redis export storage connection to store annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RedisExportStorage**](RedisExportStorage.md)|  | 

### Return type

[**RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisDelete**
> ApiStoragesExportRedisDelete(ctx, id)
Delete export storage

Delete a specific Redis export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis export storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisList**
> []RedisExportStorage ApiStoragesExportRedisList(ctx, optional)
Get all export storage

Get a list of all Redis export storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageRedisApiApiStoragesExportRedisListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageRedisApiApiStoragesExportRedisListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisPartialUpdate**
> RedisExportStorage ApiStoragesExportRedisPartialUpdate(ctx, id, data)
Update export storage

Update a specific Redis export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis export storage. | 
  **data** | [**RedisExportStorage**](RedisExportStorage.md)|  | 

### Return type

[**RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisRead**
> RedisExportStorage ApiStoragesExportRedisRead(ctx, id)
Get export storage

Get a specific Redis export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis export storage. | 

### Return type

[**RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisSyncCreate**
> RedisExportStorage ApiStoragesExportRedisSyncCreate(ctx, id, data)
Sync export storage

Sync tasks from a specific Redis export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RedisExportStorage**](RedisExportStorage.md)|  | 

### Return type

[**RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesExportRedisValidateCreate**
> RedisExportStorage ApiStoragesExportRedisValidateCreate(ctx, data)
Validate export storage

Validate a specific Redis export storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RedisExportStorage**](RedisExportStorage.md)|  | 

### Return type

[**RedisExportStorage**](RedisExportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisCreate**
> RedisImportStorage ApiStoragesRedisCreate(ctx, data)
Create import storage

Create a new Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RedisImportStorage**](RedisImportStorage.md)|  | 

### Return type

[**RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisDelete**
> ApiStoragesRedisDelete(ctx, id)
Delete import storage

Delete a specific Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis import storage. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisList**
> []RedisImportStorage ApiStoragesRedisList(ctx, optional)
Get all import storage

Get a list of all Redis import storage connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageRedisApiApiStoragesRedisListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageRedisApiApiStoragesRedisListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.Int32**| Project ID | 

### Return type

[**[]RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisPartialUpdate**
> RedisImportStorage ApiStoragesRedisPartialUpdate(ctx, id, data)
Update import storage

Update a specific Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis import storage. | 
  **data** | [**RedisImportStorage**](RedisImportStorage.md)|  | 

### Return type

[**RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisRead**
> RedisImportStorage ApiStoragesRedisRead(ctx, id)
Get import storage

Get a specific Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this redis import storage. | 

### Return type

[**RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisSyncCreate**
> RedisImportStorage ApiStoragesRedisSyncCreate(ctx, id, data)
Sync import storage

Sync tasks from a specific Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**RedisImportStorage**](RedisImportStorage.md)|  | 

### Return type

[**RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStoragesRedisValidateCreate**
> RedisImportStorage ApiStoragesRedisValidateCreate(ctx, data)
Validate import storage

Validate a specific Redis import storage connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RedisImportStorage**](RedisImportStorage.md)|  | 

### Return type

[**RedisImportStorage**](RedisImportStorage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

