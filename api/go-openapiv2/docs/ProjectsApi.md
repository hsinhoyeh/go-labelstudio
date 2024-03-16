# \ProjectsApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsCreate**](ProjectsApi.md#ApiProjectsCreate) | **Post** /api/projects/ | Create new project
[**ApiProjectsDelete**](ProjectsApi.md#ApiProjectsDelete) | **Delete** /api/projects/{id}/ | Delete project
[**ApiProjectsImportsRead**](ProjectsApi.md#ApiProjectsImportsRead) | **Get** /api/projects/{id}/imports/{import_pk}/ | Get project import info
[**ApiProjectsList**](ProjectsApi.md#ApiProjectsList) | **Get** /api/projects/ | List your projects
[**ApiProjectsPartialUpdate**](ProjectsApi.md#ApiProjectsPartialUpdate) | **Patch** /api/projects/{id}/ | Update project
[**ApiProjectsRead**](ProjectsApi.md#ApiProjectsRead) | **Get** /api/projects/{id}/ | Get project by ID
[**ApiProjectsReimportsRead**](ProjectsApi.md#ApiProjectsReimportsRead) | **Get** /api/projects/{id}/reimports/{reimport_pk}/ | Get project reimport info
[**ApiProjectsTasksDelete**](ProjectsApi.md#ApiProjectsTasksDelete) | **Delete** /api/projects/{id}/tasks/ | Delete all tasks
[**ApiProjectsTasksList**](ProjectsApi.md#ApiProjectsTasksList) | **Get** /api/projects/{id}/tasks/ | List project tasks
[**ApiProjectsValidateCreate**](ProjectsApi.md#ApiProjectsValidateCreate) | **Post** /api/projects/validate/ | Validate label config
[**ApiProjectsValidateCreate_0**](ProjectsApi.md#ApiProjectsValidateCreate_0) | **Post** /api/projects/{id}/validate/ | Validate project label config


# **ApiProjectsCreate**
> Project ApiProjectsCreate(ctx, data)
Create new project

 Create a project and set up the labeling interface in Label Studio using the API.  ```bash curl -H Content-Type:application/json -H 'Authorization: Token abc123' -X POST 'https://localhost:8080/api/projects'     --data '{\"label_config\": \"<View>[...]</View>\"}' ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsDelete**
> ApiProjectsDelete(ctx, id)
Delete project

Delete a project by specified project ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsImportsRead**
> ProjectImport ApiProjectsImportsRead(ctx, importPk, id)
Get project import info

Return data related to async project import operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importPk** | **string**|  | 
  **id** | **int32**| A unique integer value identifying this project import. | 

### Return type

[**ProjectImport**](ProjectImport.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsList**
> InlineResponse2003 ApiProjectsList(ctx, optional)
List your projects

 Return a list of the projects that you've created.  To perform most tasks with the Label Studio API, you must specify the project ID, sometimes referred to as the `pk`. To retrieve a list of your Label Studio projects, update the following command to match your own environment. Replace the domain name, port, and authorization token, then run the following from the command line: ```bash curl -X GET https://localhost:8080/api/projects/ -H 'Authorization: Token abc123' ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiApiProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiApiProjectsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **optional.String**| Which field to use when ordering the results. | 
 **ids** | **optional.String**| ids | 
 **title** | **optional.String**| title | 
 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsPartialUpdate**
> Project ApiProjectsPartialUpdate(ctx, id, data)
Update project

Update the project settings for a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
  **data** | [**Project**](Project.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsRead**
> Project ApiProjectsRead(ctx, id)
Get project by ID

Retrieve information about a project by project ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

[**Project**](Project.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsReimportsRead**
> ProjectReimport ApiProjectsReimportsRead(ctx, reimportPk, id)
Get project reimport info

Return data related to async project reimport operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reimportPk** | **string**|  | 
  **id** | **int32**| A unique integer value identifying this project reimport. | 

### Return type

[**ProjectReimport**](ProjectReimport.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsTasksDelete**
> ApiProjectsTasksDelete(ctx, id)
Delete all tasks

Delete all tasks from a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsTasksList**
> ApiProjectsTasksList(ctx, id, optional)
List project tasks

 Retrieve a paginated list of tasks for a specific project. For example, use the following cURL command: ```bash curl -X GET https://localhost:8080/api/projects/{id}/tasks/?page=1&page_size=10 -H 'Authorization: Token abc123' ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this project. | 
 **optional** | ***ProjectsApiApiProjectsTasksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiApiProjectsTasksListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| [or \&quot;start\&quot;] current page | 
 **pageSize** | **optional.Int32**| [or \&quot;length\&quot;] tasks per page, use -1 to obtain all tasks (in this case \&quot;page\&quot; has no effect and this operation might be slow) | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsValidateCreate**
> ApiProjectsValidateCreate(ctx, data)
Validate label config

Validate an arbitrary labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProjectLabelConfig**](ProjectLabelConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProjectsValidateCreate_0**
> ProjectLabelConfig ApiProjectsValidateCreate_0(ctx, data, id)
Validate project label config

 Determine whether the label configuration for a specific project is valid. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProjectLabelConfig**](ProjectLabelConfig.md)|  | 
  **id** | **int32**| A unique integer value identifying this project. | 

### Return type

[**ProjectLabelConfig**](ProjectLabelConfig.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

