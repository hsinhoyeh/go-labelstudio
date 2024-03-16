# \LabelsApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLabelLinksCreate**](LabelsApi.md#ApiLabelLinksCreate) | **Post** /api/label_links/ | Create label links
[**ApiLabelLinksDelete**](LabelsApi.md#ApiLabelLinksDelete) | **Delete** /api/label_links/{id}/ | Remove label link
[**ApiLabelLinksList**](LabelsApi.md#ApiLabelLinksList) | **Get** /api/label_links/ | List label links
[**ApiLabelLinksPartialUpdate**](LabelsApi.md#ApiLabelLinksPartialUpdate) | **Patch** /api/label_links/{id}/ | Update label link
[**ApiLabelLinksRead**](LabelsApi.md#ApiLabelLinksRead) | **Get** /api/label_links/{id}/ | Get label link
[**ApiLabelsBulkCreate**](LabelsApi.md#ApiLabelsBulkCreate) | **Post** /api/labels/bulk | Bulk update labels
[**ApiLabelsCreate**](LabelsApi.md#ApiLabelsCreate) | **Post** /api/labels/ | Create labels
[**ApiLabelsDelete**](LabelsApi.md#ApiLabelsDelete) | **Delete** /api/labels/{id}/ | Remove labels
[**ApiLabelsList**](LabelsApi.md#ApiLabelsList) | **Get** /api/labels/ | List labels
[**ApiLabelsPartialUpdate**](LabelsApi.md#ApiLabelsPartialUpdate) | **Patch** /api/labels/{id}/ | Update labels
[**ApiLabelsRead**](LabelsApi.md#ApiLabelsRead) | **Get** /api/labels/{id}/ | Get label


# **ApiLabelLinksCreate**
> LabelLink ApiLabelLinksCreate(ctx, data)
Create label links

Create label links to link new custom labels to your project labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LabelLink**](LabelLink.md)|  | 

### Return type

[**LabelLink**](LabelLink.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelLinksDelete**
> ApiLabelLinksDelete(ctx, id)
Remove label link

 Remove a label link that links custom labels to your project labeling configuration. If you remove a label link, the label stops being available for the project it was linked to. You can add a new label link at any time.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelLinksList**
> InlineResponse200 ApiLabelLinksList(ctx, optional)
List label links

List label links for a specific label and project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LabelsApiApiLabelLinksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiApiLabelLinksListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelLinksPartialUpdate**
> LabelLink ApiLabelLinksPartialUpdate(ctx, id, data)
Update label link

 Update a label link that links custom labels to a project labeling configuration, for example if the fromName,   toName, or name parameters for a tag in the labeling configuration change.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**LabelLink**](LabelLink.md)|  | 

### Return type

[**LabelLink**](LabelLink.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelLinksRead**
> LabelLink ApiLabelLinksRead(ctx, id)
Get label link

Get label links for a specific project configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**LabelLink**](LabelLink.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelsBulkCreate**
> ApiLabelsBulkCreate(ctx, )
Bulk update labels

 If you want to update the labels in saved annotations, use this endpoint. 

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

# **ApiLabelsCreate**
> []LabelCreate ApiLabelsCreate(ctx, data)
Create labels

Add labels to your project without updating the labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**[]LabelCreate**](LabelCreate.md)|  | 

### Return type

[**[]LabelCreate**](LabelCreate.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelsDelete**
> ApiLabelsDelete(ctx, id)
Remove labels

Remove labels from your project without updating the labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelsList**
> InlineResponse2001 ApiLabelsList(ctx, optional)
List labels

List all custom labels added to your project separately from the labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LabelsApiApiLabelsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiApiLabelsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| A page number within the paginated result set. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelsPartialUpdate**
> Label ApiLabelsPartialUpdate(ctx, id, data)
Update labels

Update labels used for your project without updating the labeling configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **data** | [**Label**](Label.md)|  | 

### Return type

[**Label**](Label.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiLabelsRead**
> Label ApiLabelsRead(ctx, id)
Get label

 Retrieve a specific custom label used for your project by its ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Label**](Label.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

