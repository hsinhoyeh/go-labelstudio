# \ImportApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiImportFileUploadDelete**](ImportApi.md#ApiImportFileUploadDelete) | **Delete** /api/import/file-upload/{id} | Delete file upload
[**ApiImportFileUploadPartialUpdate**](ImportApi.md#ApiImportFileUploadPartialUpdate) | **Patch** /api/import/file-upload/{id} | Update file upload
[**ApiImportFileUploadRead**](ImportApi.md#ApiImportFileUploadRead) | **Get** /api/import/file-upload/{id} | Get file upload
[**ApiProjectsFileUploadsDelete**](ImportApi.md#ApiProjectsFileUploadsDelete) | **Delete** /api/projects/{id}/file-uploads | Delete files
[**ApiProjectsFileUploadsList**](ImportApi.md#ApiProjectsFileUploadsList) | **Get** /api/projects/{id}/file-uploads | Get files list
[**ApiProjectsImportCreate**](ImportApi.md#ApiProjectsImportCreate) | **Post** /api/projects/{id}/import | Import tasks


# **ApiImportFileUploadDelete**
> ApiImportFileUploadDelete(ctx, id)
Delete file upload

Delete a specific uploaded file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this file upload. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiImportFileUploadPartialUpdate**
> FileUpload ApiImportFileUploadPartialUpdate(ctx, id, data)
Update file upload

Update a specific uploaded file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this file upload. | 
  **data** | [**FileUpload**](FileUpload.md)|  | 

### Return type

[**FileUpload**](FileUpload.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiImportFileUploadRead**
> FileUpload ApiImportFileUploadRead(ctx, id)
Get file upload

Retrieve details about a specific uploaded file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this file upload. | 

### Return type

[**FileUpload**](FileUpload.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsFileUploadsDelete**
> ApiProjectsFileUploadsDelete(ctx, id)
Delete files

 Delete uploaded files for a specific project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this file upload. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsFileUploadsList**
> []FileUpload ApiProjectsFileUploadsList(ctx, id, optional)
Get files list

 Retrieve the list of uploaded files used to create labeling tasks for a specific project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this file upload. | 
 **optional** | ***ImportApiApiProjectsFileUploadsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportApiApiProjectsFileUploadsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **all** | **optional.Bool**| Set to \&quot;true\&quot; if you want to retrieve all file uploads | 
 **ids** | [**optional.Interface of []int32**](int32.md)| Specify the list of file upload IDs to retrieve, e.g. ids&#x3D;[1,2,3] | 

### Return type

[**[]FileUpload**](FileUpload.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsImportCreate**
> TaskCreationResponse ApiProjectsImportCreate(ctx, data, id)
Import tasks

 Import data as labeling tasks in bulk using this API endpoint. You can use this API endpoint to import multiple tasks. One POST request is limited at 250K tasks and 200 MB.  **Note:** Imported data is verified against a project *label_config* and must include all variables that were used in the *label_config*. For example, if the label configuration has a *$text* variable, then each item in a data object must include a \"text\" field. <br>  ## POST requests <hr style=\"opacity:0.3\">  There are three possible ways to import tasks with this endpoint:  ### 1\\. **POST with data** Send JSON tasks as POST data. Only JSON is supported for POSTing files directly. Update this example to specify your authorization token and Label Studio instance host, then run the following from the command line.  ```bash curl -H 'Content-Type: application/json' -H 'Authorization: Token abc123' \\ -X POST 'https://localhost:8080/api/projects/1/import' --data '[{\"text\": \"Some text 1\"}, {\"text\": \"Some text 2\"}]' ```  ### 2\\. **POST with files** Send tasks as files. You can attach multiple files with different names.  - **JSON**: text files in JavaScript object notation format - **CSV**: text files with tables in Comma Separated Values format - **TSV**: text files with tables in Tab Separated Value format - **TXT**: simple text files are similar to CSV with one column and no header, supported for projects with one source only  Update this example to specify your authorization token, Label Studio instance host, and file name and path, then run the following from the command line:  ```bash curl -H 'Authorization: Token abc123' \\ -X POST 'https://localhost:8080/api/projects/1/import' -F ‘file=@path/to/my_file.csv’ ```  ### 3\\. **POST with URL** You can also provide a URL to a file with labeling tasks. Supported file formats are the same as in option 2.  ```bash curl -H 'Content-Type: application/json' -H 'Authorization: Token abc123' \\ -X POST 'https://localhost:8080/api/projects/1/import' \\ --data '[{\"url\": \"http://example.com/test1.csv\"}, {\"url\": \"http://example.com/test2.csv\"}]' ```  <br> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ImportApi**](ImportApi.md)|  | 
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

[**TaskCreationResponse**](Task creation response.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

