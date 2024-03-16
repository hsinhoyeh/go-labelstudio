# \ExportApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsExportFormatsRead**](ExportApi.md#ApiProjectsExportFormatsRead) | **Get** /api/projects/{id}/export/formats | Get export formats
[**ApiProjectsExportRead**](ExportApi.md#ApiProjectsExportRead) | **Get** /api/projects/{id}/export | Easy export of tasks and annotations
[**ApiProjectsExportsConvertCreate**](ExportApi.md#ApiProjectsExportsConvertCreate) | **Post** /api/projects/{id}/exports/{export_pk}/convert | Export conversion
[**ApiProjectsExportsCreate**](ExportApi.md#ApiProjectsExportsCreate) | **Post** /api/projects/{id}/exports/ | Create new export snapshot
[**ApiProjectsExportsDelete**](ExportApi.md#ApiProjectsExportsDelete) | **Delete** /api/projects/{id}/exports/{export_pk} | Delete export snapshot
[**ApiProjectsExportsDownloadRead**](ExportApi.md#ApiProjectsExportsDownloadRead) | **Get** /api/projects/{id}/exports/{export_pk}/download | Download export snapshot as file in specified format
[**ApiProjectsExportsList**](ExportApi.md#ApiProjectsExportsList) | **Get** /api/projects/{id}/exports/ | List all export snapshots
[**ApiProjectsExportsRead**](ExportApi.md#ApiProjectsExportsRead) | **Get** /api/projects/{id}/exports/{export_pk} | Get export snapshot by ID


# **ApiProjectsExportFormatsRead**
> []string ApiProjectsExportFormatsRead(ctx, id)
Get export formats

Retrieve the available export formats for the current project by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

**[]string**

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportRead**
> *os.File ApiProjectsExportRead(ctx, id, optional)
Easy export of tasks and annotations

 <i>Note: if you have a large project it's recommended to use  export snapshots, this easy export endpoint might have timeouts.</i><br/><br> Export annotated tasks as a file in a specific format.  For example, to export JSON annotations for a project to a file called `annotations.json`, run the following from the command line: ```bash curl -X GET https://localhost:8080/api/projects/{id}/export?exportType=JSON -H 'Authorization: Token abc123' --output 'annotations.json' ``` To export all tasks, including skipped tasks and others without annotations, run the following from the command line: ```bash curl -X GET https://localhost:8080/api/projects/{id}/export?exportType=JSON&download_all_tasks=true -H 'Authorization: Token abc123' --output 'annotations.json' ``` To export specific tasks with IDs of 123 and 345, run the following from the command line: ```bash curl -X GET https://localhost:8080/api/projects/{id}/export?ids[]=123\\&ids[]=345 -H 'Authorization: Token abc123' --output 'annotations.json' ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
 **optional** | ***ExportApiApiProjectsExportReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiApiProjectsExportReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportType** | **optional.String**| Selected export format (JSON by default) | 
 **downloadAllTasks** | **optional.String**|  If true, download all tasks regardless of status. If false, download only annotated tasks.  | 
 **downloadResources** | **optional.Bool**|  If true, download all resource files such as images, audio, and others relevant to the tasks.   | 
 **ids** | [**optional.Interface of []int32**](int32.md)|  Specify a list of task IDs to retrieve only the details for those tasks.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsConvertCreate**
> ExportConvert ApiProjectsExportsConvertCreate(ctx, data, id, exportPk)
Export conversion

 Convert export snapshot to selected format 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ExportConvert**](ExportConvert.md)|  | 
  **id** | **int32**| A unique integer value identifying this project. | 
  **exportPk** | **string**| Primary key identifying the export file. | 

### Return type

[**ExportConvert**](ExportConvert.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsCreate**
> ExportCreate ApiProjectsExportsCreate(ctx, data, id)
Create new export snapshot

 Create a new export request to start a background task and generate an export file for a specific project by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ExportCreate**](ExportCreate.md)|  | 
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

[**ExportCreate**](ExportCreate.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsDelete**
> ApiProjectsExportsDelete(ctx, id, exportPk)
Delete export snapshot

 Delete an export file by specified export ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
  **exportPk** | **string**| Primary key identifying the export file. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsDownloadRead**
> ApiProjectsExportsDownloadRead(ctx, id, exportPk, optional)
Download export snapshot as file in specified format

 Download an export file in the specified format for a specific project. Specify the project ID with the `id`  parameter in the path and the ID of the export file you want to download using the `export_pk` parameter  in the path.   Get the `export_pk` from the response of the request to [Create new export](/api#operation/api_projects_exports_create) or after [listing export files](/api#operation/api_projects_exports_list). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
  **exportPk** | **string**| Primary key identifying the export file. | 
 **optional** | ***ExportApiApiProjectsExportsDownloadReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiApiProjectsExportsDownloadReadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exportType** | **optional.String**| Selected export format | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsList**
> []Export ApiProjectsExportsList(ctx, id)
List all export snapshots

 Returns a list of exported files for a specific project by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

[**[]Export**](Export.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsExportsRead**
> Export ApiProjectsExportsRead(ctx, id, exportPk)
Get export snapshot by ID

 Retrieve information about an export file by export ID for a specific project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
  **exportPk** | **string**| Primary key identifying the export file. | 

### Return type

[**Export**](Export.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

