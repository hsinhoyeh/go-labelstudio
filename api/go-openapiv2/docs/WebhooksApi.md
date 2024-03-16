# \WebhooksApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiWebhooksCreate**](WebhooksApi.md#ApiWebhooksCreate) | **Post** /api/webhooks/ | Create a webhook
[**ApiWebhooksDelete**](WebhooksApi.md#ApiWebhooksDelete) | **Delete** /api/webhooks/{id}/ | Delete webhook info
[**ApiWebhooksInfoList**](WebhooksApi.md#ApiWebhooksInfoList) | **Get** /api/webhooks/info/ | Get all webhook actions
[**ApiWebhooksList**](WebhooksApi.md#ApiWebhooksList) | **Get** /api/webhooks/ | List all webhooks
[**ApiWebhooksPartialUpdate**](WebhooksApi.md#ApiWebhooksPartialUpdate) | **Patch** /api/webhooks/{id}/ | Update webhook info
[**ApiWebhooksRead**](WebhooksApi.md#ApiWebhooksRead) | **Get** /api/webhooks/{id}/ | Get webhook info
[**ApiWebhooksUpdate**](WebhooksApi.md#ApiWebhooksUpdate) | **Put** /api/webhooks/{id}/ | Save webhook info


# **ApiWebhooksCreate**
> Webhook ApiWebhooksCreate(ctx, data)
Create a webhook

Create a webhook for your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Webhook**](Webhook.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksDelete**
> ApiWebhooksDelete(ctx, id)
Delete webhook info



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this webhook. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksInfoList**
> ApiWebhooksInfoList(ctx, optional)
Get all webhook actions

Get descriptions of all available webhook actions to set up webhooks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiApiWebhooksInfoListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiApiWebhooksInfoListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationOnly** | **optional.Bool**| organization-only or not | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksList**
> []Webhook ApiWebhooksList(ctx, optional)
List all webhooks

List all webhooks set up for your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiApiWebhooksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiApiWebhooksListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.String**| Project ID | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksPartialUpdate**
> WebhookSerializerForUpdate ApiWebhooksPartialUpdate(ctx, id, data, url, optional)
Update webhook info



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this webhook. | 
  **data** | [**WebhookSerializerForUpdate**](WebhookSerializerForUpdate.md)|  | 
  **url** | **string**| URL of webhook | 
 **optional** | ***WebhooksApiApiWebhooksPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiApiWebhooksPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sendPayload** | **optional.Bool**| If value is False send only action | 
 **sendForAllActions** | **optional.Bool**| If value is False - used only for actions from WebhookAction | 
 **headers** | **optional.String**| Key Value Json of headers | 
 **isActive** | **optional.Bool**| If value is False the webhook is disabled | 
 **actions** | [**optional.Interface of []string**](string.md)|  | [default to []]

### Return type

[**WebhookSerializerForUpdate**](WebhookSerializerForUpdate.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksRead**
> Webhook ApiWebhooksRead(ctx, id)
Get webhook info



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this webhook. | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWebhooksUpdate**
> WebhookSerializerForUpdate ApiWebhooksUpdate(ctx, id, data, url, optional)
Save webhook info



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this webhook. | 
  **data** | [**WebhookSerializerForUpdate**](WebhookSerializerForUpdate.md)|  | 
  **url** | **string**| URL of webhook | 
 **optional** | ***WebhooksApiApiWebhooksUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiApiWebhooksUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sendPayload** | **optional.Bool**| If value is False send only action | 
 **sendForAllActions** | **optional.Bool**| If value is False - used only for actions from WebhookAction | 
 **headers** | **optional.String**| Key Value Json of headers | 
 **isActive** | **optional.Bool**| If value is False the webhook is disabled | 
 **actions** | [**optional.Interface of []string**](string.md)|  | [default to []]

### Return type

[**WebhookSerializerForUpdate**](WebhookSerializerForUpdate.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

